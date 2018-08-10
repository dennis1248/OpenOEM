package registery

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/dennis1248/OpenOEM/src/commands"
	"github.com/dennis1248/OpenOEM/src/fs"
)

// setRegValue set registery value
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

// SetSearch chang the look of search to a full search, just an icon or nothing
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

// SetTaskView set the task view icon to be visible or removed
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

// RemovePeople removes "People" from the taskbar
func RemovePeople(removeBtn bool) error {
	SetTo := "1"
	if removeBtn {
		SetTo = "0"
	}
	return setRegValue(
		"HKCU:\\Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced\\People",
		"PeopleBand",
		"DWORD",
		SetTo)
}

// RemoveJunkApps Removes the junk apps from the start menu
// Currently the apps stay installed, only the shortcuts are removed
// We tried uninstalling them but it causes some serious issues
func RemoveJunkApps(allow bool) error {
	if !allow {
		return nil
	}

	// This command will not work when just executed it in the powershell
	// remove the ` (escape characters) to run it
	_, err := commands.PSRunBypass(
		`(New-Object -Com Shell.Application).NameSpace('shell:::{4234d49b-0245-4df3-b780-3893943456e1}').Items() | %{$_.Verbs() } | ?{$_.Name -match 'Un.*pin from Start'} | %{$_.DoIt()}`)

	// This command will break your Windows installation
	// https://github.com/dennis1248/OpenOEM/issues/10
	// commands.PSRun("Get-AppXPackage | where-object {$_.name –notlike “*store*”} | Remove-AppxPackage --quiet --no-verbose >$null 2>&1")

	return err
}

// RestartUI restarts the explorer.exe in a "safe" way to avoid issues
func RestartUI() {
	commands.PSRun("Stop-Process -ProcessName explorer")
}

// SetAllRegisteryItems The main function of the file
func SetAllRegisteryItems() error {

	Package, err := fs.FindAndOpenPackageJSON()
	if err != nil {
		return err
	}

	err = RemoveJunkApps(Package.RemoveJunk)
	if err != nil {
		fmt.Println("Unable to remove apps item, Error:", err)
	}

	err = SetSearch(Package.Search)
	if err != nil {
		fmt.Println("Unable to set registery search item, Error:", err)
	}

	err = SetTaskView(Package.TaskView)
	if err != nil {
		fmt.Println("Unable to set registery task view item, Error:", err)
	}

	err = RemovePeople(Package.RemovePeople)
	if err != nil {
		fmt.Println("Unable to set registery people button, Error:", err)
	}

	RestartUI()

	return nil
}
