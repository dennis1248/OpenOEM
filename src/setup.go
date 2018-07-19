package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
)

// types
// a type pre defines what a object / array contains

// Generated json types using: https://mholt.github.io/json-to-go/
type Config struct {
	ProgramsSlash string   `json:"// programs"`
	Programs      []string `json:"programs"`
}

// Type structure to bind to config.
type Options struct {
	PackageName string
}

// used when the program needs to stop because of an error
func die() {
	fmt.Println("press any key to exit the program.")
	// Prevent the application from closing
	fmt.Scanln()
	os.Exit(0)
}

// Check if a array contains value
func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// Download a file from the internet
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
		die()
	}
}

func checkIfAdmin() bool {
	// check if the program is going to run width administative rights
	// skip this part for now, I don't know yet how to check for administrator rights
	return true
}

func checkForChoco() error {
	check := []string{"choco", "-v"}
	cmd := exec.Command(check[0], check[1:]...)
	_, err := cmd.CombinedOutput()
	return err
}

//check if chocolatey is installed or not:
func installIfNeededChocolatey() error {
	if checkForChoco() != nil {
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

		if checkForChoco() != nil {
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

func findPackageJson(toCheck []string) (string, error) {
	// this function returns the location of a found package file
	// currently this only checks in this directory and one above
	fmt.Println("Searching for config file")
	toReturn := ""
	for _, check := range toCheck {
		fullPath, _ := filepath.Abs(check)
		_, err := os.Stat(fullPath)
		if err == nil {
			toReturn = fullPath
		}
	}
	if toReturn == "" {
		return toReturn, errors.New("No " + getOptions().PackageName + " found in this directory")
	}
	return toReturn, nil
}

func openPackageJson(packageJsonFile string) (out Config, err error) {
	// returns the output of the config file
	fmt.Println("Using this config file:", packageJsonFile)
	fileContent, err := ioutil.ReadFile(packageJsonFile)
	if err != nil {
		return Config{}, err
	}
	var data Config
	json.Unmarshal([]byte(fileContent), &data)
	return data, nil
}

func installPkgList(conf Config) {
	// install all the programs
	for i, program := range conf.Programs {
		fmt.Println("Installing: [" + strconv.Itoa(i+1) + " of " + strconv.Itoa(len(conf.Programs)) + "] " + program)

		// run command to install the program,
		// dont forget to add
		// choco feature enable -n=allowGlobalConfirmation
		// otherwise it will fail

		fmt.Println("Installed: " + program)
	}
}

func installPackages() error {
	// install Chocolatey packages

	PackageName := getOptions().PackageName
	packageJson, err := findPackageJson([]string{"./" + PackageName, "./../" + PackageName})
	if err != nil {
		return err
	}

	conf, err := openPackageJson(packageJson)
	if err != nil {
		return err
	}

	fmt.Println("Insatlling programs")
	installPkgList(conf)

	return nil
}

func getOptions() Options {
	return Options{
		PackageName: "config.json"}
}

func main() {

	checkSYS()
	fmt.Println("Starting setup...")

	// check if Chocolatey is installed, if not try to install it
	err := installIfNeededChocolatey()
	if err != nil {
		fmt.Println("can't run Chocolatey installer, Error: \n", err)
		die()
	}

	// install choco packages specified in the package file
	err = installPackages()
	if err != nil {
		fmt.Println(err)
		die()
	}

	fmt.Println("Dune!")
	die()
}
