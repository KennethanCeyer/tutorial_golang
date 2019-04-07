package main

import (
	"os"
)

func openFile(filePath string) {
	_, err := os.OpenFile(filePath, os.O_RDONLY, os.FileMode(0644))

	if err != nil {
		panic(err)
	}
}

func main() {
	filePath := "noFile"
	openFile(filePath)
}
