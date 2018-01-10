package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"os"
	"regexp"
)

func deleteAccTestResources(client *api.Client, flagSet *flag.FlagSet) {
	re := regexp.MustCompile("acctest_")
	resourceTypeMap, err := client.GetAllResourceTypes()

	if err != nil {
		fmt.Println("Error getting all configuration resources: ", err)
		os.Exit(1)
	}
	for _, v := range resourceTypeMap {
		resourceType := v["name"].(string)
		allResourcesOfType, err := client.GetAllResources(resourceType)
		if err != nil {
			fmt.Println("Error getting configuration resource: ", err)
			os.Exit(1)
		}
		for _, resource := range allResourcesOfType {
			resourceName := resource["name"].(string)
			if re.MatchString(resourceName) {
				err := client.Delete(resourceType, resource["name"].(string))
				if err != nil {
					fmt.Println("Error deleting configuration resource: ", err)
					os.Exit(1)
				}
				fmt.Println("Resource Deleted: ", resourceName)
			}
		}
	}
}

func init() {
	RegisterCliCommand("delete-acctest-resources", nil, deleteAccTestResources)
}
