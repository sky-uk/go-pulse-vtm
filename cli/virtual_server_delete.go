package main

import (
	"flag"
	"github.com/sky-uk/go-rest-api"
)

var deleteVirtualServerName string

func deleteVirtualServer(client *rest.Client, flagSet *flag.FlagSet) {

}

func init() {
	deleteVirtualServerFlags := flag.NewFlagSet("virtual-server-delete", flag.ExitOnError)
	deleteVirtualServerFlags.StringVar(&deleteVirtualServerName, "name", "", "usage: -name virtual-server-name")
	RegisterCliCommand("virtual-server-delete", deleteVirtualServerFlags, deleteVirtualServer)
}
