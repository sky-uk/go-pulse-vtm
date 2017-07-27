package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/virtualserver"
	"github.com/sky-uk/go-rest-api"
	"os"
)

var deleteVirtualServerName string

func deleteVirtualServer(client *rest.Client, flagSet *flag.FlagSet) {

	if deleteVirtualServerName == "" {
		fmt.Printf("\nError: name argument required\n")
		os.Exit(1)
	}
	deleteVirtualServerAPI := virtualserver.NewDelete(deleteVirtualServerName)
	err := client.Do(deleteVirtualServerAPI)
	if err != nil {
		fmt.Printf("\nError whilst deleting virtual server %s. Error: %+v\n", deleteVirtualServerName, err)
		os.Exit(2)
	}
	fmt.Printf("\nSuccessfully deleted virtual server %s\n", deleteVirtualServerName)
}

func init() {
	deleteVirtualServerFlags := flag.NewFlagSet("virtual-server-delete", flag.ExitOnError)
	deleteVirtualServerFlags.StringVar(&deleteVirtualServerName, "name", "", "usage: -name virtual-server-name")
	RegisterCliCommand("virtual-server-delete", deleteVirtualServerFlags, deleteVirtualServer)
}
