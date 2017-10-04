package dnsZoneFile

import (
	"errors"
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

var name = api.SetTestResourceName("dns_zone_file_")
var zoneFileTemplate = []byte(`
 $TTL 3600
@                       30  IN  SOA h1ist01-v00.paas.d50.ovp.bskyb.com. hostmaster.isp.sky.com. (
                                    2017092901 ; serial
                                    3600       ; refresh after 1 hour
                                    300        ; retry after 5 minutes
                                    1209600    ; expire after 2 weeks
                                    30 )       ; minimum TTL of 30 seconds
; We may have more than one NS here.
@                       30  IN  NS  h1ist01-v00.paas.d50.ovp.bskyb.com.
h1ist01-v00                     30  IN  A   10.77.13.9
;
; Services - Each service in a location has a unique IP address. Two locations = two IPs.
;
example-service                 60  IN  A   10.100.10.5
                        60  IN  A   10.100.20.5
another-example-service             60  IN  A   10.100.10.6
                        60  IN  A   10.100.20.6
`)

func TestAll(t *testing.T) {
	testSetDNSZoneFile(t)
	testGetDNSZoneFile(t)
	testDeleteDNSZoneFile(t)
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

func testSetDNSZoneFile(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}

	// in case of file uploading no response body back..
	err = client.Set("dns_server/zone_files", name, zoneFileTemplate, nil)
	if err != nil {
		t.Fatal("Error creating a dns zone file ", err)
	}
	log.Println("Created rule ", name)
}

func testGetDNSZoneFile(t *testing.T) {
	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	client.WorkWithConfigurationResources()

	out := new([]byte)
	err = client.GetByName("dns_server/zone_files", name, out)
	if err != nil {
		t.Fatal("Error creating a resource: ", err)
	}
	assert.Equal(t, zoneFileTemplate, *out)
}

func testDeleteDNSZoneFile(t *testing.T) {
	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	err = client.Delete("dns_server/zone_files", name)
	if err != nil {
		t.Fatal("Error deleting a resource: ", err)
	} else {
		log.Printf("Resource %s deleted", name)
	}
}
