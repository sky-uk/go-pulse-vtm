package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/traffic_ip_group"
	"github.com/sky-uk/go-brocade-vtm/api/traffic_ip_group_manager"
	"github.com/sky-uk/go-rest-api"
	"os"
)

var updateTrafficIPGroupName string
var updateTrafficIPGroupObject trafficIpGroups.TrafficIPGroup
var updateTrafficIPGroupEnable, updateTrafficIPGroupHashSourcePort bool
var updateTrafficIPGroupListenIP string

func updateTrafficIPGroup(client *rest.Client, flagSet *flag.FlagSet) {

	var trafficManagers []string

	if updateTrafficIPGroupName == "" {
		fmt.Printf("\nError: name argument required\n")
		os.Exit(1)
	}

	getAllTrafficManagersAPI := trafficIpGroupManager.NewGetAll()
	err := client.Do(getAllTrafficManagersAPI)
	if err != nil {
		fmt.Printf("\nError retrieving a list of traffic managers\n")
		os.Exit(2)
	}
	trafficManagerResponse := getAllTrafficManagersAPI.ResponseObject().(*trafficIpGroupManager.TrafficManagerChildren)
	for _, trafficManager := range trafficManagerResponse.Children {
		trafficManagers = append(trafficManagers, trafficManager.Name)
	}
	updateTrafficIPGroupObject.Properties.Basic.Machines = trafficManagers
	updateTrafficIPGroupObject.Properties.Basic.Enabled = &updateTrafficIPGroupEnable
	updateTrafficIPGroupObject.Properties.Basic.HashSourcePort = &updateTrafficIPGroupHashSourcePort

	if updateTrafficIPGroupListenIP != "" {
		listenIPAddresses := make([]string, 1)
		listenIPAddresses[0] = updateTrafficIPGroupListenIP
		updateTrafficIPGroupObject.Properties.Basic.IPAddresses = listenIPAddresses
	}

	updateTrafficGroupIPAPI := trafficIpGroups.NewUpdate(updateTrafficIPGroupName, updateTrafficIPGroupObject)
	err = client.Do(updateTrafficGroupIPAPI)
	if err != nil {
		fmt.Printf("\nError while updating Traffic IP Group %s. Error %+v", updateTrafficIPGroupName, err)
		os.Exit(3)
	}
	fmt.Printf("\nSuccessfully updated Traffic IP Group %s\n", updateTrafficIPGroupName)
}

func init() {
	updateTrafficIPGroupFlags := flag.NewFlagSet("traffic-ip-group-update", flag.ExitOnError)
	updateTrafficIPGroupFlags.StringVar(&updateTrafficIPGroupName, "name", "", "usage: -name traffic-ip-group-name")
	updateTrafficIPGroupFlags.BoolVar(&updateTrafficIPGroupEnable, "enabled", false, "usage: -enabled")
	updateTrafficIPGroupFlags.BoolVar(&updateTrafficIPGroupHashSourcePort, "hash-source-port", false, "usage: -hash-source-port")
	updateTrafficIPGroupFlags.StringVar(&updateTrafficIPGroupListenIP, "listen-ip-address", "", "usage: -listen-ip-address xxx.yyy.zzz.vvv (only supports one IP)")
	updateTrafficIPGroupFlags.StringVar(&updateTrafficIPGroupObject.Properties.Basic.Mode, "mode", "", "usage: -mode singlehosted|ec2elastic|ec2vpcelastic|ec2vpcprivate|multihosted|rhi")
	updateTrafficIPGroupFlags.StringVar(&updateTrafficIPGroupObject.Properties.Basic.Multicast, "multicast-ip", "", "usage: -multicast xxx.yyy.zzz.vvv (must be a valid multicast IP)")
	RegisterCliCommand("traffic-ip-group-update", updateTrafficIPGroupFlags, updateTrafficIPGroup)
}
