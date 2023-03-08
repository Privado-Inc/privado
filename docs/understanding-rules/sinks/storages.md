# Storages

The database related rules are stored under this. There is a separate directory for each vendor such as MySQL, MongoDB, etc. At the lowest level, rules are stored as per the programming languages. There will be a file for every language. The rules which are applicable to all the programming languages are stored in the default.rule file.

Directory `rules/sinks/storages`:

```
 |__rules
   |__sinks
   |  |__storages
   |  |  |__mongodb
   |  |     |__java.yaml
   |  |     |__default.yaml
   |  |  |__mysql
   |  |     |__java.yaml
   |  |     |__default.yaml
   |  |  |__amazonS3
   |  |     |__java.yaml
   |  |     |__default.yaml
   |  |  |__arangodb
   |  |     |__java.yaml
   |  |     |__default.yaml
   |  |  |__...
   
```

Privado currently supports 23 popular databases and libraries. The full list is available HERE.
