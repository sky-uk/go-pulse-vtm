package virtualserver

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllVirtualServerAPI *GetAllVirtualServers

func setupGetAll() {
	getAllVirtualServerAPI = NewGetAll()
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllVirtualServerAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/tm/3.8/config/active/virtual_servers", getAllVirtualServerAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()

	jsonContent := `{"children":[{"name":"PaaSExampleHTTPvirtualserver","href":"/api/tm/3.8/config/active/virtual_servers/PaaSExampleHTTPvirtualserver"},{"name":"PaaSExampleHTTPvirtualserver1","href":"/api/tm/3.8/config/active/virtual_servers/PaaSExampleHTTPvirtualserver1"},{"name":"virtual_server_1","href":"/api/tm/3.8/config/active/virtual_servers/virtual_server_1"},{"name":"virtual_server_2","href":"/api/tm/3.8/config/active/virtual_servers/virtual_server_2"},{"name":"virtual_server_3","href":"/api/tm/3.8/config/active/virtual_servers/virtual_server_3"}]}`

	jsonErr := json.Unmarshal([]byte(jsonContent), getAllVirtualServerAPI.ResponseObject())
	fmt.Println(getAllVirtualServerAPI.GetResponse().Children[0].Name)
	assert.Nil(t, jsonErr)
	assert.Len(t, getAllVirtualServerAPI.GetResponse().Children, 5)
	assert.Equal(t, "PaaSExampleHTTPvirtualserver", getAllVirtualServerAPI.GetResponse().Children[0].Name)
	assert.Equal(
		t,
		"/api/tm/3.8/config/active/virtual_servers/PaaSExampleHTTPvirtualserver",
		getAllVirtualServerAPI.GetResponse().Children[0].Href,
	)
	assert.Equal(
		t,
		"PaaSExampleHTTPvirtualserver1",
		getAllVirtualServerAPI.GetResponse().Children[1].Name,
	)
	assert.Equal(
		t,
		"/api/tm/3.8/config/active/virtual_servers/PaaSExampleHTTPvirtualserver1",
		getAllVirtualServerAPI.GetResponse().Children[1].Href,
	)
}

func TestGetAllVirtualServers_GetResponse(t *testing.T) {
	setupGetAll()
	assert.IsType(t, getAllVirtualServerAPI.GetResponse(), VirtualServersList{})

}
