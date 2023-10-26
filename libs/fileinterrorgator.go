package libs

import (
	"log"
	"os"
	"path"
)

func PathExists(filepath string) (map[string]interface{}, bool) {
	fi, err := os.Stat(filepath)

	if err != nil {
		log.Printf("os.stat error:\t%s\n", err.Error())
		return nil, false
	}

	fileInfo := make(map[string]interface{})
	fileInfo["name"] = fi.Name()
	fileInfo["modified"] = fi.ModTime()
	fileInfo["mode"] = fi.Mode()
	fileInfo["size"] = fi.Size()
	fileInfo["source"] = fi.Sys()
	fileInfo["isdir"] = fi.IsDir()
	fileInfo["extension"] = path.Ext(filepath)
	fileInfo["base"] = path.Base(filepath)
	fileInfo["parent"] = path.Dir(filepath)
	fileInfo["isabs"] = path.IsAbs(filepath)

	return fileInfo, true
}

func PathIsFile(filepath string) bool {
	fi, err := os.Stat(filepath)

	if err != nil {
		log.Printf("os.stat error:\t%s\n", err.Error())
		return false
	}

	return !fi.IsDir()
}

func PathIsDir(filepath string) bool {
	fi, err := os.Stat(filepath)

	if err != nil {
		log.Printf("os.stat error:\t%s\n", err.Error())
		return false
	}

	return fi.IsDir()
}
