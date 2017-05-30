package pool

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"net/http"
)

// CreatePool - Base Struct
type CreatePoolAPI struct {
	*api.BaseAPI
}

//NewCreate

func NewCreate(poolName string , nodeList []MemberNodes, nodeMonitors []string ) *CreatePoolAPI {
	this := new(CreatePoolAPI)
	requestPayload := new(Pool)
	requestPayload.Name = poolName
	requestPayload.Properties.Basic.NodesTable = nodeList
	requestPayload.Properties.Basic.Monitors = nodeMonitors
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, "/api/tm/3.8/config/active/pools/"+poolName, requestPayload, new(string))
	return this
}

func (cp CreatePoolAPI) GetResponse() string {
	return cp.ResponseObject().(string)
}