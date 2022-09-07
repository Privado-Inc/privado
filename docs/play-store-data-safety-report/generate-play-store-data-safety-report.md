# Generate Play Store Data Safety Report

{% hint style="info" %}
Before generating the Data Safety Report, ensure that you have scanned the repository on your machine using privado CLI and the results are available on the Privado Cloud dashboard. To learn how to scan your repository, [click here](../getting-started-with-privado/running-a-scan.md).
{% endhint %}

Once your application has been scanned and results have been synced to the Privado Cloud dashboard, you can now generate Data Safety Reports for your Android application. Generating a Data Safety report may be mandatory requirement to publish your apps on the Google Play Store. To learn more about Data Safety Reports and why they might be needed for your repo, click here. Generating this report is a easy 3 step process. To begin, select your target repository → **Click Create Data Safety Report** and follow the instructions

![](<../.gitbook/assets/image (10).png>)

**1. Review Data**

Privado scans the source code of your application for over 200 data elements and looks at data collection like user forms, Android permissions, etc. to discover data elements your mobile app is collecting or sharing.

![](<../.gitbook/assets/image (16).png>)

To drill down, click **View**. On the **Code Analysis** tab, you can review individual data instances, their flows and mark them as **False Positive** to remove them from the results as well as the report if needed.

![](<../.gitbook/assets/image (13).png>)

**2. Review Data Usage**

Provide information about data collection and data security relevant to your application. Click Next to proceed.

![](<../.gitbook/assets/image (9).png>)

**3. Review Summary and Download Report**

And we are done! Review the summary and download report in CSV or PDF format

![](<../.gitbook/assets/image (5).png>)
