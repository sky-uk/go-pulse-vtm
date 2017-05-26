package pool

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getSinglePoolAPI *GetSinglePool

func setupGetSingle() {
	getSinglePoolAPI = NewGetSingle("pool_test_rui_2")
}

func TestGetSingleMethod(t *testing.T) {
	setupGetSingle()
	assert.Equal(t, http.MethodGet, getSinglePoolAPI.Method())
}

func TestGetSingleEndpoint(t *testing.T) {
	setupGetSingle()
	assert.Equal(t, "/api/tm/3.8/config/active/pools/pool_test_rui_2", getSinglePoolAPI.Endpoint())
}

func TestGetSingleUnMarshalling(t *testing.T) {
	setupGetSingle()
	fmt.Println(getSinglePoolAPI.GetResponse().Properties)
	jsonContent := []byte("{\"properties\":{\"basic\":{\"bandwidth_class\":\"\",\"failure_pool\":\"\",\"max_connection_attempts\":0,\"max_idle_connections_pernode\":50,\"max_timed_out_connection_attempts\":2,\"monitors\":[\"Ping\"],\"node_close_with_rst\":false,\"node_connection_attempts\":3,\"node_delete_behavior\":\"immediate\",\"node_drain_to_delete_timeout\":0,\"nodes_table\":[{\"node\":\"163.172.25.27:80\",\"state\":\"active\",\"weight\":1},{\"node\":\"127.0.0.1:80\",\"state\":\"active\",\"weight\":1}],\"note\":\"\",\"passive_monitoring\":true,\"persistence_class\":\"\",\"transparent\":false},\"auto_scaling\":{\"addnode_delaytime\":0,\"cloud_credentials\":\"\",\"cluster\":\"\",\"data_center\":\"\",\"data_store\":\"\",\"enabled\":false,\"external\":true,\"hysteresis\":20,\"imageid\":\"\",\"ips_to_use\":\"publicips\",\"last_node_idle_time\":3600,\"max_nodes\":4,\"min_nodes\":1,\"name\":\"\",\"port\":80,\"refractory\":180,\"response_time\":1000,\"scale_down_level\":95,\"scale_up_level\":40,\"securitygroupids\":[],\"size_id\":\"\",\"subnetids\":[]},\"connection\":{\"max_connect_time\":4,\"max_connections_per_node\":0,\"max_queue_size\":0,\"max_reply_time\":30,\"queue_timeout\":10},\"dns_autoscale\":{\"enabled\":false,\"hostnames\":[],\"port\":80},\"ftp\":{\"support_rfc_2428\":false},\"http\":{\"keepalive\":true,\"keepalive_non_idempotent\":false},\"kerberos_protocol_transition\":{\"principal\":\"\",\"target\":\"\"},\"load_balancing\":{\"algorithm\":\"round_robin\",\"priority_enabled\":false,\"priority_nodes\":1},\"node\":{\"close_on_death\":false,\"retry_fail_time\":60},\"smtp\":{\"send_starttls\":true},\"ssl\":{\"client_auth\":false,\"common_name_match\":[],\"elliptic_curves\":[],\"enable\":false,\"enhance\":false,\"send_close_alerts\":true,\"server_name\":false,\"signature_algorithms\":\"\",\"ssl_ciphers\":\"\",\"ssl_support_ssl2\":\"use_default\",\"ssl_support_ssl3\":\"use_default\",\"ssl_support_tls1\":\"use_default\",\"ssl_support_tls1_1\":\"use_default\",\"ssl_support_tls1_2\":\"use_default\",\"strict_verify\":false},\"tcp\":{\"nagle\":true},\"udp\":{\"accept_from\":\"dest_only\",\"accept_from_mask\":\"\",\"response_timeout\":0}}}")
	jsonErr := json.Unmarshal(jsonContent, getSinglePoolAPI.ResponseObject())
	assert.Nil(t, jsonErr)
}
