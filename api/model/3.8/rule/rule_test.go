package rule

import (
	"errors"
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

var name = api.SetTestResourceName("rule_")
var ruleTemplate = `if( string.ipmaskmatch( request.getremoteip(), "192.168.11.13" ) ){
    connection.discard();
} `

func TestAll(t *testing.T) {
    testSetRule(t)
    testGetRule(t)
    testDeleteRule(t)
}

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

	headers := make(map[string]string)
	headers["Content-Type"] = "application/octet-stream"
	headers["Content-Transfer-Encoding"] = "text"

	params := api.Params{
		APIVersion: "3.8",
		Server:     server,
		Username:   username,
		Password:   password,
		IgnoreSSL:  true,
		Debug:      true,
		Headers:    headers,
	}

	return api.Connect(params)
}

func testSetRule(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}

	// in case of file uploading no response body back..
	err = client.Set("rules", name, []byte(ruleTemplate), nil)
	if err != nil {
		t.Fatal("Error creating a rule ", err)
	}
	log.Println("Created rule ", name)
}

func testGetRule(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	client.WorkWithConfigurationResources()

	out := new([]byte)
	err = client.GetByName("rules", name, out)
	if err != nil {
		t.Fatal("Error creating a resource: ", err)
	}
	assert.Equal(t, []byte(ruleTemplate), *out)
}

func testDeleteRule(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	err = client.Delete("rules", name)
	if err != nil {
		t.Fatal("Error deleting a resource: ", err)
	} else {
		log.Printf("Resource %s deleted", name)
	}
}
