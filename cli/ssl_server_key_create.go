package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-brocade-vtm/api/ssl_server_key"
	"github.com/sky-uk/go-rest-api"
	"os"
)

func createSSLServerKey(client *rest.Client, flagSet *flag.FlagSet) {

	var sslServerKeyObject sslServerKey.SSLServerKey
	sslServerKeyName := flagSet.Lookup("name").Value.String()

	if sslServerKeyName == "" {
		fmt.Printf("\nError name argument required\n")
		os.Exit(1)
	}
	sslServerKeyObject.Properties.Basic.Note = flagSet.Lookup("note").Value.String()
	privateKey := retrieveSSLKeyFile(flagSet.Lookup("private-key-file").Value.String())
	certificate := retrieveSSLKeyFile(flagSet.Lookup("certificate-file").Value.String())
	csr := retrieveSSLKeyFile(flagSet.Lookup("csr-file").Value.String())

	if privateKey != "" {
		sslServerKeyObject.Properties.Basic.Private = privateKey
	}
	if certificate != "" {
		sslServerKeyObject.Properties.Basic.Public = string(certificate)
	}
	if csr != "" {
		sslServerKeyObject.Properties.Basic.Request = csr
	}

	createSSLServerKeyAPI := sslServerKey.NewCreate(sslServerKeyName, sslServerKeyObject)
	err := client.Do(createSSLServerKeyAPI)
	if err != nil {
		fmt.Printf("\nError while creating SSL Server Key %s", sslServerKeyName)
		errObj := *createSSLServerKeyAPI.ErrorObject().(*api.VTMError)
		PrettyPrintErrorObj(errObj)
		os.Exit(2)
	}
	fmt.Printf("\nSuccessfully created SSL Server Key %s\n", sslServerKeyName)
}

func init() {
	createSSLServerKeyFlags := flag.NewFlagSet("ssl-server-key-create", flag.ExitOnError)
	createSSLServerKeyFlags.String("name", "", "usage: -name ssl-server-key-name")
	createSSLServerKeyFlags.String("note", "", "usage: -note 'a note'")
	createSSLServerKeyFlags.String("private-key-file", "", "usage: -private-key-file /path/to/key")
	createSSLServerKeyFlags.String("certificate-file", "", "usage: -certificate-file /path/to/certificate")
	createSSLServerKeyFlags.String("csr-file", "", "usage: -csr-file /path/to/csr")
	RegisterCliCommand("ssl-server-key-create", createSSLServerKeyFlags, createSSLServerKey)
}
