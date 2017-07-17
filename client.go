package brocadevtm

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// NewVTMClient  Creates a new vtmClient object.
func NewVTMClient(url string, user string, password string, ignoreSSL bool, debug bool) *VTMClient {
	vtmClient := new(VTMClient)
	vtmClient.URL = url
	vtmClient.User = user
	vtmClient.Password = password
	vtmClient.IgnoreSSL = ignoreSSL
	vtmClient.debug = debug
	return vtmClient
}

// VTMClient struct.
type VTMClient struct {
	URL       string
	User      string
	Password  string
	IgnoreSSL bool
	debug     bool
}

// Do - makes the API call.
func (vtmClient *VTMClient) Do(api api.VTMApi) error {

	apiEndpoint := api.Endpoint()

	requestURL := fmt.Sprintf("%s%s", vtmClient.URL, apiEndpoint)
	var requestPayload io.Reader
	octetStreamURLS := regexp.MustCompilePOSIX(`^(/api/tm/3.8/config/active/rules/.*)`)

	// TODO: change this to JSON
	if api.RequestObject() != nil && !octetStreamURLS.MatchString(apiEndpoint) {
		fmt.Printf("Encoding req object: %v", api.RequestObject())
		requestJSONBytes, marshallingErr := json.Marshal(api.RequestObject())
		if marshallingErr != nil {
			log.Fatal(marshallingErr)
			return (marshallingErr)
		}
		if vtmClient.debug {
			log.Println("Request payload as JSON:")
			log.Println(string(requestJSONBytes))
			log.Println("--------------------------------------------------------------")
		}
		requestPayload = bytes.NewReader(requestJSONBytes)
	}

	if requestPayload == nil && api.RequestObject() != nil {
		requestPayload = bytes.NewReader(api.RequestObject().([]byte))
	}

	if vtmClient.debug {
		log.Println("requestURL:", requestURL)
	}
	req, err := http.NewRequest(api.Method(), requestURL, requestPayload)
	if err != nil {
		log.Println("ERROR building the request: ", err)
		return err
	}

	req.SetBasicAuth(vtmClient.User, vtmClient.Password)
	// TODO: remove this hardcoded value!
	if octetStreamURLS.MatchString(apiEndpoint) {
		req.Header.Set("Content-Type", "application/octet-stream")
		req.Header.Set("Content-Transfer-Encoding", "text")
	} else {
		req.Header.Set("Content-Type", "application/json")
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: vtmClient.IgnoreSSL},
	}
	httpClient := &http.Client{Transport: tr}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Println("ERROR executing request: ", err)
		return err
	}
	defer res.Body.Close()
	return vtmClient.handleResponse(api, res)
}

func (vtmClient *VTMClient) handleResponse(api api.VTMApi, res *http.Response) error {
	api.SetStatusCode(res.StatusCode)
	bodyText, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("ERROR reading response: ", err)
		return err
	}

	api.SetRawResponse(bodyText)

	if vtmClient.debug {
		log.Println(string(bodyText))
	}

	if isJSON(res.Header.Get("Content-Type")) && api.StatusCode() == 200 {
		JSONerr := json.Unmarshal(bodyText, api.ResponseObject())
		if JSONerr != nil {
			log.Println("ERROR unmarshalling response: ", JSONerr)
			return nil
		}
	} else {
		api.SetResponseObject(string(bodyText))
	}
	return nil
}

func isJSON(contentType string) bool {
	return strings.Contains(strings.ToLower(contentType), "/json")
}
