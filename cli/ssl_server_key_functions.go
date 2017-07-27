package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func retrieveSSLKeyFile(fileName string) string {
	if fileName != "" {
		fileContents, fileErr := ioutil.ReadFile(fileName)
		if fileErr != nil {
			fmt.Printf("\nError reading file %s\n", fileName)
			os.Exit(2)
		}
		return string(fileContents)
	}
	return ""
}
