# Generating ROPA Report

{% hint style="info" %}
Before generating the ROPA report, ensure that you have scanned the repository on your machine using privado CLI and the results are available on the Privado Cloud dashboard. To learn how to scan your repository, [click here](getting-started-with-privado/running-a-scan.md).
{% endhint %}

Privado’s cloud dashboard can be used to generate a Record of Processing Activity (ROPA) report using a self-serve wizard. Learn more about ROPA reports and their requirements here. Generating the ROPA report is an easy 6 step process. To access the report generator, navigate to the desired repository → Click on the **Create ROPA Report button** to start the process.

![](<.gitbook/assets/image (1).png>)

**1. Repo Details**

Enter a brieF description of the repository including business goals it follow. Click Next. Optionally you can also Save if you want to resume report creation later

![](<.gitbook/assets/image (4).png>)

**2. Data Flows**

Review all the data flows in the app. Click View → Code Analysis to drill down each flow.

![](<.gitbook/assets/image (15).png>)

If you feel certain flows are incorrect and should not be part of the repo, click the **False Positive** button in the Code Analysis View to remove them from the analysis as well as the report

![](<.gitbook/assets/image (6).png>)

**3. Data Source**

Identify the category and source of the data you are processing in this repository

![](<.gitbook/assets/image (7).png>)

**4. Data Usage**

Provide information about the purpose of this repository, the security measures in place for this repo and the duration for with data elements will be retained for.

![](<.gitbook/assets/image (17).png>)

**5. Legal**

Provide details related to specific legal questions about the repository’s use

![](<.gitbook/assets/image (2).png>)

**6. Review Summary and Download Report**

We are done! You can now Review the report and download it in CSV or PDF format from this page

![](<.gitbook/assets/image (14).png>)
