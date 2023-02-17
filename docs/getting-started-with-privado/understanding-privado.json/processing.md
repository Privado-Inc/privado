# Processing

Processing results summarises all the occurrences of detected data elements that are found in the scan.

A processing result consists of the following structure:

```json
{
    "processing": {
        "sourceId": "string",
        "occurrences": [
            {
                "sample": "string",
                "lineNumber": "int",
                "columnNumber": "int",
                "fileName": "string",
                "excerpt": "string"
            }
        ]
    }
}
```

The parameters are explained below:

| **Field**      | **Description**                                           |
| -------------- | --------------------------------------------------------- |
| `sourceId`     | ID of the source                                          |
| `occurrences`  | List of occurrences of the data element in the code       |
| `sample`       | name of the entity in which the data element is processed |
| `lineNumber`   | Line number of the occurance                              |
| `columnNumber` | Column number of the occurance                            |
| `fileName`     | Name of the file where the occurrence is detected         |
| `excerpt`      | A dump of the code around the occurrence                  |
