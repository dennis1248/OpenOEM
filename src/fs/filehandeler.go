package fs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"

	"github.com/dennis1248/Automated-Windows-10-configuration/src/functions"
	"github.com/dennis1248/Automated-Windows-10-configuration/src/options"
	"github.com/dennis1248/Automated-Windows-10-configuration/src/types"
)

var printedUsingConfigFile bool = false

func Copy(src, dst string) error {
	// copy file from source to destination
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func CheckDataFolder() {
	// check if the data folde already exsisted
	path := "C:\\ProgramData\\automated-Windows-10-configuration"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0777)
	}
}

func FindPackageJson(toCheck []string) (string, error) {
	// this function returns the location of a found package file
	// currently this only checks in this directory and one above
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
	if !printedUsingConfigFile {
		printedUsingConfigFile = true
		fmt.Println("Using this config file:", packageJsonFile)
	}
	fileContent, err := ioutil.ReadFile(packageJsonFile)
	if err != nil {
		return types.Config{}, err
	}
	var data types.Config
	json.Unmarshal([]byte(fileContent), &data)
	return data, nil
}

func FindAndOpenPackageJson() (out types.Config, err error) {
	PackageName := options.GetOptions().PackageName
	packageJsonFile, err := FindPackageJson([]string{"./" + PackageName, "./../" + PackageName})
	if err != nil {
		return types.Config{}, err
	}
	return OpenPackageJson(packageJsonFile)
}

func GetWallpaper(Package types.Config) string {

	// Copied get old wallpaper from https://github.com/reujab/wallpaper/blob/master/windows.go#L31
	var oldWallpaperPointer [256]uint16
	user32 := syscall.NewLazyDLL("user32.dll")
	systemParametersInfo := user32.NewProc("SystemParametersInfoW")
	systemParametersInfo.Call(
		uintptr(0x0073), uintptr(cap(oldWallpaperPointer)), uintptr(unsafe.Pointer(&oldWallpaperPointer[0])), uintptr(0),
	)
	oldWallpaper := strings.Trim(string(utf16.Decode(oldWallpaperPointer[:])), "\x00")

	wallpaper := Package.Wallpaper
	dst := "C:\\ProgramData\\automated-Windows-10-configuration\\wallpaper" + path.Ext(wallpaper)

	if len(wallpaper) == 0 {
		// user doesn't want wallpaper return old wallpaper
		return oldWallpaper
	}

	if _, err := os.Stat(wallpaper); err == nil {
		CheckDataFolder()
		src, err := filepath.Abs(wallpaper)
		if err != nil {
			src = wallpaper
		}
		err = Copy(src, dst)
		if err == nil {
			return dst
		}
	}

	err := funs.DownloadFile(dst, wallpaper)
	if err == nil {
		return dst
	}

	// can't set wallpaper return old wallpaper
	return oldWallpaper
}

func FinalCleanUp() {
	RemoveEdgeIcon()
	// in some cases the previous os.Remove doesn't work so just try it again
	os.Remove("installTheme.theme")
}

func RemoveEdgeIcon() error {
	// check if removing home icons is allowed by the config file
	Package, err := FindAndOpenPackageJson()
	if err != nil {
		return err
	}
	if !Package.R_EdigeIcon {
		return nil
	}

	// get the current user's home folder and remove the edge icoon
	user, err := user.Current()
	if err != nil {
		return err
	}
	tryFile1, _ := filepath.Abs(user.HomeDir + "\\Desktop\\Microsoft Edge.lnk")
	tryFile2, _ := filepath.Abs(user.HomeDir + "\\Desktop\\Microsoft Edge.lnk*")
	tryFile3, _ := filepath.Abs(user.HomeDir + "\\Desktop\\Microsoft Edge.lnk@")
	tryFile4, _ := filepath.Abs(user.HomeDir + "\\Desktop\\Microsoft Edge.lnk")
	os.RemoveAll(tryFile1)
	os.RemoveAll(tryFile2)
	os.RemoveAll(tryFile3)
	os.RemoveAll(tryFile4)
	return nil
}

func MakeFile(data string, filePath string) error {
	byteData := []byte(data)
	return ioutil.WriteFile(filePath, byteData, 0777)
}
