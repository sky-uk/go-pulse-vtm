package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/virtualserver"
	"github.com/sky-uk/go-rest-api"
	"os"
)

var createVirtualServerName, createVirtualServerListenTrafficGroup string
var createVirtualServerEnabled, createVirtualServerKeepalive, createVirtualServerListenAny bool
var createVirtualServerObject virtualserver.VirtualServer

func createVirtualServer(client *rest.Client, flagSet *flag.FlagSet) {

	if createVirtualServerName == "" {
		fmt.Printf("\nError name argument required\n")
		os.Exit(1)
	}
	if createVirtualServerObject.Properties.Basic.Pool == "" {
		fmt.Printf("\nError pool argument required\n")
		os.Exit(1)
	}
	if createVirtualServerListenTrafficGroup != "" {
		listenIPAddresses := make([]string, 1)
		listenIPAddresses[0] = createVirtualServerListenTrafficGroup
		createVirtualServerObject.Properties.Basic.ListenOnTrafficIps = listenIPAddresses
	}
	createVirtualServerObject.Properties.Basic.Enabled = &createVirtualServerEnabled
	createVirtualServerObject.Properties.Connection.Keepalive = &createVirtualServerKeepalive
	createVirtualServerObject.Properties.Basic.ListenOnAny = &createVirtualServerListenAny

	createVirtualServerAPI := virtualserver.NewCreate(createVirtualServerName, createVirtualServerObject)
	err := client.Do(createVirtualServerAPI)
	if err != nil {
		fmt.Printf("\nError whilst creating virtual server %s. Error: %+v\n", createVirtualServerName, err)
		os.Exit(2)
	}
	fmt.Printf("\nSuccessfully created virtual server %s\n", createVirtualServerName)
}

func init() {
	createVirtualServerFlags := flag.NewFlagSet("virtual-server-create", flag.ExitOnError)
	createVirtualServerFlags.StringVar(&createVirtualServerName, "name", "", "usage: -name virtual-server-name")
	createVirtualServerFlags.StringVar(&createVirtualServerListenTrafficGroup, "listen-traffic-ip-group", "", "usage: -listen-ip-traffic-group traffic-ip-group-name")
	createVirtualServerFlags.BoolVar(&createVirtualServerListenAny, "listen-on-any", false, "usage: -listen-on-any")
	createVirtualServerFlags.StringVar(&createVirtualServerObject.Properties.Basic.Pool, "pool", "", "usage: -pool pool-name")
	createVirtualServerFlags.StringVar(&createVirtualServerObject.Properties.Basic.Protocol, "protocol", "", "usage: -protocol protocol-type")
	createVirtualServerFlags.BoolVar(&createVirtualServerEnabled, "enabled", false, "usage: -enabled")
	createVirtualServerFlags.BoolVar(&createVirtualServerKeepalive, "keepalive", false, "usage: -keepalive")
	createVirtualServerFlags.UintVar(&createVirtualServerObject.Properties.Basic.Port, "port", 80, "usage: -port xx")
	createVirtualServerFlags.UintVar(&createVirtualServerObject.Properties.Connection.KeepaliveTimeout, "keepalive-timeout", 10, "usage: -keepalive-timeout xx")
	createVirtualServerFlags.UintVar(&createVirtualServerObject.Properties.Connection.Timeout, "connection-timeout", 40, "usage: -connection-timeout xx")
	RegisterCliCommand("virtual-server-create", createVirtualServerFlags, createVirtualServer)
}
