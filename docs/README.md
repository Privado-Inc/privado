# What is Privado?
Privado is an open source static code analysis tool to discover data flows in the code. It detects the personal data being processed, and further maps the journey of the data from the point of collection to going to interesting sinks such as third parties, databases, logs, and internal APIs.

# Who is it for?
1.  Privacy Engineers
2.  Data Protection Engineers
3.  Data Governance Engineers
4.  Security Engineers
5.  Mobile App Developers
6.  Developers
    
# How does it help?
Privado let’s Engineer ask contextual questions on usage of sensitive data at scale.  
Example:

> **"Hello, User Accounts API Repository. Has any sensitive data been > leaked to a log file?"**

> **“Hello, HealthCare App Repository. Can you tell me all the variables
> that process patient's medicine data? And are they shared with Google
> Analytics?”**

> **“Hey, Payments Micro-service Repository, is credit card data tokenized
> and encrypted?"**

> **“Hey, Food Delivery Mobile App Repository, can you tell me if the User
> Location data is shared with this weird S3 bucket?"**

# Use cases
1.  Generate and maintain Data map and Record of Processing Activity Reports ( Article-30 Reports )
2.  Automate the generation of the data-flow diagrams
3.  Identify and remove data leaks
4.  Improve data storage security by identifying and fixing insecure practices
5.  Finding and fixing unaccounted third-party sharing of data
6.  Establish and enforce Data Protection and Governance policies
7.  Generate Android Data Safety Report
8.  Incorporate various GDPR, CCPA, SOC, ISO, HIPAA, PCI controls
9.  Do continuous monitoring for privacy and data issues
10.  Implement Privacy by Design
    

# How does Privado work?
Privado can be run locally on your computer or in your CI/CD pipeline. During the scanning process, Privado creates a knowledge graph that answers thousands of questions about sensitive data contextually. You never have to worry about your code leaving your machine since the scan is local. An output file is stored in JSON format. The results can be viewed on Privado Cloud.

# What does the scan discover?
Upon scanning a repository, Privado will discover the following information in the code and presents it in a nice dashboard for your review.

-   Data Elements
-   Data Flow Diagrams
-   Data Inventory
-   Code Analysis
-   Issues
    
[ADD THE IMAGE HERE]

# What can I do with Privado?
Apart from getting a comprehensive outlook of your data practices for Privacy Audits, you can also use the tool to generate various privacy reports to comply with privacy laws like GDPR and CCPA.

## Record of Processing Activity ( ROPA ) Report
Our free cloud platform can be used to generate RoPA reports for a single, as well as a combination of repositories added to the platform. Check out how to create a RoPA report for your repository.

## Data Safety Report
Data Safety Report is a privacy form that is needed to publish any Android app on the Play Store. Currently, to fill the data safety form developers have to ask around in the team to find what data they are collecting, spend hours reading the documentation of SDKs to find data shared, and navigate the complex Playstore form. With our scan, we pre-fill data types collected, shared, and guide you with our wizard to generate the data safety report.

# Supported languages
Java is the first language we support. As part of the Enterprise offering, Privado supports all languages. To open source a language, the architecture change is required so that community contributions can be made easily. We are working on open sourcing support for the other languages.

# Contribute
If you love this project and would like to contribute, please check out our contribution page.

