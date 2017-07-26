package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/traffic_ip_group"
	"github.com/sky-uk/go-brocade-vtm/api/traffic_ip_group_manager"
	"github.com/sky-uk/go-rest-api"
	"os"
)

var createTrafficIPGroupName string
var createTrafficIPGroupObject trafficIpGroups.TrafficIPGroup
var createTrafficIPGroupEnable, createTrafficIPGroupHashSourcePort bool
var createTrafficIPGroupListenIP string

func createTrafficIPGroup(client *rest.Client, flagSet *flag.FlagSet) {

	var trafficManagers []string

	if createTrafficIPGroupName == "" {
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
	createTrafficIPGroupObject.Properties.Basic.Machines = trafficManagers
	createTrafficIPGroupObject.Properties.Basic.Enabled = &createTrafficIPGroupEnable
	createTrafficIPGroupObject.Properties.Basic.HashSourcePort = &createTrafficIPGroupHashSourcePort

	if createTrafficIPGroupListenIP != "" {
		listenIPAddresses := make([]string, 1)
		listenIPAddresses[0] = createTrafficIPGroupListenIP
		createTrafficIPGroupObject.Properties.Basic.IPAddresses = listenIPAddresses
	}

	createTrafficGroupIPAPI := trafficIpGroups.NewCreate(createTrafficIPGroupName, createTrafficIPGroupObject)
	err = client.Do(createTrafficGroupIPAPI)
	if err != nil {
		fmt.Printf("\nError while creating Traffic IP Group %s. Error %+v", createTrafficIPGroupName, err)
		os.Exit(3)
	}
	fmt.Printf("\nSuccessfully created Traffic IP Group %s\n", createTrafficIPGroupName)
}

func init() {
	createTrafficIPGroupFlags := flag.NewFlagSet("traffic-ip-group-create", flag.ExitOnError)
	createTrafficIPGroupFlags.StringVar(&createTrafficIPGroupName, "name", "", "usage: -name traffic-ip-group-name")
	createTrafficIPGroupFlags.BoolVar(&createTrafficIPGroupEnable, "enabled", false, "usage: -enabled")
	createTrafficIPGroupFlags.BoolVar(&createTrafficIPGroupHashSourcePort, "hash-source-port", false, "usage: -hash-source-port")
	createTrafficIPGroupFlags.StringVar(&createTrafficIPGroupListenIP, "listen-ip-address", "", "usage: -listen-ip-address xxx.yyy.zzz.vvv (only supports one IP)")
	createTrafficIPGroupFlags.StringVar(&createTrafficIPGroupObject.Properties.Basic.Mode, "mode", "", "usage: -mode singlehosted|ec2elastic|ec2vpcelastic|ec2vpcprivate|multihosted|rhi")
	createTrafficIPGroupFlags.StringVar(&createTrafficIPGroupObject.Properties.Basic.Multicast, "multicast-ip", "", "usage: -multicast xxx.yyy.zzz.vvv (must be a valid multicast IP)")
	RegisterCliCommand("traffic-ip-group-create", createTrafficIPGroupFlags, createTrafficIPGroup)
}
