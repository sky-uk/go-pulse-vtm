package main

import (
	"fmt"
	"github.com/sky-uk/go-brocade-vtm"
	"github.com/sky-uk/go-brocade-vtm/api/monitor"
)

// RunMonitorExample : run the vTM monitor example
func RunMonitorExample(vtmAddress, vtmUser, vtmPassword string, debug bool) {

	vtmClient := brocadevtm.NewVTMClient(vtmAddress, vtmUser, vtmPassword, true, debug)

	//
	// Get All Monitors. Create api object.
	//
	//
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

	//
	// Create a new monitor
	//
	fmt.Println("== Running create new monitor with name 'PaaS_Test_Monitor' ==")

	var newMonitorName = "PaaSExampleHTTPMonitor"
	newHTTPMonitor := monitor.HTTP{URIPath: "/download/private/status/check"}
	newBasicMonitor := monitor.Basic{Delay: 6, Failures: 3, Type: "http", Timeout: 4}
	newMonitorProperties := monitor.Properties{Basic: newBasicMonitor, HTTP: newHTTPMonitor}
	newMonitor := monitor.Monitor{Properties: newMonitorProperties}

	createMonitorAPI := monitor.NewCreate(newMonitorName, newMonitor)
	err = vtmClient.Do(createMonitorAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if createMonitorAPI.StatusCode() == 201 {
		fmt.Printf("Monitor %s successfully created.\n", newMonitorName)
	} else {
		fmt.Printf("Failed to create new monitor %s.\n", newMonitorName)
	}
	fmt.Println(createMonitorAPI.GetResponse())

	//
	// Read a single monitor
	//
	fmt.Println("\n\n== Reading new monitor with name 'PaaS_Test_Monitor' ==")

	err = vtmClient.Do(getAllAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// check the status code and search for example monitor
	if getAllAPI.StatusCode() == 200 {
		foundMonitor := getAllAPI.GetResponse().FilterByName("PaaSExampleHTTPMonitor")
		fmt.Printf("Found monitor:\n \tName: %-20s Href: %-20s\n", foundMonitor.Name, foundMonitor.HRef)

		getSingleMonitorAPI := monitor.NewGetSingleMonitor(foundMonitor.Name)

		err = vtmClient.Do(getSingleMonitorAPI)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Printf("Retrieved monitor values are:\n")
		fmt.Printf("\tHTTP->URIPath: %s\n", getSingleMonitorAPI.GetResponse().Properties.HTTP.URIPath)
		fmt.Printf("\tBasic->Delay: %d\n", getSingleMonitorAPI.GetResponse().Properties.Basic.Delay)
		fmt.Printf("\tBasic->Failures: %d\n", getSingleMonitorAPI.GetResponse().Properties.Basic.Failures)
		fmt.Printf("\tBasic->Type: %s\n", getSingleMonitorAPI.GetResponse().Properties.Basic.Type)
		fmt.Printf("\tBasic->Timeout: %d\n", getSingleMonitorAPI.GetResponse().Properties.Basic.Timeout)

	} else {
		fmt.Println("Status code:", getAllAPI.StatusCode())
		fmt.Println("Response: ", getAllAPI.ResponseObject())
	}

}
