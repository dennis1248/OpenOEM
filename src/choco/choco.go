package choco

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/dennis1248/Automated-Windows-10-configuration/src/fs"
	"github.com/dennis1248/Automated-Windows-10-configuration/src/functions"
	"github.com/dennis1248/Automated-Windows-10-configuration/src/options"
	"github.com/dennis1248/Automated-Windows-10-configuration/src/types"
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

		fmt.Println("Installing Chocolatey [1 of 2] Downloading installer..")
		ChocoInstallFile := "chocoSetup.ps1"
		err := funs.DownloadFile(ChocoInstallFile, "https://chocolatey.org/install.ps1")
		if err != nil {
			return err
		}

		// fmt.Println("Installing Chocolatey [2 of 3] adding go to user path")
		// skip this command for now because it might break the path variable :(
		// cmd = exec.Command("cmd", "/c", "set", "PATH=" + os.Getenv("PATH") + ";%ALLUSERSPROFILE%\\chocolatey\\bin")
		// cmd.CombinedOutput()

		fmt.Println("Installing Chocolatey [2 of 2] Running installer..")
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
				Chocolatey is installed but is not added to path 
				try restarting the program or
				run the installer manually:
				https://chocolatey.org/install
			`)
		}

	} else {
		fmt.Println("Chocolatey is already installed")
	}
	return nil
}

func InstallPkgList(conf types.Config) {
	// install all the programs

	// setting flags
	fmt.Println("Configuring Chocolatey..")
	cmd := exec.Command(
		"choco",
		"feature", "enable", "-n=allowGlobalConfirmation")
	_, err := cmd.CombinedOutput()
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

			cmd := exec.Command(
				"choco",
				"install", program,
				"--yes", "--force")
			_, err := cmd.CombinedOutput()
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

func PkgChecks(pkg string) error {
	// check if the package exists in the Chocolatey repos
	if len(pkg) == 0 {
		return errors.New("Package name must not be blank")
	}
	// check if the package exists in the Chocolatey repos
	cmd := exec.Command(
		"choco",
		"search", pkg)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	r, _ := regexp.Compile(strings.ToLower(pkg))
	check := r.MatchString(strings.ToLower(string(output)))
	if !check {
		return errors.New("Package not found, check the availability and proper naming of the package at https://chocolatey.org/")
	}
	// check if the package is already installed
	cmd = exec.Command(
		"choco",
		"search", "--lo", pkg)
	output, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}
	r, _ = regexp.Compile(strings.ToLower(pkg))
	check = r.MatchString(strings.ToLower(string(output)))
	if check {
		return errors.New("Package is already installed")
	}
	return nil
}

func InstallPackages() error {
	// install Chocolatey packages

	PackageName := options.GetOptions().PackageName
	packageJson, err := fs.FindPackageJson([]string{"./" + PackageName, "./../" + PackageName})
	if err != nil {
		return err
	}

	conf, err := fs.OpenPackageJson(packageJson)
	if err != nil {
		return err
	}

	fmt.Println("Installing packages..")
	InstallPkgList(conf)

	return nil
}
