# Generate Play Store Data Safety Report

## Run Scan <a href="#run-scan" id="run-scan"></a>

To generate the data safety report, browse to the folder containing your Android app code(`cd ~/Android/myapp/`) and run `privado scan .`



## Generating Data Safety Report

Once the scan finishes, we will open the results in your browser.

![](<../.gitbook/assets/Initial Summary (2).png>)

Click on, Create Data Safety Report button and answer a few questions in an easy to use & guided wizard.

### Review Data Elements

Privado scans the source code of your application for over 200 data elements and looks at data collection like user forms, Android permissions, etc. to discover data elements your mobile app is collecting or sharing.

![](<../.gitbook/assets/Base (8).png>)

Click on the View button to see where we found the data element, in case you find a discovery is wrong, mark it false positive and we will remove the data element from the report.

### Review SDKs

Our scan also discovers all third-party SDKs, libraries that your mobile app uses. For popular SDKs we have a database of default data collected, shared and purposes & we will fill these details automatically. For SDKs, where we do not have the information, please select the values in the next steps.

![](<../.gitbook/assets/Base (1) (1).png>)

Click on the View button to see where we found the SDKs, in case you find a discovery is wrong, mark it false positive and we will remove corresponding values from the report.

### Define Data Usage

In this section, we will ask you some questions about the data types & SDKs we have discovered. First up, for your application let us know details around data deletion, encryption.

![](<../.gitbook/assets/Base (2) (1).png>)

In the next step, for data types tell us if you are sharing them with any third party. Please note data type collected is checked by default.

![](<../.gitbook/assets/Base (3) (1).png>)

Finally, for each data type mark them if they are Required meaning app users have to provide that data, or Optional meaning app owners can choose to not share that data with your app.\


![](<../.gitbook/assets/Base (4) (1) (1).png>)

### Describe Purposes

For each Data Type that you are collecting, mark the purposes. Please note one data type could have multiple purposes of collection.

![](<../.gitbook/assets/Base (5) (1).png>)

For data types, you are sharing mark the Purposes of Sharing. Please note only data types that were marked as being shared will be in this list.

![](<../.gitbook/assets/Base (6) (1).png>)

You are done with the data safety report!

### Reviewing Results

Congratulations, you have completed the form. Quickly review what you have filled and make sure the details are filled correctly. This will be added to your app’s data safety section and is important for you to be transparent to build trust with your app users.

![](<../.gitbook/assets/Base (7) (1).png>)

Click on, Download Report as CSV to generate a file that you can import in Play Store.\
