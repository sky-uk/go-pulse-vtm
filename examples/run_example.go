package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Printf("syntax error\nUsages: %s [https://load_balancer_address] [username] [password] [example name]\n\n", os.Args[0])
		os.Exit(1)
	}

	vtmAddress := os.Args[1]
	vtmUser := os.Args[2]
	vtmPassword := os.Args[3]
	exampleName := os.Args[4]
	debug := false

	if len(os.Args) == 6 && os.Args[5] == "true" {
		debug = true
	}

	fmt.Printf("Running %s with VTM addr: %s, VTM user: %s, VTM password: %s, debug: %v\n", exampleName, vtmAddress, vtmUser, vtmPassword, debug)

	switch exampleName {
	case "monitor":
		RunMonitorExample(vtmAddress, vtmUser, vtmPassword, debug)
		return
	case "pool":
		RunPoolExample(vtmAddress, vtmUser, vtmPassword, debug)
		return
	case "virtual_server":
		RunVirtualServerExample(vtmAddress, vtmUser, vtmPassword, debug)
		return
	}
	fmt.Println("Example not implemented.")
}
