package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/monitor"
	"github.com/sky-uk/go-rest-api"
	"os"
)

var showMonitorName string

func showMonitor(client *rest.Client, flagSet *flag.FlagSet) {

	if showMonitorName == "" {
		fmt.Printf("\nError: name argument required\n")
		os.Exit(1)
	}

	readMonitorAPI := monitor.NewGet(showMonitorName)
	err := client.Do(readMonitorAPI)
	if err != nil {
		fmt.Printf("\nError whilst retrieving monitor %s\n", showMonitorName)
		os.Exit(2)
	}
	response := readMonitorAPI.ResponseObject().(*monitor.Monitor)
	row := map[string]interface{}{}
	row["Name"] = showMonitorName
	row["Delay"] = response.Properties.Basic.Delay
	row["Timeout"] = response.Properties.Basic.Timeout
	row["Failures"] = response.Properties.Basic.Failures
	row["Verbose"] = *response.Properties.Basic.Verbose
	row["UseSSL"] = *response.Properties.Basic.UseSSL
	row["HTTP-Host-Header"] = response.Properties.HTTP.HostHeader
	row["HTTP-Path"] = response.Properties.HTTP.URIPath
	row["HTTP-Authentication"] = response.Properties.HTTP.Authentication
	row["HTTP-Body-Regex"] = response.Properties.HTTP.BodyRegex
	PrettyPrintSingle(row)
}

func init() {
	showMonitorFlags := flag.NewFlagSet("monitor-show", flag.ExitOnError)
	showMonitorFlags.StringVar(&showMonitorName, "name", "", "usage: -name monitor-name")
	RegisterCliCommand("monitor-show", showMonitorFlags, showMonitor)
}
