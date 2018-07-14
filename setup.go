package main

import (
	"fmt"
	"runtime"
	"os"
)

func contains(arr []string, str string) bool {
	for _, a := range arr {
		 if a == str {
				return true
		 }
	}
	return false
}

func checkSYS() {
	// check youre system if everything is supported 
	status := true
	args := os.Args[1:]
	skip := contains(args, "--skipChecks") || contains(args, "-s")
	if runtime.GOOS != "windows" {
		fmt.Println("This script can only be ran on a windows PC")
		status = false
	}
	if !status && !skip {
		fmt.Println("Use --skipChecks or -s to skip the checks")
		os.Exit(0)
	}
}

func main() {
	checkSYS()
	fmt.Println("Starting setup...")
	
	// do stuff
	fmt.Println("Dune!")

	// force the program to run forever
	fmt.Scanln()
}