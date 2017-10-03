package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"os"
)

func getAllResourceTypes(client *api.Client, flagSet *flag.FlagSet) {

	// work with an environment
	client.WorkWithConfigurationResources()
	res, err := client.GetAllResourceTypes()
	if err != nil {
		fmt.Println("Error getting all configuration resources: ", err)
		os.Exit(2)
	}
	PrettyPrintMany([]string{"name", "href"}, res)
}

func init() {
	RegisterCliCommand("get-all-resource-types", nil, getAllResourceTypes)
}
