package assets

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/CrowBits/gobuild/packages/configuration"
)

// CheckForChange will return true if there are chagnes the config dose not know about
func CheckForChange(confList []configuration.AssetsFilesStruct) bool {
	currList := getCurrentList("./assets")

	// Check Len
	if len(currList) != len(confList) {
		return true
	}

	for _, item := range confList {

		// Check the path is in current
		if _, err := currList[item.Path]; !err {
			fmt.Println("NOTFOUND")
			return true
		}

		if item.Md5 != currList[item.Path] {
			return true
		}

		delete(currList, item.Path)

	}

	if len(currList) > 0 {
		return true
	}

	return false
}

// BuildAssets will Build a new Assets File
func BuildAssets() error {
	_, err := exec.Command("go-bindata", "-o", "./assets/assets.go", "-prefix", "./assets/", "./assets/...").Output()
	if err != nil {
		return err
	}
	return fixAssetsFilePackage()
}

//GetFilesForConfig will return an updated file list for the config
func GetFilesForConfig() (list []configuration.AssetsFilesStruct) {
	currList := getCurrentList("./assets")
	var tmpFile configuration.AssetsFilesStruct

	for path, md5 := range currList {
		tmpFile.Path = path
		tmpFile.Md5 = md5
		list = append(list, tmpFile)
	}
	return
}

func getCurrentList(pathAssets string) map[string]string {
	fList := make(map[string]string)
	// Walk the current Dir and build list
	filepath.Walk(pathAssets, func(path string, fi os.FileInfo, err error) (e error) {
		if !fi.IsDir() {
			// Read in file
			fileBytes, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			// Gen MD5 Sum
			md5sum := fmt.Sprintf("%x", md5.Sum(fileBytes))

			// Add to current List
			fList[path] = md5sum
		}
		return nil
	})
	return fList
}

func fixAssetsFilePackage() error {
	input, err := ioutil.ReadFile("./assets/assets.go")
	if err != nil {
		return err
	}
	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.HasPrefix(line, "package main") {
			lines[i] = "package assets"
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("./assets/assets.go", []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}
