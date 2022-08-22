# Understanding Rules



Privado has rules to answer the contextual questions related to the personal data. The journey of tracking data starts from “sources”. “sources” are where data dictionary is defined. Privado identifies the variables, classes and structures matching sources and then further tracks the flows to third parties, databases and leakages. Threats are code or configuration implementation which has direct impact on the data security and privacy. Policies allow you to enforce compliance and data governance rules.

Rules directory structure:

```
|__rules
   |__sources
   |  |__contact_data.yaml
   |  |__account_data.yaml
   |  |__personal_identification.yaml
   |  |__ ...
   |__sinks
   |  |__storages
   |  |  |__mongodb
   |  |     |__java.yaml
   |  |     |__python.yaml
   |  |     |__cpp.yaml
   |  |     |__default.yaml
   |  |  |__mysql
   |  |     |__java.yaml
   |  |     |__python.yaml
   |  |     |__cpp.yaml
   |  |  |__ ...
   |  |__leakages
   |  |  |__logs
   |  |     |__java.yaml
   |  |     |__python.yaml
   |  |     |__cpp.yaml
   |  |__third_parties
   |  |  |__api
   |  |        |_java.yaml
   |  |        |__python.yaml
   |  |        |__cpp.yaml
   |  |        |__default.yaml
   |  |  |__sdk
   |  |     |__slack
   |  |        |__java.yaml
   |  |        |__python.yaml
   |  |        |__cpp.yaml 
   |  |     |__jira
   |  |        |__java.yaml
   |  |        |__python.yaml
   |  |        |__cpp.yaml
   |  |        |__default.yaml
   |__collections
   |  |__annotations
   |  |  |__java.yaml
   |  |  |__python.yaml
   |  |  |__default.yaml
   |__threats
   |  |__collection.yaml
   |  |__configuration.yaml
   |  |__leakage.yaml
   |  |__sharing.yaml
   |  |__storage.yaml
   |__policies
   |  |__restrict_data_elements.yaml
   |  |__allow_data_elements.yaml
   |  |__ai_governance.yaml
```
