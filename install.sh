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

    mkdir -p ~/.privado/bin
    if [[ "$OS" == "windows" ]]; then
	curl -L "$BASE_URL$OS-$ARCH.zip" -o /tmp/privado-$OS-$ARCH.zip
	curl -L "$BASE_URL$OS-$ARCH.zip.md5" -o /tmp/privado-$OS-$ARCH.zip.md5
	MD5_ACTUAL=$(certutil -hashfile /tmp/privado-$OS-$ARCH.zip MD5)
	MD5_EXPECTED=$(cat /tmp/privado-$OS-$ARCH.zip.md5)
    else
	curl -L "$BASE_URL$OS-$ARCH.tar.gz" -o /tmp/privado-$OS-$ARCH.tar.gz
	curl -L "$BASE_URL$OS-$ARCH.tar.gz.md5" -o /tmp/privado-$OS-$ARCH.tar.gz.md5
	if [[ "$OS" == "linux" ]]; then
		MD5_ACTUAL=$(md5sum /tmp/privado-$OS-$ARCH.tar.gz | awk '{print $1}')
	else
		MD5_ACTUAL=$(md5 /tmp/privado-$OS-$ARCH.tar.gz | awk '{print $4}')
	fi
	MD5_EXPECTED=$(cat /tmp/privado-$OS-$ARCH.tar.gz.md5)
    fi

    if [[ "$MD5_EXPECTED" != "$MD5_ACTUAL" ]]; then
	echo "Error in downloading the file. Please retry"
	exit 10
    fi

    if [[ "$OS" == "windows" ]]; then
	unzip /tmp/privado-$OS-$ARCH.zip -d ~/.privado/bin
    elif [[ "$OS" == "darwin" ]]; then
        tar -xf /tmp/privado-$OS-$ARCH.tar.gz -C ~/.privado/bin
    else
	tar -xf /tmp/privado-$OS-$ARCH.tar.gz -C ~/.privado/bin
    fi

    PROFILE_PATH="~/.bashrc"

    for EACH_PROFILE in ".profile" ".bashrc" ".bash_profile" ".zshrc"
    do
      if [[ -f $HOME/$EACH_PROFILE ]]; then
	cat $HOME/$EACH_PROFILE | grep "/.privado" || echo "export PATH=\$PATH:~/.privado/bin" >> $HOME/$EACH_PROFILE
	PROFILE_PATH=$HOME/$EACH_PROFILE
        break
      fi
    done

    echo "Installation is complete. Please open a new session or run the command"
    echo ". $PROFILE_PATH"
    echo "In your existing session"  
}

findOS
findArch
downloadAndInstallLatestVersion
