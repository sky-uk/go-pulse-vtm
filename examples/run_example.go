package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Printf("syntax error\nUsages: %s [https://load_balancer_address] [username] [password] \n\n", os.Args[0])
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

	switch exampleName {
	case "monitor":
		fmt.Println("Running monitor with: ", vtmAddress, vtmUser, vtmPassword, exampleName, debug)
		RunMonitorExample(vtmAddress, vtmUser, vtmPassword, debug)
		return
    case "virtual_server":
        fmt.Println("Running virtual server with: ", vtmAddress, vtmUser, vtmPassword, exampleName, debug)
        RunVirtualServerExample(vtmAddress, vtmUser, vtmPassword, debug)
        return
    }
	fmt.Println("Example not implemented.")
}
