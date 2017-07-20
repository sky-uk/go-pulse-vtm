# go-rest-api

A fairly generic HTTP API.

Supports both HTTP and TLS (https).
Encoding schemes supported: json/xml.


## Importing

```
import(
    "rest"
    "rest/api"
    "net/http"
)
```

## Usage

### Get a REST Client object

```
    client := NewClient(url, user, password, ignoreSSL, debug, headers) 

    // for a simple http client...
    httpClient := NewClient(url, "", "", false, false, nil)

```

### Perform a request

```
    // Prepare a request...
    api := api.NewRestAPI(
        http.MethodGet,         // request method
        "/",                    // request path
        nil,                    // request payload object
        new(string),            // (pointer to) response payload object
        nil,                    // (pointer to) error object
    )

    // Perform the request...
    err := httpClient.Do(api)
    if err != nil {
        // handle errors....
    }
```

### Getting the response object

```

    // example (json payload)
    type JSONFoo struct {
	    Fields map[string]string `json:"fields"`
    }

    // Prepare a request...
    api := api.NewRestAPI(
        http.MethodGet,         // request method
        "/",                    // request path
        nil,                    // request payload object
        new(JSONFoo),           // (pointer to) response payload object
        nil,                    // (pointer to) error object
    )

    //...perform request...

    respObj := *api.ResponseObject().(*JSONFoo)
```

More examples for the supported encodings in the client_test.go module.

### Getting response status code

```
    status := api.StatusCode()
```

### Getting the raw response as a byte stream

```
    raw := api.RawResponse()
```
