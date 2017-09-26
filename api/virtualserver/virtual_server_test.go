package virtualserver

import (
	"encoding/json"
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const testVirtualServerEndpoint = "/api/tm/3.8/config/active/virtual_servers/"

var createVirtualServerAPI, getAllVirtualServerAPI, getVirtualServerAPI, updateVirtualServerAPI, deleteVirtualServerAPI *rest.BaseAPI
var testVirtualServerName string
var testVirtualServerObject VirtualServer
var testVirtualServerMarshalExpectedJSON string
var testVirtualServerGetAllUnmarshallJSON, testVirtualServerGetUnmarshalJSON []byte

func setupVirtualServerTest() {
	testVirtualServerName = "test-virtual-server"
	enabled := false
	testVirtualServerBasic := Basic{
		Enabled:  enabled,
		Pool:     "test-pool",
		Port:     80,
		Protocol: "http",
	}
	testVirtualServerProperties := Properties{Basic: testVirtualServerBasic}
	testVirtualServerObject = VirtualServer{Properties: testVirtualServerProperties}

	testVirtualServerMarshalExpectedJSON = `{"properties":{"basic":{"add_cluster_ip":false,"add_x_forwarded_for":false,"add_x_forwarded_proto":false,"autodetect_upgrade_headers":false,"close_with_rst":false,"connect_timeout":0,"enabled":false,"ftp_force_server_secure":false,"listen_on_any":false,"mss":0,"pool":"test-pool","port":80,"protocol":"http","so_nagle":false,"ssl_decrypt":false,"transparent":false},"aptimizer":{"enabled":false},"connection":{"keepalive":false,"keepalive_timeout":0,"max_client_buffer":0,"max_server_buffer":0,"max_transaction_duration":0,"timeout":0},"connection_errors":{},"cookie":{},"dns":{"edns_client_subnet":false,"edns_udpsize":0,"max_udpsize":0,"verbose":false},"ftp":{"data_source_port":0,"force_client_secure":false,"port_range_high":0,"port_range_low":0,"ssl_data":false},"gzip":{"compress_level":0,"enabled":false,"etag_rewrite":"","include_mime":null,"max_size":0,"min_size":0,"no_size":false},"http":{"mime_detect":false},"http2":{"connect_timeout":0,"data_frame_size":0,"enabled":false,"header_table_size":0,"headers_index_default":false,"idle_timeout_no_streams":0,"idle_timeout_open_streams":0,"max_concurrent_streams":0,"max_frame_size":0,"max_header_padding":0,"merge_cookie_headers":false,"stream_window_size":0},"kerberos_protocol_transition":{"enabled":false},"log":{"always_flush":false,"client_connection_failures":false,"enabled":false,"save_all":false,"server_connection_failures":false,"session_persistence_verbose":false,"ssl_failures":false},"recent_connections":{"enabled":false,"save_all":false},"request_tracing":{"enabled":false,"trace_io":false},"rtsp":{"streaming_port_range_high":0,"streaming_port_range_low":0,"streaming_timeout":0},"sip":{"follow_route":false,"max_connection_mem":0,"rewrite_uri":false,"streaming_port_range_high":0,"streaming_port_range_low":0,"streaming_timeout":0,"timeout_messages":false,"transaction_timeout":0},"smtp":{"expect_starttls":false},"ssl":{"add_http_headers":false,"ocsp_enable":false,"oscp_max_response_age":0,"oscp_stapling":false,"oscp_time_tolerance":0,"oscp_timeout":0,"prefer_sslv3":false,"send_close_alerts":false,"trust_magic":false},"sys_log":{"enabled":false,"msg_len_limit":0},"tcp":{"proxy_close":false},"udp":{"end_point_persistence":false,"port_smp":false,"response_datagrams_expected":0,"timeout":0},"web_cache":{"enabled":false,"error_page_time":0,"max_time":0,"refresh_time":0}}}`

	testVirtualServerGetAllUnmarshallJSON = []byte(`{"children":[{"name":"PaaSExampleHTTPvirtualserver","href":"/api/tm/3.8/config/active/virtual_servers/PaaSExampleHTTPvirtualserver"},{"name":"PaaSExampleHTTPvirtualserver1","href":"/api/tm/3.8/config/active/virtual_servers/PaaSExampleHTTPvirtualserver1"},{"name":"virtual_server_1","href":"/api/tm/3.8/config/active/virtual_servers/virtual_server_1"},{"name":"virtual_server_2","href":"/api/tm/3.8/config/active/virtual_servers/virtual_server_2"},{"name":"virtual_server_3","href":"/api/tm/3.8/config/active/virtual_servers/virtual_server_3"}]}`)

	testVirtualServerGetUnmarshalJSON = []byte(`{"properties":{"basic":{"enabled":false,"pool":"pool_test","port":80,"protocol":"http"},"aptimizer":{},"connection":{},"connection_errors":{},"cookie":{},"dns":{},"ftp":{"ssl_data":false},"gzip":{},"ssl":{}}}`)

	createVirtualServerAPI = NewCreate(testVirtualServerName, testVirtualServerObject)
	createVirtualServerAPI.SetResponseObject(&testVirtualServerObject)

	getAllVirtualServerAPI = NewGetAll()

	getVirtualServerAPI = NewGet(testVirtualServerName)

	updateVirtualServerAPI = NewUpdate(testVirtualServerName, testVirtualServerObject)
	updateVirtualServerAPI.SetResponseObject(&testVirtualServerObject)

	deleteVirtualServerAPI = NewDelete(testVirtualServerName)
}

func TestVirtualServerNewCreateMethod(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, http.MethodPut, createVirtualServerAPI.Method())
}

func TestVirtualServerNewCreateEndpoint(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, testVirtualServerEndpoint+testVirtualServerName, createVirtualServerAPI.Endpoint())
}

func TestVirtualServerNewCreateMarshal(t *testing.T) {
	setupVirtualServerTest()
	jsonBytes, err := json.Marshal(createVirtualServerAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, testVirtualServerMarshalExpectedJSON, string(jsonBytes))
}

func TestVirtualServerNewGetAllMethod(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, http.MethodGet, getAllVirtualServerAPI.Method())
}

func TestVirtualServerNewGetAllEndpoint(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, testVirtualServerEndpoint, getAllVirtualServerAPI.Endpoint())
}

func TestVirtualServerNewGetAllUnmarshal(t *testing.T) {
	setupVirtualServerTest()
	jsonErr := json.Unmarshal(testVirtualServerGetAllUnmarshallJSON, getAllVirtualServerAPI.ResponseObject())
	response := *getAllVirtualServerAPI.ResponseObject().(*VirtualServersList)

	assert.Nil(t, jsonErr)
	assert.Len(t, response.Children, 5)
	assert.Equal(t, "PaaSExampleHTTPvirtualserver", response.Children[0].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/virtual_servers/PaaSExampleHTTPvirtualserver", response.Children[0].Href)
	assert.Equal(t, "virtual_server_3", response.Children[4].Name)
	assert.Equal(t, "/api/tm/3.8/config/active/virtual_servers/virtual_server_3", response.Children[4].Href)
}

func TestVirtualServerNewGetMethod(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, http.MethodGet, getVirtualServerAPI.Method())
}

func TestVirtualServerNewGetEndpoint(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, testVirtualServerEndpoint+testVirtualServerName, getVirtualServerAPI.Endpoint())
}

func TestVirtualServerNewGetUnmarshal(t *testing.T) {
	setupVirtualServerTest()
	jsonErr := json.Unmarshal(testVirtualServerGetUnmarshalJSON, getVirtualServerAPI.ResponseObject())
	response := *getVirtualServerAPI.ResponseObject().(*VirtualServer)

	assert.Nil(t, jsonErr)
	assert.Equal(t, false, response.Properties.Basic.Enabled)
	assert.Equal(t, "pool_test", response.Properties.Basic.Pool)
	assert.Equal(t, uint(80), response.Properties.Basic.Port)
	assert.Equal(t, "http", response.Properties.Basic.Protocol)
}

func TestVirtualServerNewUpdateMethod(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, http.MethodPut, updateVirtualServerAPI.Method())
}

func TestVirtualServerNewUpdateEndpoint(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, testVirtualServerEndpoint+testVirtualServerName, updateVirtualServerAPI.Endpoint())
}

func TestVirtualServerNewUpdateMarshal(t *testing.T) {
	setupVirtualServerTest()
	jsonBytes, err := json.Marshal(updateVirtualServerAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, testVirtualServerMarshalExpectedJSON, string(jsonBytes))
}

func TestVirtualServerNewDeleteMethod(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, http.MethodDelete, deleteVirtualServerAPI.Method())
}

func TestVirtualServerNewDeleteEndpoint(t *testing.T) {
	setupVirtualServerTest()
	assert.Equal(t, testVirtualServerEndpoint+testVirtualServerName, deleteVirtualServerAPI.Endpoint())
}
