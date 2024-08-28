private val logger = LoggerFactory.getLogger(getClass)

  private def processCPG(
    xtocpg: Try[codepropertygraph.Cpg],
    privadoInput: PrivadoInput,
    ruleCache: RuleCache,
    sourceRepoLocation: String,
    dataFlowCache: DataFlowCache,
    auditCache: AuditCache,
    s3DatabaseDetailsCache: S3DatabaseDetailsCache,
    appCache: AppCache,
    propertyFilterCache: PropertyFilterCache
  ): Either[String, Unit] = {
    xtocpg match {
      case Success(cpg) =>
        try {
          statsRecorder.initiateNewStage("Additional passes")
          new NaiveCallLinker(cpg).createAndApply()
          if (privadoInput.enableLambdaFlows)
            new ExperimentalLambdaDataFlowSupportPass(cpg).createAndApply()
          statsRecorder.endLastStage()
          statsRecorder.setSupressSubstagesFlag(false)
          statsRecorder.initiateNewStage("Privado source passes")
          new HTMLParserPass(cpg, sourceRepoLocation, ruleCache, privadoInputConfig = privadoInput)
            .createAndApply()
          if (privadoInput.assetDiscovery) {
            new JsonPropertyParserPass(cpg, s"$sourceRepoLocation/${Constants.generatedConfigFolderName}")
              .createAndApply()
            new JsConfigPropertyPass(cpg).createAndApply()
          } else
            new PropertyParserPass(
              cpg,
              sourceRepoLocation,
              ruleCache,
              Language.JAVASCRIPT,
              propertyFilterCache,
              privadoInput
            )
              .createAndApply()
          new JSPropertyLinkerPass(cpg).createAndApply()
          new SQLParser(cpg, sourceRepoLocation, ruleCache).createAndApply()
          new DBTParserPass(cpg, sourceRepoLocation, ruleCache).createAndApply()
          new AndroidXmlParserPass(cpg, sourceRepoLocation, ruleCache).createAndApply()
          statsRecorder.endLastStage()
          // Unresolved function report
          if (privadoInput.showUnresolvedFunctionsReport) {
            val path = s"${privadoInput.sourceLocation.head}/${Constants.outputDirectoryName}"
            UnresolvedReportUtility.reportUnresolvedMethods(xtocpg, path, Language.JAVASCRIPT)
          }
          logger.info("=====================")

          // Run tagger
          statsRecorder.initiateNewStage("Tagger ...")
          val taggerCache = TaggerCache()
          cpg.runTagger(ruleCache, taggerCache, privadoInputConfig = privadoInput, dataFlowCache, appCache)
          statsRecorder.endLastStage()
          statsRecorder.initiateNewStage("Finding data flows ...")
          val dataflowMap =
            Dataflow(cpg, statsRecorder).dataflow(privadoInput, ruleCache, dataFlowCache, auditCache, appCache)
          statsRecorder.endLastStage()
          statsRecorder.justLogMessage(s"Processed final flows - ${dataFlowCache.getDataflowAfterDedup.size}")
          statsRecorder.initiateNewStage("Brewing result")
          MetricHandler.setScanStatus(true)
          val errorMsg = new ListBuffer[String]()
          // Exporting
          JSONExporter.fileExport(
            cpg,
            outputFileName,
            sourceRepoLocation,
            dataflowMap,
            ruleCache,
            taggerCache,
            dataFlowCache.getDataflowAfterDedup,
            privadoInput,
            List(),
            s3DatabaseDetailsCache,
            appCache,
            propertyFilterCache
          ) match {
            case Left(err) =>
              MetricHandler.otherErrorsOrWarnings.addOne(err)
              errorMsg += err
            case Right(_) =>
              println(s"Successfully exported output to '${appCache.localScanPath}/$outputDirectoryName' folder")
              logger.debug(
                s"Total Sinks identified : ${cpg.tag.where(_.nameExact(Constants.catLevelOne).valueExact(CatLevelOne.SINKS.name)).call.tag.nameExact(Constants.id).value.toSet}"
              )
              val codelist = cpg.call
                .whereNot(_.methodFullName(Operators.ALL.asScala.toSeq: _*))
                .map(item => (item.methodFullName, item.location.filename))
                .dedup
                .l
              logger.debug(s"size of code : ${codelist.size}")
              codelist.foreach(item => logger.debug(item._1, item._2))
              logger.debug("Above we printed methodFullName")

              Right(())
          }

          // Exporting the Audit report
          if (privadoInput.generateAuditReport) {
            ExcelExporter.auditExport(
              outputAuditFileName,
              AuditReportEntryPoint.getAuditWorkbookJS(xtocpg, taggerCache, sourceRepoLocation, auditCache, ruleCache),
              sourceRepoLocation
            ) match {
              case Left(err) =>
                MetricHandler.otherErrorsOrWarnings.addOne(err)
                errorMsg += err
              case Right(_) =>
                println(
                  s"${Calendar.getInstance().getTime} - Successfully exported Audit report to '${appCache.localScanPath}/$outputDirectoryName' folder..."
                )
            }
          }

          // Exporting the Intermediate report
          if (privadoInput.testOutput || privadoInput.generateAuditReport) {
            JSONExporter.IntermediateFileExport(
              outputIntermediateFileName,
              sourceRepoLocation,
              dataFlowCache.getJsonFormatDataFlow(dataFlowCache.getIntermediateDataFlow())
            ) match {
              case Left(err) =>
                MetricHandler.otherErrorsOrWarnings.addOne(err)
                errorMsg += err
              case Right(_) =>
                println(
                  s"${Calendar.getInstance().getTime} - Successfully exported intermediate output to '${appCache.localScanPath}/${Constants.outputDirectoryName}' folder..."
                )
            }

            // Exporting the Unresolved report
            JSONExporter.UnresolvedFlowFileExport(
              outputUnresolvedFilename,
              sourceRepoLocation,
              dataFlowCache.getJsonFormatDataFlow(auditCache.unfilteredFlow)
            ) match {
              case Left(err) =>
                MetricHandler.otherErrorsOrWarnings.addOne(err)
                errorMsg += err
              case Right(_) =>
                println(
                  s"${Calendar.getInstance().getTime} - Successfully exported Unresolved flow output to '${appCache.localScanPath}/${Constants.outputDirectoryName}' folder..."
                )
            }
          }
          statsRecorder.endLastStage()
          // Check if any of the export failed
          if (errorMsg.toList.isEmpty)
            Right(())
          else
            Left(errorMsg.toList.mkString("\n"))
        } finally {
          cpg.close()
          import java.io.File
          val cpgOutputPath = s"$sourceRepoLocation/$outputDirectoryName/$cpgOutputFileName"
          val cpgFile       = new File(cpgOutputPath)
          println(s"\n\n\nBinary file size -- ${cpgFile.length()} in Bytes - ${cpgFile.length() * 0.000001} MB\n\n\n")
        }
      case Failure(exception) =>
        logger.error("Error while parsing the source code!")
        logger.debug("Error : ", exception)
        MetricHandler.setScanStatus(false)
        Left("Error while parsing the source code: " + exception.toString)
    }
  }

  /** Create cpg using Javascript Language
    *
    * @param sourceRepoLocation
    * @param lang
    * @return
    */
  def createJavaScriptCpg(): Either[String, Unit] = {
    statsRecorder.justLogMessage("Processing source code using Javascript engine")
    statsRecorder.initiateNewStage("Base source processing")
    val cpgOutputPath = s"$sourceRepoLocation/$outputDirectoryName/$cpgOutputFileName"
    // Create the .privado folder if not present
    createCpgFolder(sourceRepoLocation);

    // Need to convert path to absolute path as javaScriptCpg need abolute path of repo
    val absoluteSourceLocation = File(sourceRepoLocation).path.toAbsolutePath.normalize().toString
    val excludeFileRegex       = ruleCache.getExclusionRegex
    val cpgconfig =
        .withInputPath(absoluteSourceLocation)
        .withOutputPath(cpgOutputPath)
        .withIgnoredFilesRegex(excludeFileRegex)
    val xtocpg = new JsSrc2Cpg().createCpgWithAllOverlays(cpgconfig)
    statsRecorder.endLastStage()
    processCPG(
      xtocpg,
      privadoInput,
)