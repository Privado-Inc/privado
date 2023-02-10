# Policies

Policies in Privado are used to govern the usage of personal data in the application and its integration points. Incompatible data usages are a major reason for various privacy fines. To protect data, data governance teams want to define and enforce controls at the time of collection and processing. This can be achieved using Policies.

### Examples

* Usage of Race, Ethnicity and Nationality in your Machine Learning and Artificial Intelligence model may create a bias in the results, towards certain individuals. You can have a policy that prohibits the use of Race, Ethnicity and Nationality in Artificial Intelligence and Machine Learning models.

```yaml
policies: 

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
       laws: GDPR, CCPA2
```

* Business intelligence dashboards are popular among business analysts to run queries and get insights to help customers and grow the business. Exposing personal data as part of these dashboards can lead to unauthorized access and breaches. You can have a policy to deny access of personal data such as name, email, mobile, address to your business dashboard applications.

```yaml
policies:

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
```

### Organization

Policies are present in `rules/policies` and are organized as follows,

```
|__rules
   |__policies                   
   |  |__disallow_personal_data_in_business_dashboards.yaml
   |  |__ai_governance.yaml
```
