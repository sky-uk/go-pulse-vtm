package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/traffic_ip_group"
	"github.com/sky-uk/go-rest-api"
	"os"
)

func deleteTrafficIPGroup(client *rest.Client, flagSet *flag.FlagSet) {

	trafficIPGroupName := flagSet.Lookup("name").Value.String()
	if trafficIPGroupName == "" {
		fmt.Printf("\nError name argument required\n")
		os.Exit(1)
	}

	if apiVersion != "" {
		trafficIpGroups.TrafficIPGroupEndpoint = "/api/tm/" + apiVersion + "/config/active/traffic_ip_groups/"
	}

	deleteTrafficIPGroupAPI := trafficIpGroups.NewDelete(trafficIPGroupName)
	err := client.Do(deleteTrafficIPGroupAPI)
	if err != nil {
		fmt.Printf("\nError whilst deleting Traffic IP Group %s\n", trafficIPGroupName)
		os.Exit(2)
	}
	fmt.Printf("\nSuccessfully deleted Traffic IP Group %s\n", trafficIPGroupName)
}

func init() {
	deleteTrafficIPGroupFlags := flag.NewFlagSet("traffic-ip-group-delete", flag.ExitOnError)
	deleteTrafficIPGroupFlags.String("name", "", "usage: -name traffic-ip-group-name")
	deleteTrafficIPGroupFlags.StringVar(&apiVersion, "apiversion", "", "usage: -apiversion 3.8")
	RegisterCliCommand("traffic-ip-group-delete", deleteTrafficIPGroupFlags, deleteTrafficIPGroup)
}
