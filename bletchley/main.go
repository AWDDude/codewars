package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	rleFilePath := flag.String("path", "attachments/PC-POW-corrected.txt", "path to the input Run Length Encoded file")
	flag.Parse()

	rawRLE := getFileContents(*rleFilePath)
	fmt.Print(rawRLE)
}

func getFileContents(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
