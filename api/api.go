package api

import (
	"fmt"
	"github.com/sky-uk/go-rest-api"
	"log"
	"net/http"
	"sort"
	"time"
)

const defaultAPIVersion = "3.8"
const configPath = "config/active"
const apiPrefix = "/api/tm"

// Params - connection parameters
type Params struct {
	Username   string
	Password   string
	Server     string
	APIVersion string
	IgnoreSSL  bool
	Debug      bool
	Timeout    time.Duration
}

// Client - the Brocade vTM Client struct
type Client struct {
	VersionsSupported []string
	restClient        rest.Client
	currentVersion    string
	RootPath          string
	currentServer     string
}

// WorkWithConfigurationResources - set current path to work with
// configuration resources
func (client *Client) WorkWithConfigurationResources() {
	client.RootPath = apiPrefix + "/" + client.currentVersion + "/" + configPath
	log.Println("Current Path: ", client.RootPath)
}

// Connect - connect to the Brocade REST API server
// and get the list of supported API versions
// Returns a new client object if everything is fine
func Connect(params Params) (*Client, error) {
	client := new(Client)
	client.currentVersion = params.APIVersion

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	client.restClient = rest.Client{
		URL:       "https://" + params.Server,
		User:      params.Username,
		Password:  params.Password,
		IgnoreSSL: params.IgnoreSSL,
		Debug:     params.Debug,
		Headers:   headers,
		Timeout:   params.Timeout,
	}

	supportedVersionsMap := make(map[string]interface{})
	var errStr VTMError
	api := rest.NewBaseAPI(
		http.MethodGet,
		apiPrefix,
		nil,
		&supportedVersionsMap,
		&errStr,
	)
	err := client.restClient.Do(api)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("Got:\n%+v", supportedVersionsMap)

	versions := make([]string, 0)
	for _, version := range supportedVersionsMap["children"].([]interface{}) {
		if vAsMap, ok := version.(map[string]interface{}); ok {
			versions = append(versions, vAsMap["name"].(string))
		}
	}
	client.VersionsSupported = versions
	if client.currentVersion == "" {
		sort.Sort(sort.Reverse(sort.StringSlice(client.VersionsSupported)))
		client.currentVersion = client.VersionsSupported[0]
		log.Println("Working with REST API Version: ", client.currentVersion)
	}

	return client, nil
}

// GetAllResourceTypes - returns the list of all types of configuration resources
func (client Client) GetAllResourceTypes() ([]map[string]interface{}, error) {
	path := client.RootPath
	res := make(map[string]interface{}, 0)
	var tmErr *VTMError

	log.Println("Going to get all resource types, using PATH:\n", path)
	api := rest.NewBaseAPI(
		http.MethodGet,
		path,
		nil,
		&res,
		tmErr,
	)
	err := client.restClient.Do(api)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if tmErr != nil {
		log.Println("Error setting a resource: ", tmErr.ErrorText)
		return nil, fmt.Errorf(tmErr.ErrorText)
	}
	resTypes := make([]map[string]interface{}, 0)
	for _, item := range res["children"].([]interface{}) {
		resTypes = append(resTypes, item.(map[string]interface{}))
	}
	return resTypes, nil
}

// TraverseTree - retrieves a resource and eventually keep doing it
// for each nested resource
// Fill up the passed slice of resources, returns the first error it
// eventually bumps into
func (client Client) TraverseTree(url string, resources map[string]interface{}) error {
	var tmErr *VTMError
	res := make(map[string]interface{})

	if url == "" {
		return fmt.Errorf("Invalid path")
	}

	log.Println("Going to get PATH: ", url)
	api := rest.NewBaseAPI(
		http.MethodGet,
		url,
		nil,
		&res,
		tmErr,
	)
	err := client.restClient.Do(api)
	if err != nil {
		log.Println(err)
		return err
	}
	if tmErr != nil {
		log.Println("Error getting path: ", tmErr.ErrorText)
		return fmt.Errorf(tmErr.ErrorText)
	}

	if children, exists := res["children"]; exists {
		log.Println("Going deeper...")
		for _, item := range children.([]interface{}) {
			if itemAsMap, ok := item.(map[string]interface{}); ok {
				err = client.TraverseTree(itemAsMap["href"].(string), resources)
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("Strange...I expected a slice of maps")
			}
		}
	} else {
		if _, exists := res["properties"]; exists {
			resources[url] = res
		}
	}

	return nil
}

// GetAllResources - returns all resources of the specified type
func (client Client) GetAllResources(resType string) ([]map[string]interface{}, error) {
	path := client.RootPath + "/" + resType
	res := make(map[string]interface{})
	var tmErr *VTMError

	log.Println("Going to get all resources, using PATH: ", path)
	api := rest.NewBaseAPI(
		http.MethodGet,
		path,
		nil,
		&res,
		tmErr,
	)
	err := client.restClient.Do(api)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if tmErr != nil {
		log.Println("Error setting a resource: ", tmErr.ErrorText)
		return nil, fmt.Errorf(tmErr.ErrorText)
	}
	resources := make([]map[string]interface{}, 0)
	if list, exists := res["children"].([]interface{}); exists {
		for _, item := range list {
			resources = append(resources, item.(map[string]interface{}))
		}
	} else {
		log.Println("Key \"children\" not found in response")
	}
	return resources, nil
}

// GetByName - gets a resource profile given its type and name
func (client Client) GetByName(resType, resName string) (map[string]interface{}, error) {
	path := client.RootPath + "/" + resType + "/" + resName
	res := make(map[string]interface{})
	var tmErr *VTMError

	log.Println("Going to get a resource, using PATH: ", path)
	api := rest.NewBaseAPI(
		http.MethodGet,
		path,
		nil,
		&res,
		tmErr,
	)
	err := client.restClient.Do(api)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if tmErr != nil {
		log.Println("Error getting a resource: ", tmErr.ErrorText)
		return nil, fmt.Errorf(tmErr.ErrorText)
	}
	return res, nil
}

// GetByURL - gets a resource profile given its type and URL
func (client Client) GetByURL(resURL string) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	var tmErr *VTMError

	log.Println("Going to get a resource, using PATH: ", resURL)
	api := rest.NewBaseAPI(
		http.MethodGet,
		resURL,
		nil,
		&res,
		tmErr,
	)
	err := client.restClient.Do(api)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if tmErr != nil {
		log.Println("Error setting a resource: ", tmErr.ErrorText)
		return nil, fmt.Errorf(tmErr.ErrorText)
	}
	return res, nil
}

// SetResource - Sets a resource
// This works only in Configuration environment (statistics/information resources can't be set)
// A new resources gets created if not existent or an existent resource gets updated
// The restClient.StatusCode is set properly to http.StatusCreated or http.StatusOK accordingly
// Returns the created/updated object or an error
func (client Client) Set(resType, name string, profile interface{}) (map[string]interface{}, error) {
	var tmErr *VTMError
	res := make(map[string]interface{})

    // you can only set configuration resources...
    client.WorkWithConfigurationResources()

	path := client.RootPath + "/" + resType + "/" + name
	api := rest.NewBaseAPI(
		http.MethodPut,
		path,
		profile,
		&res,
		tmErr,
	)
	err := client.restClient.Do(api)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if tmErr != nil {
		log.Println("Error setting a resource: ", tmErr.ErrorText)
		return nil, fmt.Errorf(tmErr.ErrorText)
	}

	return res, nil
}
