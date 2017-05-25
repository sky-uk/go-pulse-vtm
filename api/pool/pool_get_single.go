package pool

import ("net/http"
	"github.com/sky-uk/go-brocade-vtm/api"

)


type GetSinglePool struct {
	*api.BaseAPI
}

func NewGetSingle(poolName string) *GetSinglePool {
	this := new(GetSinglePool)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/tm/3.8/config/active/pools/"+poolName, nil, new(Pool))
	return this

}

// GetResponse returns ResponseObject of GetSinglePool
func (gsp GetSinglePool) GetResponse() *Pool {
	return gsp.ResponseObject().(*Pool)
}
