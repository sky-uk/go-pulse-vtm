package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/pool"
	"github.com/sky-uk/go-rest-api"
	"os"
	"strings"
)

var poolName, poolNodes, poolNodesState string
var createPoolObject pool.Pool
var poolNodesPriority, poolNodesWeight int

func createPool(client *rest.Client, flagSet *flag.FlagSet) {

	//TODO Not working yet


	if poolName == "" {
		fmt.Printf("\nError -name argument required\n")
		os.Exit(1)
	}
	poolNodeList := strings.Split(poolNodes, ",")
	nodesTable := make([]pool.MemberNode, len(poolNodeList))

	for idx, node := range poolNodeList {
		nodesTable[idx].Node = node
		nodesTable[idx].State = poolNodesState
		nodesTable[idx].Priority = poolNodesPriority
		nodesTable[idx].Weight = poolNodesWeight
	}

	createPoolObject.Properties.Basic.NodesTable = nodesTable

	fmt.Printf("\nNodes table is %+v\n", createPoolObject.Properties.Basic.NodesTable)

	createPoolAPI := pool.NewCreate(poolName, createPoolObject)
	err := client.Do(createPoolAPI)
	if err != nil {
		fmt.Printf("\nError creating pool %s. Error: %+v\n", poolName, err)
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
	RegisterCliCommand("pool-create", createPoolFlags, createPool)
}
