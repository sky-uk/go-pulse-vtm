package rule

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getRuleAPI *GetRuleAPI
var getRuleName string
var getRuleResponseObject []byte

func setupGet() {
	getRuleName = "test-rule"
	getRuleResponseObject = []byte(`if( http.responseHeaderExists( "Keep-Alive" ) == 0) {
  		http.removeResponseHeader( "Keep-Alive" );
	}`)
	getRuleAPI = NewGetRule(getRuleName)
	getRuleAPI.SetRawResponse(getRuleResponseObject)
}

func TestGetMethod(t *testing.T) {
	setupGet()
	assert.Equal(t, http.MethodGet, getRuleAPI.Method())
}

func TestGetEndpoint(t *testing.T) {
	setupGet()
	assert.Equal(t, "/api/tm/3.8/config/active/rules/"+getRuleName, getRuleAPI.Endpoint())
}

func TestGetResponse(t *testing.T) {
	setupGet()
	ruleGetResponse := getRuleAPI.GetResponse()
	assert.Equal(t, string(getRuleResponseObject), ruleGetResponse)
}
