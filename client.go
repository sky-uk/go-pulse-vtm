package brocadevtm

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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
	requestURL := fmt.Sprintf("%s%s", vtmClient.URL, api.Endpoint())
	var requestPayload io.Reader

	// TODO: change this to JSON
	if api.RequestObject() != nil {
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
	req.Header.Set("Content-Type", "application/json")

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

func (vtmClient *VTMClient) handleResponse(apiObj api.VTMApi, res *http.Response) error {
	apiObj.SetStatusCode(res.StatusCode)
	bodyText, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("ERROR reading response: ", err)
		return err
	}

	apiObj.SetRawResponse(bodyText)

	if vtmClient.debug {
		log.Println(string(bodyText))
	}

	if isJSON(res.Header.Get("Content-Type")) {
		if apiObj.StatusCode() >= http.StatusOK && apiObj.StatusCode() < http.StatusBadRequest {
			if len(bodyText) > 0 {
				JSONerr := json.Unmarshal(bodyText, apiObj.ResponseObject())
				if JSONerr != nil {
					log.Println("ERROR unmarshalling response: ", JSONerr)
					return JSONerr
				}
			}
			return nil
		}

		if len(bodyText) > 0 {
			var errObj api.ReqError
			err := json.Unmarshal(bodyText, &errObj)
			if err != nil {
				log.Printf("Error unmarshalling error response:\n%v", err)
			}
			apiObj.SetResponseObject(errObj)
			return errors.New(errObj.Error.ErrorText)
		}
	}
	return nil
}

func isJSON(contentType string) bool {
	return strings.Contains(strings.ToLower(contentType), "/json")
}
