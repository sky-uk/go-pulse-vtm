package pool

import (
	"encoding/json"
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const testPoolEndpoint = "/api/tm/3.8/config/active/pools"

var testPoolName, marshallingTestExpectedJSON string
var createPoolAPI, getAllPoolsAPI, getPoolAPI, updatePoolAPI, deletePoolAPI *rest.BaseAPI
var getAllUnmarshallingTestJSON, getUnmarshallingTestJSON []byte

func setupPoolTest() {
	testPoolName = "test-pool"
	node := NewMemberNode("127.0.0.1:80", 1, "active", 1)
	node2 := NewMemberNode("127.0.0.1:81", 1, "active", 1)
	testPool := Pool{}
	testPool.Properties.Basic.NodesTable = []MemberNode{node, node2}
	testPool.Properties.Basic.Monitors = []string{"ping"}
	createPoolAPI = NewCreate(testPoolName, testPool)
	createPoolAPI.SetResponseObject(&testPool)

	marshallingTestExpectedJSON = `{"properties":{"basic":{"monitors":["ping"],"node_drain_to_delete_timeout":0,"nodes_table":[{"node":"127.0.0.1:80","priority":1,"state":"active","weight":1},{"node":"127.0.0.1:81","priority":1,"state":"active","weight":1}]},"connection":{},"http":{},"load_balancing":{},"node":{},"ssl":{},"tcp":{}}}`

	getAllPoolsAPI = NewGetAll()
	getAllUnmarshallingTestJSON = []byte(`{"children":[{"name":"pool_test_1","href":"/api/tm/3.8/config/active/pools/pool_test_1"},{"name":"pool_test_2","href":"/api/tm/3.8/config/active/pools/pool_test_2"}]}`)

	getPoolAPI = NewGet(testPoolName)
	getUnmarshallingTestJSON = []byte(`{"properties":{"basic":{"bandwidth_class":"","failure_pool":"","max_connection_attempts":0,"max_idle_connections_pernode":50,"max_timed_out_connection_attempts":2,"monitors":["Ping"],"node_close_with_rst":false,"node_connection_attempts":3,"node_delete_behavior":"immediate","node_drain_to_delete_timeout":0,"nodes_table":[{"node":"163.172.25.27:80","state":"active","weight":1},{"node":"127.0.0.1:80","state":"active","weight":1}],"note":"","passive_monitoring":true,"persistence_class":"","transparent":false},"auto_scaling":{"addnode_delaytime":0,"cloud_credentials":"","cluster":"","data_center":"","data_store":"","enabled":false,"external":true,"hysteresis":20,"imageid":"","ips_to_use":"publicips","last_node_idle_time":3600,"max_nodes":4,"min_nodes":1,"name":"","port":80,"refractory":180,"response_time":1000,"scale_down_level":95,"scale_up_level":40,"securitygroupids":[],"size_id":"","subnetids":[]},"connection":{"max_connect_time":4,"max_connections_per_node":0,"max_queue_size":0,"max_reply_time":30,"queue_timeout":10},"dns_autoscale":{"enabled":false,"hostnames":[],"port":80},"ftp":{"support_rfc_2428":false},"http":{"keepalive":true,"keepalive_non_idempotent":false},"kerberos_protocol_transition":{"principal":"","target":""},"load_balancing":{"algorithm":"round_robin","priority_enabled":false,"priority_nodes":1},"node":{"close_on_death":false,"retry_fail_time":60},"smtp":{"send_starttls":true},"ssl":{"client_auth":false,"common_name_match":[],"elliptic_curves":[],"enable":false,"enhance":false,"send_close_alerts":true,"server_name":false,"signature_algorithms":"","ssl_ciphers":"","ssl_support_ssl2":"use_default","ssl_support_ssl3":"use_default","ssl_support_tls1":"use_default","ssl_support_tls1_1":"use_default","ssl_support_tls1_2":"use_default","strict_verify":false},"tcp":{"nagle":true},"udp":{"accept_from":"dest_only","accept_from_mask":"","response_timeout":0}}}`)

	updatePoolAPI = NewUpdate(testPoolName, testPool)

	deletePoolAPI = NewDelete(testPoolName)
}

func TestNewCreateMethod(t *testing.T) {
	setupPoolTest()
	assert.Equal(t, http.MethodPut, createPoolAPI.Method())
}

func TestNewCreateEndpoint(t *testing.T) {
	setupPoolTest()
	assert.Equal(t, testPoolEndpoint+"/"+testPoolName, createPoolAPI.Endpoint())
}

func TestNewCreateMarshalling(t *testing.T) {
	setupPoolTest()
	jsonBytes, err := json.Marshal(createPoolAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, marshallingTestExpectedJSON, string(jsonBytes))
}

func TestNewGetAllMethod(t *testing.T) {
	setupPoolTest()
	assert.Equal(t, http.MethodGet, getAllPoolsAPI.Method())
}

func TestNewGetAllEndpoint(t *testing.T) {
	setupPoolTest()
	assert.Equal(t, testPoolEndpoint, getAllPoolsAPI.Endpoint())
}

func TestNewGetAllUnmarshalling(t *testing.T) {
	setupPoolTest()
	jsonErr := json.Unmarshal(getAllUnmarshallingTestJSON, getAllPoolsAPI.ResponseObject())
	response := *getAllPoolsAPI.ResponseObject().(*LBPoolList)

	assert.Nil(t, jsonErr)
	assert.Len(t, response.ChildPools, 2)
	assert.Equal(t, response.ChildPools[0].Name, "pool_test_1")
	assert.Equal(t, response.ChildPools[0].Href, "/api/tm/3.8/config/active/pools/pool_test_1")
	assert.Equal(t, response.ChildPools[1].Name, "pool_test_2")
	assert.Equal(t, response.ChildPools[1].Href, "/api/tm/3.8/config/active/pools/pool_test_2")
}

func TestNewGetMethod(t *testing.T) {
	setupPoolTest()
	assert.Equal(t, http.MethodGet, getPoolAPI.Method())
}

func TestNewGetEndpoint(t *testing.T) {
	setupPoolTest()
	assert.Equal(t, testPoolEndpoint+"/"+testPoolName, getPoolAPI.Endpoint())
}

func TestNewGetUnmarshalling(t *testing.T) {
	setupPoolTest()
	jsonErr := json.Unmarshal(getUnmarshallingTestJSON, getPoolAPI.ResponseObject())
	response := *getPoolAPI.ResponseObject().(*Pool)

	assert.Nil(t, jsonErr)
	assert.Equal(t, 0, response.Properties.Basic.MaxConnectionAttempts)
	assert.Equal(t, 50, response.Properties.Basic.MaxIdleConnectionsPerNode)
	assert.Equal(t, 2, response.Properties.Basic.MaxTimeoutConnectionAttempts)
	assert.Equal(t, "Ping", response.Properties.Basic.Monitors[0])
	assert.Equal(t, 3, response.Properties.Basic.NodeConnectionAttempts)
	assert.Equal(t, "163.172.25.27:80", response.Properties.Basic.NodesTable[0].Node)
	assert.Equal(t, "127.0.0.1:80", response.Properties.Basic.NodesTable[1].Node)
}

func TestNewUpdateMethod(t *testing.T) {
	setupPoolTest()
	assert.Equal(t, http.MethodPut, updatePoolAPI.Method())
}

func TestNewUpdateEndpoint(t *testing.T) {
	setupPoolTest()
	assert.Equal(t, testPoolEndpoint+"/"+testPoolName, updatePoolAPI.Endpoint())
}

func TestNewUpdateMarshalling(t *testing.T) {
	setupPoolTest()
	jsonBytes, err := json.Marshal(updatePoolAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, marshallingTestExpectedJSON, string(jsonBytes))
}

func TestNewDeleteMethod(t *testing.T) {
	setupPoolTest()
	assert.Equal(t, http.MethodDelete, deletePoolAPI.Method())
}

func TestNewDeleteEndpoint(t *testing.T) {
	setupPoolTest()
	assert.Equal(t, testPoolEndpoint+"/"+testPoolName, deletePoolAPI.Endpoint())
}
