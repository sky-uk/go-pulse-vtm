package pool

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"net/http"

)

// CreatePool - Base Struct
type CreatePoolAPI struct {
	*api.BaseAPI
}

//NewCreate - Creates a new pool
func NewCreate(poolName string , nodeList []MemberNodes, nodeMonitors []string) *CreatePoolAPI {

	return execCreateUpdate(poolName,nodeList,nodeMonitors)
}

func NewUpdate(poolName string , nodeList []MemberNodes, nodeMonitors []string) *CreatePoolAPI {
	return execCreateUpdate(poolName,nodeList,nodeMonitors)
}

func execCreateUpdate(poolName string , nodeList []MemberNodes, nodeMonitors []string) *CreatePoolAPI {
	this := new(CreatePoolAPI)
	requestPayload := new(Pool)
	requestPayload.Properties.Basic.NodesTable = nodeList
	requestPayload.Properties.Basic.Monitors = nodeMonitors
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, "/api/tm/3.8/config/active/pools/"+poolName, requestPayload, new(Pool))
	return this
}

func (cp CreatePoolAPI) GetResponse() *Pool {
	return cp.ResponseObject().(*Pool)
}