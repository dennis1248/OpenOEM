package commands

import (
	"os/exec"
)

func Run(commands ...string) (output []byte, err error) {
	// a simple version of exec.Command + cmd.CombinedOutput
	cmd := exec.Command(commands[0], commands[1:]...)
	return cmd.CombinedOutput()
}

func ChocoRun(commands ...string) (output []byte, err error) {
	// run chocolatey command with full choco path
	chocoDir := "C:\\ProgramData\\chocolatey\\choco.exe"
	fullCommand := append([]string{chocoDir}, commands...)
	return Run(fullCommand...)
}

func PSRun(command string) (output []byte, err error) {
	// run a command inside powershell
	return Run(
		"C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe",
		"-NoProfile",
		"-InputFormat", "None",
		"-ExecutionPolicy", "Bypass",
		"-Command", command)
}
