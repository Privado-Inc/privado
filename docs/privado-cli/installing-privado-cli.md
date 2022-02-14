# Installing Privado CLI

To start off, make sure `docker` is installed. To install docker, you can follow the steps stated in the [official documentation](https://docs.docker.com/engine/install/). You can install Privado CLI in multiple manners:  

- [Using `curl`](#install-using-curl)
- [Using `go`](#install-using-go)
- [Using releases](#install-release-manually)


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

This will place the `privado` binary in your `GOPATH`'s bin directory. This directory must be added to the `$PATH` environment variable. You can learn more [here](https://www.digitalocean.com/community/tutorial_series/how-to-install-and-set-up-a-local-programming-environment-for-go). 

## Install Release Manually
We use [GitHub Releases](https://github.com/Privado-Inc/privado/releases) to ship versioned `privado` releases for supported platforms. You can download a executable of Privado CLI for your platform.

For example, to setup `privado` for MAC-OSX (arm64), you will download `privado-darwin-amd64.tar.gz`, and run:
```
tar -xf ~/.privado/privado-darwin-amd64.tar.gz
chmod +x privado
mv privado /usr/bin/privado
```


{% hint style="info" %}
**Good to know:** your product docs aren't just a reference of all your features! use them to encourage folks to perform certain actions and discover the value in your product.
{% endhint %}
