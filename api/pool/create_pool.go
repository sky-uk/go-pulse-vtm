package pool

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"net/http"
)

<<<<<<< HEAD

=======
>>>>>>> 8b461f7e26a3bc1f8cd0128d196fa7e2f0107ef6
// CreatePoolAPI - Base Struct
type CreatePoolAPI struct {
	*api.BaseAPI
}

//NewCreate - Creates a new pool
func NewCreate(poolName string, pool Pool) *CreatePoolAPI {
<<<<<<< HEAD
	return execCreateUpdate(poolName, pool)
}


func NewUpdate(poolName string, pool Pool) *CreatePoolAPI {
	return execCreateUpdate(poolName, pool)

=======

	return execCreateUpdate(poolName, pool)
}

//NewUpdate - Placeholder to create
func NewUpdate(poolName string, pool Pool) *CreatePoolAPI {
	return execCreateUpdate(poolName, pool)
>>>>>>> 8b461f7e26a3bc1f8cd0128d196fa7e2f0107ef6
}

func execCreateUpdate(poolName string, pool Pool) *CreatePoolAPI {
	this := new(CreatePoolAPI)
<<<<<<< HEAD
=======
	/*requestPayload := new(Pool)
	requestPayload.Properties.Basic.NodesTable = nodeList
	requestPayload.Properties.Basic.Monitors = nodeMonitors*/
>>>>>>> 8b461f7e26a3bc1f8cd0128d196fa7e2f0107ef6
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, "/api/tm/3.8/config/active/pools/"+poolName, pool, new(Pool))
	return this
}

// GetResponse - Returns the http call response
func (cp CreatePoolAPI) GetResponse() *Pool {
	return cp.ResponseObject().(*Pool)
}
<<<<<<< HEAD

=======
>>>>>>> 8b461f7e26a3bc1f8cd0128d196fa7e2f0107ef6
