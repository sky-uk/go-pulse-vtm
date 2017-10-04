package dnsZone

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestDNSZone(t *testing.T) {

	name := api.SetTestResourceName("dns_zone_")
	setDNSZone(name, t)
	getDNSZone(name, t)
	deleteDNSZone(name, t)
}

func setDNSZone(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}

	resource := DNSZone{}
	resource.Properties.Basic.Origin = "test-example.com"
	resource.Properties.Basic.ZoneFile = "test-example.com.db"

	newDNSZone := DNSZone{}
	err = client.Set("dns_server/zones", name, resource, &newDNSZone)
	if err != nil {
		t.Fatal("Error creating resource: ", err)
	}
	log.Println("Created DNS Zone ", name)

	assert.Equal(t, "test-example.com", newDNSZone.Properties.Basic.Origin)
	assert.Equal(t, "test-example.com.db", newDNSZone.Properties.Basic.ZoneFile)
}

func getDNSZone(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	client.WorkWithConfigurationResources()

	dnsZone := DNSZone{}
	err = client.GetByName("dns_server/zones", name, &dnsZone)
	if err != nil {
		t.Fatal("Error getting resource: ", err)
	}
	log.Println("Found DNS Zone: ", dnsZone)

	assert.Equal(t, "test-example.com", dnsZone.Properties.Basic.Origin)
	assert.Equal(t, "test-example.com.db", dnsZone.Properties.Basic.ZoneFile)
}

func deleteDNSZone(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	err = client.Delete("dns_server/zones", name)
	if err != nil {
		t.Fatal("Error deleting resource: ", err)
	} else {
		log.Printf("Resource %s deleted", name)
	}

}
