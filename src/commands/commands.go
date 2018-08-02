package commands

import (
	"os/exec"
)

func Run(commands ...string) (output []byte, err error) {
	// a simpeler version of exec.Command + cmd.CombinedOutput
	cmd := exec.Command(commands[0], commands[1:]...)
	return cmd.CombinedOutput()
}

func ChocoRun(commands ...string) (output []byte, err error) {
	// run chocolatery command with full choco path
	chocoDir := "C:\\ProgramData\\chocolatey\\choco.exe"
	fullCommand := append([]string{chocoDir}, commands...)
	return Run(fullCommand...)
}
