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
	tipgName := "cdu16-test-group"

	fmt.Println("\n--- Create Example ---")

	tipgIPAddresses := []string{"172.0.0.1"}
	tipgBasic := trafficIpGroups.Basic{IPAddresses: tipgIPAddresses, Location: 0, Mode: "rhi", Note: "", RhiOspfv2MetricBase: 10, RhiOspfv2PassiveMetricOffset: 10}
	tipgProperties := trafficIpGroups.Properties{tipgBasic}
	tipgTrafficIPGroup := trafficIpGroups.TrafficIPGroup{tipgProperties}

	createTrafficIPGroupAPI := trafficIpGroups.NewCreate(tipgName, tipgTrafficIPGroup)

	err := vtmClient.Do(createTrafficIPGroupAPI)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	if createTrafficIPGroupAPI.StatusCode() == 201 {

		fmt.Println("Succesfully created: ", tipgName)
		fmt.Println(createTrafficIPGroupAPI.StatusCode())
		spew.Dump(createTrafficIPGroupAPI.ResponseObject())
	} else {
		fmt.Println("Status code: ", createTrafficIPGroupAPI.StatusCode())
		fmt.Println("Response: ", createTrafficIPGroupAPI.ResponseObject())
	}

	fmt.Println("\n--- Update Example ---")

	keeptogether := true
	tipgUpdateBasic := trafficIpGroups.Basic{KeepTogether: &keeptogether}
	tipgUpdateProperties := trafficIpGroups.Properties{tipgUpdateBasic}
	tipgUpdateTrafficIPGroup := trafficIpGroups.TrafficIPGroup{tipgUpdateProperties}

	updateTrafficIPGroupAPI := trafficIpGroups.NewUpdate(tipgName, tipgUpdateTrafficIPGroup)

	err = vtmClient.Do(updateTrafficIPGroupAPI)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	if updateTrafficIPGroupAPI.StatusCode() == 200 {
		fmt.Println("Succesfully updated: ", tipgName)
		spew.Dump(updateTrafficIPGroupAPI.ResponseObject())
	} else {
		fmt.Println("Status code: ", updateTrafficIPGroupAPI.StatusCode())
		fmt.Println("Response: ", updateTrafficIPGroupAPI.ResponseObject())
	}

	fmt.Println("\n--- Get All Example ---")

	getAllAPI := trafficIpGroups.NewGetAll()

	// make api call.
	err = vtmClient.Do(getAllAPI)

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

	fmt.Println("\n--- Get Single Example ---")

	getSingleAPI := trafficIpGroups.NewGetSingle(tipgName)

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

	fmt.Println("\n--- Delete Example ---")

	deleteTrafficIPGroupAPI := trafficIpGroups.NewDelete(tipgName)

	err = vtmClient.Do(deleteTrafficIPGroupAPI)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	if deleteTrafficIPGroupAPI.StatusCode() == 204 {
		fmt.Println("Succesfully deleted: ", tipgName)
	} else {
		fmt.Println("Status code:", deleteTrafficIPGroupAPI.StatusCode())
		fmt.Println("Response: ", deleteTrafficIPGroupAPI.ResponseObject())
	}
}
