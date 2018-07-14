package main

import (
	"fmt"
	"runtime"
	"os"
	"os/exec"
	"io"
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
	// check youre system if everything is supported 
	status := true
	args := os.Args[1:]
	skip := contains(args, "--skipChecks") || contains(args, "-s")
	if runtime.GOOS != "windows" {
		fmt.Println("This program can only be ran on a windows PC")
		status = false
	}
	if !checkIfAdmin() {
		fmt.Println("This program can only be ran with administrator rights")
	}
	if !status && !skip {
		fmt.Println("Use --skipChecks or -s to skip the checks")
		os.Exit(0)
	}
}

func checkIfAdmin() bool {
	// check if the program gets ran as admin
	// skip this part for now because i don't know how to check for administrator writes
	return true
}

func installIfNeededChocolatey() error {
	cmd := exec.Command("choco", "-v")
	_, err := cmd.CombinedOutput()
	if err != nil {
		// choco is not installed run the choco setup
		fmt.Println("Installing Chocolatey...")
		ChocoInstallFile := "chocoSetup.ps1"
		err := DownloadFile(ChocoInstallFile, "https://chocolatey.org/install.ps1")
		if err != nil {
			return err
		}
		cmd1 := exec.Command(
			"C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe", 
			"-NoProfile", 
			"-InputFormat", "None", 
			"-ExecutionPolicy", "Bypass", 
			"-file", ChocoInstallFile)
		_, cmd1err := cmd1.CombinedOutput()
		if cmd1err != nil {
			return cmd1err
		}
		fmt.Println("Installed choco... setting env variables")
		// cmd2 := exec.Command("SET", "\"PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin\"")
		// if 
	} else {
		fmt.Println("Chocolatey is already installed")
	}
	return nil
}

func main() {
	checkSYS()
	fmt.Println("Starting setup...")
	
	// install choco
	err := installIfNeededChocolatey()
	if err != nil {
		fmt.Println("can't run Chocolatey installer, Error: %s\n", err)
	}

	fmt.Println("Dune!")

	// force the program to run forever
	fmt.Scanln()
}