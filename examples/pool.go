package main

import (
	"fmt"
	"github.com/sky-uk/go-brocade-vtm"
	"github.com/sky-uk/go-brocade-vtm/api/pool"
)

func RunPoolExample(vtmAddress, vtmUser, vtmPassword string, debug bool) {
	vtmClient := brocadevtm.NewVTMClient(vtmAddress, vtmUser, vtmPassword, true, debug)

	//
	// Get All Services.
	//
	// Create api object.
	getAllAPI := pool.NewGetAll()

	// make api call.
	err := vtmClient.Do(getAllAPI)

	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// check the status code and proceed accordingly.
	if getAllAPI.StatusCode() == 200 {
		AllPools := getAllAPI.GetResponse().ChildPools
		for _, pool := range AllPools {
			fmt.Printf("Name: %-20s HRef: %-20s\n", pool.Name, pool.Href)
		}
	} else {
		fmt.Println("Status code:", getAllAPI.StatusCode())
		fmt.Println("Response: ", getAllAPI.ResponseObject())
	}

}