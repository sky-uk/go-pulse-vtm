package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-brocade-vtm/api/pool"
	"github.com/sky-uk/go-rest-api"
	"os"
)

var showPoolName string

func showPool(client *rest.Client, flagSet *flag.FlagSet) {

	if showPoolName == "" {
		fmt.Printf("\nError name argument required. Usage: -name pool-name\n")
		os.Exit(1)
	}

	if apiVersion != "" {
		pool.PoolEndpoint = "/api/tm/" + apiVersion + "/config/active/pools/"
	}

	readPoolAPI := pool.NewGet(showPoolName)
	err := client.Do(readPoolAPI)
	if err != nil {
		fmt.Printf("\nError whilst retrieving pool %s\n", showMonitorName)
		errObj := *readPoolAPI.ErrorObject().(*api.VTMError)
		PrettyPrintErrorObj(errObj)
		os.Exit(2)
	}

	response := readPoolAPI.ResponseObject().(*pool.Pool)
	row := map[string]interface{}{}
	row["Name"] = showPoolName

	var nodeTableEntry string
	for _, nodeItem := range response.Properties.Basic.NodesTable {
		if nodeTableEntry != "" {
			nodeTableEntry = fmt.Sprintf("%s |", nodeTableEntry)
		}
		nodeTableEntry = fmt.Sprintf("%s %s, %s, %d, %d", nodeTableEntry, nodeItem.Node, nodeItem.State, nodeItem.Priority, nodeItem.Weight)
	}
	row["Node Table (Node, State, Priority, Weight)"] = nodeTableEntry
	PrettyPrintSingle(row)
}

func init() {
	showPoolFlags := flag.NewFlagSet("pool-show", flag.ExitOnError)
	showPoolFlags.StringVar(&showPoolName, "name", "", "usage: -name pool-name")
	showPoolFlags.StringVar(&apiVersion, "apiversion", "", "usage: -apiversion 3.8")
	RegisterCliCommand("pool-show", showPoolFlags, showPool)
}
