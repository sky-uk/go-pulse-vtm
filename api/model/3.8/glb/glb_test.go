package glb

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/stretchr/testify/assert"
	"log"
	"sort"
	"testing"
)

func TestGLB(t *testing.T) {

	name := api.SetTestResourceName("glb_")
	setGLB(name, t)
	getGLB(name, t)
	deleteGLB(name, t)
}

// setGLB : create/update GLB
func setGLB(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}

	domains := []string{"example.com", "another-example.com"}
	sort.Strings(domains)

	sslKeys := []string{"another-example.com", "example.com"}
	dnsSecKey := DNSSecKey{"example.com", sslKeys}
	dnsSecKeys := make([]DNSSecKey, 0)
	dnsSecKeys = append(dnsSecKeys, dnsSecKey)

	locationOneMonitors := []string{"glb-example-monitor", "glb-example-monitor2"}
	locationTwoMonitors := []string{"glb-example-monitor"}

	locationOneIPS := []string{"192.168.234.56", "192.0.2.2"}
	locationTwoIPS := []string{"192.168.17.56", "192.168.8.22"}

	locationOneSettings := LocationSetting{Location: "example-location-one", Weight: 34, IPS: locationOneIPS, Monitors: locationOneMonitors}
	locationTwoSettings := LocationSetting{Location: "example-location-two", Weight: 66, IPS: locationTwoIPS, Monitors: locationTwoMonitors}
	locationSettingsList := make([]LocationSetting, 0)
	locationSettingsList = append(locationSettingsList, locationOneSettings)
	locationSettingsList = append(locationSettingsList, locationTwoSettings)

	resource := GLB{}
	resource.Properties.Basic = Basic{
		Algorithm:            "weighted_random",
		AllMonitorsNeeded:    true,
		AutoRecovery:         true,
		ChainedAutoFailback:  true,
		ChainedLocationOrder: []string{"example-location-one", "example-location-two"},
		DisableOnFailure:     true,
		DNSSecKeys:           dnsSecKeys,
		Domains:              domains,
		Enabled:              true,
		GeoEffect:            10,
		LastResortResponse:   []string{"192.168.12.10", "192.168.120.10"},
		LocationDraining:     []string{"example-location-one"},
		LocationSettings:     locationSettingsList,
		ReturnIPSOnFail:      true,
		Rules:                []string{"ruleOne", "ruleTwo"},
		TTL:                  30,
	}
	resource.Properties.Log = Log{
		Enabled:  true,
		Filename: "/var/log/brocadevtm/test.log",
		Format:   "%g, %n, %d, %a, %t, %s, %l, %q",
	}

	newGLB := GLB{}
	err = client.Set("glb_services", name, resource, &newGLB)
	if err != nil {
		t.Fatal("Error creating a resource: ", err)
	}
	log.Println("Created GLB ", name)

	assert.Equal(t, "weighted_random", newGLB.Properties.Basic.Algorithm)
	assert.Equal(t, true, newGLB.Properties.Basic.AllMonitorsNeeded)
	assert.Equal(t, true, newGLB.Properties.Basic.AutoRecovery)
	assert.Equal(t, true, newGLB.Properties.Basic.ChainedAutoFailback)
	assert.Equal(t, []string{"example-location-one", "example-location-two"}, newGLB.Properties.Basic.ChainedLocationOrder)
	assert.Equal(t, true, newGLB.Properties.Basic.DisableOnFailure)
	assert.Equal(t, 1, len(newGLB.Properties.Basic.DNSSecKeys))
	assert.Equal(t, dnsSecKeys[0].Domain, newGLB.Properties.Basic.DNSSecKeys[0].Domain)
	assert.Equal(t, 2, len(newGLB.Properties.Basic.DNSSecKeys[0].SSLKeys))
	assert.Equal(t, dnsSecKeys[0].SSLKeys[0], newGLB.Properties.Basic.DNSSecKeys[0].SSLKeys[0])
	assert.Equal(t, dnsSecKeys[0].SSLKeys[1], newGLB.Properties.Basic.DNSSecKeys[0].SSLKeys[1])
	assert.Equal(t, domains, newGLB.Properties.Basic.Domains)
	assert.Equal(t, true, newGLB.Properties.Basic.Enabled)
	assert.Equal(t, uint(10), newGLB.Properties.Basic.GeoEffect)
	assert.Equal(t, []string{"192.168.12.10", "192.168.120.10"}, newGLB.Properties.Basic.LastResortResponse)
	assert.Equal(t, []string{"example-location-one"}, newGLB.Properties.Basic.LocationDraining)
	assert.Equal(t, 2, len(newGLB.Properties.Basic.LocationSettings))
	assert.Equal(t, true, newGLB.Properties.Basic.ReturnIPSOnFail)
	assert.Equal(t, []string{"ruleOne", "ruleTwo"}, newGLB.Properties.Basic.Rules)
	assert.Equal(t, true, newGLB.Properties.Log.Enabled)
	assert.Equal(t, "/var/log/brocadevtm/test.log", newGLB.Properties.Log.Filename)
	assert.Equal(t, "%g, %n, %d, %a, %t, %s, %l, %q", newGLB.Properties.Log.Format)
}

// getGLB : retrieve a GLB
func getGLB(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	client.WorkWithConfigurationResources()

	glb := GLB{}
	err = client.GetByName("glb_services", name, &glb)
	if err != nil {
		t.Fatal("Error creating a resource: ", err)
	}
	log.Println("Found GLB: ", glb)

	domains := []string{"example.com", "another-example.com"}
	sort.Strings(domains)

	sslKeys := []string{"another-example.com", "example.com"}
	dnsSecKey := DNSSecKey{"example.com", sslKeys}
	dnsSecKeys := make([]DNSSecKey, 0)
	dnsSecKeys = append(dnsSecKeys, dnsSecKey)

	assert.Equal(t, "weighted_random", glb.Properties.Basic.Algorithm)
	assert.Equal(t, true, glb.Properties.Basic.AllMonitorsNeeded)
	assert.Equal(t, true, glb.Properties.Basic.AutoRecovery)
	assert.Equal(t, true, glb.Properties.Basic.ChainedAutoFailback)
	assert.Equal(t, []string{"example-location-one", "example-location-two"}, glb.Properties.Basic.ChainedLocationOrder)
	assert.Equal(t, true, glb.Properties.Basic.DisableOnFailure)
	assert.Equal(t, 1, len(glb.Properties.Basic.DNSSecKeys))
	assert.Equal(t, dnsSecKeys[0].Domain, glb.Properties.Basic.DNSSecKeys[0].Domain)
	assert.Equal(t, 2, len(glb.Properties.Basic.DNSSecKeys[0].SSLKeys))
	assert.Equal(t, dnsSecKeys[0].SSLKeys[0], glb.Properties.Basic.DNSSecKeys[0].SSLKeys[0])
	assert.Equal(t, dnsSecKeys[0].SSLKeys[1], glb.Properties.Basic.DNSSecKeys[0].SSLKeys[1])
	assert.Equal(t, domains, glb.Properties.Basic.Domains)
	assert.Equal(t, true, glb.Properties.Basic.Enabled)
	assert.Equal(t, uint(10), glb.Properties.Basic.GeoEffect)
	assert.Equal(t, []string{"192.168.12.10", "192.168.120.10"}, glb.Properties.Basic.LastResortResponse)
	assert.Equal(t, []string{"example-location-one"}, glb.Properties.Basic.LocationDraining)
	assert.Equal(t, 2, len(glb.Properties.Basic.LocationSettings))
	assert.Equal(t, true, glb.Properties.Basic.ReturnIPSOnFail)
	assert.Equal(t, []string{"ruleOne", "ruleTwo"}, glb.Properties.Basic.Rules)
	assert.Equal(t, true, glb.Properties.Log.Enabled)
	assert.Equal(t, "/var/log/brocadevtm/test.log", glb.Properties.Log.Filename)
	assert.Equal(t, "%g, %n, %d, %a, %t, %s, %l, %q", glb.Properties.Log.Format)
}

// deleteGLB : delete a GLB
func deleteGLB(name string, t *testing.T) {

	client, err := api.GetClient()
	if err != nil {
		t.Fatal("Connection error: ", err)
	}
	err = client.Delete("glb_services", name)
	if err != nil {
		t.Fatal("Error deleting a resource: ", err)
	} else {
		log.Printf("Resource %s deleted", name)
	}
}
