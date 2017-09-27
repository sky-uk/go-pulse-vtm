package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-brocade-vtm/api/pool"
	"github.com/sky-uk/go-rest-api"
	"os"
	"strings"
)

var updatePoolName, updatePoolNodes, updatePoolNodeState string
var updatePoolNodePriority, updatePoolNodeWeight int

func updatePool(client *rest.Client, flagSet *flag.FlagSet) {

	if updatePoolName == "" {
		fmt.Printf("\nError name argument required. Usage: -name pool-name\n")
		os.Exit(1)
	}

	if apiVersion != "" {
		pool.PoolEndpoint = "/api/tm/" + apiVersion + "/config/active/pools/"
	}

	updatePoolObject := new(pool.Pool)
	poolNodeList := strings.Split(updatePoolNodes, ",")

	// The state, priority and weight are set the same for all nodes.
	var nodesTable []pool.MemberNode
	for _, node := range poolNodeList {
		memberNode := new(pool.MemberNode)
		memberNode.Node = node
		memberNode.State = updatePoolNodeState
		memberNode.Weight = updatePoolNodeWeight
		memberNode.Priority = updatePoolNodePriority

		nodesTable = append(nodesTable, *memberNode)
	}
	updatePoolObject.Properties.Basic.NodesTable = nodesTable

	updatePoolAPI := pool.NewCreate(updatePoolName, *updatePoolObject)
	err := client.Do(updatePoolAPI)
	if err != nil {
		fmt.Printf("\nError updating pool %s. Error: %+v\n", updatePoolName, err)
		errObj := *updatePoolAPI.ErrorObject().(*api.VTMError)
		PrettyPrintErrorObj(errObj)
		os.Exit(2)
	}
	fmt.Printf("\nSuccessfully updated pool %s\n", updatePoolName)
}

func init() {
	updatePoolFlags := flag.NewFlagSet("pool-update", flag.ExitOnError)
	updatePoolFlags.StringVar(&updatePoolName, "name", "", "usage: -name pool-name")
	updatePoolFlags.StringVar(&updatePoolNodes, "nodes", "", "usage: -nodes node1:80,node2:80")
	updatePoolFlags.StringVar(&updatePoolNodeState, "state", "active", "usage: -state active (all nodes)")
	updatePoolFlags.IntVar(&updatePoolNodePriority, "priority", 10, "usage: -priority 10 (all nodes)")
	updatePoolFlags.IntVar(&updatePoolNodeWeight, "weight", 1, "usage: -weight 1 (all nodes")
	updatePoolFlags.StringVar(&apiVersion, "apiversion", "", "usage: -apiversion 3.8")
	RegisterCliCommand("pool-update", updatePoolFlags, updatePool)
}
