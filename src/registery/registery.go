package registery

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/dennis1248/Automated-Windows-10-configuration/src/commands"
	"github.com/dennis1248/Automated-Windows-10-configuration/src/fs"
)

// edit registery things

func setRegValue(Path string, Item string, ValueType string, Newval string) error {

	regex := "^DWORD$"
	r, _ := regexp.Compile(regex)
	check := r.MatchString(ValueType)
	if !check {
		return errors.New("ValueType wrong, input: " + ValueType + " Regex to match: " + regex)
	}

	commandToRun := "New-ItemProperty -Path " + Path + " -Name " + Item + " -Value " + Newval + " -PropertyType " + ValueType + " -Force | Out-Null"
	// if the command is not working echo out the line under here to view the command
	// fmt.Println(commandToRun)
	_, err := commands.PSRun(commandToRun)

	return err
}

func SetSearch(setTo string) error {

	SetToRegistery := "2"
	if setTo == "icon" {
		SetToRegistery = "1"
	} else if setTo == "hidden" {
		SetToRegistery = "0"
	}

	return setRegValue(
		"HKCU:\\Software\\Microsoft\\Windows\\CurrentVersion\\Search",
		"SearchboxTaskbarMode",
		"DWORD",
		SetToRegistery)
}

func SetTaskView(SetTo bool) error {
	SetToRegistery := "1"
	if !SetTo {
		SetToRegistery = "0"
	}
	return setRegValue(
		"HKCU:\\Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced",
		"ShowTaskViewButton",
		"DWORD",
		SetToRegistery)
}

func RestartUI() {
	commands.PSRun("Stop-Process -ProcessName explorer")
}

func SetAllRegisteryItems() error {

	Package, err := fs.FindAndOpenPackageJson()
	if err != nil {
		return err
	}

	err = SetSearch(Package.Search)
	if err != nil {
		fmt.Println("can't set registery search item, Error:", err)
	}

	err = SetTaskView(Package.TaskView)
	if err != nil {
		fmt.Println("can't set registery task view item, Error:", err)
	}

	RestartUI()

	return nil
}
