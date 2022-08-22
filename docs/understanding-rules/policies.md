# Policies

Incompatible data usages are major reason for the privacy fines. To protect data, data governance teams want to define and enforce controls at the time of collection and processing. Policies allows user to govern  the usage of personal data.

Example:

Usage of Race, Ethnicity and Nationality in your Machine Learning and Artificial Intelligence model may make them bias towards certain individuals. You can have policy that prohibits the use of Race, Ethnicity and Nationality in Artificial Intelligence and Machine Learning models.


Business intelligence dashboards are popular among business analysts to run queries and get insights to help customers and grow the business. Exposing personal data as part of these dashboards can lead to unauthorised access and breaches. You can have a policy to deny access of personal data such as name, email, mobile, address to your business dashboard applications.


Directory `rules/policies/`:

    |__rules
       |__policies                   
       |  |__disallow_personal_data_in_business_dashboards.yaml
       |  |__ai_governance.yaml
