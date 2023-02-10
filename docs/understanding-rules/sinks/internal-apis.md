# Internal APIs

A micro-service architecture based application is made of hundreds of internal APIs that talk to each other and exchange data including personal data. Tracking data flow within these internal APIs is important for data safety and ensuring data usage is compatible with your Data Processing Agreements (DPA) and Privacy policies. Privado tracks data flows to internal APIs giving you visibility to enforce data safety controls.

Directory `rules/internal_apis/api/`:

    |__rules
       |__sinks
       |  |__internal_apis
       |  |  |__api
       |  |     |__java.yaml
       |  |     |__default.yaml
