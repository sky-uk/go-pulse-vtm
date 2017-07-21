package main

import (
	"fmt"
	"github.com/sky-uk/go-brocade-vtm"
	"github.com/sky-uk/go-brocade-vtm/api/monitor"
)

// RunMonitorExample : run the vTM monitor example
func RunMonitorExample(vtmAddress, vtmUser, vtmPassword string, debug bool) {

	vtmClient := brocadevtm.NewVTMClient(vtmAddress, vtmUser, vtmPassword, true, debug, nil)

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
	var exampleMonitorName = "PaaSExampleHTTPMonitor"
	fmt.Printf("\n== Running create new monitor with name %s ==\n", exampleMonitorName)

	newHTTPMonitor := monitor.HTTP{URIPath: "/download/private/status/check"}
	monitorVerbosity := true
	newBasicMonitor := monitor.Basic{Delay: 6, Failures: 3, Type: "http", Timeout: 4, Verbose: &monitorVerbosity}
	newMonitorProperties := monitor.Properties{Basic: newBasicMonitor, HTTP: newHTTPMonitor}
	newMonitor := monitor.Monitor{Properties: newMonitorProperties}

	createMonitorAPI := monitor.NewCreate(exampleMonitorName, newMonitor)
	err = vtmClient.Do(createMonitorAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if createMonitorAPI.StatusCode() == 201 {
		fmt.Printf("Monitor %s successfully created.\n", exampleMonitorName)
		fmt.Printf("\tMonitor GetResponse was: %+v\n", createMonitorAPI.GetResponse())
	} else {
		fmt.Printf("Failed to create new monitor %s.\n", exampleMonitorName)
	}

	//
	// Read a single monitor
	//
	fmt.Printf("\n== Reading new monitor with name %s ==\n", exampleMonitorName)

	err = vtmClient.Do(getAllAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// check the status code and search for example monitor
	if getAllAPI.StatusCode() == 200 {
		foundMonitor := getAllAPI.GetResponse().FilterByName(exampleMonitorName)
		fmt.Printf("Found monitor:\n \tName: %-20s Href: %-20s\n", foundMonitor.Name, foundMonitor.HRef)

		getSingleMonitorAPI := monitor.NewGet(foundMonitor.Name)

		err = vtmClient.Do(getSingleMonitorAPI)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Printf("\nRetrieved monitor values for %s are:\n", foundMonitor.Name)
		fmt.Printf("\tHTTP->URIPath: %s\n", getSingleMonitorAPI.GetResponse().Properties.HTTP.URIPath)
		fmt.Printf("\tBasic->Delay: %d\n", getSingleMonitorAPI.GetResponse().Properties.Basic.Delay)
		fmt.Printf("\tBasic->Failures: %d\n", getSingleMonitorAPI.GetResponse().Properties.Basic.Failures)
		fmt.Printf("\tBasic->Type: %s\n", getSingleMonitorAPI.GetResponse().Properties.Basic.Type)
		fmt.Printf("\tBasic->Timeout: %d\n", getSingleMonitorAPI.GetResponse().Properties.Basic.Timeout)

	} else {
		fmt.Printf("Status code: %+v\n", getAllAPI.StatusCode())
		fmt.Printf("Response: +%v\n", getAllAPI.ResponseObject())
	}

	//
	// Update a monitor
	//
	fmt.Printf("\n\n== Updating monitor with name %s to use /private/status/check ==\n", exampleMonitorName)

	var updateMonitor monitor.Monitor

	updateMonitor.Properties.HTTP.URIPath = "/private/status/check"
	monitorVerbosity = false
	updateMonitor.Properties.Basic.Verbose = &monitorVerbosity
	updateMonitorAPI := monitor.NewUpdate(exampleMonitorName, updateMonitor)
	err = vtmClient.Do(updateMonitorAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if updateMonitorAPI.StatusCode() == 200 {
		fmt.Printf("Successfully updated monitor %s to use /private/status/check\n", exampleMonitorName)
		getSingleMonitorAPI := monitor.NewGet(exampleMonitorName)
		err = vtmClient.Do(getSingleMonitorAPI)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Printf("New value for HTTP->URIPath is %s.\n", getSingleMonitorAPI.GetResponse().Properties.HTTP.URIPath)
	}

	//
	// Delete a monitor
	//
	fmt.Printf("\n== Deleting monitor with name %s ==\n", exampleMonitorName)

	deleteAPI := monitor.NewDelete(exampleMonitorName)
	err = vtmClient.Do(deleteAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if deleteAPI.StatusCode() == 204 {
		fmt.Printf("Monitor %s was successfully deleted\n", exampleMonitorName)
	} else {
		fmt.Printf("Monitor %s wasn't deleted\n", exampleMonitorName)
	}
}
