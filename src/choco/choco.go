package choco

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/dennis1248/Automated-Windows-10-configuration/src/functions"
)

// this file contains all the chocolatery functions

func CheckForChoco() error {
	check := []string{"choco", "-v"}
	cmd := exec.Command(check[0], check[1:]...)
	_, err := cmd.CombinedOutput()
	return err
}

//check if chocolatey is installed or not:
func InstallIfNeededChocolatey() error {
	if CheckForChoco() != nil {
		// If chocolatey is not installed run the following:

		fmt.Println("Installing Chocolatey [1 of 2] Downloading installer")
		ChocoInstallFile := "chocoSetup.ps1"
		err := funs.DownloadFile(ChocoInstallFile, "https://chocolatey.org/install.ps1")
		if err != nil {
			return err
		}

		// fmt.Println("Installing Chocolatey [2 of 3] adding go to user path")
		// skip this command for now because it might break the path variable :(
		// cmd = exec.Command("cmd", "/c", "set", "PATH=" + os.Getenv("PATH") + ";%ALLUSERSPROFILE%\\chocolatey\\bin")
		// cmd.CombinedOutput()

		fmt.Println("Installing Chocolatey [2 of 2] run installer")
		cmd := exec.Command(
			"C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe",
			"-NoProfile",
			"-InputFormat", "None",
			"-ExecutionPolicy", "Bypass",
			"-file", ChocoInstallFile)
		_, err = cmd.CombinedOutput()
		if err != nil {
			return err
		}

		if CheckForChoco() != nil {
			return errors.New(`
				chocolatery is installed but is not added to path 
				try to restart the program or
				run the installer yourself:
				https://chocolatey.org/install
			`)
		}

	} else {
		fmt.Println("Chocolatey is already installed")
	}
	return nil
}
