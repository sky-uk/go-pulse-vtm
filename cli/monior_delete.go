package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/monitor"
	"github.com/sky-uk/go-rest-api"
	"net/http"
	"os"
)

func deleteMonitor(client *rest.Client, flagSet *flag.FlagSet) {

	monitorName := flagSet.Lookup("name").Value.String()
	if monitorName == "" {
		fmt.Printf("\nError: -name argument required")
		os.Exit(1)
	}
	deleteMonitorAPI := monitor.NewDelete(monitorName)
	err := client.Do(deleteMonitorAPI)
	if err != nil && deleteMonitorAPI.StatusCode() != http.StatusNotFound {
		fmt.Printf("\nError creating monitor %s. Error: %v", monitorName, err)
		os.Exit(2)
	}
	fmt.Printf("\nSuccessfully deleted monitor %s", monitorName)

}

func init() {
	deleteMonitorFlags := flag.NewFlagSet("monitor-delete", flag.ExitOnError)
	deleteMonitorFlags.String("name", "", "usage: -name monitor-name")
	RegisterCliCommand("monitor-delete", deleteMonitorFlags, deleteMonitor)
}
