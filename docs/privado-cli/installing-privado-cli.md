# Installing Privado CLI

To start off, make sure `docker` is installed. To install docker, you can follow the steps stated in the [official documentation](https://docs.docker.com/engine/install/). You can install Privado CLI in multiple manners:

* [Using `curl`](installing-privado-cli.md#install-using-curl)
* [Using `go`](installing-privado-cli.md#install-using-go)
* [Using releases](installing-privado-cli.md#install-release-manually)
* [Build locally](installing-privado-cli.md#build-privado-cli-locally)

## Install using `curl`:

The installation script will download and setup the latest stable release for you as per your OS and arch. Run:

```
curl -o- https://raw.githubusercontent.com/Privado-Inc/privado/main/install.sh | bash
```

To uninstall, simply delete `~/.privado/bin`.

## Install using Go

If you are a GoLang fan, you can use the `go install` command to install the Privado CLI:

```
go install github.com/Privado-Inc/privado@latest
```

This will place the `privado` binary in your `GOPATH`'s bin directory. This directory must be added to the `$PATH` environment variable. You can learn more [here](https://www.digitalocean.com/community/tutorial\_series/how-to-install-and-set-up-a-local-programming-environment-for-go).

## Install Release Manually

We use [GitHub Releases](https://github.com/Privado-Inc/privado/releases) to ship versioned `privado` releases for supported platforms. You can download a executable of Privado CLI for your platform.   

To know your architecture, you can run: 
```
$ uname -m
```
 
### MacOSX
#### ARM64 (M1 Chip)
To setup `privado` for macOS (arm64) i.e. Macbook with M1 chip, download `privado-darwin-arm64.tar.gz` from the [latest release](https://github.com/Privado-Inc/privado/releases/tag/latest). 

Navigate to the download directory and run:

```
$ tar -xf ~/.privado/privado-darwin-arm64.tar.gz
$ chmod +x privado
$ mv privado /usr/local/bin/
```

#### AMD64 (Intel Chip)
To setup `privado` for macOS (amd64), download `privado-darwin-amd64.tar.gz` from the [latest release](https://github.com/Privado-Inc/privado/releases/tag/latest). 

Navigate to the download directory and run:

```
$ tar -xf ~/.privado/privado-darwin-amd64.tar.gz
$ chmod +x privado
$ mv privado /usr/local/bin/
```

### Linux
To setup `privado` on your linux system, download the respective zip from [latest release](https://github.com/Privado-Inc/privado/releases/tag/latest) for your platform. Navigate to the download directory and run the following commands:

#### ARM64
```
$ tar -xf ~/.privado/privado-linux-arm64.tar.gz
$ chmod +x privado
$ mv privado /usr/bin/privado
```

#### AMD64
```
$ tar -xf ~/.privado/privado-linux-amd64.tar.gz
$ chmod +x privado
$ mv privado /usr/bin/privado
```

### Windows
To setup `privado` on your windows system, download `privado-windows-amd64.zip` from [latest release](https://github.com/Privado-Inc/privado/releases/tag/latest). Navigate to the download directory and run the following `bash` commands:

```
$ mkdir -p $HOME/.privado/bin
$ unzip -o privado-windows-amd64.zip -d $HOME/.privado/bin
$ chmod +x $HOME/.privado/bin/privado
$ echo "export PATH=\$PATH:$HOME/.privado/bin" >> $HOME/.bashrc
```

Open a new session or source profile for effects to take place in the same session:
```
$ source $HOME/.bashrc
```

When using [WSL](https://docs.microsoft.com/en-us/windows/wsl/), we recommend moving the binary to `/usr/bin` instead for optimal experience across users. Refer to [Installation for Linux](#linux) for more information.

## Build Privado CLI Locally
If you do not wish to use the pre-built binaries shipped in releases, you can choose to build Privado CLI locally. To do this, make sure that [GoLang](https://go.dev/doc/install) is installed and follow the following steps:

1. Clone the repository:
    ```
    git clone https://github.com/Privado-Inc/privado.git
    ```

2. Change directory: `cd privado`

3. To build the [latest stable release](https://github.com/Privado-Inc/privado/releases/tag/latest), checkout the `latest` tag:
    > Skip this step if you intend to build the `main` branch.

    ```
    git checkout latest
    ```

4. Build with Go
    ```
    go build
    ```
5. You can now run `./privado`.   

For convenience, we recommend moving `privado` to a `$PATH` directory. You can refer to manual installation steps for more details.
