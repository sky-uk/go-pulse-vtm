package location

import (
	//"encoding/json"
	"errors"
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func getClient() (*api.Client, error) {

	server, ok := os.LookupEnv("BROCADEVTM_SERVER")
	if ok == false || server == "" {
		return nil, errors.New("BROCADEVTM_SERVER env var not set")
	}

	username, ok := os.LookupEnv("BROCADEVTM_USERNAME")
	if ok == false {
		return nil, errors.New("BROCADEVTM_USERNAME env var not set")
	}

	password, ok := os.LookupEnv("BROCADEVTM_PASSWORD")
	if ok == false {
		return nil, errors.New("BROCADEVTM_PASSWORD env var not set")
	}

	params := api.Params{
		//APIVersion: "3.8",
		Server:    server,
		Username:  username,
		Password:  password,
		IgnoreSSL: true,
		Debug:     true,
	}

	return api.Connect(params)

}

var name = "glb_" + strconv.Itoa(rand.Int())

func TestSetLocation(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	resource := Location{}
	resource.Properties.Basic = Basic{
		ID:        32001,
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
	client, err := getClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}

	location := Location{}
	err = client.GetByName("locations", name, &location)
	if err != nil {
		t.Fatal("Error creating a resource: ", err)
	}
	log.Println("Found Location: ", location)
	assert.Equal(t, "test location", location.Properties.Basic.Note)
}

func TestDeleteLocation(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
    err = client.Delete("locations", name)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Resource %s deleted", name)
	}

}
