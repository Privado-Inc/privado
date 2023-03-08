# Add a new Source

Source is a data element definition. To add a new data element, you first need to decide category. If you go to directory [`rules/sources`](https://github.com/Privado-Inc/privado/tree/main/rules/sources), you can see 22 files, each representing a category. The full list of Privado source categories can be found HERE.

You can either existing category or use your own custom category. If you use existing category, you can open the corresponding file and add entry for the data element.

**List of fields for defining a data element:**

| Field         | Description                                                                                                                                                                                                        |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `id`          | It is unique identifier for the data element. It has format “Data.Sensitive." + category name ( without spaces and special characters ) + data element name ( without spaces and special characters )              |
| `name`        | It is name of the data element                                                                                                                                                                                     |
| `category`    | It is category of the data element                                                                                                                                                                                 |
| `isSensitive` | It is a boolean flag to indicate if the data element is sensitive                                                                                                                                                  |
| `sensitivity` | It is a flag to indicate sensitivity level of the data element. It can have values “low”, “medium”, “high”                                                                                                         |
| `patterns`    | It is an array of regex patterns for the data element. This regex will be used to search variable names. Matching variables will be tagged as the source for this data element                                     |
| `tags`        | <p>It’s an object of key-value pairs. This is useful to group and filter data elements. Example: you can tag applicable laws for the data element.<br><br><code>tags:</code><br><code>laws: GDPR, HIPAA</code></p> |

High level key is `data_dictionary` which is an array of data elements. Once the data element object is defined, we can add it to the array of `data_dictionary`.

For a custom category, you can create a new file with the category name ( all small case, replace space with `_` ) in directory `rules/sources` and add the details to it.

Once the new data element is added, Privado will detect and track data flows for this new data element.
