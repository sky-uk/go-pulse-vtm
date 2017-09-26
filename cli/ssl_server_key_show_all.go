package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/ssl_server_key"
	"github.com/sky-uk/go-rest-api"
	"os"
)

func showAllSSLServerKey(client *rest.Client, flagSet *flag.FlagSet) {

	if apiVersion != "" {
		sslServerKey.SSLServerKeyEndpoint = "/api/tm/" + apiVersion + "/config/active/server_keys/"
	}

	sslServerKeyShowAllAPI := sslServerKey.NewGetAll()
	err := client.Do(sslServerKeyShowAllAPI)
	if err != nil {
		fmt.Printf("\nError retrieving the SSL Server Key list: %+v", err)
		os.Exit(1)
	}

	sslServerKeyList := *sslServerKeyShowAllAPI.ResponseObject().(*sslServerKey.SSLServerKeysList)
	rows := []map[string]interface{}{}
	headers := []string{"Name", "HREF"}

	for _, sslServerKey := range sslServerKeyList.Children {
		row := map[string]interface{}{}
		row["Name"] = sslServerKey.Name
		row["HREF"] = sslServerKey.HRef
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)
}

func init() {
	showAllSSLServerKeyFlags := flag.NewFlagSet("ssl-server-key-show-all", flag.ExitOnError)
	showAllSSLServerKeyFlags.StringVar(&apiVersion, "apiversion", "", "usage: -apiversion 3.8")
	RegisterCliCommand("ssl-server-key-show-all", showAllSSLServerKeyFlags, showAllSSLServerKey)
}
