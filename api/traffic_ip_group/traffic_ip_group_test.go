package trafficIpGroups

import (
	"encoding/json"
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const trafficIPGroupTestEndpoint = "/api/tm/3.8/config/active/traffic_ip_groups"

var createTrafficIPGroupAPI, getAllTrafficIPGroupAPI, getTrafficIPGroupAPI, updateTrafficIPGroupAPI, deleteTrafficIPGroupAPI *rest.BaseAPI
var trafficIPGroupTestName, expectedMarshallTestJSON string
var testTrafficIPGroup TrafficIPGroup
var testTrafficIPGroupGetAllJSON, testTrafficIPGroupGetJSON []byte

func setupTrafficIPGroupTest() {
	trafficIPGroupTestName = "test-traffic-ip-group"
	trafficIPGroupAddresses := []string{"10.0.0.1"}
	trafficIPGroupEnable := true
	trafficIPGroupBasic := Basic{Enabled: &trafficIPGroupEnable, IPAddresses: trafficIPGroupAddresses, Location: 0, Mode: "rhi", Note: "", RhiOspfv2MetricBase: 10, RhiOspfv2PassiveMetricOffset: 10}
	testTrafficIPGroup.Properties = Properties{trafficIPGroupBasic}

	expectedMarshallTestJSON = `{"properties":{"basic":{"enabled":true,"ipaddresses":["10.0.0.1"],"mode":"rhi","rhi_ospfv2_metric_base":10,"rhi_ospfv2_passive_metric_offset":10}}}`

	createTrafficIPGroupAPI = NewCreate(trafficIPGroupTestName, testTrafficIPGroup)
	createTrafficIPGroupAPI.SetResponseObject(&testTrafficIPGroup)

	getAllTrafficIPGroupAPI = NewGetAll()
	testTrafficIPGroupGetAllJSON = []byte(`{"children":[{"name":"test-group-1","href":"/api/tm/3.8/config/active/traffic_ip_groups/test-group-1"},{"name":"test-group-2","href":"/api/tm/3.8/config/active/traffic_ip_groups/test-group-2"}]}`)

	getTrafficIPGroupAPI = NewGet(trafficIPGroupTestName)
	testTrafficIPGroupGetJSON = []byte(`{"properties":{"basic":{"enabled": true,"hash_source_port": true,"ip_assignment_mode": "alphabetic","ip_mapping":[{"ip": "172.0.0.5", "traffic_manager": "172.0.0.10" }],"ipaddresses":["172.0.0.1", "172.0.0.2"],"keeptogether": false,"location": 3,"machines":["172.0.0.8", "172.0.0.9"],"mode": "singlehosted","multicast": "224.1.2.3","note":"testnote","rhi_bgp_metric_base": 5,"rhi_bgp_passive_metric_offset": 5,"rhi_ospfv2_metric_base": 7,"rhi_ospfv2_passive_metric_offset": 5,"rhi_protocols": "ospf","slaves":[]}}}`)

	updateTrafficIPGroupAPI = NewUpdate(trafficIPGroupTestName, testTrafficIPGroup)
	updateTrafficIPGroupAPI.SetResponseObject(&testTrafficIPGroup)

	deleteTrafficIPGroupAPI = NewDelete(trafficIPGroupTestName)

}

func TestTrafficIPGroupCreateMethod(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, http.MethodPut, createTrafficIPGroupAPI.Method())
}

func TestTrafficIPGroupCreateEndpoint(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, trafficIPGroupTestEndpoint+"/"+trafficIPGroupTestName, createTrafficIPGroupAPI.Endpoint())
}

func TestTrafficIPGroupCreateMarshalling(t *testing.T) {
	setupTrafficIPGroupTest()
	jsonBytes, err := json.Marshal(createTrafficIPGroupAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedMarshallTestJSON, string(jsonBytes))
	assert.Equal(t, *createTrafficIPGroupAPI.ResponseObject().(*TrafficIPGroup), testTrafficIPGroup)
}

func TestTrafficIPGroupGetAllMethod(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, http.MethodGet, getAllTrafficIPGroupAPI.Method())
}

func TestTrafficIPGroupGetAllEndpoint(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, trafficIPGroupTestEndpoint, getAllTrafficIPGroupAPI.Endpoint())
}

func TestTrafficIPGroupGetAllUnmarshalling(t *testing.T) {
	setupTrafficIPGroupTest() //
	jsonErr := json.Unmarshal(testTrafficIPGroupGetAllJSON, getAllTrafficIPGroupAPI.ResponseObject())
	response := *getAllTrafficIPGroupAPI.ResponseObject().(*TrafficIPGroupList)

	assert.Nil(t, jsonErr)
	assert.Equal(t, 2, len(response.Children))
	assert.Equal(t, "test-group-1", response.Children[0].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/traffic_ip_groups/test-group-1", response.Children[0].HRef)
	assert.Equal(t, "test-group-2", response.Children[1].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/traffic_ip_groups/test-group-2", response.Children[1].HRef)
}

func TestTrafficIPGroupGetMethod(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, http.MethodGet, getTrafficIPGroupAPI.Method())
}

func TestTrafficIPGroupGetEndpoint(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, trafficIPGroupTestEndpoint+"/"+trafficIPGroupTestName, getTrafficIPGroupAPI.Endpoint())
}

func TestTrafficIPGroupGetUnmarshalling(t *testing.T) {
	setupTrafficIPGroupTest()
	jsonErr := json.Unmarshal(testTrafficIPGroupGetJSON, getTrafficIPGroupAPI.ResponseObject())
	response := *getTrafficIPGroupAPI.ResponseObject().(*TrafficIPGroup)

	assert.Nil(t, jsonErr)
	assert.Equal(t, true, *response.Properties.Basic.Enabled)
	assert.Equal(t, true, *response.Properties.Basic.HashSourcePort)
	assert.Equal(t, "alphabetic", response.Properties.Basic.IPAssignmentMode)
	assert.Equal(t, "172.0.0.5", response.Properties.Basic.IPMapping[0].IP)
	assert.Equal(t, "172.0.0.10", response.Properties.Basic.IPMapping[0].TrafficManager)
	assert.Equal(t, "172.0.0.1", response.Properties.Basic.IPAddresses[0])
	assert.Equal(t, "172.0.0.2", response.Properties.Basic.IPAddresses[1])
	assert.Equal(t, "172.0.0.8", response.Properties.Basic.Machines[0])
	assert.Equal(t, "172.0.0.9", response.Properties.Basic.Machines[1])
	assert.Equal(t, "singlehosted", response.Properties.Basic.Mode)
	assert.Equal(t, "224.1.2.3", response.Properties.Basic.Multicast)
}

func TestTrafficIPGroupUpdateMethod(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, http.MethodPut, updateTrafficIPGroupAPI.Method())
}

func TestTrafficIPGroupUpdateEndpoint(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, trafficIPGroupTestEndpoint+"/"+trafficIPGroupTestName, updateTrafficIPGroupAPI.Endpoint())
}

func TestTrafficIPGroupUpdateMarshalling(t *testing.T) {
	setupTrafficIPGroupTest()
	jsonBytes, err := json.Marshal(updateTrafficIPGroupAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedMarshallTestJSON, string(jsonBytes))
	assert.Equal(t, *updateTrafficIPGroupAPI.ResponseObject().(*TrafficIPGroup), testTrafficIPGroup)
}

func TestTrafficIPGroupDeleteMethod(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, http.MethodDelete, deleteTrafficIPGroupAPI.Method())
}

func TestTrafficIPGroupDeleteEndpoint(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, trafficIPGroupTestEndpoint+"/"+trafficIPGroupTestName, deleteTrafficIPGroupAPI.Endpoint())
}
