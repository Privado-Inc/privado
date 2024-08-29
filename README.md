# What is Privado?
[![slack](https://img.shields.io/badge/slack-privado-5A34D9.svg?logo=slack)](https://join.slack.com/t/privado-community/shared_invite/zt-yk5zcxh3-gj8sS9w6SvL5lNYZLMbIpw)
[![docs](https://img.shields.io/badge/docs-gitbook-brightgreen.svg?logo=gitbook)](https://docs.privado.ai)

Privado is an open-source static code analysis tool to discover data flows in the code. It detects more than 110 [personal data elements](docs/extra/data%20element%20list.csv) being processed and further maps the data flow from the point of collection to "sinks" such as external third parties, databases, logs, and internal APIs.


<img src="https://user-images.githubusercontent.com/80044360/186333819-779bfff5-d7a2-4bba-88e9-0ca866e1ee81.gif" width="600px">

# Supported languages
We support Java and Python in GA. Our Enterprise offering covers all programming languages, and we're working on adding support for more languages to OSS. Support for JS/TS is coming soon!


# Quick Start

First, make sure you have [Docker](https://docs.docker.com/get-docker/) installed on your system, then follow these simple steps to get started with Privado.

### Download the Privado CLI

```
curl -o- https://raw.githubusercontent.com/Privado-Inc/privado-cli/main/install.sh | bash
```

### Clone our test repo
We recommend using [this sample app](https://github.com/saurabh-sudo/BankingSystem-Backend) to get started with Privado.
```
git clone git@github.com:saurabh-sudo/BankingSystem-Backend.git
```

### Scan your repository

```
privado scan <source directory>
```


### Get results

The results are generated at `<source directory>/.privado/privado.json` and a preview will be shown in your terminal.

### Visualize results

To visualize the results and generate reports, you can create a free account at the end of a successful scan. Once a scan is complete, it will ask your permission to synchronize the generated results with Privado Cloud Dashboard. Note that **no code is sent to the cloud**–only the JSON output generated by the scan. Upon successful sync, you can view the results on our free platform.

<img src="https://user-images.githubusercontent.com/80044360/186335775-82139291-4edc-4750-85bf-18b26d5655d3.png" width="600px">

# Who is it for?
1. Privacy Engineers
2. Data Protection Engineers
3. Data Governance Engineers
4. Security Engineers
5. Mobile App Developers
6. Developers 
    
# How does it help?
Privado lets Engineers ask contextual questions about the usage of sensitive data at scale.  
Examples:

<img src="https://user-images.githubusercontent.com/80044360/186449832-5799854c-a8e0-43c6-bf43-4b47ff5d23a5.jpeg" width="600px">

# Use cases
1. Generate and maintain Data maps and Record of Processing Activity Reports ( RoPA / Article-30 Reports )
2. Automate the generation of the data-flow diagrams
3. Identify and remove data leaks
4. Improve data storage security by identifying and fixing insecure practices
5. Finding and fixing unaccounted third-party sharing of data
6. Establish and enforce Data Protection and Governance policies
7. Generate Android Data Safety Report
8. Incorporate various GDPR, CCPA, SOC, ISO, HIPAA, and PCI controls
9. Do continuous monitoring for privacy and data issues
10. Implement Privacy by Design

# How does Privado work?
Privado can be run locally on your computer or in your CI/CD pipeline. Privado creates a knowledge graph during the scanning process that contextually answers thousands of questions about sensitive data. Since the scan is local, you never have to worry about your code leaving your machine. An output file is stored in JSON format, and the results can be viewed on Privado Cloud.

# What does the scan discover?
Privado will discover the following information in the code during scanning and present it in a dashboard for your review.

-   Data Elements
-   Data Flow Diagrams
-   Data Inventory
-   Code Analysis
-   Issues
    
# What can I do with Privado?
Apart from getting a comprehensive outlook of your data practices for Privacy Audits, you can also use the tool to generate various privacy reports to comply with privacy laws like GDPR and CCPA.

## Record of Processing Activity ( ROPA ) Report

Our free cloud platform can be used to generate RoPA reports for one or more synced repositories.

## Data Safety Report 
A Data Safety Report is a privacy form needed to publish any Android app on the Play Store. Most of the time, filling out a report means developers asking around the team to find what data they're collecting, spending hours reading SDK docs to see where information gets shared and navigating the complex Playstore form. With our scan, we pre-fill data types that are collected and shared, and our wizard guides you through generating the report.

# Contribute
Please check out our [contribution page](https://docs.privado.ai/extra/contributing) if you love this project and would like to contribute.
