package virtualserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteVirtualServerAPI *DeleteVirtualServerAPI

func setupDelete() {
	deleteVirtualServerAPI = NewDelete("test-delete-virtual-server")
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteVirtualServerAPI.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/tm/3.8/config/active/virtual_servers/test-delete-virtual-server", deleteVirtualServerAPI.Endpoint())
}
