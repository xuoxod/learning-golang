package libs

import (
	"os"
)

func PathExists(filepath string) (map[string]interface{}, bool) {
	fi, err := os.Stat(filepath)

	if err != nil {
		return nil, false
	}

	fileInfo := make(map[string]interface{})
	fileInfo["name"] = fi.Name()
	fileInfo["modified"] = fi.ModTime()
	fileInfo["mode"] = fi.Mode()
	fileInfo["size"] = fi.Size()
	fileInfo["source"] = fi.Sys()
	fileInfo["isdir"] = fi.IsDir()

	return fileInfo, true
}

func PathIsFile(filepath string) bool {
	fi, err := os.Stat(filepath)

	if err != nil {
		return false
	}

	return !fi.IsDir()
}

func PathIsDir(filepath string) bool {
	fi, err := os.Stat(filepath)

	if err != nil {
		return false
	}

	return fi.IsDir()
}
