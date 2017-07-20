package pool

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// GetSinglePool - Base Struct
type GetSinglePool struct {
	*rest.BaseAPI
}

// NewGetSingle - Returns a single pool
func NewGetSingle(poolName string) *GetSinglePool {
	this := new(GetSinglePool)
	this.BaseAPI = rest.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/pools/"+poolName, nil, new(Pool), new(api.VTMError))
	return this

}

// GetResponse returns ResponseObject of GetSinglePool
func (gsp GetSinglePool) GetResponse() Pool {
	return *gsp.ResponseObject().(*Pool)
}
