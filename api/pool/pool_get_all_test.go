package pool

import (
	//"encoding/json"
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
