package dwpath

import (
	"fmt"
	"io/fs"
	"os"
)

func IsPathExists(path string) (bool, fs.FileInfo, error) {
	fileinfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil, nil
		} else {
			return false, nil, err
		}
	}
	return true, fileinfo, nil
}

func IsDirectoryExists(dir string) (bool, error) {
	exists, fileinfo, err := IsPathExists(dir)
	if err != nil {
		return false, err
	}

	if !exists {
		return false, nil
	}

	if !fileinfo.IsDir() {
		return false, fmt.Errorf("path '%s' bukan direktori", dir)
	}

	return true, nil
}

func IsFileExists(path string) (bool, fs.FileInfo, error) {
	exists, fileinfo, err := IsPathExists(path)
	if err != nil {
		return false, nil, err
	}

	if !exists {
		return false, nil, nil
	}

	if fileinfo.IsDir() {
		return false, fileinfo, fmt.Errorf("path '%s' bukan file", path)
	}

	return true, fileinfo, nil
}
