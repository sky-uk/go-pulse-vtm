package main

import (
	"flag"
	"github.com/sky-uk/go-rest-api"
)

var createVirtualServerName string

func createVirtualServer(client *rest.Client, flagSet *flag.FlagSet) {

}

func init() {
	createVirtualServerFlags := flag.NewFlagSet("virtual-server-create", flag.ExitOnError)
	createVirtualServerFlags.StringVar(&createVirtualServerName, "name", "", "usage: -name virtual-server-name")
	RegisterCliCommand("virtual-server-create", createVirtualServerFlags, createVirtualServer)
}
