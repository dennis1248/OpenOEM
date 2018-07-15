package main

import (
	"os"
	"io"
	"fmt"
	"runtime"
	"os/exec"
	"net/http"
)

func contains(arr []string, str string) bool {
	for _, a := range arr {
		 if a == str {
				return true
		 }
	}
	return false
}

func DownloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func checkSYS() {
	// check if your system is supported
	status := true
	args := os.Args[1:]
	skip := contains(args, "--skipChecks") || contains(args, "-s")
	if runtime.GOOS != "windows" {
		fmt.Println("This applications appears to be running on a non Windows system.")
		status = false
	}
	if !checkIfAdmin() {
		fmt.Println("This application requires administative rights to run.")
	}
	if !status && !skip {
		fmt.Println("Use --skipChecks or -s to skip checks.")
		os.Exit(0)
	}
}

func checkIfAdmin() bool {
	// check if the program is going to run width administative rights
	// skip this part for now, I don't know yet how to check for administrator rights
	return true
}

//check if chocolatey is installed or not:
func installIfNeededChocolatey() error {
	check := []string{"choco", "-v"}
	cmd := exec.Command(check[0], check[1:]...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		// If chocolatey is not installed run the following:

		fmt.Println("Installing Chocolatey [1 of 2] Downloading installer")
		ChocoInstallFile := "chocoSetup.ps1"
		err := DownloadFile(ChocoInstallFile, "https://chocolatey.org/install.ps1")
		if err != nil {
			return err
		}

		// fmt.Println("Installing Chocolatey [2 of 3] adding go to user path")
		// skip this command for now because it might break the path variable :(
		// cmd = exec.Command("cmd", "/c", "set", "PATH=" + os.Getenv("PATH") + ";%ALLUSERSPROFILE%\\chocolatey\\bin")
		// cmd.CombinedOutput()

		fmt.Println("Installing Chocolatey [2 of 2] run installer")
		cmd = exec.Command(
			"C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe", 
			"-NoProfile", 
			"-InputFormat", "None", 
			"-ExecutionPolicy", "Bypass", 
			"-file", ChocoInstallFile)
		_, err = cmd.CombinedOutput()
		if err != nil {
			return err
		}

		fmt.Println("Check if Chocolatey was installed successfull")
		cmd = exec.Command(check[0], check[1:]...)
		_, err = cmd.CombinedOutput()
		if err != nil {
			return err
		}

	} else {
		fmt.Println("Chocolatey is already installed")
	}
	return nil
}

func installPackages() {
	// install Chocolatey packages
	fmt.Println("Insatlling programs")
}

func main() {
	checkSYS()
	fmt.Println("Starting setup...")
	
	// do the Chocolatey stuff
	err := installIfNeededChocolatey()
	if err != nil {
		fmt.Println("can't run Chocolatey installer, Error: \n", err)
	} else {
		installPackages()
	}

	fmt.Println("Dune!")

	// Prevent the application from closing
	fmt.Scanln()
}
