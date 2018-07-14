# Automated Windows 10 configuration

## What is this?!
It's still far from done but the ultimate goal is to be able to run an .exe (Or whatever works) and make it automatically install software and configure my fresh Windows 10 intallation.

## Requirements
- Windows 10 x64
- Chocolatey *is not needed later*  

## Run
- Run the (base/choco_systemconfig_software.bat)[./base/choco_systemconfig_software.bat] file

## Devolopment
### Setup
- install [golang](https://golang.org/dl/)
- `$ go get github.com/akavel/rsrc`
- Windows: Add `%USERPROFILE%\go\bin` to System Variables ([how to](https://www.java.com/en/download/help/path.xml))
  Linux:
### Run
- `$ go run setup.go`
### Build program
- execute `$ sh buildSetup.sh` using [git bash](https://git-scm.com/downloads) or [bash (Ubuntu, fedora, etc)](https://docs.microsoft.com/en-us/windows/wsl/install-win10)
