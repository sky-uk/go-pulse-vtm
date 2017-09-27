package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-brocade-vtm/api/monitor"
	"github.com/sky-uk/go-rest-api"
	"net/http"
	"os"
)

func deleteMonitor(client *rest.Client, flagSet *flag.FlagSet) {

	monitorName := flagSet.Lookup("name").Value.String()
	if monitorName == "" {
		fmt.Printf("\nError: -name argument required\n")
		os.Exit(1)
	}

	if apiVersion != "" {
		monitor.MonitorEndpoint = "/api/tm/" + apiVersion + "/config/active/monitors/"
	}

	deleteMonitorAPI := monitor.NewDelete(monitorName)
	err := client.Do(deleteMonitorAPI)
	if err != nil && deleteMonitorAPI.StatusCode() != http.StatusNotFound {
		fmt.Printf("\nError deleting monitor %s. Error: %v\n", monitorName, err)
		errObj := *deleteMonitorAPI.ErrorObject().(*api.VTMError)
		PrettyPrintErrorObj(errObj)
		os.Exit(2)
	}
	fmt.Printf("\nSuccessfully deleted monitor %s\n", monitorName)

}

func init() {
	deleteMonitorFlags := flag.NewFlagSet("monitor-delete", flag.ExitOnError)
	deleteMonitorFlags.String("name", "", "usage: -name monitor-name")
	deleteMonitorFlags.StringVar(&apiVersion, "apiversion", "", "usage: -apiversion 3.8")
	RegisterCliCommand("monitor-delete", deleteMonitorFlags, deleteMonitor)
}
