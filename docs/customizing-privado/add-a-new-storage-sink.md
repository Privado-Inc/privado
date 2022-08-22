# Add a new Storage Sink

Storage sinks are the databases, file systems or cloud services where data is stored. Example: MySQL, MongoDB, AWS S3, etc.

**List of fields for the definition of storage sink**:


|  Field | Description |
| ------------ | ------------ |
|  id |It is unique identifier for the storage sink. It has format “Storages“ + Vendor Name ( without spaces and special characters ) + “Read“ or “Write“ <br> </brExample:>Example: For MongoDB system, you can have “Storages.MongoDB.Read” or “Storages.MongoDB.Write”  |
| name  | It is name of the storage Example: MongoDB ( Read )  |
| patterns  | It is an array of regex patterns for the storage sink. This regex will be used to search method names and to further check if data elements are going to the identified methods. Matching methods with data flows will be tagged for this storage sink. <br></br>Example: Mark specific method from a known class</br>class name: `com.privado.MySinkClass`</br>method name: `mySinkMethod()`</br>pattern: `com.privado.MySinkClass.mySinkMethod`</br></br>Example: Mark all methods from a known class<br>class name: `com.privado.MySinkClass`</br>method one: `mySinkMethod1()`</br>method two: `mySinkMethod2()`</br>pattern: `com.privado.MySinkClass.*`<br></br>Example: Mark a specific method across the classes</br>class name: `com.privado.MySinkClass`</br>method one: `mySinkMethod()`</br>class name: `com.privado.MySinkClass2`<br>method one: `mySinkMethod()`<br>pattern: `.*mySinkMethod`  |
|  tags | It's an object of key-value pairs. This is useful to group and filter data leakages.</br>Example: you can tag applicable laws for the data leakages. <br></br>```tags:```<br>`      laws: GDPR, HIPAA `|



High level key is `sinks` which is an array of storages. Once the storage `sink` object is defined, we can add it to the array of sinks.

For a new vendor, you can create sub-directory with the vendor name under directory `rules/sinks/storages/`. You can create a language specific file - `java.yaml` and add the storage sink definition to it.

Once the new storage sink is added, Privado will detect and track data flows to this storage sink.
