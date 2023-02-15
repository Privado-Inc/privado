# Sources

Sources are the starting points of a data element in a data flow. This includes places like user inputs, database queries, or API collections.

A source result consists of the following structure:

```json
{
    "sources": {
        "sourceType": "string",
        "id": "string",
        "name": "string",
        "category": "string",
        "sensitivity": "string",
        "isSensitive": "boolean",
        "tags": [
            {
                "key": "value"
            }
        ]
    }
}
```

The parameters are explained below:

| **Field**     | **Description**                                         |
| ------------- | ------------------------------------------------------- |
| `sourceType`  | The type of source detected                             |
| `id`          | ID of the source                                        |
| `name`        | Name of the source                                      |
| `category`    | category of the source                                  |
| `sensitivity` | Sensitivity of the source                               |
| `isSensitive` | Boolean indicating if the field is sensitive            |
| `tags`        | Key value pairs to attach further context to the source |
