package choco

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dennis1248/Automated-Windows-10-configuration/src/commands"
	"github.com/dennis1248/Automated-Windows-10-configuration/src/fs"
	"github.com/dennis1248/Automated-Windows-10-configuration/src/functions"
	"github.com/dennis1248/Automated-Windows-10-configuration/src/types"
)

// this file contains all the chocolatey functions

// CheckForChoco checks if choco works
func CheckForChoco() error {
	_, err := commands.ChocoRun("-v")
	return err
}

// InstallIfNeededChocolatey checks if chocolatey is installed or not:
func InstallIfNeededChocolatey() error {
	if CheckForChoco() != nil {
		// If chocolatey is not installed run the following:

		fmt.Println("Installing Chocolatey [1 of 2] Downloading installer..")
		ChocoInstallFile := "chocoSetup.ps1"
		err := funs.DownloadFile(ChocoInstallFile, "https://chocolatey.org/install.ps1")
		if err != nil {
			return err
		}

		fmt.Println("Installing Chocolatey [2 of 2] Running installer..")
		_, err = commands.Run(
			"C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe",
			"-NoProfile",
			"-InputFormat", "None",
			"-ExecutionPolicy", "Bypass",
			"-file", ChocoInstallFile)
		if err != nil {
			return err
		}

		if CheckForChoco() != nil {
			return errors.New(`
				Chocolatey isn't installed
				try running the program again or
				run the installer manually:
				https://chocolatey.org/install
			`)
		}

	} else {
		fmt.Println("Chocolatey is already installed")
	}
	return nil
}

// InstallPkgList installs all the programs included in the config file
func InstallPkgList(conf types.Config) {

	// setting flags
	fmt.Println("Configuring Chocolatey..")
	_, err := commands.ChocoRun(
		"feature", "enable", "-n=allowGlobalConfirmation")
	if err != nil {
		fmt.Println("Unable to enable Chocolatey feature, Error:", err)
		funs.Die()
	}

	for i, program := range conf.Programs {
		fmt.Println("Installing: [" + strconv.Itoa(i+1) + " of " + strconv.Itoa(len(conf.Programs)) + "] " + program)

		// run package checks
		err = PkgChecks(program)
		if err == nil {
			// Install the package

			_, err := commands.ChocoRun(
				"install", program,
				"--yes", "--force", "--allowunofficial", "--use-system-powershell")
			if err != nil {
				fmt.Println("Unable to install:", program, "Reason:", err)
			} else {
				fmt.Println("Installed: " + program)
			}
		} else {
			fmt.Println("Skipping:", program, err)
		}
	}
}

// PkgChecks checks if the package is installeble
func PkgChecks(pkg string) error {

	// check if the package exists in the Chocolatey repos
	if len(pkg) == 0 {
		return errors.New("Package name must not be blank")
	}

	// check if the package exists in the Chocolatey repos
	output, err := commands.ChocoRun(
		"search", pkg)
	if err != nil {
		return err
	}

	r, _ := regexp.Compile(strings.ToLower(pkg) + "\\s")
	check := r.MatchString(strings.ToLower(string(output)))
	if !check && strings.Contains(pkg, "-") {
		// replace "-" with "" because sometimes the "-" breaks choco search
		output, err = commands.ChocoRun("search", strings.Replace(pkg, "-", "", -1))
		if err != nil {
			return err
		}
		r, _ = regexp.Compile(strings.ToLower(pkg) + "\\s")
		check = r.MatchString(strings.ToLower(string(output)))
	}

	if !check {
		return errors.New("Package not found, check the availability and proper naming of the package at https://chocolatey.org/packages")
	}

	// check if the package is already installed
	output, err = commands.ChocoRun("search", "--lo", pkg)
	if err != nil {
		return err
	}

	r, _ = regexp.Compile(strings.ToLower(pkg) + "\\s")
	check = r.MatchString(strings.ToLower(string(output)))
	if check {
		return errors.New("Package is already installed")
	}

	return nil
}

// InstallPackages install Chocolatey packages
func InstallPackages() error {

	conf, err := fs.FindAndOpenPackageJSON()
	if err != nil {
		return err
	}

	fmt.Println("Installing packages..")
	InstallPkgList(conf)

	return nil
}
