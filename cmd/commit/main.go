package main

import (
	"fmt"
	"os"

	"github.com/jwhittle933/profile/internal/fs"
)

func main() {
	readme, err := fs.Open("README.md")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}
	defer readme.Close()

	if err := readme.Flush(); err != nil {
		fmt.Printf("ERROR flushing: %s\n", err.Error())
	}
}
