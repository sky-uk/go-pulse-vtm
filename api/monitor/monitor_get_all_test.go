package monitor

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllMonitorAPI *GetAllMonitors

func setupGetAll() {
	getAllMonitorAPI = NewGetAll()
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllMonitorAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/tm/3.8/config/active/monitors", getAllMonitorAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	jsonContent := []byte("{\"children\":[{\"name\":\"MonitorOne\",\"href\":\"/api/tm/3.8/config/active/monitors/MonitorOne\"},{\"name\":\"MonitorTwo\",\"href\":\"/api/tm/3.8/config/active/monitors/MonitorTwo\"}]}")
	jsonErr := json.Unmarshal(jsonContent, getAllMonitorAPI.ResponseObject())

	assert.Nil(t, jsonErr)
	assert.Len(t, getAllMonitorAPI.GetResponse().Children, 2)
	assert.Equal(t, "MonitorOne", getAllMonitorAPI.GetResponse().Children[0].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/monitors/MonitorOne", getAllMonitorAPI.GetResponse().Children[0].HRef)
	assert.Equal(t, "MonitorTwo", getAllMonitorAPI.GetResponse().Children[1].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/monitors/MonitorTwo", getAllMonitorAPI.GetResponse().Children[1].HRef)
}
