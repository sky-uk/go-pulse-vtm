package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/ssl_server_key"
	"github.com/sky-uk/go-rest-api"
	"os"
)

func deleteSSLServerKey(client *rest.Client, flagSet *flag.FlagSet) {

	sslServerKeyName := flagSet.Lookup("name").Value.String()
	if sslServerKeyName == "" {
		fmt.Printf("\nError: -name argument required")
		os.Exit(1)
	}
	deleteSSLServerKeyAPI := sslServerKey.NewDelete(sslServerKeyName)
	err := client.Do(deleteSSLServerKeyAPI)
	if err != nil {
		fmt.Printf("\nError deleting SSL server key %s. Error: %+v", sslServerKeyName, err)
		os.Exit(2)
	}
	fmt.Printf("\nSuccessfully deleted SSL server key %s\n", sslServerKeyName)
}

func init() {
	deleteSSLServerKeyFlags := flag.NewFlagSet("ssl-server-key-delete", flag.ExitOnError)
	deleteSSLServerKeyFlags.String("name", "", "usage: -name ssl-server-key-name")
	RegisterCliCommand("ssl-server-key-delete", deleteSSLServerKeyFlags, deleteSSLServerKey)
}
