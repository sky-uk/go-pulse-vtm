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
| Virtual Server          |   Y    |   Y   |    Y    |   Y    |
| Rule                    |   Y    |   Y   |    Y    |   Y    |

### Notes
[1] : Currently only HTTP monitoring is supported


## Usage

For usage please see the wiki https://github.com/sky-uk/go-brocade-vtm/wiki
