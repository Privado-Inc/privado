# Add a new Leakage Sink

Leakage sinks are the places where personal data should not go in the plain text format. Privado currently support log sinks. There may be a use case for adding support for file types, sockets, IO streams, caches, etc.

**List of fields for the definition of leakage:**

|  Field | Description |
| ------------ | ------------ |
|  id |It is unique identifier for the data leakage. It has format “Leakages“  + leakage type name ( without spaces and special characters )  <br> Example: For file system, you will have “Leakages.FileSystem”  |
| name  | It is name of the data leakage. Example: File System  |
| patterns  | It is an array of regex patterns for the data leakage. This regex will be used to search method names and to further check if data elements are going to the identified methods. Matching methods with data flows will be tagged for this data leakage. <br></br>Example: Mark specific method from a known class</br>class name: `com.privado.MySinkClass`</br>method name: `mySinkMethod()`</br>pattern: `com.privado.MySinkClass.mySinkMethod`</br></br>Example: Mark all methods from a known class<br>class name: `com.privado.MySinkClass`</br>method one: `mySinkMethod1()`</br>method two: `mySinkMethod2()`</br>pattern: `com.privado.MySinkClass.*`<br></br>Example: Mark a specific method across the classes</br>class name: `com.privado.MySinkClass`</br>method one: `mySinkMethod()`</br>class name: `com.privado.MySinkClass2`<br>method one: `mySinkMethod()`<br>pattern: `.*mySinkMethod`  |
|  tags | It’s an object of key-value pairs. This is useful to group and filter data leakages. </br>Example: you can tag applicable laws for the data leakages. <br></br>```tags:```<br>`      laws: GDPR, HIPAA `|

High level key is `sinks` which is an array of leakages. Once the leakage object is defined, we can add it to the array of `sinks`.

For a new type of leakage, you can create a new file with the leakage name ( all small case, replace space with `_` ) in directory `rules/sinks/leakages` and add the details to it.

Once the new leakage is added, Privado will detect and track data flows to this leakage sinks.

