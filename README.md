<p align="center">
  <img src="https://uploads-ssl.webflow.com/5ec4dc8db66aa018f257c00f/611c2d51857811724661bcda_Full%20Logo%2032px.svg" />
</p>

# Privado CLI
[Privado](https://www.privado.ai/) CLI scans &amp; monitors your repositories to build privacy, transparency reports &amp; finds privacy issues.

[![Slack](https://img.shields.io/badge/Slack%20Workspace-Join%20now!-36C5F0?logo=slack)](https://join.slack.com/t/devprivops/shared_invite/zt-yk5zcxh3-gj8sS9w6SvL5lNYZLMbIpw)
[![GitHub stars](https://img.shields.io/github/stars/Privado-Inc/privado?style=social)](https://github.com/Privado-Inc/privado)


## About Privado
We are building Privado with a fundamental belief that privacy is the defining problem of this decade and developers want to build Privacy first products. However, current solutions take a top-down checklist kind of compliance approach. Our vision is to enable developers to embed privacy in their products. Our first product is a code scanner built grounds up for privacy which will allow developers to get visibility into privacy issues in their products and will generate transparency reports for the users of these products.  

## Introducing Code Scanning for Privacy
Code is where the business logic of collecting, sharing and processing of personal data lives. We are scanning the source code to discover personal data, data flows, 3rd party integrations automatically. These scans enable developers to build transparency reports with very little effort and also surfaces privacy issues in the code.

This project is under active development, we encourage you to join our [slack channel](https://join.slack.com/t/devprivops/shared_invite/zt-yk5zcxh3-gj8sS9w6SvL5lNYZLMbIpw) and collaborate with us. 


### Privacy Reports:
1. Transparency Report like Privacy Policy
2. Privacy compliance reports like GDPR Article 30 or RoPA report
3. Apple Privacy Nutrition Label report
4. Google Safety Section report

### Privacy Issues
We are building a OWASP like list of common privacy issues and will open-source them soon. Some examples of privacy issues are:
1. Dark Patterns: A pre-checked consent box
2. Personal data going to logs

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
