package pool

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"strconv"
	"testing"
)

var name = "pool_" + strconv.Itoa(rand.Int())

func TestSetPool(t *testing.T) {
	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	resource := Pool{}
	tf := true
	ff := false

	priority := uint(1)
	weight := 1
	maxConnectionAttempts := uint(10)
	maxConnectTime := uint(60)
	maxConnectionsPerNode := uint(10)
	maxQueueSize := uint(20)
	maxReplyTime := uint(60)
	queueTimeout := uint(60)
	priorityNodes := uint(8)

	resource.Properties.Basic = Basic{
		Monitors: []string{"ping"},
		NodesTable: []MemberNode{{
			Node:     "127.0.0.1:80",
			Priority: &priority,
			State:    "active",
			Weight:   &weight,
		}},
		MaxConnectionAttempts:        &maxConnectionAttempts,
		MaxIdleConnectionsPerNode:    20,
		MaxTimeoutConnectionAttempts: 20,
		NodeCloseWithReset:           true,
	}
	resource.Properties.Connection.MaxConnectTime = &maxConnectTime
	resource.Properties.Connection.MaxConnectionsPerNode = &maxConnectionsPerNode
	resource.Properties.Connection.MaxQueueSize = &maxQueueSize
	resource.Properties.Connection.MaxReplyTime = &maxReplyTime
	resource.Properties.Connection.QueueTimeout = &queueTimeout

	resource.Properties.HTTP.HTTPKeepAlive = &ff
	resource.Properties.HTTP.HTTPKeepAliveNonIdempotent = &ff
	resource.Properties.LoadBalancing.PriorityEnabled = &ff

	resource.Properties.LoadBalancing.PriorityNodes = &priorityNodes

	resource.Properties.LoadBalancing.Algorithm = "least_connections"
	resource.Properties.TCP.Nagle = &tf
	resource.Properties.DNSAutoScale.Enabled = &ff
	resource.Properties.DNSAutoScale.Hostnames = []string{}

	newPool := Pool{}
	err = client.Set("pools", name, resource, &newPool)
	if err != nil {
		t.Fatal("Error creating a resource: ", err)
	}
	log.Println("Resource created: ", name)
	assert.Equal(t, uint(10), *newPool.Properties.Basic.MaxConnectionAttempts)
}

func TestGetPool(t *testing.T) {
	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	client.WorkWithConfigurationResources()

	pool := Pool{}
	err = client.GetByName("pools", name, &pool)
	if err != nil {
		t.Fatal("Error getting a resource: ", err)
	}
	log.Println("Resource found: ", pool)
	assert.Equal(t, uint(10), *pool.Properties.Basic.MaxConnectionAttempts)
}

func TestDeletePool(t *testing.T) {
	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	err = client.Delete("pools", name)
	if err != nil {
		t.Fatal("Error deleting a resource: ", err)
	} else {
		log.Printf("Resource %s deleted", name)
	}
}
