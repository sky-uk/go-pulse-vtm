package pool

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"net/http"
)

// DeletePool - Base Struct
type DeletePool struct {
	*api.BaseAPI
}

// NewDelete - Deletes a pool
func NewDelete(poolName string) *DeletePool {
	this := new(DeletePool)
	this.BaseAPI = api.NewBaseAPI(http.MethodDelete, "/api/tm/3.8/config/active/pools/"+poolName, nil, new(Pool))
	return this
}
