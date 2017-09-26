package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-brocade-vtm/api/ssl_server_key"
	"github.com/sky-uk/go-rest-api"
	"os"
)

func updateSSLServerKey(client *rest.Client, flagSet *flag.FlagSet) {
	var sslServerKeyObject sslServerKey.SSLServerKey
	sslServerKeyName := flagSet.Lookup("name").Value.String()

	if sslServerKeyName == "" {
		fmt.Printf("\nError name argument required\n")
		os.Exit(1)
	}

	if apiVersion != "" {
		sslServerKey.SSLServerKeyEndpoint = "/api/tm/" + apiVersion + "/config/active/server_keys/"
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

	updateSSLServerKeyAPI := sslServerKey.NewUpdate(sslServerKeyName, sslServerKeyObject)
	err := client.Do(updateSSLServerKeyAPI)
	if err != nil {
		fmt.Printf("\nError while updating SSL Server Key %s", sslServerKeyName)
		errObj := *updateSSLServerKeyAPI.ErrorObject().(*api.VTMError)
		PrettyPrintErrorObj(errObj)
		os.Exit(2)
	}
	fmt.Printf("\nSuccessfully updated SSL Server Key %s\n", sslServerKeyName)
}

func init() {
	updateSSLServerKeyFlags := flag.NewFlagSet("ssl-server-key-update", flag.ExitOnError)
	updateSSLServerKeyFlags.String("name", "", "usage: -name ssl-server-key-name")
	updateSSLServerKeyFlags.String("note", "", "usage: -note 'a note'")
	updateSSLServerKeyFlags.String("private-key-file", "", "usage: -private-key-file /path/to/key")
	updateSSLServerKeyFlags.String("certificate-file", "", "usage: -certificate-file /path/to/certificate")
	updateSSLServerKeyFlags.String("csr-file", "", "usage: -csr-file /path/to/csr")
	updateSSLServerKeyFlags.StringVar(&apiVersion, "apiversion", "", "usage: -apiversion 3.8")
	RegisterCliCommand("ssl-server-key-update", updateSSLServerKeyFlags, updateSSLServerKey)
}
