# Third Parties

Modern day applications are assembled by using many third party libraries.  Example: use SendGrid for sending automated emails to customer or Stripe for payments. Personal data may flow into the third parties which may not be accounted for. The third parties are integrated either via api-to-api mechanism or by using specific SDKs. Privado let’s you detect data flows to third parties. 

The rules for third parties are under directory `rules/sinks/third_parties/`:

    |__rules
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
       |  |        |__default.rule
       |  |     |...

Privado currently has coverage for the popular third parties including AWS, GCP, Azure, Stripe, Twilio, etc. The list is available [here](https://github.com/Privado-Inc/privado/tree/main/rules/sinks/third_parties).
