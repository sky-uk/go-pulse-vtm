package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-brocade-vtm/api/monitor"
	"github.com/sky-uk/go-rest-api"
	"os"
)

var showMonitorName string

func showMonitor(client *rest.Client, flagSet *flag.FlagSet) {

	if showMonitorName == "" {
		fmt.Printf("\nError: name argument required\n")
		os.Exit(1)
	}

	readMonitorAPI := monitor.NewGet(showMonitorName)
	err := client.Do(readMonitorAPI)
	if err != nil {
		fmt.Printf("\nError whilst retrieving monitor %s. Error: %+v\n", showMonitorName, err)
		errObj := *readMonitorAPI.ErrorObject().(*api.VTMError)
		PrettyPrintErrorObj(errObj)
		os.Exit(2)
	}
	response := readMonitorAPI.ResponseObject().(*monitor.Monitor)
	row := map[string]interface{}{}
	row["Name"] = showMonitorName
	row["Type"] = response.Properties.Basic.Type
	row["Delay"] = response.Properties.Basic.Delay
	row["Timeout"] = response.Properties.Basic.Timeout
	row["Failures"] = response.Properties.Basic.Failures
	row["Verbose"] = *response.Properties.Basic.Verbose
	row["UseSSL"] = *response.Properties.Basic.UseSSL
	row["HTTP-Host-Header"] = response.Properties.HTTP.HostHeader
	row["HTTP-Path"] = response.Properties.HTTP.URIPath
	row["HTTP-Authentication"] = response.Properties.HTTP.Authentication
	row["HTTP-Body-Regex"] = response.Properties.HTTP.BodyRegex
	row["RTSP-Body-Regex"] = response.Properties.RTSP.BodyRegex
	row["RTSP-Path"] = response.Properties.RTSP.URIPath
	row["RTSP-Status-regex"] = response.Properties.RTSP.StatusRegex
	row["SCRIPT-Arguments"] = response.Properties.SCRIPT.Arguments
	row["SCRIPT-Program"] = response.Properties.SCRIPT.Program
	row["SIP-Body-Regex"] = response.Properties.SIP.BodyRegex
	row["SIP-Status-Regex"] = response.Properties.SIP.StatusRegex
	row["SIP-Transport"] = response.Properties.SIP.Transport
	row["TCP-Close-String"] = response.Properties.TCP.CloseString
	row["TCP-Max-Response-Len"] = response.Properties.TCP.MaxResponseLen
	row["TCP-Response-Regex"] = response.Properties.TCP.ResponseRegex
	row["TCP-Write-String"] = response.Properties.TCP.WriteString
	row["UDP-Accept-All"] = response.Properties.UDP.AcceptAll
	PrettyPrintSingle(row)
}

func init() {
	showMonitorFlags := flag.NewFlagSet("monitor-show", flag.ExitOnError)
	showMonitorFlags.StringVar(&showMonitorName, "name", "", "usage: -name monitor-name")
	RegisterCliCommand("monitor-show", showMonitorFlags, showMonitor)
}
