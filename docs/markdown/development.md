# Development
### NOTES
Please carefully read the following before you start playing around with it;
- Do not use git clone because that will give errors, [why](https://stackoverflow.com/questions/26942150/importing-go-files-in-same-folder)
-  As-is it can only be compiled on a Linux-based client, use [git bash](https://git-scm.com/downloads) or [bash (Ubuntu, fedora, etc)](https://docs.microsoft.com/en-us/windows/wsl/install-win10)
- Do not store files or modify files in the build directory because then you will likely end up breaking everything and files might be overwritten
### Setup
- install [Golang](https://golang.org/dl/)
- Add GOPATH to the system variables
  - Windows: Add `%USERPROFILE%\go\bin` to System Variables ([how to](https://www.java.com/en/download/help/path.xml))
  - Linux: Execute `$ echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc && echo 'PATH=$PATH:$GOPATH/bin' >> $HOME/.bashrc`
- `$ go get github.com/akavel/rsrc`
- `$ go get github.com/dennis1248/OpenOEM` or use your repo if you have forked this project
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
Read: [docs/markdown/testing.md](https://github.com/dennis1248/OpenOEM/blob/master/docs/markdown/testing.md)

### VScode And Linux Bug
if you get this error under the tab PROBLEMS:  
```
go build errors: mkdir /usr/lib/golang/pkg/windows_386/: permission denied
go build internal/race: mkdir /usr/lib/golang/pkg/windows_386: permission denied
...
```  
Do this:
- `$ su -`
- `$ cd /usr/lib/golang/pkg`
- `$ ln -s linux_amd64 windows_386`
- `$ chmod 777 linux_amd64 -R`
