package monitor

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createMonitorAPI *CreateMonitorAPI
var createMonitorName = "exampleMonitor"
var createMonitor Monitor

func createSetup() {

	newHTTPMonitor := HTTP{URIPath: "/download/private/status/check"}
	monitorVerbosity := true
	newBasicMonitor := Basic{
		Delay:    6,
		Failures: 3,
		Type:     "http",
		Timeout:  4,
		Verbose:  &monitorVerbosity,
	}
	createMonitorProperties := Properties{Basic: newBasicMonitor, HTTP: newHTTPMonitor}
	createMonitor = Monitor{Properties: createMonitorProperties}

	createMonitorAPI = NewCreate(createMonitorName, createMonitor)
	createMonitorAPI.SetResponseObject(&createMonitor)
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPut, createMonitorAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/tm/3.8/config/active/monitors/"+createMonitorName, createMonitorAPI.Endpoint())
}

func TestCreateRequestMarshalling(t *testing.T) {
	createSetup()
	expectedJSON := "{\"properties\":{\"basic\":{\"delay\":6,\"failures\":3,\"type\":\"http\",\"timeout\":4,\"verbose\":true},\"http\":{\"path\":\"/download/private/status/check\"}}}"
	jsonBytes, err := json.Marshal(createMonitorAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}

func TestGetResponse(t *testing.T) {
	createSetup()
	monsList := createMonitorAPI.GetResponse()
	assert.Equal(t, monsList, createMonitor)
}
