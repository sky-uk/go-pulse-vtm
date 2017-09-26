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

var poolName, poolNodes, poolNodesState string
var poolNodesPriority, poolNodesWeight int

func createPool(client *rest.Client, flagSet *flag.FlagSet) {

	if poolName == "" {
		fmt.Printf("\nError -name argument required\n")
		os.Exit(1)
	}

	if apiVersion != "" {
		pool.PoolEndpoint = "/api/tm/" + apiVersion + "/config/active/pools/"
	}

	createPoolObject := new(pool.Pool)
	poolNodeList := strings.Split(poolNodes, ",")

	// The state, priority and weight are set the same for all nodes.
	var nodesTable []pool.MemberNode
	for _, node := range poolNodeList {
		memberNode := new(pool.MemberNode)
		memberNode.Node = node
		memberNode.State = poolNodesState
		memberNode.Weight = poolNodesWeight
		memberNode.Priority = poolNodesPriority

		nodesTable = append(nodesTable, *memberNode)
	}
	createPoolObject.Properties.Basic.NodesTable = nodesTable

	createPoolAPI := pool.NewCreate(poolName, *createPoolObject)
	err := client.Do(createPoolAPI)
	if err != nil {
		fmt.Printf("\nError creating pool %s. Error: %+v\n", poolName, err)
		errObj := *createPoolAPI.ErrorObject().(*api.VTMError)
		PrettyPrintErrorObj(errObj)
		os.Exit(2)
	}
	fmt.Printf("\nSuccessfully created pool %s\n", poolName)
}

func init() {
	createPoolFlags := flag.NewFlagSet("pool-create", flag.ExitOnError)
	createPoolFlags.StringVar(&poolName, "name", "", "usage: -name pool-name")
	createPoolFlags.StringVar(&poolNodes, "nodes", "", "usage: -nodes node1:80,node2:80")
	createPoolFlags.StringVar(&poolNodesState, "state", "active", "usage: -state active (all nodes)")
	createPoolFlags.IntVar(&poolNodesPriority, "priority", 10, "usage: -priority 10 (all nodes)")
	createPoolFlags.IntVar(&poolNodesWeight, "weight", 1, "usage: -weight 1 (all nodes")
	createPoolFlags.StringVar(&apiVersion, "apiversion", "", "usage: -apiversion 3.8")
	RegisterCliCommand("pool-create", createPoolFlags, createPool)
}
