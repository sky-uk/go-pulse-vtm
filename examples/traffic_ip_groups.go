package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/sky-uk/go-brocade-vtm"
	"github.com/sky-uk/go-brocade-vtm/api/traffic_ip_group"
)

// RunTrafficIPGroupsExample : run traffic ip group example
func RunTrafficIPGroupsExample(vtmAddress, vtmUser, vtmPassword string, debug bool) {
	vtmClient := brocadevtm.NewVTMClient(vtmAddress, vtmUser, vtmPassword, true, debug)

	//
	// Get All Services.
	//
	// Create api object.

	getAllAPI := trafficIpGroups.NewGetAll()

	// make api call.
	err := vtmClient.Do(getAllAPI)

	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// check the status code and proceed accordingly.
	if getAllAPI.StatusCode() == 200 {
		AllTrafficIPGroups := getAllAPI.GetResponse().Children
		for _, trafficIPGroup := range AllTrafficIPGroups {
			fmt.Printf("Name: %-20s HRef: %-20s\n", trafficIPGroup.Name, trafficIPGroup.HRef)
		}
	} else {
		fmt.Println("Status code:", getAllAPI.StatusCode())
		fmt.Println("Response: ", getAllAPI.ResponseObject())
	}

	tipg := "craig-test-group-1"
	getSingleAPI := trafficIpGroups.NewGetSingle(tipg)

	err = vtmClient.Do(getSingleAPI)

	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if getSingleAPI.StatusCode() == 200 {
		singleTrafficGroup := getSingleAPI.GetResponse().Properties.Basic
		spew.Dump(singleTrafficGroup)
	} else {
		fmt.Println("Status code:", getSingleAPI.StatusCode())
		fmt.Println("Response: ", getSingleAPI.ResponseObject())
	}

	deleteTrafficIPGroupAPI := trafficIpGroups.NewDelete(tipg)

	err = vtmClient.Do(deleteTrafficIPGroupAPI)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	if deleteTrafficIPGroupAPI.StatusCode() == 204 {
		fmt.Println("Succesfully deleted: ", tipg)
	} else {
		fmt.Println("Status code:", deleteTrafficIPGroupAPI.StatusCode())
		fmt.Println("Response: ", deleteTrafficIPGroupAPI.ResponseObject())
	}
}
