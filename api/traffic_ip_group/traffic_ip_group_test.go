package trafficIpGroups

import (
	"github.com/sky-uk/go-rest-api"
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
)

const trafficIPGroupTestEndpoint = "/api/tm/3.8/config/active/traffic_ip_groups"

var createTrafficIPGroupAPI, getAllTrafficIPGroupAPI, getTrafficIPGroupAPI, updateTrafficIPGroupAPI, deleteTrafficIPGroupAPI *rest.BaseAPI
var trafficIPGroupTestName string
var testTrafficIPGroup TrafficIPGroup

func setupTrafficIPGroupTest() {
	trafficIPGroupTestName = "test-traffic-ip-group"
	trafficIPGroupAddresses := []string{"10.0.0.1", "10.0.0.2"}
	trafficIPGroupEnable := true
	trafficIPGroupBasic := Basic{Enabled: &trafficIPGroupEnable, IPAddresses: trafficIPGroupAddresses, Location: 0, Mode: "rhi", Note: "", RhiOspfv2MetricBase: 10, RhiOspfv2PassiveMetricOffset: 10}
	testTrafficIPGroup.Properties = Properties{trafficIPGroupBasic}


	createTrafficIPGroupAPI = NewCreate(trafficIPGroupTestName, testTrafficIPGroup)
	createTrafficIPGroupAPI.SetResponseObject(&testTrafficIPGroup)

	getAllTrafficIPGroupAPI = NewGetAll()

	getTrafficIPGroupAPI = NewGet(trafficIPGroupTestName)

	updateTrafficIPGroupAPI = NewUpdate(trafficIPGroupTestName, testTrafficIPGroup)

	deleteTrafficIPGroupAPI = NewDelete(trafficIPGroupTestName)

}

func TestTrafficIPGroupCreateMethod(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, http.MethodPut, createTrafficIPGroupAPI.Method())
}

func TestTrafficIPGroupCreateEndpoint(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, trafficIPGroupTestEndpoint + "/" + trafficIPGroupTestName, createTrafficIPGroupAPI.Endpoint())
}

func TestTrafficIPGroupCreateMarshalling(t *testing.T) {
	setupTrafficIPGroupTest()
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
	setupTrafficIPGroupTest()
}

func TestTrafficIPGroupGetMethod(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, http.MethodGet, getTrafficIPGroupAPI.Method())
}

func TestTrafficIPGroupGetEndpoint(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, trafficIPGroupTestEndpoint + "/" + trafficIPGroupTestName, getTrafficIPGroupAPI.Endpoint())
}

func TestTrafficIPGroupGetUnmarshalling(t *testing.T) {
	setupTrafficIPGroupTest()
}

func TestTrafficIPGroupUpdateMethod(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, http.MethodPut, updateTrafficIPGroupAPI.Method())
}

func TestTrafficIPGroupUpdateEndpoint(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, trafficIPGroupTestEndpoint + "/" + trafficIPGroupTestName, updateTrafficIPGroupAPI.Endpoint())
}

func TestTrafficIPGroupUpdateMarshalling(t *testing.T) {
	setupTrafficIPGroupTest()
}

func TestTrafficIPGroupDeleteMethod(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, http.MethodDelete, deleteTrafficIPGroupAPI.Method())
}

func TestTrafficIPGroupDeleteEndpoint(t *testing.T) {
	setupTrafficIPGroupTest()
	assert.Equal(t, trafficIPGroupTestEndpoint + "/" + trafficIPGroupTestName, deleteTrafficIPGroupAPI.Endpoint())
}


