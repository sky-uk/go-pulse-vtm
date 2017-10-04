package trafficIpGroups

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-brocade-vtm/api/model/3.8/traffic_manager"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestTrafficIPGroup(t *testing.T) {

	name := api.SetTestResourceName("traffic_ip_group_")
	setTrafficIPGroup(name, t)
	getTrafficIPGroup(name, t)
	deleteTrafficIPGroup(name, t)
}

func setTrafficIPGroup(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}

	setTrue := true
	trafficManagers := getTrafficManagers(t)

	resource := TrafficIPGroup{}
	resource.Properties.Basic.Enabled = &setTrue
	resource.Properties.Basic.Note = "go-brocade-vtm test traffic IP group"
	resource.Properties.Basic.HashSourcePort = &setTrue
	resource.Properties.Basic.IPAssignmentMode = "alphabetic"
	resource.Properties.Basic.IPAddresses = []string{"10.0.34.5"}
	resource.Properties.Basic.Machines = trafficManagers
	resource.Properties.Basic.Mode = "singlehosted"

	newTrafficIPGroup := TrafficIPGroup{}
	err = client.Set("traffic_ip_groups", name, resource, &newTrafficIPGroup)
	if err != nil {
		t.Fatal("Error creating resource: ", err)
	}
	log.Println("Created traffic IP group ", name)

	assert.Equal(t, true, *newTrafficIPGroup.Properties.Basic.Enabled)
	assert.Equal(t, "go-brocade-vtm test traffic IP group", newTrafficIPGroup.Properties.Basic.Note)
	assert.Equal(t, true, *newTrafficIPGroup.Properties.Basic.HashSourcePort)
	assert.Equal(t, "alphabetic", newTrafficIPGroup.Properties.Basic.IPAssignmentMode)
	assert.Equal(t, []string{"10.0.34.5"}, newTrafficIPGroup.Properties.Basic.IPAddresses)
	assert.Equal(t, trafficManagers, newTrafficIPGroup.Properties.Basic.Machines)
	assert.Equal(t, "singlehosted", newTrafficIPGroup.Properties.Basic.Mode)
}

func getTrafficIPGroup(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	client.WorkWithConfigurationResources()

	trafficIPGroup := TrafficIPGroup{}
	err = client.GetByName("traffic_ip_groups", name, &trafficIPGroup)
	if err != nil {
		t.Fatal("Error getting resource: ", err)
	}
	log.Println("Found Traffic IP Group: ", trafficIPGroup)


	assert.Equal(t, true, *trafficIPGroup.Properties.Basic.Enabled)
	assert.Equal(t, "go-brocade-vtm test traffic IP group", trafficIPGroup.Properties.Basic.Note)
	assert.Equal(t, true, *trafficIPGroup.Properties.Basic.HashSourcePort)
	assert.Equal(t, "alphabetic", trafficIPGroup.Properties.Basic.IPAssignmentMode)
	assert.Equal(t, []string{"10.0.34.5"}, trafficIPGroup.Properties.Basic.IPAddresses)
	assert.Equal(t, getTrafficManagers(t), trafficIPGroup.Properties.Basic.Machines)
	assert.Equal(t, "singlehosted", trafficIPGroup.Properties.Basic.Mode)
}

func deleteTrafficIPGroup(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	err = client.Delete("traffic_ip_groups", name)
	if err != nil {
		t.Fatal("Error deletingresource: ", err)
	} else {
		log.Printf("Resource %s deleted", name)
	}
}

func getTrafficManagers(t *testing.T) []string {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	client.WorkWithConfigurationResources()

	trafficManagers := trafficManager.TrafficManagers{}
	err = client.GetByName("traffic_managers", "", &trafficManagers)
	if err != nil {
		t.Fatal("Error getting resource: ", err)
	}
	log.Println("Found Traffic Managers: ", trafficManagers)

	trafficManagerList := make([]string, 0)
	for _, trafficManager := range trafficManagers.Children {
		trafficManagerList = append(trafficManagerList, trafficManager.Name)
	}

	return trafficManagerList
}