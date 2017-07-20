package pool

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// DeletePoolAPI - Base Struct
type DeletePoolAPI struct {
	*rest.BaseAPI
}

// NewDelete - Deletes a pool
func NewDelete(poolName string) *DeletePoolAPI {
	this := new(DeletePoolAPI)
	this.BaseAPI = rest.NewBaseAPI(http.MethodDelete, "/api/tm/3.8/config/active/pools/"+poolName, nil, new(Pool), new(api.VTMError))
	return this
}
