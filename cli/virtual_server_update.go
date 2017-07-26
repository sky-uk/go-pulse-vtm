package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/virtualserver"
	"github.com/sky-uk/go-rest-api"
	"os"
)

var updateVirtualServerName, updateVirtualServerListenTrafficGroup string
var updateVirtualServerEnabled, updateVirtualServerKeepalive, updateVirtualServerListenAny bool
var updateVirtualServerObject virtualserver.VirtualServer

func updateVirtualServer(client *rest.Client, flagSet *flag.FlagSet) {

	if updateVirtualServerName == "" {
		fmt.Printf("\nError name argument required\n")
		os.Exit(1)
	}
	if updateVirtualServerListenTrafficGroup != "" {
		listenIPAddresses := make([]string, 1)
		listenIPAddresses[0] = updateVirtualServerListenTrafficGroup
		updateVirtualServerObject.Properties.Basic.ListenOnTrafficIps = listenIPAddresses
	}
	updateVirtualServerObject.Properties.Basic.Enabled = &updateVirtualServerEnabled
	updateVirtualServerObject.Properties.Connection.Keepalive = &updateVirtualServerKeepalive
	updateVirtualServerObject.Properties.Basic.ListenOnAny = &updateVirtualServerListenAny

	updateVirtualServerAPI := virtualserver.NewUpdate(updateVirtualServerName, updateVirtualServerObject)
	err := client.Do(updateVirtualServerAPI)
	if err != nil {
		fmt.Printf("\nError whilst updating virtual server %s. Error: %+v\n", updateVirtualServerName, err)
		os.Exit(2)
	}
	fmt.Printf("\nSuccessfully updated virtual server %s\n", updateVirtualServerName)
}

func init() {
	updateVirtualServerFlags := flag.NewFlagSet("virtual-server-update", flag.ExitOnError)
	updateVirtualServerFlags.StringVar(&updateVirtualServerName, "name", "", "usage: -name virtual-server-name")
	updateVirtualServerFlags.StringVar(&updateVirtualServerListenTrafficGroup, "listen-traffic-ip-group", "", "usage: -listen-ip-traffic-group traffic-ip-group-name")
	updateVirtualServerFlags.BoolVar(&updateVirtualServerListenAny, "listen-on-any", false, "usage: -listen-on-any")
	updateVirtualServerFlags.StringVar(&updateVirtualServerObject.Properties.Basic.Pool, "pool", "", "usage: -pool pool-name")
	updateVirtualServerFlags.StringVar(&updateVirtualServerObject.Properties.Basic.Protocol, "protocol", "", "usage: -protocol protocol-type")
	updateVirtualServerFlags.BoolVar(&updateVirtualServerEnabled, "enabled", false, "usage: -enabled")
	updateVirtualServerFlags.BoolVar(&updateVirtualServerKeepalive, "keepalive", false, "usage: -keepalive")
	updateVirtualServerFlags.UintVar(&updateVirtualServerObject.Properties.Basic.Port, "port", 80, "usage: -port xx")
	updateVirtualServerFlags.UintVar(&updateVirtualServerObject.Properties.Connection.KeepaliveTimeout, "keepalive-timeout", 10, "usage: -keepalive-timeout xx")
	updateVirtualServerFlags.UintVar(&updateVirtualServerObject.Properties.Connection.Timeout, "connection-timeout", 40, "usage: -connection-timeout xx")
	RegisterCliCommand("virtual-server-update", updateVirtualServerFlags, updateVirtualServer)
}
