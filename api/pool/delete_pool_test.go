package pool

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var DeleteAPI *DeletePoolAPI

func setupDelete() {
	DeleteAPI = NewDelete("pool_test_rui_2")
}

func TestNewDelete(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, DeleteAPI.Method())
	assert.Equal(t,"/api/tm/3.8/config/active/pools/pool_test_rui_2",DeleteAPI.Endpoint())
	assert.Equal(t,0,DeleteAPI.StatusCode())

}
