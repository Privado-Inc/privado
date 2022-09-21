# Sources

The journey of tracking personal data starts from “sources”. They are variables, classes and structures in the code which represent personal data. Sources are divided into 22 categories. There is a separate file for each category. You can see them in the [rules/sources](https://github.com/Privado-Inc/privado/tree/main/rules/sources) folder. Each file in the folder contains the rules for data elements belonging to that category.

### Example

From the `account_data.yaml` [file](https://github.com/Privado-Inc/privado/blob/main/rules/sources/account_data.yaml)

```yaml
sources:
  
  - id: Data.Sensitive.FinancialData.BankAccountDetails
    name: Bank Account Details
    category: Financial Data
    isSensitive: False
    sensitivity: high
    patterns:
      - "(?i).*((?<!question)bank[^\\s/(;)#|,=!>]*(?:name|account|details|detail|address|country|(swift|bic)-code|(swift|bic)_code)|bank[^\\s/(;)#|,=!>]*account[^\\s/(;)#|,=!>]*details|(swift|bic)[-_]code|(swift|bic)code)"
    tags:
      law: GDPR
```

### Organization

Sources are present in [`rules/sources`](https://github.com/Privado-Inc/privado/tree/main/rules/sources) directory and are organized as follows,

```
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
```

### Execution

When the code is scanned, Privado first marks the “sources”, then tracks it’s journey to sinks such as third parties, databases, and logs.
