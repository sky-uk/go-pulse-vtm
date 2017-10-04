package trafficManager

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"log"
	"testing"
)

func TestGetTrafficManager(t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	client.WorkWithConfigurationResources()

	trafficManagers := TrafficManagers{}
	err = client.GetByName("traffic_managers", "", &trafficManagers)
	if err != nil {
		t.Fatal("Error getting resource: ", err)
	}
	log.Println("Found Traffic Managers: ", trafficManagers)
}
