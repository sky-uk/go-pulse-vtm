package virtualserver

import (
	"encoding/json"
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const testVirtualServerEndpoint = "/api/tm/3.8/config/active/virtual_servers/"

var createVirtualServerAPI, getAllVirtualServerAPI, getVirtualServerAPI, updateVirtualServerAPI, deleteVirtualServerAPI *rest.BaseAPI
var testVirtualServerName string
var testVirtualServerObject VirtualServer
var testVirtualServerMarshalExpectedJSON string
var testVirtualServerGetAllUnmarshallJSON, testVirtualServerGetUnmarshalJSON []byte

func setupVirtualServerTest() {
	testVirtualServerName = "test-virtual-server"
	enabled := false
	var port uint = 80
	testVirtualServerBasic := Basic{
		Enabled:  &enabled,
		Pool:     "test-pool",
		Port:     &port,
		Protocol: "http",
	}
	testVirtualServerProperties := Properties{Basic: testVirtualServerBasic}
	testVirtualServerObject = VirtualServer{Properties: testVirtualServerProperties}

	testVirtualServerMarshalExpectedJSON = `{"properties":{"basic":{"enabled":false,"pool":"test-pool","port":80,"protocol":"http"},"aptimizer":{},"connection":{},"connection_errors":{},"cookie":{},"dns":{},"ftp":{},"gzip":{},"http":{},"http2":{},"kerberos_protocol_transition":{},"log":{},"recent_connections":{},"request_tracing":{},"rtsp":{},"sip":{},"smtp":{},"ssl":{},"syslog":{},"tcp":{},"udp":{},"web_cache":{}}}`

	testVirtualServerGetAllUnmarshallJSON = []byte(`{"children":[{"name":"PaaSExampleHTTPvirtualserver","href":"/api/tm/3.8/config/active/virtual_servers/PaaSExampleHTTPvirtualserver"},{"name":"PaaSExampleHTTPvirtualserver1","href":"/api/tm/3.8/config/active/virtual_servers/PaaSExampleHTTPvirtualserver1"},{"name":"virtual_server_1","href":"/api/tm/3.8/config/active/virtual_servers/virtual_server_1"},{"name":"virtual_server_2","href":"/api/tm/3.8/config/active/virtual_servers/virtual_server_2"},{"name":"virtual_server_3","href":"/api/tm/3.8/config/active/virtual_servers/virtual_server_3"}]}`)

	testVirtualServerGetUnmarshalJSON = []byte(`{"properties":{"basic":{"enabled":false,"pool":"pool_test","port":80,"protocol":"http"},"aptimizer":{},"connection":{},"connection_errors":{},"cookie":{},"dns":{},"ftp":{"ssl_data":false},"gzip":{},"ssl":{}}}`)

	createVirtualServerAPI = NewCreate(testVirtualServerName, testVirtualServerObject)
	createVirtualServerAPI.SetResponseObject(&testVirtualServerObject)

	getAllVirtualServerAPI = NewGetAll()

	getVirtualServerAPI = NewGet(testVirtualServerName)

	updateVirtualServerAPI = NewUpdate(testVirtualServerName, testVirtualServerObject)
	updateVirtualServerAPI.SetResponseObject(&testVirtualServerObject)

	deleteVirtualServerAPI = NewDelete(testVirtualServerName)
}

func TestVirtualServerNewCreateMethod(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, http.MethodPut, createVirtualServerAPI.Method())
}

func TestVirtualServerNewCreateEndpoint(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, testVirtualServerEndpoint+testVirtualServerName, createVirtualServerAPI.Endpoint())
}

func TestVirtualServerNewCreateMarshal(t *testing.T) {
	setupVirtualServerTest()
	jsonBytes, err := json.Marshal(createVirtualServerAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, testVirtualServerMarshalExpectedJSON, string(jsonBytes))
}

func TestVirtualServerNewGetAllMethod(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, http.MethodGet, getAllVirtualServerAPI.Method())
}

func TestVirtualServerNewGetAllEndpoint(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, testVirtualServerEndpoint, getAllVirtualServerAPI.Endpoint())
}

func TestVirtualServerNewGetAllUnmarshal(t *testing.T) {
	setupVirtualServerTest()
	jsonErr := json.Unmarshal(testVirtualServerGetAllUnmarshallJSON, getAllVirtualServerAPI.ResponseObject())
	response := *getAllVirtualServerAPI.ResponseObject().(*VirtualServersList)

	assert.Nil(t, jsonErr)
	assert.Len(t, response.Children, 5)
	assert.Equal(t, "PaaSExampleHTTPvirtualserver", response.Children[0].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/virtual_servers/PaaSExampleHTTPvirtualserver", response.Children[0].Href)
	assert.Equal(t, "virtual_server_3", response.Children[4].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/virtual_servers/virtual_server_3", response.Children[4].Href)
}

func TestVirtualServerNewGetMethod(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, http.MethodGet, getVirtualServerAPI.Method())
}

func TestVirtualServerNewGetEndpoint(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, testVirtualServerEndpoint+testVirtualServerName, getVirtualServerAPI.Endpoint())
}

func TestVirtualServerNewGetUnmarshal(t *testing.T) {
	setupVirtualServerTest()
	jsonErr := json.Unmarshal(testVirtualServerGetUnmarshalJSON, getVirtualServerAPI.ResponseObject())
	response := *getVirtualServerAPI.ResponseObject().(*VirtualServer)
	expectedEnabled := false
	var expectedPort uint = 80
	assert.Nil(t, jsonErr)
	assert.Equal(t, &expectedEnabled, response.Properties.Basic.Enabled)
	assert.Equal(t, "pool_test", response.Properties.Basic.Pool)
	assert.Equal(t, &expectedPort, response.Properties.Basic.Port)
	assert.Equal(t, "http", response.Properties.Basic.Protocol)
}

func TestVirtualServerNewUpdateMethod(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, http.MethodPut, updateVirtualServerAPI.Method())
}

func TestVirtualServerNewUpdateEndpoint(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, testVirtualServerEndpoint+testVirtualServerName, updateVirtualServerAPI.Endpoint())
}

func TestVirtualServerNewUpdateMarshal(t *testing.T) {
	setupVirtualServerTest()
	jsonBytes, err := json.Marshal(updateVirtualServerAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, testVirtualServerMarshalExpectedJSON, string(jsonBytes))
}

func TestVirtualServerNewDeleteMethod(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, http.MethodDelete, deleteVirtualServerAPI.Method())
}

func TestVirtualServerNewDeleteEndpoint(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, testVirtualServerEndpoint+testVirtualServerName, deleteVirtualServerAPI.Endpoint())
}
