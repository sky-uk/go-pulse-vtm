package pool

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"net/http"
)

<<<<<<< HEAD

=======
>>>>>>> 8b461f7e26a3bc1f8cd0128d196fa7e2f0107ef6
// DeletePoolAPI - Base Struct
type DeletePoolAPI struct {
	*api.BaseAPI
}

// NewDelete - Deletes a pool
func NewDelete(poolName string) *DeletePoolAPI {
	this := new(DeletePoolAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodDelete, "/api/tm/3.8/config/active/pools/"+poolName, nil, new(Pool))
	return this
}
