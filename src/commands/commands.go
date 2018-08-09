package commands

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Run a simple version of exec.Command + cmd.CombinedOutput
func Run(commands ...string) (output []byte, err error) {
	cmd := exec.Command(commands[0], commands[1:]...)
	return cmd.CombinedOutput()
}

// ChocoRun run chocolatey command with full choco path
func ChocoRun(commands ...string) (output []byte, err error) {
	chocoDir := "C:\\ProgramData\\chocolatey\\choco.exe"
	fullCommand := append([]string{chocoDir}, commands...)
	return Run(fullCommand...)
}

// PSRun run a command inside powershell
func PSRun(command string) (output []byte, err error) {
	return Run(
		"C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe",
		"-NoProfile",
		"-InputFormat", "None",
		"-ExecutionPolicy", "Bypass",
		"-Command", command)
}

// PSRunBypass bypassess the security that may block commands
func PSRunBypass(command string) (output []byte, err error) {

	filename, err := filepath.Abs("commands.ps1")
	if err != nil {
		return nil, err
	}

	toRun := "Start-Process -WindowStyle Minimized -Wait -FilePath Powershell -ArgumentList " +
		"\"-NoProfile\", \"-InputFormat\", \"None\", \"-ExecutionPolicy\", \"Bypass\"," +
		"\"-Command\", \"" + strings.Replace(command, "$", "`$", -1) + "\""

	// write the command to a powershell file
	err = ioutil.WriteFile(filename, []byte(toRun), 0777)
	if err != nil {
		return nil, err
	}

	// execute the powershell command
	out, err := Run("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe",
		"-NoProfile",
		"-InputFormat", "None",
		"-ExecutionPolicy", "Bypass",
		"-file", filename)

	// remove the powershell file
	os.Remove(filename)

	return out, err
}
