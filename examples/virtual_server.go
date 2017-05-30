package main

import (
	"encoding/json"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm"
	"github.com/sky-uk/go-brocade-vtm/api/virtualserver"
	"log"
	"os"
)

// RunVirtualServerExample : run virtualserver example
func RunVirtualServerExample(vtmAddress, vtmUser, vtmPassword string, debug bool) {

	vtmClient := brocadevtm.NewVTMClient(vtmAddress, vtmUser, vtmPassword, true, debug)

	//
	// Get All Services.
	//
	// Create api object.
	getAllAPI := virtualserver.NewGetAll()

	fmt.Println("Get all virtual servers")
	fmt.Println("-------------------------------------------------------------------------")
	// make api call.
	err := vtmClient.Do(getAllAPI)

	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// check the status code and proceed accordingly.
	if getAllAPI.StatusCode() == 200 {
		Allvirtualservers := getAllAPI.GetResponse().Children
		for _, virtualserver := range Allvirtualservers {
			fmt.Printf("Name: %-20s HRef: %-20s\n", virtualserver.Name, virtualserver.HRef)
		}
	} else {
		fmt.Println("Status code:", getAllAPI.StatusCode())
		fmt.Println("Response: ", getAllAPI.ResponseObject())
	}

	fmt.Println("== Running Create new virtualserver with name 'PaaS_Test_virtualserver' ==")

	var newvirtualserverName = "PaaSExampleHTTPvirtualserver"

	// ------------------  Deleting first the resource...
	log.Print("Trying first to delete virtual server with name: ", newvirtualserverName)
	deleteAPI := virtualserver.NewDelete(newvirtualserverName)
	err = vtmClient.Do(deleteAPI)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Resource successfully deleted")
	}

	// ------------------ then creating a new one...
	newBasicvirtualserver := virtualserver.Basic{
		Enabled:            false,
		DefaultTrafficPool: "pool_test_rui",
		Port:               80,
		Protocol:           "http",
	}
	newvirtualserverProperties := virtualserver.Properties{Basic: newBasicvirtualserver}
	newvirtualserver := virtualserver.VirtualServer{Properties: newvirtualserverProperties}

	// trying to encode to json... -----------------------------------------
	json_str, e := json.Marshal(newvirtualserver)
	if e != nil {
		fmt.Println("Error encoding structure to json: ", e)
	} else {
		//fmt.Printf("New virtual server json: \n%v", json_str )
		fmt.Println("New Virtual Server:")
		os.Stdout.Write(json_str)
		fmt.Println("\n")
	}
	//-----------------------------------------------------------------------

	createvirtualserverAPI := virtualserver.NewCreate(newvirtualserverName, newvirtualserver)
	err = vtmClient.Do(createvirtualserverAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if createvirtualserverAPI.StatusCode() == 201 {
		fmt.Printf("virtualserver %s successfully created.\n", newvirtualserverName)
		if debug {
			fmt.Println(createvirtualserverAPI.GetResponse())
		}
	} else {
		fmt.Printf("Failed to create new virtualserver %s.\n", newvirtualserverName)
	}
}
