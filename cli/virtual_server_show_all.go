package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/virtualserver"
	"github.com/sky-uk/go-rest-api"
	"os"
)

func showAllVirtualServer(client *rest.Client, flagSet *flag.FlagSet) {

	if apiVersion != "" {
		virtualserver.VirtualServerEndpoint = "/api/tm/" + apiVersion + "/config/active/virtual_servers/"
	}

	getAllVirtualServerAPI := virtualserver.NewGetAll()
	err := client.Do(getAllVirtualServerAPI)
	if err != nil {
		fmt.Printf("\nError retreiving the list of virtual servers")
		os.Exit(1)
	}
	virtualServerList := getAllVirtualServerAPI.ResponseObject().(*virtualserver.VirtualServersList)
	rows := []map[string]interface{}{}
	headers := []string{"Name", "HRef"}

	for _, virtualServerItem := range virtualServerList.Children {
		row := map[string]interface{}{}
		row["Name"] = virtualServerItem.Name
		row["HRef"] = virtualServerItem.Href
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)
}

func init() {
	showAllVirtualServerFlags := flag.NewFlagSet("virtual-server-show-all", flag.ExitOnError)
	showAllVirtualServerFlags.StringVar(&apiVersion, "apiversion", "", "usage: -apiversion 3.8")
	RegisterCliCommand("virtual-server-show-all", showAllVirtualServerFlags, showAllVirtualServer)
}
