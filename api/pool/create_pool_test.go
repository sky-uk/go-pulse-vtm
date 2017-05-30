package pool

import (
	//"encoding/json"
	//"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

)

var createPool *CreatePoolAPI
var node MemberNodes
var node2 MemberNodes

func setupCreatePool() {
	node = NewMemberNodes("127.0.0.1:80",1,"active",1 )
	node2 = NewMemberNodes("127.0.0.1:81",1,"active",1 )
	createPool = NewCreate("pool_test_rui_4", []MemberNodes{node,node2}, []string{"ping"})
}

func TestNewCreate(t *testing.T) {
	setupCreatePool()
	assert.Equal(t,http.MethodPut, createPool.Method())
	assert.Equal(t,"/api/tm/3.8/config/active/pools/pool_test_rui_4", createPool.Endpoint())
	assert.Equal(t,0,createPool.StatusCode())
	assert.Nil(t,createPool.GetResponse().Properties.Basic.NodesTable)
}

func TestNewUpdate(t *testing.T) {
	setupCreatePool()
	assert.Equal(t,http.MethodPut, createPool.Method())
	assert.Equal(t,"/api/tm/3.8/config/active/pools/pool_test_rui_4", createPool.Endpoint())
	assert.Equal(t,0,createPool.StatusCode())
	assert.Nil(t,createPool.GetResponse().Properties.Basic.NodesTable)

}