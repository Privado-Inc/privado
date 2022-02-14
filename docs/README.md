# Privado User Documentation

## What is Privado? <a href="#what-is-privado" id="what-is-privado"></a>

Privado is a static code scanning tool to find, fix and remediate privacy issues in your products & applications. Our scan discovers what personal data(as defined by GDPR, other laws) your app is processing, third-party integrations, data flows. With our scan results, we generate privacy reports for your apps as mandated by laws like GDPR or platforms like Apple and keep them in sync with code changes.

## How Privado CLI handles your data?

Privado CLI tool was engineered with security in mind. Our tool runs the scan locally on your machine and your code never leaves your system.

## What does our scan discover? <a href="#what-does-our-scan-discover" id="what-does-our-scan-discover"></a>

1. Data Elements: These are personal data that your app is collecting, sharing, processing. Here is a list of data elements that we are discovering.
2. Third-Parties: Any third-party integrations inside your code, via APIs or SDKs/libraries.
3. APIs: We also discover any internal APIs that your app is connected with.
4. Datastores(not released yet, still beta): Identify the databases where you are sourcing the data from or sending the data.
5. Privacy Vulnerabilities(not released yet, still beta): Code issues that exist which can lead to privacy vulnerabilities

## What can I do with Privado?

### Generate Play Store Data Safety Report

This is the first use case that we are live with. Currently, to fill the data safety form Android developers have to ask around in the team to find what data they are collecting, spend hours reading the documentation of SDKs to find data shared, and navigate the complex Playstore form. With our scan, we pre-fill data types collected, shared, and guide you with our wizard to generate the data safety report.

### Privacy Audits <a href="#privacy-audits" id="privacy-audits"></a>

Privacy Engineers can use our CLI tool as an MRI for products, applications and find out privacy risks. With our scans, privacy engineers save the time they have to spend chasing engineers with assessments and can directly start prescribing privacy controls for data minimization, sharing, etc.

#### We have the following use cases on our Roadmap: <a href="#we-have-the-following-use-cases-on-our-roadmap" id="we-have-the-following-use-cases-on-our-roadmap"></a>

1. Generating Apple Nutrition Label Report
2. Generating privacy compliance reports like GDPR Article 30 or RoPA report
3. Detecting Privacy Vulnerabilities in current code implementation
4. Privacy Policy Generator

## What does it cost?

Privado is free for:

1. Open Source Projects
2. For individual developers and small teams.
