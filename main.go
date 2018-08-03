package main

import (
	"fmt"

	"github.com/dennis1248/Automated-Windows-10-configuration/src/choco"
	"github.com/dennis1248/Automated-Windows-10-configuration/src/functions"
	"github.com/dennis1248/Automated-Windows-10-configuration/src/registery"
	"github.com/dennis1248/Automated-Windows-10-configuration/src/theme"
)

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
	err = choco.InstallPackages()
	if err != nil {
		fmt.Println(err)
		funs.Die()
	}

	err = theme.SetTheme()
	if err != nil {
		fmt.Println("Error while installing theme, Error:")
		fmt.Println(err)
		fmt.Println()
	}

	err = registery.SetAllRegisteryItems()
	if err != nil {
		fmt.Println("Error while changing registery items, Error:")
		fmt.Println(err)
		fmt.Println()
	}

	fmt.Println("Done!")

	funs.EndTips()

	funs.Die()
}
