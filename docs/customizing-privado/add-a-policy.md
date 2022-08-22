# Add a new Policy
Policies allows user to govern the usage of personal data.

**List of fields for the definition of policy:**

|  Field | Description |
| ------------ | ------------ |
|  id |It is unique identifier for the policy. It has the format: “Policy“ + “Allow” or “Deny” + “Processing“ or “Sharing“ + unique action.  |
| description  | It is description of the policy  |
| patterns  | There are two possible values. <br><br> 1. “Deny”: It means that if the mentioned sources or data flows are detected for the repositories mentioned, it is a violation. The respective repositories are not supposed to process the data elements or have the data flows. <br><br> 2. “Allow“: It means that if the mentioned sources or data flows are detected for the repositories other than mentioned, it is a violation. The mentioned data elements or data flows are only allowed for the repository mentioned.  |
| dataFlow | It is mandatory object. It has two keys - “sources” and “sinks”. <br><br> If only “sources” are specified, it means the rule is to either allow or deny processing of data elements. If both “sources” and “sinks” are specified, it means the rule is to either allow or deny data flows. |
|sources| It is mandatory key. It is array of regular expression for the data elements ids.|
|sinks| It is optional key. It is array of regular expression for the sink ids.|
|repositories|It is mandatory key. It specifies the repositories for which to evaluate the policy. Use `.*` for specifying all repositories.|
|  tags | It’s an object of key-value pairs. This is useful to group and filter issues resulting from the policies.</br><br>Example: you can tag applicable laws for the policies. <br></br>```tags:```<br>`      laws: GDPR, HIPAA `|

High level key is `policies` which is an array of policy definitions. Once the policy object is defined, we can add it to the array of `policies`.

You can either add policy definition to an existing a file or create a new yaml file. The policies are located at directory: `rules/policies`.

Once the new policy is added, Privado will evaluate the policy on the data elements and data flows after the code scan, and it will create the issues in case of policy violations.


Examples:

a) Usage of Race, Ethnicity and Nationality in your Machine Learning and Artificial Intelligence model may make them bias towards certain individuals. You can have policy that prohibits the use of Race, Ethnicity and Nationality in Artificial Intelligence and Machine Learning models.
Policy: Ethical AI usage policy


    - id: Policy.Deny.Processing.EthicalUsageForAI
        name : "Ethical AI Usage Policy"
        type: Compliance
        description: "Don't use ethnicity, race and nationality for machine learning and AI"
        fix: "Talk to the Privacy Engineering team: privacy-engineering@org.com"
        action: Deny    
        dataFlow:      
          sources:
                 - "Data.Sensitive.PersonalIdentification.Ethnicity"
                 - "Data.Sensitive.PersonalIdentification.Race"
                 - "Data.Sensitive.PersonalIdentification.Nationality"
        repositories: 
                 - curate-offers-machine-learning
                 - track-engagement-ai
        tags:
           laws: GDPR, CCPA



b) Business intelligence dashboards are popular among business analysts to run queries and get insights to help customers and grow the business. Exposing personal data as part of these dashboards can lead to unauthorised access and breaches. You can have a policy to deny access of personal data such as name, email, mobile, address to your business dashboard applications.
Policy: Restrict usage of personal data for Business Intelligence dashboards

    - id: Policy.Deny.Processing.NoPersonalDataInBIReports
        name : "Restrict usage of personal data in BI reports"
        type: Compliance
        description: "Personal data in BI reports may expose it to unathorized access."
        fix: "Talk to the Data Protection team: data-protection@org.com"
        action: Deny    
        dataFlow:      
          sources:
                 - "Data.Sensitive.PersonalIdentification.*"
        repositories: 
                 - business-intelligence-dashoboard
                 - business-intelligence-advanced-reports
        tags:
           laws: GDPR, CCPA
		   

c) Developers use Slack for monitoring the applications. Logs, critical errors and events are sent. Often customer data is sent which results in an unauthorised access. You can have a policy to deny sharing any personal data on Slack.
Policy: Restrict sharing of personal data on Slack

    - id: Policy.Deny.Sharing.DontShareContactDataToSlack
        name : "Restrict sharing customer contact data on Slack"
        type: Compliance
        description: "Customer contact shared on Slack can lead to data breaches and unathorized access"
        fix: "Talk to the Data Protection team: data-protection@org.com"
        action: Deny    
        dataFlow:      
          sources:
            - "Data.Sensitive.ContactData.*"
            - "Data.Sensitive.PersonalIdentification.PhoneNumber"
          
          sinks:
            - "Data.Sharing.Slack"
        repositories: 
                  - "**"
        tags:
           laws: GDPR, CCPA
     

