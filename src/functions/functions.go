package funs

// this file contains contains a lot of random functions

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

// used when the program needs to stop because of an error
func Die() {
	fmt.Println("press any key (besides the space bar..) to exit the application")
	// Prevent the application from closing
	fmt.Scanln()
	os.Exit(0)
}

// Check if a array contains value
func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func checkIfAdmin() bool {
	// check if the program is going to run width administative rights
	// skip this part for now, I don't know yet how to check for administrator rights
	return true
}

func CheckSYS() {
	// check if your system is supported
	status := true
	args := os.Args[1:]
	skip := Contains(args, "--skipChecks") || Contains(args, "-s")
	if runtime.GOOS != "windows" {
		fmt.Println("This applications appears to be running on a non-Windows system")
		status = false
	}
	if !checkIfAdmin() {
		fmt.Println("This application requires administative rights to run")
	}
	if !status && !skip {
		fmt.Println("Use --skipChecks or -s to skip checks")
		Die()
	}
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
