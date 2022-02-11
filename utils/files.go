package utils

import (
	"os"
)

func DirExists(dirName string) bool {
	_, folderErr := os.Stat(dirName)
	if os.IsNotExist(folderErr) {
		return false
	}
	return true
}

func FileExists(fileName string) bool {
	_, fileErr := os.Open(fileName)
	if os.IsNotExist(fileErr) {
		return false
	}
	return true
}
