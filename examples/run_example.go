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
		fmt.Println("running monitor with: ", vtmAddress, vtmUser, vtmPassword, exampleName, debug)
		RunMonitorExample(vtmAddress, vtmUser, vtmPassword, debug)
		return
	case "pool":
		fmt.Println("running pool with:", vtmAddress, vtmUser, vtmPassword, exampleName, debug)
		RunPoolExample(vtmAddress, vtmUser, vtmPassword, debug)
	case "traffic_ip_groups":
		fmt.Println("running traffic_ip_groups with: ", vtmAddress, vtmUser, vtmPassword, exampleName, debug)
		RunTrafficIPGroupsExample(vtmAddress, vtmUser, vtmPassword, debug)
		return
	}

	fmt.Println("Example not implemented.")
}
