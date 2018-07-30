package fs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dennis1248/Automated-Windows-10-configuration/src/options"
	"github.com/dennis1248/Automated-Windows-10-configuration/src/types"
)

func FindPackageJson(toCheck []string) (string, error) {
	// this function returns the location of a found package file
	// currently this only checks in this directory and one above
	fmt.Println("Searching for config file")
	toReturn := ""
	for _, check := range toCheck {
		fullPath, _ := filepath.Abs(check)
		_, err := os.Stat(fullPath)
		if err == nil {
			toReturn = fullPath
		}
	}
	if toReturn == "" {
		return toReturn, errors.New("No " + options.GetOptions().PackageName + " found in this directory")
	}
	return toReturn, nil
}

func OpenPackageJson(packageJsonFile string) (out types.Config, err error) {
	// returns the output of the config file
	fmt.Println("Using this config file:", packageJsonFile)
	fileContent, err := ioutil.ReadFile(packageJsonFile)
	if err != nil {
		return types.Config{}, err
	}
	var data types.Config
	json.Unmarshal([]byte(fileContent), &data)
	return data, nil
}
