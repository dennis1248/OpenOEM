# Automated Windows 10 configuration

## What we are trying to achieve
The ultimate goal of this project is to be able to run an .exe (Or whatever works) and make it automatically install software and configure your fresh Windows 10 intallation.

## Requirements
- Windows 10 x64

## Use the program
**!! Disclaimer !!**  
**This is a personal project and it is FAR from done, in it's current state you can't simply download and run it.**
- Download the latest [release zip](https://github.com/dennis1248/Automated-Windows-10-configuration/releases)
- Unpack the files
- Edit the config.json
- dubble click the setup.exe

## Devolopment
### NOTES:
Kinda important to read before doing anything
- Do **NOT** use git clone because that will give errors, [why](https://stackoverflow.com/questions/26942150/importing-go-files-in-same-folder)!!
- You can't use cmd and powershell because of the linux style script files, use [git bash](https://git-scm.com/downloads) or [bash (Ubuntu, fedora, etc)](https://docs.microsoft.com/en-us/windows/wsl/install-win10)
### Setup
- install [golang](https://golang.org/dl/)
- `$ go get github.com/akavel/rsrc`
- Add GOPATH to the system variables
  - Windows: Add `%USERPROFILE%\go\bin` to System Variables ([how to](https://www.java.com/en/download/help/path.xml))
  - Linux: Execute `$ echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc && PATH=$PATH:$GOPATH/bin >> $HOME/.bashrc`
- `$ go get github.com/dennis1248/Automated-Windows-10-configuration` or use youre repo if you have forked the repo
- In the output of the last command will be the direcotry to where the project is cloned to
### Build and run the program
- `$ cd scripts`
- execute `$ sh buildDev.sh` (The output exe file wil run automaticly look in task bar for the adminministrator tile)
- If you are on a linux system use: `$ sh buildRelease.sh`
### Make a release
NOTE: The release does NOT use the config.json, it uses the config.example.json as config.json so you can test the build with a modified config.json without having to worry about releasing a wrong config.
- `$ cd scripts`
- `$ sh buildRelease.sh`
