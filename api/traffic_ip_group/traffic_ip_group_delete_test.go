package trafficIpGroups

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteTrafficIPGroupAPI *DeleteTrafficIPGroupAPI

func setupDeleteTrafficIPGroup() {
	deleteTrafficIPGroupAPI = NewDelete("test-group-1")
}

func TestDeleteTrafficIPGroupMethod(t *testing.T) {
	setupDeleteTrafficIPGroup()
	assert.Equal(t, http.MethodDelete, deleteTrafficIPGroupAPI.Method())
}

func TestDeleteTrafficIPGroupEndpoint(t *testing.T) {
	setupDeleteTrafficIPGroup()
	assert.Equal(t, "/api/tm/3.8/config/active/traffic_ip_groups/test-group-1", deleteTrafficIPGroupAPI.Endpoint())
}
