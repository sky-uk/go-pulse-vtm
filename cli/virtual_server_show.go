package main

import (
	"flag"
	"github.com/sky-uk/go-rest-api"
)

var showVirtualServerName string

func showVirtualServer(client *rest.Client, flagSet *flag.FlagSet) {

}

func init() {
	showVirtualServerFlags := flag.NewFlagSet("virtual-server-show", flag.ExitOnError)
	showVirtualServerFlags.StringVar(&showVirtualServerName, "name", "", "usage: -name virtual-server-name")
	RegisterCliCommand("virtual-server-show", showVirtualServerFlags, showVirtualServer)
}
