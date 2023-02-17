# Understanding Sinks

Sinks are the destinations where personal data is being sent in the code. They are further categorized into storages, leakages, third parties, and internal apis. These top-level categories are aligned with the usages of the data. The top-level categories are further divided into sub-categories. For example, Storages are divided into MongoDB, MySQL, PSQL, etc. At the lowest level, rules are stored as per the programming languages. There will be a file for every language. The rules which are applicable to all the programming languages are stored in the `default.yaml` file.

### Example

```yaml
sinks:

  - id: Storages.AmazonS3.Read
    name: Amazon S3(Read)
    domains:
      - aws.amazon.com
    patterns: 
      - "(?i).*(?:AmazonS3ClientBuilder|S3Client[.]builder|AmazonS3EncryptionClient|software.amazon.awssdk.services.s3).*(?:get|list|head|select).*"
    tags:

  - id: Storages.AmazonS3.Write
    name: Amazon S3(Write)
    domains:
      - aws.amazon.com
    patterns: 
      - "(?i).*(?:AmazonS3ClientBuilder|S3Client[.]builder|AmazonS3EncryptionClient|software.amazon.awssdk.services.s3).*(?:abortMultipartUpload|completeMultipartUpload|copy|create|delete|put|uploadPart).*"
    tags:
```

### Organization

Sinks are present in [`rules/sinks`](https://github.com/Privado-Inc/privado/tree/main/rules/sinks) directory and are organized as follows,

```
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
```
