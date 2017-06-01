package trafficIpGroups

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createTrafficIPGroupAPI *CreateUpdateTrafficIPGroupAPI
var updateTrafficIPGroupAPI *CreateUpdateTrafficIPGroupAPI

func setupCreateTrafficIPGroup() {
	tipgName := "test-tipg"
	tipgIPAddresses := []string{"172.0.0.1"}
	tipgEnabled := true
	tipgBasic := Basic{Enabled: &tipgEnabled, IPAddresses: tipgIPAddresses, Location: 0, Mode: "rhi", Note: "", RhiOspfv2MetricBase: 10, RhiOspfv2PassiveMetricOffset: 10}
	tipgProperties := Properties{tipgBasic}
	tipgTrafficIPGroup := TrafficIPGroup{tipgProperties}

	createTrafficIPGroupAPI = NewCreate(tipgName, tipgTrafficIPGroup)
}

func setupUpdateTrafficIPGroup() {
	tipgName := "test-tipg"
	tipgIPAddresses := []string{"172.0.0.1"}
	tipgEnabled := true
	tipgBasic := Basic{Enabled: &tipgEnabled, IPAddresses: tipgIPAddresses, Location: 0, Mode: "rhi", Note: "", RhiOspfv2MetricBase: 10, RhiOspfv2PassiveMetricOffset: 10}
	tipgProperties := Properties{tipgBasic}
	tipgTrafficIPGroup := TrafficIPGroup{tipgProperties}

	updateTrafficIPGroupAPI = NewUpdate(tipgName, tipgTrafficIPGroup)
}

func TestNewCreateMethod(t *testing.T) {
	setupCreateTrafficIPGroup()
	assert.Equal(t, http.MethodPut, createTrafficIPGroupAPI.Method())
}

func TestNewCreateEndpoint(t *testing.T) {
	setupCreateTrafficIPGroup()
	assert.Equal(t, "/api/tm/3.8/config/active/traffic_ip_groups/test-tipg", createTrafficIPGroupAPI.Endpoint())
}

func TestNewCreateMarshalling(t *testing.T) {
	setupCreateTrafficIPGroup()
	expectedJSON := "{\"properties\":{\"basic\":{\"enabled\":true,\"ipaddresses\":[\"172.0.0.1\"],\"mode\":\"rhi\",\"rhi_ospfv2_metric_base\":10,\"rhi_ospfv2_passive_metric_offset\":10}}}"
	jsonBytes, err := json.Marshal(createTrafficIPGroupAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}

func TestNewUpdateMethod(t *testing.T) {
	setupUpdateTrafficIPGroup()
	assert.Equal(t, http.MethodPut, updateTrafficIPGroupAPI.Method())
}

func TestNewUpdateEndpoint(t *testing.T) {
	setupCreateTrafficIPGroup()
	assert.Equal(t, "/api/tm/3.8/config/active/traffic_ip_groups/test-tipg", updateTrafficIPGroupAPI.Endpoint())
}

func TestNewUpdateMarshalling(t *testing.T) {
	setupUpdateTrafficIPGroup()
	expectedJSON := "{\"properties\":{\"basic\":{\"enabled\":true,\"ipaddresses\":[\"172.0.0.1\"],\"mode\":\"rhi\",\"rhi_ospfv2_metric_base\":10,\"rhi_ospfv2_passive_metric_offset\":10}}}"
	jsonBytes, err := json.Marshal(updateTrafficIPGroupAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}
