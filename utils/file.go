package utils

import (
	"os"
	"path/filepath"
)

// Exist file exist
func Exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}

		if os.IsNotExist(err) {
			return false
		}

		return false
	}

	return true
}

// CurDir current command dir
func CurDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}

// MakeSureDirExist make sure directory exists
func MakeSureDirExist(path string) {
	dir := filepath.Dir(path)
	if !Exist(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}
}
