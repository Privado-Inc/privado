#!/bin/bash

OS=""
ARCH=""
BASE_URL="https://github.com/Privado-Inc/privado/releases/download/latest/privado-"

function findOS {
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        OS="linux"
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        OS="darwin"
    elif [[ "$OSTYPE" == "msys" ]]; then 
        OS="windows"
    else
	echo "Unsupported OS"
	exit 10
    fi
}

function findArch {
    ARCH_STR=$(uname -m)
    if [[ "$ARCH_STR" == "x86_64" ]]; then
	ARCH="amd64"
    elif [[ "$ARCH_STR" == "arm64" ]]; then
	ARCH="arm64"
    else
	echo "Unsupported Architecture"
        exit 10
    fi
}

function downloadAndInstallLatestVersion {
    if [[ "$OS" == "" || "$ARCH" == "" ]]; then
        echo "Unsupported OS or Arch type. Please visit https://privado.ai/cli for more information."
        exit 1
    fi

    mkdir -p ~/privado
    if [[ "$OS" == "windows" ]]; then
	curl -L "$BASE_URL$OS-$ARCH.zip" -o ~/privado/privado-$OS-$ARCH.zip
	curl -L "$BASE_URL$OS-$ARCH.zip.md5" -o ~/privado/privado-$OS-$ARCH.zip.md5
	MD5_ACTUAL=$(certutil -hashfile ~/privado/privado-$OS-$ARCH.zip MD5)
	MD5_EXPECTED=$(cat ~/privado/privado-$OS-$ARCH.zip.md5)
    else
	curl -L "$BASE_URL$OS-$ARCH.tar.gz" -o ~/privado/privado-$OS-$ARCH.tar.gz
	echo "$BASE_URL$OS-$ARCH.tar.gz.md5"
	curl -L "$BASE_URL$OS-$ARCH.tar.gz.md5" -o ~/privado/privado-$OS-$ARCH.tar.gz.md5
	if [[ "$OS" == "linux" ]]; then
		MD5_ACTUAL=$(md5sum ~/privado/privado-$OS-$ARCH.tar.gz | awk '{print $1}')
	else
		MD5_ACTUAL=$(md5 ~/privado/privado-$OS-$ARCH.tar.gz)
	fi
	MD5_EXPECTED=$(cat ~/privado/privado-$OS-$ARCH.tar.gz.md5)
    fi

    if [[ "$MD5_EXPECTED" != "$MD5_ACTUAL" ]]; then
	echo "Error in downloading the file. Please retry"
	exit 10
    fi

    if [[ "$OS" == "windows" ]]; then
	unzip ~/privado/privado-$OS-$ARCH.zip -d /usr/local/bin
    elif [[ "$OS" == "darwin" ]]; then
        tar -xf ~/privado/privado-$OS-$ARCH.tar.gz -C /usr/local/bin
    else
	tar -xf ~/privado/privado-$OS-$ARCH.tar.gz -C /usr/local/bin
    fi
}

findOS
findArch
downloadAndInstallLatestVersion
