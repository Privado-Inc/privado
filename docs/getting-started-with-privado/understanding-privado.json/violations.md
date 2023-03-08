# Violations

A Violation is a privacy risk identified within the code that could potentially cause a data leak. They are also detected based on certain custom rules declared during the scan. To know more about the rules, [click here](broken-reference).

A violation result consists of the following structure:

```json
"violations": [
    {
        "policyId" : "string",
          "policyDetails" : {
            "name" : "string",
            "policyType" : "string",
            "description" : "string",
            "fix" : "string",
            "tags" : {
                "key": "value"
            }
        }
    }
]
```

The parameters are explained below:

| **Field**                          | **Description**                                             |
| ---------------------------------- | ----------------------------------------------------------- |
| `policyId`                         | ID of the policy which is violated                          |
| `policyDetails`                    | Object containing details about the violation               |
| `name`                             | name of the violation which is detected                     |
| <pre><code>policyType</code></pre> | Type of the policy which can be either threat or compliance |
| `description`                      | A brief description about the violation                     |
| `fix`                              | A detailed description of how to resolve the violation      |
| `tags`                             | Key-value pairs to add more context to the violation        |
