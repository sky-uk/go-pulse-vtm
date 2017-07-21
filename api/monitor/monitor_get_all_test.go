package monitor

import (
	"encoding/json"
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllMonitorAPI *rest.BaseAPI

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

	response := getAllMonitorAPI.ResponseObject().(*MonitorsList)
	assert.Len(t, response.Children, 2)
	assert.Equal(t, "MonitorOne", response.Children[0].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/monitors/MonitorOne", response.Children[0].HRef)
	assert.Equal(t, "MonitorTwo", response.Children[1].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/monitors/MonitorTwo", response.Children[1].HRef)
}
