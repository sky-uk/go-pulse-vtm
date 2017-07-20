package rule

import (
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateRuleAPI *rest.BaseAPI
var updateRuleName = "testRule"

func updateSetup() {
	trafficScript := []byte(`
	if( string.ipmaskmatch( request.getremoteip(), "10.12.13.10" ) ){
		connection.discard();
	}`)
	updateRuleAPI = NewUpdate(updateRuleName, trafficScript)
}

func TestUpdateMethod(t *testing.T) {
	updateSetup()
	assert.Equal(t, http.MethodPut, updateRuleAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	updateSetup()
	assert.Equal(t, "/api/tm/3.8/config/active/rules/"+updateRuleName, updateRuleAPI.Endpoint())
}
