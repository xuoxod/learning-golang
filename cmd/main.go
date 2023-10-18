package main

import (
	"flag"
	"fmt"

	"github.com/xuoxod/mylibs/libs"
	"github.com/xuoxod/mylibs/utils"
)

func main() {
	filepath := flag.String("path", "./test.txt", "Expects a file path string")
	flag.Parse()

	fileInfo, exists := libs.PathExists(*filepath)

	if exists {
		fmt.Printf("File %s exists? %t\n\nFile Details\n", *filepath, exists)
		fmt.Println("Name:\t", fileInfo["name"])
		fmt.Println("Modified:\t", fileInfo["modified"])
		fmt.Println("Mode:\t", fileInfo["mode"])
		fmt.Println("Size:\t", fileInfo["size"])
		fmt.Println("Source:\t", fileInfo["source"])
		fmt.Println("Is Dir:\t", fileInfo["isdir"])

		list := libs.ReadFileToList(*filepath)

		utils.ReadStringListToConsole(list)

	} else {
		fmt.Printf("File path [%s] does not exists\n", *filepath)
	}
}
