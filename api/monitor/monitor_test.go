package monitor

import (
	"encoding/json"
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const testMonitorEndpoint = "/api/tm/3.8/config/active/monitors/"

var createMonitorAPI, updateMonitorAPI, getAllMonitorAPI, getMonitorAPI, deleteMonitorAPI *rest.BaseAPI
var testMonitor Monitor
var testMonitorName, marshallingTestExpectedJSON string
var getAllUnmarshallingTestJSON, getUnmarshallingTestJSON []byte

//var SetArguments map[string]string

func setupMonitorTest() {
	testHTTPMonitor := HTTP{URIPath: "/my-app/healthcheck"}
	monitorVerbosity := true
	argument1 := NewArgumentIssue("test1", "paas test1", "dns.pl")
	argument2 := NewArgumentIssue("test2", "paas test2", "snmp.pl")
	monitorUDPAcceptAll := false
	testBasicMonitor := Basic{
		Delay:    6,
		Failures: 3,
		Type:     "http",
		Timeout:  4,
		Verbose:  &monitorVerbosity,
	}
	testRTSPMonitor := RTSP{
		BodyRegex:   "",
		URIPath:     "/",
		StatusRegex: "^[234][0-9][0-9]$",
	}
	testSCRIPTMonitor := SCRIPT{
		Arguments: []ArgumentIssue{argument1, argument2},
		Program:   "",
	}
	testSIPMonitor := SIP{
		BodyRegex:   "",
		StatusRegex: "^[234][0-9][0-9]$",
		Transport:   "udp",
	}
	testTCPMonitor := TCP{
		CloseString:    "",
		MaxResponseLen: 2048,
		ResponseRegex:  ".+",
		WriteString:    "",
	}
	testUDPMonitor := UDP{
		AcceptAll: &monitorUDPAcceptAll,
	}
	testMonitorProperties := Properties{Basic: testBasicMonitor, HTTP: testHTTPMonitor, RTSP: testRTSPMonitor, SCRIPT: testSCRIPTMonitor, SIP: testSIPMonitor, TCP: testTCPMonitor, UDP: testUDPMonitor}
	testMonitor = Monitor{Properties: testMonitorProperties}

	marshallingTestExpectedJSON = `{"properties":{"basic":{"delay":6,"failures":3,"type":"http","timeout":4,"verbose":true},"http":{"path":"/my-app/healthcheck"},"rtsp":{"path":"/","status_regex":"^[234][0-9][0-9]$"},"script":{"arguments":[{"name":"test1","description":"paas test1","value":"dns.pl"},{"name":"test2","description":"paas test2","value":"snmp.pl"}]},"sip":{"status_regex":"^[234][0-9][0-9]$","transport":"udp"},"tcp":{"max_response_len":2048,"response_regex":".+"},"udp":{"accept_all":false}}}`
	getAllUnmarshallingTestJSON = []byte(`{"children":[{"name":"MonitorOne","href":"/api/tm/3.8/config/active/monitors/MonitorOne"},{"name":"MonitorTwo","href":"/api/tm/3.8/config/active/monitors/MonitorTwo"}]}`)
	getUnmarshallingTestJSON = []byte(`{"properties":{"basic":{"delay":12,"failures":2,"type":"http","timeout":7,"verbose":true},"http":{"path":"/my-other-app/healthcheck"},"rtsp":{},"script":{},"sip":{},"tcp":{},"udp":{}}}`)

	createMonitorAPI = NewCreate(testMonitorName, testMonitor)
	createMonitorAPI.SetResponseObject(&testMonitor)

	getMonitorAPI = NewGet(testMonitorName)

	getAllMonitorAPI = NewGetAll()

	updateMonitorAPI = NewUpdate(testMonitorName, testMonitor)
	updateMonitorAPI.SetResponseObject(&testMonitor)

	deleteMonitorAPI = NewDelete(testMonitorName)
}

func TestNewCreateMethod(t *testing.T) {
	setupMonitorTest()
	assert.Equal(t, http.MethodPut, createMonitorAPI.Method())
}

func TestNewCreateEndpoint(t *testing.T) {
	setupMonitorTest()
	assert.Equal(t, testMonitorEndpoint+testMonitorName, createMonitorAPI.Endpoint())
}

func TestNewCreateRequestMarshalling(t *testing.T) {
	setupMonitorTest()
	jsonBytes, err := json.Marshal(createMonitorAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, marshallingTestExpectedJSON, string(jsonBytes))
}

func TestNewUpdateMethod(t *testing.T) {
	setupMonitorTest()
	assert.Equal(t, http.MethodPut, createMonitorAPI.Method())
}

func TestNewUpdateEndpoint(t *testing.T) {
	setupMonitorTest()
	assert.Equal(t, testMonitorEndpoint+testMonitorName, updateMonitorAPI.Endpoint())
}

func TestNewUpdateRequestMarshalling(t *testing.T) {
	setupMonitorTest()
	jsonBytes, err := json.Marshal(updateMonitorAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, marshallingTestExpectedJSON, string(jsonBytes))
}

func TestNewGetAllMethod(t *testing.T) {
	setupMonitorTest()
	assert.Equal(t, http.MethodGet, getAllMonitorAPI.Method())
}

func TestNewGetAllEndpoint(t *testing.T) {
	setupMonitorTest()
	assert.Equal(t, testMonitorEndpoint, getAllMonitorAPI.Endpoint())
}

func TestNewGetAllUnmarshalling(t *testing.T) {
	setupMonitorTest()
	jsonErr := json.Unmarshal(getAllUnmarshallingTestJSON, getAllMonitorAPI.ResponseObject())
	response := *getAllMonitorAPI.ResponseObject().(*MonitorsList)

	assert.Nil(t, jsonErr)
	assert.Len(t, response.Children, 2)
	assert.Equal(t, "MonitorOne", response.Children[0].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/monitors/MonitorOne", response.Children[0].HRef)
	assert.Equal(t, "MonitorTwo", response.Children[1].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/monitors/MonitorTwo", response.Children[1].HRef)
}

func TestNewGetMethod(t *testing.T) {
	setupMonitorTest()
	assert.Equal(t, http.MethodGet, getMonitorAPI.Method())
}

func TestNewGetEndpoint(t *testing.T) {
	setupMonitorTest()
	assert.Equal(t, testMonitorEndpoint+testMonitorName, getMonitorAPI.Endpoint())
}

func TestNewGetUnmarshalling(t *testing.T) {
	setupMonitorTest()
	jsonErr := json.Unmarshal(getUnmarshallingTestJSON, getMonitorAPI.ResponseObject())
	response := *getMonitorAPI.ResponseObject().(*Monitor)

	assert.Nil(t, jsonErr)
	assert.Equal(t, uint(12), response.Properties.Basic.Delay)
	assert.Equal(t, uint(2), response.Properties.Basic.Failures)
	assert.Equal(t, "http", response.Properties.Basic.Type)
	assert.Equal(t, uint(7), response.Properties.Basic.Timeout)
	assert.Equal(t, true, *response.Properties.Basic.Verbose)
	assert.Equal(t, "/my-other-app/healthcheck", response.Properties.HTTP.URIPath)
}

func TestNewDeleteMethod(t *testing.T) {
	setupMonitorTest()
	assert.Equal(t, http.MethodDelete, deleteMonitorAPI.Method())
}

func TestNewDeleteEndpoint(t *testing.T) {
	setupMonitorTest()
	assert.Equal(t, testMonitorEndpoint+testMonitorName, deleteMonitorAPI.Endpoint())
}
