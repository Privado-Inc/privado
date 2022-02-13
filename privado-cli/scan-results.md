# Scan Results

Once you authenticate, change to the project you want to scan and run command, `privado scan`.

CLI image

We will open up the results in the browser.

## Overview <a href="#overview" id="overview"></a>

![Privado Scan Summary](<../.gitbook/assets/Initial Summary (1) (1).png>)

You can see the high-level summary of Privado’s scan. Start creating a Data Safety report or navigate to Data Elements or Third Party by clicking on View All.

## Data Elements <a href="#data-elements" id="data-elements"></a>

Privado discovers all personal data being processed by your products or applications and categorizes them into data categories. We also assign a sensitivity to each data element that goes in the overall risk calculation of the repository.

At the top, we give an overview of the total data elements discovered along with a count of data categories and sensitive data elements. The data element table gives you the following attributes to sort, filter.\
\


![](<../.gitbook/assets/01 Initial Screen (1).png>)



1. Data Element: The data element we discovered in the code repository
2. Category: The category to which data element belongs to, this is configurable. This category is used in reporting like Records of Processing Activity(RoPA) report or a customer-facing privacy policy
3. Confidence: Categorizes the discovered data elements in two buckets, High and Low confidence. High confidence data elements are accurate most of the time and Low confidence data elements are wrong most of the time
4. Sensitivity: A configurable value to denote the risk of data elements. This allows you to add context to data elements like location data is personal data under GDPR but in your company's context could be of high risk to your users and you should mark it as High Sensitivity.&#x20;
5. Occurrence: These are locations in your code where we discovered the data element. You can click on it to see the exact location of the data element in your code. We label the identified code as yellow, you can copy and share occurrences internally for further investigation

### Viewing data element in code:

Click on the occurrence of the data element you want to see and it will open the code location where we found the data element:\


![](<../.gitbook/assets/-\_9KEyP6CHSl4tQuW7Z\_lH4ujp1AYYacGA (1).png>)

You can remove data elements by marking them as False Positive and give us feedback by marking them as Confirmed.

## Third-Party

Privado automatically discovers any third-party integration in your product or applications. Navigate to the Third-Party section by clicking on the Third Party tab or by clicking View all from summary page.

![](<../.gitbook/assets/Tw\_4-fVyUEHfYniZZMuPGhK7Os3Sfi5R-g (1).png>)



1. Logo: We display the logo of the third party for easy identification
2. Vendor Name: This is the name of the third party we have discovered
3. Domain: Here we display the URL that was integrated into your code or the SDK that was part of your code because of which this third party was discovered
4. Category: Some examples, Information Technology, Finance, Infrastructure
5. Description: Summary of the parent company of the third party
6. Occurrence: These are locations in your code where we discovered the third parties. You can click on it to see the exact location of the third parties in your code. We label the identified code as yellow, you can copy and share occurrences internally for further investigation

### Viewing Third Party in code:

Click on the occurrence of the third party you want to see and it will open the code location where we found the data element:



You can remove third parties by marking them as False Positive and give us feedback by marking them as Confirmed.

\


\
\


