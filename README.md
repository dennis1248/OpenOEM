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
### Run
- `$ go run setup.go`
### Make FILE.syso
The FILE.syso contains extra info about the program in this case it's will specifiy that the program needs to be ran with administrator writes
- `$ rsrc -manifest APP.EXE.manifest -o FILE.syso`
