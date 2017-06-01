package trafficIpGroups

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getSingleTrafficIPGroupAPI *GetSingleTrafficIPGroupAPI

func setupGetSingleTrafficIPGroup() {
	getSingleTrafficIPGroupAPI = NewGetSingle("test-group-1")
}

func TestGetSingleTrafficIPGroupMethod(t *testing.T) {
	setupGetSingleTrafficIPGroup()
	assert.Equal(t, http.MethodGet, getSingleTrafficIPGroupAPI.Method())
}

func TestGetSingleTrafficIPGroupEndpoint(t *testing.T) {
	setupGetSingleTrafficIPGroup()
	assert.Equal(t, "/api/tm/3.8/config/active/traffic_ip_groups/test-group-1", getSingleTrafficIPGroupAPI.Endpoint())
}

func TestGetSingleTrafficIPGroupUnMarshalling(t *testing.T) {
	setupGetSingleTrafficIPGroup()
	jsonContent := []byte("{\"properties\":{\"basic\":{\"enabled\": true,\"hash_source_port\": true,\"ip_assignment_mode\": \"alphabetic\",\"ip_mapping\":[{\"ip\": \"172.0.0.5\", \"traffic_manager\": \"172.0.0.10\" }],\"ipaddresses\":[\"172.0.0.1\", \"172.0.0.2\"],\"keeptogether\": false,\"location\": 3,\"machines\":[\"172.0.0.8\", \"172.0.0.9\"],\"mode\": \"singlehosted\",\"multicast\": \"172.0.0.99\",\"note\":\"testnote\",\"rhi_bgp_metric_base\": 5,\"rhi_bgp_passive_metric_offset\": 5,\"rhi_ospfv2_metric_base\": 7,\"rhi_ospfv2_passive_metric_offset\": 5,\"rhi_protocols\": \"ospf\",\"slaves\":[]}}}")
	jsonErr := json.Unmarshal(jsonContent, getSingleTrafficIPGroupAPI.ResponseObject())
	assert.Nil(t, jsonErr)
}

func TestGetSingleTrafficIPGroupAPIGetResponse(t *testing.T) {
	setupGetSingleTrafficIPGroup()
	assert.IsType(t, getSingleTrafficIPGroupAPI.GetResponse(), &TrafficIPGroup{})
}
