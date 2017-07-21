package rule

import (
	"github.com/sky-uk/go-rest-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getRuleAPI *rest.BaseAPI
var getRuleName string
var getRuleResponseObject []byte

func setupGetRule() {
	getRuleName = "test-rule"
	getRuleResponseObject = []byte(`if( http.responseHeaderExists( "Keep-Alive" ) == 0) {
  		http.removeResponseHeader( "Keep-Alive" );
	}`)
	getRuleAPI = NewGetRule(getRuleName)
	getRuleAPI.SetRawResponse(getRuleResponseObject)
}

func TestGetMethod(t *testing.T) {
	setupGetRule()
	assert.Equal(t, http.MethodGet, getRuleAPI.Method())
}

func TestGetEndpoint(t *testing.T) {
	setupGetRule()
	assert.Equal(t, "/api/tm/3.8/config/active/rules/"+getRuleName, getRuleAPI.Endpoint())
}
