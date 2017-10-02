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
	resource.Properties.Basic = Basic{
        Monitors: []string{"ping"},
        NodesTable: []MemberNode{{
            Node: "127.0.0.1:80",
            Priority: 1,
            State: "active",
            Weight: 1,
        }},
        MaxConnectionAttempts: 10,
        MaxIdleConnectionsPerNode: 20,
        MaxTimeoutConnectionAttempts: 20,
        NodeCloseWithReset: &tf,
	}
	resource.Properties.Connection.MaxConnectTime = 60
	resource.Properties.Connection.MaxConnectionsPerNode = 10
	resource.Properties.Connection.MaxQueueSize = 20
	resource.Properties.Connection.MaxReplyTime = 60
	resource.Properties.Connection.QueueTimeout = 60
	resource.Properties.HTTP.HTTPKeepAlive = &ff
	resource.Properties.HTTP.HTTPKeepAliveNonIdempotent = &ff
	resource.Properties.LoadBalancing.PriorityEnabled = &ff
    resource.Properties.LoadBalancing.PriorityNodes = 8
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
	assert.Equal(t, uint(10), newPool.Properties.Basic.MaxConnectionAttempts)
}

func TestGetPool(t *testing.T) {
	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}

	pool := Pool{}
	err = client.GetByName("pools", name, &pool)
	if err != nil {
		t.Fatal("Error getting a resource: ", err)
	}
	log.Println("Resource found: ", pool)
	assert.Equal(t, uint(10), pool.Properties.Basic.MaxConnectionAttempts)
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
