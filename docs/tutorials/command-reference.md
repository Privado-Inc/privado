# How to scan a repository using Privado

## Prerequisites <a href="#prerequisites" id="prerequisites"></a>

This tutorial assumes that you have the following setup ready:

* Git, to clone a repository. To install Git, [click here](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
* Docker (Make sure that the Docker engine is running). To install Docker, [click here](https://docs.docker.com/engine/install/)
* Privado OSS project. To install Privado, follow [these steps](https://docs.privado.ai/getting-started-with-privado/getting-started-with-privado)
* A code repository to scan. For this tutorial, we will use [BankingSystem-Backend](https://github.com/saurabh-sudo/BankingSystem-Backend)

### Clone the repository <a href="#clone-the-repository" id="clone-the-repository"></a>

To clone the repository, use the following command

```
git clone https://github.com/saurabh-sudo/BankingSystem-Backend
```

It should be something similar to the following result:

```
Cloning into 'BankingSystem-Backend'...
remote: Enumerating objects: 198, done.
remote: Counting objects: 100% (198/198), done.
remote: Compressing objects: 100% (115/115), done.
remote: Total 198 (delta 42), reused 186 (delta 37), pack-reused 0
Receiving objects: 100% (198/198), 99.97 KiB | 10.00 MiB/s, done.
Resolving deltas: 100% (42/42), done.1Cloning into 'BankingSystem-Backend'... 2remote: Enumerating objects: 198, done. 3remote: Counting objects: 100% (198/198), done. 4remote: Compressing objects: 100% (115/115), done. 5remote: Total 198 (delta 42), reused 186 (delta 37), pack-reused 0 6Receiving objects: 100% (198/198), 99.97 KiB | 10.00 MiB/s, done. 7Resolving deltas: 100% (42/42), done.
```

#### Not working? <a href="#not-working" id="not-working"></a>

If you do not get the above result, check out the [troubleshooting section](command-reference.md#troubleshooting) for help

### Running the scan <a href="#running-the-scan" id="running-the-scan"></a>

To start the scan, run the following command

```
privado scan BankingSystem-Backend
```

This will start the Privado scan and perform the static code analysis of the project and its dependencies. You will get the following result:

```
> Scanning directory: /Users/anujagrawal/Desktop/Projects/Privado/repos/BankingSystem-Backend

> Pulling the latest image: public.ecr.aws/privado/privado:latest
latest: Pulling from privado/privado
Digest: sha256:47f9bd5a32ff4dbea131d39ed355ada0e9190416ffb61b70a2ecd686fa6278ba
Status: Image is up to date for public.ecr.aws/privado/privado:latest

> Starting container with the latest image
> Container ID: 5108f009875a373ca940022dda554ab9bde86218f298de2dfcfd9fb9d9bae7a8

> Waiting for process to complete:
Privado CLI Version: v2.1.0
Privado Core Version: 1.1.0
Privado Main Version: 1.1.0

Configuration parsed...
Guessing source code language...
Detected language 'Java'
Processing source code using JAVASRC engine
Downloading dependencies and Parsing source code...
Tagging source code with rules...
Finding source to sink flow of data...
Deduplicating data flows...
Brewing result...

-----------------------------------------------------------
SUMMARY
-----------------------------------------------------------

Privado discovers data elements that are being collected, processed, or shared in the code.

DATA ELEMENTS  |  11 |
THIRD PARTY    |  1 |
STORAGES       |  3 |
ISSUES         |  0 |

---------------------------------------------------------
11 DATA ELEMENTS
Here is a list of data elements discovered in the code along with details on data flows to third parties, databases and leakages to logs.

1. PASSPORT
	Sharing         ->  fast2sms.com
	Storage         ->  AccountDao JBDC Connector, JPA Repository(Read), JPA Repository(Write)
	Leakage         ->  1
	Collections     ->  /accounts/add, /accounts/update/{id}
	Processing      ->  5

2. AGE
	Sharing         ->  fast2sms.com
	Storage         ->  AccountDao JBDC Connector, JPA Repository(Read), JPA Repository(Write)
	Leakage         ->  1
	Collections     ->  /accounts/add, /accounts/update/{id}
	Processing      ->  5

3. PHONE NUMBER
	Sharing         ->  fast2sms.com
	Storage         ->  AccountDao JBDC Connector, JPA Repository(Read), JPA Repository(Write)
	Leakage         ->  1
	Collections     ->  /accounts/add, /accounts/update/{id}, /transfer/betweenAccounts
	Processing      ->  6

4. DATE OF BIRTH
	Sharing         ->  fast2sms.com
	Storage         ->  AccountDao JBDC Connector, JPA Repository(Read), JPA Repository(Write)
	Leakage         ->  1
	Collections     ->  /accounts/add, /accounts/update/{id}
	Processing      ->  5

5. ACCOUNT PASSWORD
	Sharing         ->  fast2sms.com
	Storage         ->  AccountDao JBDC Connector, JPA Repository(Read), JPA Repository(Write)
	Leakage         ->  1
	Collections     ->  /accounts/add, /accounts/update/{id}
	Processing      ->  10

6. ACCOUNT NAME
	Sharing         ->  fast2sms.com
	Storage         ->  AccountDao JBDC Connector, JPA Repository(Read), JPA Repository(Write)
	Leakage         ->  3
	Collections     ->  /accounts/add, /accounts/update/{id}
	Processing      ->  13

7. LAST NAME
	Sharing         ->  fast2sms.com
	Storage         ->  AccountDao JBDC Connector, JPA Repository(Read), JPA Repository(Write)
	Leakage         ->  1
	Collections     ->  /accounts/add, /accounts/update/{id}
	Processing      ->  4

8. EMAIL ADDRESS
	Sharing         ->  fast2sms.com
	Storage         ->  AccountDao JBDC Connector, JPA Repository(Read), JPA Repository(Write)
	Leakage         ->  1
	Collections     ->  /accounts/add, /accounts/update/{id}
	Processing      ->  5

9. ACCOUNT ID
	Sharing         ->  fast2sms.com
	Storage         ->  AccountDao JBDC Connector, JPA Repository(Read), JPA Repository(Write)
	Leakage         ->  1
	Collections     ->  /accounts/add, /transfer/balanceAmountOnly/{accountId}, /transfer/balance/{accountId}, /accounts/update/{id}, /transfer/betweenAccounts, /transfer/transactionHistory/{accountId}
	Processing      ->  14

10. FIRST NAME
	Sharing         ->  fast2sms.com
	Storage         ->  AccountDao JBDC Connector, JPA Repository(Read), JPA Repository(Write)
	Leakage         ->  1
	Collections     ->  /accounts/add, /accounts/update/{id}
	Processing      ->  4

11. LANGUAGE PREFERENCES
	Sharing         ->  fast2sms.com
	Leakage         ->  1
	Processing      ->  1

Successfully exported output to '/Users/anujagrawal/Desktop/Projects/Privado/repos/BankingSystem-Backend/.privado' folder
```

On the console, you can see data elements and corresponding third parties, storages, leakages, collection points, and processing instances detected during the scan. A detailed report is also generated at `BankingSystem-Backend/.privado/privado.json`.

{% hint style="info" %}
Note that the actual result can be a bit different from the one shown above. It will depend on the version of Privado OSS installed and the repository that is being scanned
{% endhint %}

The scan usually runs for less than a minute, depending on the size of the repositories and dependencies.

#### Not working? <a href="#not-working-.1" id="not-working-.1"></a>

If you do not get the above result, check out the [troubleshooting section](command-reference.md#troubleshooting) for help

### Analyzing the result <a href="#analyzing-the-result" id="analyzing-the-result"></a>

After the scan is completed, the results will be stored in the `/.privado/privado.json` file inside the repository folder (`BankingSystem-Backend` in our case)

You can also look at the sample `privado.json` generated during a scan on 27th Sep 2022.

{% file src="../.gitbook/assets/privado.json" %}
Sample privado.json generated from the scan
{% endfile %}

## Troubleshooting <a href="#troubleshooting" id="troubleshooting"></a>

### Cloning the repository <a href="#cloning-the-repository" id="cloning-the-repository"></a>

If you are facing errors while cloning the repository, it can be due to the following reasons:

**Git not installed**

```
zsh: command not found: git
```

If you get the above message, it means that Git is not installed. Follow [these steps](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) to install Git

**Poor network connectivity**

```
Cloning into 'BankingSystem-Backend'...
fatal: unable to access 'https://github.com/saurabh-sudo/BankingSystem-Backend/': Could not resolve host: github.com
```

Make sure you have a stable internet connection and/or your firewall does not block GitHub repository cloning

### Running the scan <a href="#running-the-scan.1" id="running-the-scan.1"></a>

If the above command does not start the scan, it can be due to the following reasons:

**Docker engine not running**

```
> Scanning directory: /Users/anujagrawal/Desktop/Projects/Privado/repos/BankingSystem-Backend

> Pulling the latest image: public.ecr.aws/privado/cli:latest
Received error: Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
```

Make sure that Docker is installed and running on your machine.

**Unsupported languages**

```
> Scanning directory: /Users/anujagrawal/Desktop/Projects/Privado/repos/python-example

> Pulling the latest image: public.ecr.aws/privado/cli:niagara-dev
niagara-dev: Pulling from privado/cli
Digest: sha256:5b41906451148f8f9c9676a545f157016d71dfb241eb5d17d9690c7f5d1d8531
Status: Image is up to date for public.ecr.aws/privado/cli:niagara-dev

> Starting container with the latest image
> Container ID: e8cf930370a0232300109c1b5271797265d0c9946a9950138b4071d8e343326c

> Waiting for process to complete:
Privado CLI Version: dev
Privado Core Version: 0.0.169
Privado Main Version: 0.0.109

Configuration parsed...
Guessing source code language...
As of now we only support privacy code scanning for 'Java' code base.
We detected this code base of 'PYTHONSRC'.
```

While scanning any repository, make sure that the language is supported by Privado. You can find the languages supported by Privado by [clicking here](https://docs.privado.ai/#supported-languages).
