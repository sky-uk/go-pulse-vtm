package monitor

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateMonitorAPI *UpdateMonitorAPI
var updateMonitorName = "updateMonitor"

func setup() {
	newHTTPMonitor := HTTP{URIPath: "/private/status/check"}
	newBasicMonitor := Basic{Delay: 6, Failures: 3, Type: "http", Timeout: 4}
	newMonitorProperties := Properties{Basic: newBasicMonitor, HTTP: newHTTPMonitor}
	newMonitor := Monitor{Properties: newMonitorProperties}

	updateMonitorAPI = NewUpdate(updateMonitorName, newMonitor)
	updateMonitorAPI.SetResponseObject("/private/status/check")
}

func TestUpdateMethod(t *testing.T) {
	setup()
	assert.Equal(t, http.MethodPut, updateMonitorAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	setup()
	assert.Equal(t, "/api/tm/3.8/config/active/monitors/"+updateMonitorName, updateMonitorAPI.Endpoint())
}

func TestUpdateMarshalling(t *testing.T) {
	setup()
	expectedJSON := "{\"properties\":{\"basic\":{\"delay\":6,\"failures\":3,\"type\":\"http\",\"timeout\":4},\"http\":{\"path\":\"/private/status/check\"}}}"
	jsonBytes, err := json.Marshal(updateMonitorAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}

func TestUpdateGetResponse(t *testing.T) {
	setup()
	getResponse := updateMonitorAPI.GetResponse()
	assert.Equal(t, getResponse, "/private/status/check")

}
