# go-brocadev-tm - Go Bindings for the Brocade Virtual Traffic Manager (vTM)

## Overview

This is the GoLang API wrapper for Brocade Virtual Traffic Manager (vTM).
This wrapper uses the REST API interface provided by the vTM, currently version 3.8.


## Handled Resources

| Resource                | Create | Read  | Update  | Delete |
|-------------------------|--------|-------|---------|--------|
| Monitor [1]             |   Y    |   Y   |    Y    |   Y    |
| Pool                    |   Y    |   Y   |    Y    |   Y    |
| SSL Server Key          |   Y    |   Y   |    Y    |   Y    |
| Traffic IP Group        |   Y    |   Y   |    Y    |   Y    |

### Notes
[1] : Currently only HTTP monitoring is supported


## Usage

### Import library

```
import(
    "github.com/sky-uk/go-brocade-vtm"
)
```

### Get a client object

In order to get a client object authentication credentials needs to be
provided. The REST API transport procotol is by default HTTPS and hence
SSL encryption is used by default. Certificate handling behaviour can be 
controlled via the ignoreSSL flag.


```
    client := brocadevtm.NewVTMClient(
        serverUrl,  // https://<server>:<port>
        username,
        password,
        ignoreSSL,  // if server should ignore SSL certificate checking
        debug,      // if enhanced trace log is needed...
        header,     // a set of header key/value pairs to use, not mandatory
    )
```

### Handling Monitors

#### Importing proper packages

In order to work with monitors, import the monitor package section:

```
import (
    "github.com/sky-uk/go-brocade-vtm/api/monitor"
)
```


#### Getting the list of active monitors

```
    api := monitor.NewGetAll()
    err := client.Do(api)
    if err != nil {
        log.Fatal(err)
    }

    monitors := api.GetResponse()
```


#### Retrieving a single monitor

```
    api := monitor.NewGetSingleMonitor( name )
    err := client.Do(api)
    if err != nil {
        log.Fatal(err)
    }

    monitor := api.GetResponse()
```

#### Creating a monitor


```
    newMonitor := monitor.Monitor{
        ...
    }

    api := monitor.NewCreate( name, newMonitor )
    err := client.Do(api)
    if err != nil {
        log.Fatal(err)
    }

    monCreated := api.GetResponse()
```

#### Updating a monitor


```
    updateMonitor := monitor.Monitor{
        ...
    }

    api := monitor.NewUpdate( name, updateMonitor )
    err := client.Do(api)
    if err != nil {
        log.Fatal(err)
    }

    monUpdated := api.GetResponse()
```

#### Deleting a monitor

```
    api := monitor.NewDelete( name )
    err := client.Do(api)
    if err != nil {
        log.Fatal(err)
    } else {
        log.Printf("Monitor %s deleted", name)
    }
```

