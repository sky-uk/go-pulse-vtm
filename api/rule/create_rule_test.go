package rule

import (
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createRuleAPI *rest.BaseAPI
var ruleName = "testRule"

func createSetup() {
	trafficScript := []byte(`
	if( string.ipmaskmatch( request.getremoteip(), "192.168.123.10" ) ){
		connection.discard();
	}`)
	createRuleAPI = NewCreate(ruleName, trafficScript)
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPut, createRuleAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/tm/3.8/config/active/rules/"+ruleName, createRuleAPI.Endpoint())
}
