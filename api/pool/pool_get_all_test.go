package pool

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllPoolAPI *GetAllPools

func setupGetAll() {
	getAllPoolAPI = NewGetAll()
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllPoolAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/tm/3.8/config/active/pools", getAllPoolAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	jsonContent := []byte("{\"children\":[{\"name\":\"pool_test_rui\",\"href\":\"/api/tm/3.8/config/active/pools/pool_test_rui\"},{\"name\":\"pool_test_rui_2\",\"href\":\"/api/tm/3.8/config/active/pools/pool_test_rui_2\"}]}")
	jsonErr := json.Unmarshal(jsonContent, getAllPoolAPI.ResponseObject())
	fmt.Println(getAllPoolAPI.GetResponse().ChildPools[0].Name)
	assert.Nil(t, jsonErr)
	assert.Len(t, getAllPoolAPI.GetResponse().ChildPools, 2)
	assert.Equal(t, "pool_test_rui", getAllPoolAPI.GetResponse().ChildPools[0].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/pools/pool_test_rui", getAllPoolAPI.GetResponse().ChildPools[0].Href)
	assert.Equal(t, "pool_test_rui_2", getAllPoolAPI.GetResponse().ChildPools[1].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/pools/pool_test_rui_2", getAllPoolAPI.GetResponse().ChildPools[1].Href)
}

func TestGetAllPools_GetResponse(t *testing.T) {
	setupGetAll()
	assert.IsType(t, getAllPoolAPI.GetResponse(), LBPoolList{})

}
