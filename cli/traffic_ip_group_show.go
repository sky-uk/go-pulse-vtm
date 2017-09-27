package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/traffic_ip_group"
	"github.com/sky-uk/go-rest-api"
	"os"
)

func showTrafficIPGroup(client *rest.Client, flagSet *flag.FlagSet) {

	trafficIPGroupName := flagSet.Lookup("name").Value.String()

	if trafficIPGroupName == "" {
		fmt.Printf("\nError name argument required\n")
		os.Exit(1)
	}

	if apiVersion != "" {
		trafficIpGroups.TrafficIPGroupEndpoint = "/api/tm/" + apiVersion + "/config/active/traffic_ip_groups/"
	}

	getTrafficIPGroupAPI := trafficIpGroups.NewGet(trafficIPGroupName)
	err := client.Do(getTrafficIPGroupAPI)
	if err != nil {
		fmt.Printf("\nError retrieving Traffic IP Group %s\n", trafficIPGroupName)
		os.Exit(2)
	}
	response := getTrafficIPGroupAPI.ResponseObject().(*trafficIpGroups.TrafficIPGroup)
	row := map[string]interface{}{}
	row["Name"] = trafficIPGroupName
	row["Enabled"] = *response.Properties.Basic.Enabled
	row["Hash-source-port"] = *response.Properties.Basic.HashSourcePort
	if len(response.Properties.Basic.IPAddresses) > 0 {
		row["listen-ip-address"] = response.Properties.Basic.IPAddresses[0]
	}
	row["Mode"] = response.Properties.Basic.Mode
	row["multicast-ip"] = response.Properties.Basic.Multicast
	PrettyPrintSingle(row)
}

func init() {
	showTrafficIPGroupFlags := flag.NewFlagSet("traffic-ip-group-show", flag.ExitOnError)
	showTrafficIPGroupFlags.String("name", "", "usage: -name traffic-ip-group-name")
	showTrafficIPGroupFlags.StringVar(&apiVersion, "apiversion", "", "usage: -apiversion 3.8")
	RegisterCliCommand("traffic-ip-group-show", showTrafficIPGroupFlags, showTrafficIPGroup)
}
