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
			fmt.Printf("Name: %-20s Href: %-20s\n", virtualserver.Name, virtualserver.Href)
		}
	} else {
		fmt.Println("Status code:", getAllAPI.StatusCode())
		fmt.Println("Response: ", getAllAPI.ResponseObject())
	}

	fmt.Println("== Running Create new virtualserver with name 'PaaS_Test_virtualserver' ==")

	var newvirtualserverName = "PaaSExampleHTTPvirtualserver"

	//
	// Delete a virtual server
	// ------------------  Deleting first the resource...
	log.Print("Trying first to delete virtual server with name: ", newvirtualserverName)
	deleteAPI := virtualserver.NewDelete(newvirtualserverName)
	err = vtmClient.Do(deleteAPI)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Resource successfully deleted")
	}

	//
	// Create a virtual server
	// ------------------ then creating a new one...
	newBasicvirtualserver := virtualserver.Basic{
		Enabled:  false,
		Pool:     "pool_test_rui",
		Port:     80,
		Protocol: "http",
	}
	newvirtualserverProperties := virtualserver.Properties{Basic: newBasicvirtualserver}
	newvirtualserver := virtualserver.VirtualServer{Properties: newvirtualserverProperties}

	// trying to encode to json... -----------------------------------------
	jsonStr, e := json.Marshal(newvirtualserver)
	if e != nil {
		fmt.Println("Error encoding structure to json: ", e)
	} else {
		//fmt.Printf("New virtual server json: \n%v", jsonStr )
		fmt.Printf("New Virtual Server: ------- %s --------\n", newvirtualserverName)
		os.Stdout.Write(jsonStr)
		fmt.Println()
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
		fmt.Printf("Status Code:%d\n", createvirtualserverAPI.StatusCode())
		fmt.Printf("Error: %+v\n", createvirtualserverAPI.Error())
	}

	//
	// Get a single virtual server
	//
	fmt.Printf("\n== Reading the Virtual Server with name %s ==\n", newvirtualserverName)
	getSingleVirtualServerAPI := virtualserver.NewGetSingle(newvirtualserverName)
	err = vtmClient.Do(getSingleVirtualServerAPI)
	if err != nil {
		fmt.Println("Error getting single virtual server: ", err)
	}

	// check status code and what we got back...
	if getSingleVirtualServerAPI.StatusCode() == 200 {
		var foo *virtualserver.VirtualServer
		foo = getSingleVirtualServerAPI.GetResponse()
		fmt.Printf("virtual server pool: %s\n", foo.Properties.Basic.Pool)
	} else {
		fmt.Printf("Error: Status code: %+v\n", getSingleVirtualServerAPI.StatusCode())
		fmt.Printf("Error: Response: +%v\n", getSingleVirtualServerAPI.ResponseObject())
	}

	//
	// Update a virtual server...
	fmt.Printf("\n== Updating a Virtual Server ==\n")
	var updatevirtualserverName = "PaaSExampleHTTPvirtualserver"

	updateBasicvirtualserver := virtualserver.Basic{
		Enabled:  false,
		Pool:     "pool_test_rui",
		Port:     90,
		Protocol: "http",
	}
	updatevirtualserverProperties := virtualserver.Properties{Basic: updateBasicvirtualserver}
	updatevirtualserver := virtualserver.VirtualServer{Properties: updatevirtualserverProperties}

	updateVirtualServerAPI := virtualserver.NewUpdate(
		updatevirtualserverName,
		updatevirtualserver,
	)
	err = vtmClient.Do(updateVirtualServerAPI)
	if err != nil {
		fmt.Println("Error updating virtual server: ", err)
	} else {
		fmt.Println("Virtual server updated succesfully")
	}

}
