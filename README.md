# Automated Windows 10 configuration

## What we are trying to achieve
The ultimate goal of this project is to be able to run an .exe (Or whatever works) and make it automatically install software and configure my fresh Windows 10 intallation.

## Requirements
- Windows 10 x64

## Use the program
**!! Disclaimer !!**  
**This is a personal project and it is FAR from done, in it's current state you can't simply download and run it.**
- Download the latest [release zip](https://github.com/dennis1248/Automated-Windows-10-configuration/releases)
- Unpack the files
- Eddit the config.json
- dubble click the setup.exe

## Devolopment
Note: all devlopment is dune in the [src](./src) dir  
### Setup
- install [golang](https://golang.org/dl/)
- `$ go get github.com/akavel/rsrc`
- Windows: Add `%USERPROFILE%\go\bin` to System Variables ([how to](https://www.java.com/en/download/help/path.xml))
- Linux: `$ echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc && PATH=$PATH:$GOPATH/bin >> $HOME/.bashrc`
### Build and run the program
- execute `$ sh buildSetup.sh` using [git bash](https://git-scm.com/downloads) or [bash (Ubuntu, fedora, etc)](https://docs.microsoft.com/en-us/windows/wsl/install-win10)
- Open the setup.exe file inside the build dir
### Make a relaese
NOTE: the release does NOT use the config.json, it uses the config.example.json as config.json because then you can test the build with a modifyied config.json and don't have to worry about releaseing a wrong config 
- run `$ sh buildRelease.sh`
