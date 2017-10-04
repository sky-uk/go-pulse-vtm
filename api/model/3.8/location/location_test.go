package location

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"strconv"
	"testing"
)

var name = "glb_" + strconv.Itoa(rand.Int())

func TestSetLocation(t *testing.T) {
	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	resource := Location{}
	resource.Properties.Basic = Basic{
		ID:        32002,
		Latitude:  -36.353417,
		Longitude: 146.687568,
		Note:      "test location",
		Type:      "config",
	}

	newLocation := Location{}
	err = client.Set("locations", name, resource, &newLocation)
	if err != nil {
		t.Fatal("Error creating a resource: ", err)
	}
	log.Println("Created location ", name)
	assert.Equal(t, "test location", newLocation.Properties.Basic.Note)
}

func TestGetLocation(t *testing.T) {
	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	client.WorkWithConfigurationResources()

	location := Location{}
	err = client.GetByName("locations", name, &location)
	if err != nil {
		t.Fatal("Error getting a resource: ", err)
	}
	log.Println("Found Location: ", location)
	assert.Equal(t, "test location", location.Properties.Basic.Note)
}

func TestDeleteLocation(t *testing.T) {
	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	err = client.Delete("locations", name)
	if err != nil {
		t.Fatal("Error deleting a resource: ", err)
	} else {
		log.Printf("Resource %s deleted", name)
	}
}
