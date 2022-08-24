# What are Rules?

Privado has rules to answer the contextual questions related to  personal data. The journey of tracking data starts from "sources". Sources are where data dictionary is defined. Privado identifies the variables, classes and structures matching sources and then further tracks the flows to third parties, databases and leakages which are called as "sinks". Threats are code or configuration implementation which has direct impact on the data security and privacy. Policies allow you to enforce compliance and data governance rules. The rules presents a single _common language_ which embeds all the knowledge of a privacy and data researcher about sources, sinks, data policies, threats to drive the code analysis engine

### Rule Structure

All Privado rules are defined in YAML format and generally have the following structure:

<figure><img src="../.gitbook/assets/Rule YAML.jpg" alt=""><figcaption></figcaption></figure>

The structure of rule varies a bit based on the types of rules that are defined. For example, the Source rule contains `isSensitive` as well as `sensitivity` keys so that based on the values set, the source data is tagged appropriately. Simlarly,  policy rules contain `description` that is needed for the issue that gets created when policy is violated. It also contains `dataflow` as well as `repositories` on which the policy will be applied. To learn more about rules, you can review the [rules](https://github.com/Privado-Inc/privado/tree/main/rules) directory on Github

### Organization

Rules are organized in the privado repository under [`privado/rules`](https://github.com/Privado-Inc/privado/tree/main/rules)  directory. The structure provides a logical way of how rules can be arranged. If you come up with some cool rules, this is where you can drop them in. You may also consider [contributing](../extra/contributing.md) them upstream :handshake:

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
