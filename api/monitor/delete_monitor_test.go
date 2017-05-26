package monitor

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteMonitorAPI *DeleteMonitorAPI

func setupDelete() {
	deleteMonitorAPI = NewDelete("test-delete-monitor")
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteMonitorAPI.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/tm/3.8/config/active/monitors/test-delete-monitor", deleteMonitorAPI.Endpoint())
}
