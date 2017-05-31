package pool

import (

	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createPool *CreatePoolAPI
var updatePool *CreatePoolAPI
var node MemberNode
var node2 MemberNode

func setupCreatePool() {
	node = NewMemberNode("127.0.0.1:80", 1, "active", 1)
	node2 = NewMemberNode("127.0.0.1:81", 1, "active", 1)
	testPool := Pool{}
	testPool.Properties.Basic.NodesTable = []MemberNode{node, node2}
	testPool.Properties.Basic.Monitors = []string{"ping"}
	createPool = NewCreate("pool_test_rui_4", testPool )
}
func setupUpdatePool() {
	node = NewMemberNode("127.0.0.1:80", 1, "active", 1)
	node2 = NewMemberNode("127.0.0.1:81", 1, "active", 1)
	testPool := Pool{}
	testPool.Properties.Basic.NodesTable = []MemberNode{node, node2}
	testPool.Properties.Basic.Monitors = []string{"ping"}
	updatePool = NewUpdate("pool_test_rui_4", testPool )
}



func TestNewCreate(t *testing.T) {
	setupCreatePool()
	assert.Equal(t, http.MethodPut, createPool.Method())
	assert.Equal(t, "/api/tm/3.8/config/active/pools/pool_test_rui_4", createPool.Endpoint())
	assert.Equal(t, 0, createPool.StatusCode())
	assert.Nil(t, createPool.GetResponse().Properties.Basic.NodesTable)
}

func TestNewUpdate(t *testing.T) {
	setupUpdatePool()
	assert.Equal(t, http.MethodPut, updatePool.Method())
	assert.Equal(t, "/api/tm/3.8/config/active/pools/pool_test_rui_4", updatePool.Endpoint())
	assert.Equal(t, 0, updatePool.StatusCode())
	assert.Nil(t, updatePool.GetResponse().Properties.Basic.NodesTable)

}
