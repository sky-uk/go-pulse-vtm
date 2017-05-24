package main

import (
	"github.com/sky-uk/go-brocade-vtm"
	"github.com/sky-uk/go-brocade-vtm/api/monitor"
	"fmt"
)

func RunMonitorExample(vtmAddress, vtmUser, vtmPassword string, debug bool) {

	vtmClient := brocadevtm.NewVTMClient(vtmAddress, vtmUser, vtmPassword, true, debug)

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
		for _, monitor := range AllMonitors {
			fmt.Printf("Name: %-20s HRef: %-20s\n", monitor.Name, monitor.HRef)
		}
	} else {
		fmt.Println("Status code:", getAllAPI.StatusCode())
		fmt.Println("Response: ", getAllAPI.ResponseObject())
	}

	fmt.Println("== Running Create new Monitor with name 'PaaS_Test_Monitor' ==")

	var newMonitorName string = "PaaSExampleHTTPMonitor"
	newHTTPMonitor := monitor.MonitorHTTP{URIPath: "/download/private/status/check"}
	newBasicMonitor := monitor.MonitorBasic{Delay: 6, Failures: 3, Type: "http", Timeout: 4}
	newMonitorProperties := monitor.MonitorProperties{Basic: newBasicMonitor, Http: newHTTPMonitor}
	newMonitor := monitor.Monitor{Properties: newMonitorProperties}

	createMonitorAPI := monitor.NewCreate(newMonitorName, newMonitor)
	err = vtmClient.Do(createMonitorAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if createMonitorAPI.StatusCode() == 201 {
		fmt.Printf("Monitor %s successfully created.\n", newMonitorName)
	} else {
		fmt.Println("Failed to create new monitor %s", newMonitorName)
	}
	fmt.Println(createMonitorAPI.GetResponse())
}
