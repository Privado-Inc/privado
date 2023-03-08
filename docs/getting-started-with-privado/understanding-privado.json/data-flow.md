# Data Flow

Data flows are the journey of a data element within the codebase. It maps out the journey of a data element from a source to a sink detected via static code analysis.

A data flow result consists of the following structure:

```json
{
    "dataFlow": {
        "third_parties": <DataFlow>,
        "leakages": <DataFlow>,
        "storages": <DataFlow>,
        "internal_apis": <DataFlow>,
        "miscellaneous": <DataFlow>
    }
}
```

The structure of all data flow representations is similar. As an example, the following is the structure of a storage sink:

```json
{
    "storages": [
        "sourceId": "string",
        "sinks": [
            {
                "sinkType": "string",
                "id": "string",
                "name": "string",
                "isSensitive": "boolean",
                "paths": [
                    {
                        "pathId": "string",
                        "path": [
                            {
                                "sample": "string",
                                "lineNumber": "int",
                                "columnNumber": "int",
                                "fileName": "string",
                                "excerpt": "string"
                            }
                        ]
                    }
                ]
            }
        ]
    ]
}
```

The parameters of the results are explained below:

| **Field**       | **Description**                                |
| --------------- | ---------------------------------------------- |
| `third_parties` | Third parties that are consuming data elements |
| `leakages`      | Leakages that are consuming data elements      |
| `storages`      | Databases that are consuming data elements     |
| `internal_apis` | Internal APIs that are consuming data elements |
| `miscellaneous` | Miscellaneous data flows                       |

The parameters of a data flow results are explained below:

| **Field**      | **Description**                                                |
| -------------- | -------------------------------------------------------------- |
| `sourceId`     | ID of the source which is processed                            |
| `sinks`        | A list of sinks that are detected in a particular type of sink |
| `sinkType`     | The type of sink                                               |
| `sinks.Id`     | ID of the sink                                                 |
| `name`         | Name of the sink                                               |
| `isSensitive`  | Boolean value indicating if the data element is sensitive      |
| `paths`        | A list of paths defining the data flow of the element          |
| `pathId`       | Unique ID of the path                                          |
| `path`         | An occurrence of a data element                                |
| `sample`       | name of the entity in which the data element is processed      |
| `lineNumber`   | Line number of the occurance                                   |
| `columnNumber` | Column number of the occurance                                 |
| `fileName`     | Name of the file where the occurrence is detected              |
| `excerpt`      | A dump of the code around the occurrence                       |
