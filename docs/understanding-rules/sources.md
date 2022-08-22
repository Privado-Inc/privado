# Sources

The journey of tracking the personal data starts from “sources”. They are variables, classes and structures in the code which represents the personal data. Sources are divided into 22 categories. There is a separate file for each category. The file contains the rules for data elements belonging to that category.

Directory `rules/sources`:

    |__rules
       |__sources
       |  |__account_data.yaml
       |  |__audio_visual_sensory_data.yaml
       |  |__background_check_data.yaml
       |  |__biometric_data.yaml
       |  |__contact_data.yaml
       |  |__education_background_data.yaml
       |  |__financial_data.yaml
       |  |__health_data.yaml
       |  |__location_data.yaml
       |  |__national_identification_numbers.yaml
       |  |__online_identifiers.yaml
       |  |__personal_characteristics.yaml
       |  |__personal_identification.yaml
       |  |__professional_employment_background_data.yaml
       |  |__purchase_data.yaml
       |  |__social_media_data.yaml
       |  |__spouse_family_dependend_data.yaml
       |  |__technical_data.yaml
       |  |__usage_data.yaml
       |  |__user_content_data.yaml
       |  |__vehicle_data.yaml
       |  |__workplace_monitoring_data.yaml


When the code is scanned, Privado first marks the “sources”, then tracks it’s journey to sinks such as third parties, databases, and logs.




