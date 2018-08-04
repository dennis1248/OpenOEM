# Automated Windows 10 configuration
![Project logo](/src/icon-small.png?raw=true "Project logo")

While you drink a cup of coffee we do the rest  
This automated tool will Install programs from chocolatey, give windows your favourite colour, set a wallpaper, remove annoying UI elements and much more from just a config file

## What can this tool do
- :heavy_check_mark: **Install** programs from **chocolatey** 
- :heavy_check_mark: Setting **theme colour** and **wallpaper**  
- :heavy_check_mark: Change or **remove** the **task view** and **search**
- If you have any suggestions create an [issue](https://github.com/dennis1248/Automated-Windows-10-configuration/issues)

## Requirements
- Windows 10 x64  

## Use the program
**!! Disclaimer !!**  
**This application has not been properly tested yet, you could encounter bugs and unexpected behaviour.**
- Use the [winconfig tool](https://winconfig.masterracegeeks.eu/#generator)
**OR**
- Download the latest [release zip](https://github.com/dennis1248/Automated-Windows-10-configuration/releases)
- Unpack the files
- Edit the config.json
- Double click the setup.exe or save it for later use :)

## Development
### NOTES
Please carefully read the following before you start playing around with it;
- Do not use git clone because that will give errors, [why](https://stackoverflow.com/questions/26942150/importing-go-files-in-same-folder)
-  As-is it can only be compiled on a Linux-based client, use [git bash](https://git-scm.com/downloads) or [bash (Ubuntu, fedora, etc)](https://docs.microsoft.com/en-us/windows/wsl/install-win10)
- Do not store files or modify files in the build directory because then you will likely end up breaking everything and files might be overwritten
### Setup
- install [Golang](https://golang.org/dl/)
- `$ go get github.com/akavel/rsrc`
- Add GOPATH to the system variables
  - Windows: Add `%USERPROFILE%\go\bin` to System Variables ([how to](https://www.java.com/en/download/help/path.xml))
  - Linux: Execute `$ echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc && PATH=$PATH:$GOPATH/bin >> $HOME/.bashrc`
- `$ go get github.com/dennis1248/Automated-Windows-10-configuration` or use your repo if you have forked this project
- The output of the last command will contain the location of the directory where the project is cloned to
### Compile the code
make sure you are inside the scripts folder: `$ cd scripts`  

Exec | Output
--- | ---
`$ sh buildSetup.sh` | Build just the setup.exe file for development  
`$ sh buildDev.sh` | Build the program and execute  
`$ sh buildRelease.sh` | Build a release setup.exe file  
`$ sh buildTest.sh` | Build and add the testing wallpaper and config
  
### Testing
Read: [docs/markdown/testing.md](https://github.com/dennis1248/Automated-Windows-10-configuration/blob/master/docs/markdown/testing.md)
