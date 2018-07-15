# Automated Windows 10 configuration

## Goal of this project?!
It's still far from done but the ultimate goal is to be able to run an .exe (Or whatever works) and make it automatically install software and configure my fresh Windows 10 intallation.

## Requirements
- Windows 10 x64

## Devolopment
Note all devlopment is dune in the [GOlang_compile](./GOlang_compile) dir
### Setup
- install [golang](https://golang.org/dl/)
- `$ go get github.com/akavel/rsrc`
- Windows: Add `%USERPROFILE%\go\bin` to System Variables ([how to](https://www.java.com/en/download/help/path.xml))
- Linux: `$ echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc && PATH=$PATH:$GOPATH/bin >> $HOME/.bashrc`
### Run
- `$ go run setup.go`
### Build program
- execute `$ sh buildSetup.sh` using [git bash](https://git-scm.com/downloads) or [bash (Ubuntu, fedora, etc)](https://docs.microsoft.com/en-us/windows/wsl/install-win10)

## Disclaimer
This is a personal project and it is FAR from done, in it's current state you can't simply download and run it.
