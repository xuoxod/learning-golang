package libs

import (
	"fmt"
	"io"
	"os"
)

func ReadFileToConsole(filepath string) {
	var mode os.FileMode = 0700
	// file, err := os.Open(filepath)
	file, err := os.OpenFile(filepath, os.O_RDONLY, mode)

	if err != nil {
		fmt.Println("Error opening file:\t", err.Error())
	}

	defer file.Close()

	const maxSz = 4096

	// make a buffer
	buf := make([]byte, maxSz)

	for {
		readTotal, err := file.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
			}
			break
		}
		fmt.Println(string(buf[:readTotal]))
	}

}

func ReadFileToList(filepath string) []string {
	var mode os.FileMode = 0700
	// file, err := os.Open(filepath)
	file, err := os.OpenFile(filepath, os.O_RDONLY, mode)

	if err != nil {
		fmt.Println("Error opening file:\t", err.Error())
	}

	defer file.Close()

	content := []string{}

	const maxSz = 4096

	// make a buffer
	buf := make([]byte, maxSz)

	for {
		readTotal, err := file.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
			}
			break
		}
		content = append(content, string(buf[:readTotal]))
	}
	return content
}
