package dwpath

import (
	"fmt"
	"io/fs"
	"os"
)

func IsPathExists(path string) (bool, fs.FileInfo, error) {
	fileinfo, err := os.Stat(path)
	if err == nil {
		return true, fileinfo, err
	}
	if os.IsNotExist(err) {
		return false, nil, fmt.Errorf("path '%s' tidak ditemukan", path)
	}
	return false, nil, err
}

func IsDirectoryExists(dir string) (bool, error) {
	exists, fileinfo, err := IsPathExists(dir)
	if !exists {
		if err != nil {
			return false, err
		} else {
			return false, nil
		}
	} else {
		if !fileinfo.IsDir() {
			return false, fmt.Errorf("path '%s' bukan direktori", dir)
		}
	}
	return true, nil
}

func IsFileExists(path string) (bool, fs.FileInfo, error) {
	exists, fileinfo, err := IsPathExists(path)
	if err != nil {
		return false, fileinfo, err
	}

	if fileinfo.IsDir() {
		return false, fileinfo, fmt.Errorf("path '%s' bukan file", path)
	}

	if !exists {
		return false, nil, fmt.Errorf("file '%s' tidak ditemukan", path)
	}

	return true, fileinfo, nil
}
