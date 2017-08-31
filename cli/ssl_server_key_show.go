package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-brocade-vtm/api/ssl_server_key"
	"github.com/sky-uk/go-rest-api"
	"os"
)

func showSSLServerKey(client *rest.Client, flagSet *flag.FlagSet) {

	var sslServerKeyObject sslServerKey.SSLServerKey

	sslServerKeyName := flagSet.Lookup("name").Value.String()

	if sslServerKeyName == "" {
		fmt.Printf("\nError name argument required\n")
		os.Exit(1)
	}

	sslServerKeyShowAPI := sslServerKey.NewGet(sslServerKeyName)
	err := client.Do(sslServerKeyShowAPI)
	if err != nil {
		fmt.Printf("\nError retrieving SSL server key %s from API. Error %+v\n", sslServerKeyName, err)
		errObj := *sslServerKeyShowAPI.ErrorObject().(*api.VTMError)
		PrettyPrintErrorObj(errObj)
		os.Exit(2)
	}
	sslServerKeyObject = *sslServerKeyShowAPI.ResponseObject().(*sslServerKey.SSLServerKey)
	row := map[string]interface{}{}
	row["Name"] = sslServerKeyName
	row["Note"] = sslServerKeyObject.Properties.Basic.Note
	row["Certificate"] = sslServerKeyObject.Properties.Basic.Public
	row["CSR"] = sslServerKeyObject.Properties.Basic.Request
	row["Private"] = "Not available"
	PrettyPrintSingle(row)
}

func init() {
	sslServerKeyShowFlags := flag.NewFlagSet("ssh-server-key-show", flag.ExitOnError)
	sslServerKeyShowFlags.String("name", "", "usage: -name ssl-server-key-name")
	RegisterCliCommand("ssh-server-key-show", sslServerKeyShowFlags, showSSLServerKey)
}
