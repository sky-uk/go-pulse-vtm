# go-brocadev-tm - Go Bindings for the Brocade Virtual Traffic Manager (vTM)

## Overview

This is the GoLang API wrapper for Brocade Virtual Traffic Manager (vTM).
Starting from version 0.4.0 the API has been redesigned.

# API usage

## Importing the API

```
    import("github.com/sky-uk/go-brocade-vtm/api")
```

## Connecting to a BrocadevTM server
The Connect API returns a client object or an error in case of issues.
The APIVersion parameter can be used to specify a server API version to use;
in case it is not provided, the list of available API versions is retrieved
from the server and the highest available is used.
Custom headers can be passed at connection time, if not passed the content type
is anyhow set to "application/json"

```
    headers := make(map[string]string)
    // set your custom headers...
	params := Params{
		APIVersion: "3.8",
		Server:    server,
		Username:  username,
		Password:  password,
		IgnoreSSL: true,
		Debug:     true,
        Headers:   headers,
	}

	client, err := Connect(params)
```

## Dealing with configuration resources

### Setting (creating/updating) a resource
The Set API is provided to both create and update a configuration resource.
The last input paramenter is a pointer to a map (or struct) where the new
created/updated resource is provided back.

```
    err := client.Set(<resource type>,<resource name>, <resource profile>, <pointer to map/struct>)

    // example
	profile := make(map[string]interface{})
    name := "new_virtual_server_8347"
    // fill the profile...
    //...

	updatedRes := make(map[string]interface{})
	err = client.Set("virtual_servers", name, profile, &updatedRes)
```


### Getting a resource

#### Retrieve a resource by its name
```
    err := client.GetByName(<resource type>, <resource name>, <pointer to map/struct>)
```

#### Retrieve a resource by its url
```
    err := client.GetByURL(<resource type>, <resource name>, <pointer to map/struct>)
```


### Deleting a resource

```
    err := client.Delete(<resource type>, <resource name>)
```

### Getting all resource types
The ``GetAllResourceTypes`` returns the list of all resource types
```
    types, err := client.GetAllResourceTypes()
```


### Getting all resources of a type
The ``GetAllResources`` API returns a list of maps containing names/urls of all 
resources of the provided type.

```
    objs, err := client.GetAllResources(<resource type>)
```

## Dealing with information resources
The ``GetInformation`` API returns the information section for the passed server name
```
    info, err := client.GetInformation(<server name>)
```

## Dealing with statistics
The ``GetStatistics`` API returns all the statistics resources for a server node name
```
    stats, err := client.GetStatistics(<server name>)
```

## Getting the server state
The ``GetState`` API returns the current state for the passed server node name
```
    state, err := client.GetState(<server name>)
```

