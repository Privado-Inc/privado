# Privado User Documentation

## What is Privado? <a href="#what-is-privado" id="what-is-privado"></a>

Privado is a static code scanning tool to find, fix and remediate privacy issues in your products & applications. Our scan discovers what personal data(as defined by GDPR, other laws) your app is processing, third-party integrations, data flows. With our scan results, we generate privacy reports for your apps as mandated by laws like GDPR or platforms like Apple and keep them in sync with code changes.

## What does our scan discover? <a href="#what-does-our-scan-discover" id="what-does-our-scan-discover"></a>

1. Data Elements: These are personal data that your app is collecting, sharing, processing. Here is a list of data elements that we are discovering.
2. Third-Parties: Any third-party integrations inside your code, via APIs or SDKs/libraries.
3. APIs: We also discover any internal APIs that your app is connected with.
4. Datastores(not released yet, still beta): Identify the databases where you are sourcing the data from or sending the data.
5. Privacy Vulnerabilities(not released yet, still beta): Code issues that exist which can lead to privacy vulnerabilities

## What can I do with Privado? <a href="#what-can-i-do-with-privado" id="what-can-i-do-with-privado"></a>

### Generate Play Store Data Safety Report <a href="#generate-play-store-data-safety-report" id="generate-play-store-data-safety-report"></a>

This is the first use case that we are live with. Currently, to fill the data safety form Android developers have to ask around in the team to find what data they are collecting, spend hours reading the documentation of SDKs to find data shared, and navigate the complex Playstore form. With our scan, we pre-fill data types collected, shared, and guide you with our wizard to generate the data safety report.

### Privacy Audits <a href="#privacy-audits" id="privacy-audits"></a>

Privacy Engineers can use our CLI tool as an MRI for products, applications and find out privacy risks. With our scans, privacy engineers save the time they have to spend chasing engineers with assessments and can directly start prescribing privacy controls for data minimization, sharing, etc.

### We have the following use cases on our Roadmap: <a href="#we-have-the-following-use-cases-on-our-roadmap" id="we-have-the-following-use-cases-on-our-roadmap"></a>

1. Generating Apple Nutrition Label Report
2. Generating privacy compliance reports like GDPR Article 30 or RoPA report
3. Detecting Privacy Vulnerabilities in current code implementation
4. Privacy Policy Generator


## Installation

To start off, make sure `docker` is installed. To install docker, you can follow the steps stated in the [official documentation](https://docs.docker.com/engine/install/). You can install Privado CLI in multiple manners:  

- [Using `curl` or `wget`](#install-using-curl-or-wget)
- [Using `go`](#install-using-go)
- [Using releases](#install-release-manually)


### Install using `curl`:
The installation script will download and setup the latest stable release for you as per your OS and arch. Run: 

```
curl -o- https://raw.githubusercontent.com/Privado-Inc/privado/main/install.sh | bash
```

To uninstall, simply delete `~/.privado/bin`.

### Install using Go
If you are a GoLang fan, you can use the `go install` to install the Privado CLI:

```
go install github.com/Privado-Inc/privado@latest
```

This will place the `privado` binary in your `GOPATH`'s bin directory. This directory must be added to the `$PATH` environment variable. You can learn more [here](https://www.digitalocean.com/community/tutorial_series/how-to-install-and-set-up-a-local-programming-environment-for-go). 

### Install Release Manually
We use [GitHub Releases](https://github.com/Privado-Inc/privado/releases) to ship versioned `privado` releases for supported platforms. You can download a executable of Privado CLI for your platform.

For example, to setup `privado` for MAC-OSX (arm64), you will download `privado-darwin-amd64.tar.gz`, and run:
```
tar -xf ~/.privado/privado-darwin-amd64.tar.gz
chmod +x privado
mv privado /usr/bin/privado
```

## Authenticating
Privado CLI requires a license key to run scans. To generate a license, run the following command:
```
privado auth <your_email@example.com>
```
A copy of the license will be emailed to you. 

To authenticate and bootstrap the app using the generated license, run:
```
privado bootstrap </path/to/privado-license.json>
```

and done! You are all set to scan your projects and generate compliance reports.

> Please note that generated licenses are valid for 1 year from the date of issue.    
> For more information about licensing, feel free to get in touch with us on [Slack](https://join.slack.com/t/devprivops/shared_invite/zt-yk5zcxh3-gj8sS9w6SvL5lNYZLMbIpw) or [Email](mailto:support@privado.ai).


## Running a Scan
Privado CLI works on the client-end and does not share any files, code-snippets, results, or reports during the complete lifecycle.


To scan a repository, simply run:
```
privado scan <path/to/repository>
```
Depending on repository size and system configuration, scanning can take upto 5-10 minutes. Post completion, the results can be viewed on [localhost:3000](http://localhost:3000).   

<img width="456" alt="privado-scan-completion-snapshot" src="https://user-images.githubusercontent.com/35270649/153745682-3a97e6eb-0696-4536-b8ca-562d31ead1f9.png">

To use a different port, simply use the `-p` (or  `--port`) flag:
```
privado scan <path/to/repository> -p 5001
```

Results and reports (if generated), are saved to `repository/.privado`. We encourage keeping `.privado` folder as a part of your repository to facilitate report collaboration and share privacy discovery.


## Loading Results
At any point in time, you can directly load the existing results without running the entire scan and continue to generate or modify reports:
```
privado load <path/to/repository>
```
This is also helpful for huge codebases and projects with multiple collaborators.


## Command Reference 
The section contains detailed reference to `privado` commands.

### Privado CLI Global Flags
| Flag             	| Description                                                                                    	|
|--------------------------	|---------------------------------------------------------------------------------------------------------------------------	|
| `-h, --help` 	| Help about any command, or sub-command 	|
| `-l, --license <string>` 	| The license file to be used. Overrides the default bootstrapped license (default "`/Users/ojaswa/.privado/license.json`") 	|


### Privado CLI Commands
| Command      	| Description                                                                                                                                                                           	| Usage                          	| Supported Flags                                                                                                                                                                                                                                                                                              	|
|--------------	|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|--------------------------------	|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|
| `auth`       	| Generate license for Privado                                                                                                                                                          	| `privado auth `                	| -                                                                                                                                                                                                                                                                                                            	|
| `bootstrap`  	| Authenticates Privado using the requested license and generates required configurations                                                                                               	| `privado bootstrap  [flags]`   	| `--overwrite`: <br>Overwrites the existing license fil (if any)                                                                                                                                                                                                                                              	|
| `completion` 	| Generate the autocompletion script for privado for the specified shell. See each sub-command's help for details on how to use the generated script.                                   	| `privado completion [command]` 	| -                                                                                                                                                                                                                                                                                                            	|
| `help`       	| Help about any command                                                                                                                                                                	| `privado help [command]`       	| -                                                                                                                                                                                                                                                                                                            	|
| `scan`       	| Scan a codebase or repository to identify privacy issues and generate compliance reports                                                                                              	| `privado scan  [flags]`        	| `-o, --overwrite`: <br>If specified, the warning prompt for existing scan results is disabled and any existing results are overwritten<br><br> `-p, --port `: <br>The port t be used to render HTML results (default 3000) <br><br>`--debug`: <br>To enable underlying process output for debugging purposes 	|
| `load`       	| Load a scanned codebase or repository and continue generating compliance reports. It skips privacy scan and loads the results present in the target repository (`.privado` directory) 	| `privado load  [flags]`        	| -p, --port    : The port t be used to render HTML results (default 3000)   --debug : To enable underlying process output for debugging purposes                                                                                                                                                              	|


## What does it cost? <a href="#what-does-it-cost" id="what-does-it-cost"></a>

Privado is free for:

1. Open Source Projects
2. For individual developers and small teams.

## How Privado CLI handles your data? <a href="#how-privado-cli-handles-your-data" id="how-privado-cli-handles-your-data"></a>

Privado CLI tool was engineered with security in mind. Our tool runs the scan locally on your machine and your code never leaves your system.
