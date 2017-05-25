package pool

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"net/http"
)

type GetAllPools struct {
	*api.BaseAPI
}

func NewGetAll() *GetAllPools {
	this := new(GetAllPools)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/pools", nil, new(PoolList))
	return this

}

// GetResponse returns ResponseObject of GetAllPools.
func (gap GetAllPools) GetResponse() *PoolList {
	return gap.ResponseObject().(*PoolList)
}
