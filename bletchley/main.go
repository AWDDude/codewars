package main

import (
	"errors"
	"flag"
	"log"
	"os"
)

func main() {
	rleFilePath := flag.String("path", "attachments/PC-POW-corrected.txt", "path to the input Run Length Encoded file")
	flag.Parse()

	rawRLE := getFileContents(*rleFilePath)
}

func getFileContents(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(errors.Join())
	}
}
