package sslServerKey

import (
	"encoding/json"
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// NewUpdate : creates a new object of type UpdateMonitorAPI
func NewUpdate(name string, sslServerKey SSLServerKey) *rest.BaseAPI {
	sslServerKeyUpdateAPI := rest.NewBaseAPI(http.MethodPut, "/api/tm/3.8/config/active/ssl/server_keys/"+name, sslServerKey, new(json.RawMessage), new(api.VTMError))
	return sslServerKeyUpdateAPI
}
