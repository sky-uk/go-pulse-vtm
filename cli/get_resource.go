package main

import (
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/sky-uk/go-pulse-vtm/api"
	"os"
)

func getResource(client *api.Client, flagSet *flag.FlagSet) {

	objType := flagSet.Lookup("type").Value.String()
	objName := flagSet.Lookup("name").Value.String()
	objHref := flagSet.Lookup("href").Value.String()

	// work with an environment
	client.WorkWithConfigurationResources()

	res := make(map[string]interface{})
	var err error
	if objType != "" && objName != "" {
		err = client.GetByName(objType, objName, &res)
	} else {
		err = client.GetByURL(objHref, &res)
	}
	if err != nil {
		fmt.Println("Error getting a resource: ", err)
		os.Exit(2)
	}
	spew.Dump(res)
}

func init() {
	getResourceFlags := flag.NewFlagSet("get-resource", flag.ExitOnError)
	getResourceFlags.String("type", "", "usage: -type <resource type>")
	getResourceFlags.String("name", "", "usage: -name <resource name>")
	getResourceFlags.String("href", "", "usage: -href <resource url>")
	RegisterCliCommand("get-resource", getResourceFlags, getResource)
}
