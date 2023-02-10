# Collections

Collections indicate all the API data collections that are present in the codebase.

A collection result consists of the following structure:

```json
{
    "collections": [
        {
            "collectionId": "string",
            "name": "string",
            "isSensitive": "boolean",
            "collections": [
                {
                    "sourceId": "string",
                    "occurences": [
                        {
                            "endPoint": "string",
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
}
```

The parameters are explained below:

| **Field**      | **Description**                                           |
| -------------- | --------------------------------------------------------- |
| `collectionId` | ID of the collection which is detected                    |
| `name`         | Name of the collection                                    |
| `isSensitive`  | Boolean indicating if the field is sensitive              |
| `collections`  | List of collections of the data element found in the code |
| `sourceId`     | ID of the source which is collected                       |
| `occurences`   | List of occurrences of the data element in the code       |
| `endPoint`     | Endpoint of the API                                       |
| `sample`       | name of the entity which is processed in API              |
| `lineNumber`   | Line number of the occurance                              |
| `columnNumber` | Column number of the occurance                            |
| `fileName`     | Name of the file where the occurrence is detected         |
| `excerpt`      | A dump of the code around the occurrence                  |
