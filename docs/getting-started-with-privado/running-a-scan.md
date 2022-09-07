# Running a Scan

{% hint style="info" %}
**Requirements:** Privado requires Docker and an Internet connection for fetching its scanning engine (provided as Docker images) for the first time. To install docker, you can follow the steps stated in the [official documentation](https://docs.docker.com/engine/install/). Linux users should also follow docker [post installation steps](https://docs.docker.com/engine/install/linux-postinstall/#manage-docker-as-a-non-root-user) in order to run Privado CLI without root (`sudo`) privileges.
{% endhint %}

The simplest way to run a Privado scan is by just doing,

```bash
privado scan <source directory>
```

This begins analyzing the app, fetching/updating the analysis engine and then performing the scan of target source directory locally.

<figure><img src="../.gitbook/assets/2x 60fps latest.gif" alt=""><figcaption></figcaption></figure>

Upon completion, results are generated locally under `<source directory>/.privado/privado.json` file. This file contains all the information such as discovered data elements, data inventory, data flows to 3rd party, logs, and other sensitive sinks etc. Optionally, when configured, the scan can also send its results to Privado dashboard for visualisation. To know more about how to view results on dashboard, click here.

## Advanced Scan Options

You can view all advanced options using `privado scan --help` and can be run using `privado scan <source directory> [OPTIONS]` A few of the options are summarized below

| `--config`                   | Specifies the config (with rules) directory to be passed to privado-core for scanning. These external rules and configurations are merged with the default set that Privado defines                                                                                                                                                | To know more about what Privado Rules are, click here.             |
| ---------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------ |
| `--debug`                    | Shows debug information while running                                                                                                                                                                                                                                                                                              |                                                                    |
| `--disable-deduplication`    | When specified, the engine does not remove duplicate and subset dataflows. This option is useful if you wish to review all flows (including duplicates) manually. When specified, the engine does not remove duplicate and subset dataflows. This option is useful if you wish to review all flows (including duplicates) manually | To know more about how to view and understand results, click here. |
| `--ignore-default-rules`     | If specified, the default rules are ignored and only the specified rules via configuration are considered                                                                                                                                                                                                                          | To know more about what Privado Rules are, click here.             |
| `--overwrite`                | If specified, the warning prompt for existing scan results is disabled and any existing results are overwritten                                                                                                                                                                                                                    |                                                                    |
| `--skip-dependency-download` | Occasionally for Java projects, Privado would require dependencies to be downloaded for a more deeper analysis if local dependencies are not available. If an analysis takes unacceptable amount of time due to network latency, this option can be used to skip dependency downloads at the cost of incomplete results            |                                                                    |

## Troubleshooting

### Scan Time Out

While Privado scans rarely take more than 7 minutes even for very large java repositories, static analysis due to its nature can be a resource intensive process. If your scan is timing out, try the following approaches to manage resources used by Docker and JVM:

1. Limit the RAM consumed by docker. Reference: Runtime options with Memory, CPUs, and GPUs]\(https://docs.docker.com/config/containers/resource\_constraints/)
2. Specific to Java, we can also set the env variables of docker according to host machine configuration, Reference: How To Configure Java Heap Size Inside a Docker Container | Baeldung]\(https://www.baeldung.com/ops/docker-jvm-heap-size)
