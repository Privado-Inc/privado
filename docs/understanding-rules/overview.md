# Rules Overview
Privado has rules to answer the contextual questions related to the personal data. The journey of tracking data starts from “sources”. “sources” are where data dictionary is defined. Privado identifies the variables, classes and structures matching sources and then further tracks the flows to third parties, databases and leakages. Threats are code or configuration implementation which has direct impact on the data security and privacy. Policies allow you to enforce compliance and data governance rules.

Rules directory structure:

    |__rules
       |__sources
       |  |__contact_data.rule
       |  |__account_data.rule
       |  |__personal_identification.rule
       |  |__ ...
       |__sinks
       |  |__storages
       |  |  |__mongodb
       |  |     |__java.rule
       |  |     |__python.rule
       |  |     |__cpp.rule
       |  |     |__default.rule
       |  |  |__mysql
       |  |     |__java.rule
       |  |     |__python.rule
       |  |     |__cpp.rule
       |  |  |__ ...
       |  |__leakages
       |  |  |__logs
       |  |     |__java.rule
       |  |     |__python.rule
       |  |     |__cpp.rule  
       |  |__third_parties
       |  |  |__api
       |  |        |_java.rule
       |  |        |__python.rule
       |  |        |__cpp.rule
       |  |        |__default.rule
       |  |  |__sdk
       |  |     |__slack
       |  |        |__java.rule
       |  |        |__python.rule
       |  |        |__cpp.rule  
       |  |     |__jira
       |  |        |__java.rule
       |  |        |__python.rule
       |  |        |__cpp.rule
       |  |        |__default.rule
       |__collections
       |  |__annotations
       |  |  |__java.rule
       |  |  |__python.rule
       |  |  |__default.rule                    
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
