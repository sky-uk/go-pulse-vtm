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
	f := false
	resource.Properties.Basic = Basic{
        Monitors: []string{"ping"},
        NodesTable: []MemberNode{{
            Node: "127.0.0.1:80",
            Priority: 1,
            State: "active",
            Weight: 1,
        }},

	}
	resource.Properties.HTTP = HTTP{
		HostHeader:     "some_header",
		Authentication: "some_authentication",
		BodyRegex:      "^healthy",
		URIPath:        "/some/other/status/page",
	}
	resource.Properties.RTSP = RTSP{
		StatusRegex: "^[234][0-9][0-9]$",
		URIPath:     "/",
		BodyRegex:   "something",
	}

	newPool := Pool{}
	err = client.Set("pools", name, resource, &newPool)
	if err != nil {
		t.Fatal("Error creating a resource: ", err)
	}
	log.Println("Resource created: ", name)
	assert.Equal(t, uint(9), newPool.Properties.Basic.Failures)
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
	assert.Equal(t, uint(9), pool.Properties.Basic.Failures)
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
