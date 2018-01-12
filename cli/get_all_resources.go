package cli

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-pulse-vtm/api"
	"os"
)

func getAllResources(client *api.Client, flagSet *flag.FlagSet) {

	objType := flagSet.Lookup("type").Value.String()

	// work with an environment
	client.WorkWithConfigurationResources()
	res, err := client.GetAllResources(objType)
	if err != nil {
		fmt.Println("Error getting all configuration resources: ", err)
		os.Exit(2)
	}
	PrettyPrintMany([]string{"name", "href"}, res)
}

func init() {
	getAllResourcesFlags := flag.NewFlagSet("get-all-resources", flag.ExitOnError)
	getAllResourcesFlags.String("type", "", "usage: -type <resource type>")
	RegisterCliCommand("get-all-resources", getAllResourcesFlags, getAllResources)
}
