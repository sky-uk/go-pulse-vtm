package monitor

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"strconv"
	"testing"
)

var name = "monitor_" + strconv.Itoa(rand.Int())

func TestSetMonitor(t *testing.T) {
	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	resource := Monitor{}
	f := false
	resource.Properties.Basic = Basic{
		Delay:     5,
		Timeout:   5,
		Failures:  9,
		Verbose:   &f,
		UseSSL:    &f,
		CanUseSSL: &f,
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

	newMonitor := Monitor{}
	err = client.Set("monitors", name, resource, &newMonitor)
	if err != nil {
		t.Fatal("Error creating a resource: ", err)
	}
	log.Println("Resource created: ", name)
	assert.Equal(t, uint(9), newMonitor.Properties.Basic.Failures)
}

func TestGetMonitor(t *testing.T) {
	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}

	monitor := Monitor{}
	err = client.GetByName("monitors", name, &monitor)
	if err != nil {
		t.Fatal("Error getting a resource: ", err)
	}
	log.Println("Resource found: ", monitor)
	assert.Equal(t, uint(9), monitor.Properties.Basic.Failures)
}

func TestDeleteMonitor(t *testing.T) {
	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	err = client.Delete("monitors", name)
	if err != nil {
		t.Fatal("Error deleting a resource: ", err)
	} else {
		log.Printf("Resource %s deleted", name)
	}
}
