package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-brocade-vtm/api/virtualserver"
	"github.com/sky-uk/go-rest-api"
	"os"
)

var showVirtualServerName string

func showVirtualServer(client *rest.Client, flagSet *flag.FlagSet) {

	if showVirtualServerName == "" {
		fmt.Printf("\nError: name argument required\n")
		os.Exit(1)
	}

	if apiVersion != "" {
		virtualserver.VirtualServerEndpoint = "/api/tm/" + apiVersion + "/config/active/virtual_servers/"
	}

	readVirtualServerAPI := virtualserver.NewGet(showVirtualServerName)
	err := client.Do(readVirtualServerAPI)
	if err != nil {
		fmt.Printf("\nError whilst reading virtual server %s. Error: %+v\n", showVirtualServerName, err)
		errObj := *readVirtualServerAPI.ErrorObject().(*api.VTMError)
		PrettyPrintErrorObj(errObj)
		os.Exit(2)
	}
	response := readVirtualServerAPI.ResponseObject().(*virtualserver.VirtualServer)

	row := map[string]interface{}{}
	row["Name"] = showVirtualServerName
	row["Listen-Traffic-IP-Group"] = response.Properties.Basic.ListenOnTrafficIps
	row["Listen-on-any"] = response.Properties.Basic.ListenOnAny
	row["Port"] = response.Properties.Basic.Port
	row["Pool"] = response.Properties.Basic.Pool
	row["Protocol"] = response.Properties.Basic.Protocol
	row["Enabled"] = response.Properties.Basic.Enabled
	row["Keepalive"] = response.Properties.Connection.Keepalive
	row["KeepAlive-Timeout"] = response.Properties.Connection.KeepaliveTimeout
	row["Connection-Timeout"] = response.Properties.Connection.Timeout
	PrettyPrintSingle(row)
}

func init() {
	showVirtualServerFlags := flag.NewFlagSet("virtual-server-show", flag.ExitOnError)
	showVirtualServerFlags.StringVar(&showVirtualServerName, "name", "", "usage: -name virtual-server-name")
	showVirtualServerFlags.StringVar(&apiVersion, "apiversion", "", "usage: -apiversion 3.8")
	RegisterCliCommand("virtual-server-show", showVirtualServerFlags, showVirtualServer)
}
