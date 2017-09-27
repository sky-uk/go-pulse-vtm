package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/traffic_ip_group"
	"github.com/sky-uk/go-rest-api"
	"os"
)

func showAllTrafficIPGroup(client *rest.Client, flagSet *flag.FlagSet) {

	if apiVersion != "" {
		trafficIpGroups.TrafficIPGroupEndpoint = "/api/tm/" + apiVersion + "/config/active/traffic_ip_groups/"
	}

	getAllTrafficIPGroupsAPI := trafficIpGroups.NewGetAll()
	err := client.Do(getAllTrafficIPGroupsAPI)
	if err != nil {
		fmt.Printf("\nError retrieving a list of Traffic IP Groups\n")
		os.Exit(1)
	}

	trafficIPGroupList := getAllTrafficIPGroupsAPI.ResponseObject().(*trafficIpGroups.TrafficIPGroupList)
	rows := []map[string]interface{}{}
	headers := []string{"Name", "HRef"}

	for _, trafficIPGroupItem := range trafficIPGroupList.Children {
		row := map[string]interface{}{}
		row["Name"] = trafficIPGroupItem.Name
		row["HRef"] = trafficIPGroupItem.HRef
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)

}

func init() {
	showAllTrafficIPGroupFlags := flag.NewFlagSet("traffic-ip-group-show-all", flag.ExitOnError)
	showAllTrafficIPGroupFlags.StringVar(&apiVersion, "apiversion", "", "usage: -apiversion 3.8")
	RegisterCliCommand("traffic-ip-group-show-all", showAllTrafficIPGroupFlags, showAllTrafficIPGroup)
}
