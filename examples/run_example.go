package main

import (
	"fmt"
	"github.com/sky-uk/go-brocade-vtm"
	"github.com/sky-uk/go-brocade-vtm/api/monitor"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("syntax error\nUsages: %s [https://load_balancer_address] [username] [password] \n\n", os.Args[0])
		os.Exit(1)
	}

	vtmAddress := os.Args[1]
	vtmUser := os.Args[2]
	vtmPassword := os.Args[3]
	debug := false

	if len(os.Args) == 5 && os.Args[4] == "true" {
		debug = true
	}

	vtmClient := go_brocade_vtm.NewVTMClient(vtmAddress, vtmUser, vtmPassword, true, debug)

	//
	// Get All Services.
	//
	// Create api object.
	getAllAPI := monitor.NewGetAll()

	// make api call.
	err := vtmClient.Do(getAllAPI)

	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// check the status code and proceed accordingly.
	if getAllAPI.StatusCode() == 200 {
		AllMonitors := getAllAPI.GetResponse().Children
		for _, service := range AllMonitors {
			fmt.Printf("Name: %-20s HRef: %-20s\n", service.Name, service.HRef)
		}
	} else {
		fmt.Println("Status code:", getAllAPI.StatusCode())
		fmt.Println("Response: ", getAllAPI.ResponseObject())
	}

}
