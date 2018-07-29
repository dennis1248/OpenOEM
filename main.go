package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dennis1248/Automated-Windows-10-configuration/src/choco"
	"github.com/dennis1248/Automated-Windows-10-configuration/src/functions"
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

	funs.CheckSYS()
	fmt.Println("Starting setup...")

	// check if Chocolatey is installed, if not try to install it
	err := choco.InstallIfNeededChocolatey()
	if err != nil {
		fmt.Println("can't run Chocolatey installer, Error: \n", err)
		funs.Die()
	}

	// install choco packages specified in the package file
	err = installPackages()
	if err != nil {
		fmt.Println(err)
		funs.Die()
	}

	fmt.Println("Dune!")
	funs.Die()
}
