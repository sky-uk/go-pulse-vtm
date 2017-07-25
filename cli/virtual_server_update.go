package main

import (
	"flag"
	"github.com/sky-uk/go-rest-api"
)

var updateVirtualServerName string

func updateVirtualServer(client *rest.Client, flagSet *flag.FlagSet) {

}

func init() {
	updateVirtualServerFlags := flag.NewFlagSet("virtual-server-update", flag.ExitOnError)
	updateVirtualServerFlags.StringVar(&updateVirtualServerName, "name", "", "usage: -name virtual-server-name")
	RegisterCliCommand("virtual-server-update", updateVirtualServerFlags, updateVirtualServer)
}
