package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-brocade-vtm/api/monitor"
	"github.com/sky-uk/go-rest-api"
	"os"
)

var updateMonitorName string
var updateMonitorStruct monitor.Monitor
var updateUseSSL, updateVerbose bool

func updateMonitor(client *rest.Client, flagSet *flag.FlagSet) {

	if updateMonitorName == "" {
		fmt.Printf("\nError -name argument required. Usage: -name monitor-name\n")
		os.Exit(1)
	}

	if apiVersion != "" {
		monitor.MonitorEndpoint = "/api/tm/" + apiVersion + "/config/active/monitors/"
	}

	updateMonitorStruct.Properties.Basic.UseSSL = &updateUseSSL
	updateMonitorStruct.Properties.Basic.Verbose = &updateVerbose
	updateMonitorAPI := monitor.NewUpdate(updateMonitorName, updateMonitorStruct)
	err := client.Do(updateMonitorAPI)
	if err != nil {
		fmt.Printf("\nError updating monitor %sError: %+v\n", updateMonitorName, err)
		errObj := *updateMonitorAPI.ErrorObject().(*api.VTMError)
		PrettyPrintErrorObj(errObj)
		os.Exit(2)
	}
	fmt.Printf("\nSuccessfully updated monitor %s\n", updateMonitorName)
}

func init() {
	updateMonitorFlags := flag.NewFlagSet("monitor-update", flag.ExitOnError)
	updateMonitorFlags.StringVar(&updateMonitorName, "name", "", "usage: -name monitor-name")
	updateMonitorFlags.StringVar(&updateMonitorStruct.Properties.Basic.Type, "type", "http", "usage: -type monitor-type")
	updateMonitorFlags.UintVar(&updateMonitorStruct.Properties.Basic.Delay, "delay", 3, "usage: -delay 3")
	updateMonitorFlags.UintVar(&updateMonitorStruct.Properties.Basic.Timeout, "timeout", 3, "usage: -timeout 3")
	updateMonitorFlags.UintVar(&updateMonitorStruct.Properties.Basic.Failures, "failures", 3, "usage: -failures 3")
	updateMonitorFlags.BoolVar(&updateVerbose, "verbose", false, "usage: -verbose")
	updateMonitorFlags.BoolVar(&updateUseSSL, "use-ssl", false, "usage: -use-ssl")
	updateMonitorFlags.StringVar(&updateMonitorStruct.Properties.HTTP.HostHeader, "http-host-header", "", "usage: -http-host-header a-header")
	updateMonitorFlags.StringVar(&updateMonitorStruct.Properties.HTTP.URIPath, "http-path", "/", "usage: -http-path /healthcheck")
	updateMonitorFlags.StringVar(&updateMonitorStruct.Properties.HTTP.Authentication, "authentication", "", "usage: -authentication basic-auth-string")
	updateMonitorFlags.StringVar(&updateMonitorStruct.Properties.HTTP.BodyRegex, "http-body-regex", `^[234][0-9][0-9]$`, `usage: -http-body-regex [234][0-9][0-9]$`)
	updateMonitorFlags.StringVar(&apiVersion, "apiversion", "", "usage: -apiversion 3.8")
	RegisterCliCommand("monitor-update", updateMonitorFlags, updateMonitor)
}
