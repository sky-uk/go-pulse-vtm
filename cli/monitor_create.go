package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-brocade-vtm/api/monitor"
	"github.com/sky-uk/go-rest-api"
	"os"
)

var monitorName string
var createMonitorStruct monitor.Monitor
var monitorUseSSL, monitorVerbose bool

func createMonitor(client *rest.Client, flagSet *flag.FlagSet) {

	if monitorName == "" {
		fmt.Printf("\nError -name argument required\n")
		os.Exit(1)
	}

	createMonitorStruct.Properties.Basic.UseSSL = &monitorUseSSL
	createMonitorStruct.Properties.Basic.Verbose = &monitorVerbose
	createMonitorAPI := monitor.NewCreate(monitorName, createMonitorStruct)
	err := client.Do(createMonitorAPI)
	if err != nil {
		fmt.Printf("\nError creating monitor %s. Error: %+v\n", monitorName, err)
		errObj := *createMonitorAPI.ErrorObject().(*api.VTMError)
		PrettyPrintErrorObj(errObj)
		os.Exit(2)
	}
	fmt.Printf("\nSuccessfully created monitor %s\n", monitorName)

}

func init() {
	createMonitorFlags := flag.NewFlagSet("monitor-create", flag.ExitOnError)
	createMonitorFlags.StringVar(&monitorName, "name", "", "usage: -name monitor-name")
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.Basic.Type, "type", "http", "usage: -type monitor-type")
	createMonitorFlags.UintVar(&createMonitorStruct.Properties.Basic.Delay, "delay", 3, "usage: -delay 3")
	createMonitorFlags.UintVar(&createMonitorStruct.Properties.Basic.Timeout, "timeout", 3, "usage: -timeout 3")
	createMonitorFlags.UintVar(&createMonitorStruct.Properties.Basic.Failures, "failures", 3, "usage: -failures 3")
	createMonitorFlags.BoolVar(&monitorVerbose, "verbose", false, "usage: -verbose")
	createMonitorFlags.BoolVar(&monitorUseSSL, "use-ssl", false, "usage: -use-ssl")
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.HTTP.HostHeader, "http-host-header", "", "usage: -http-host-header a-header")
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.HTTP.URIPath, "http-path", "/", "usage: -http-path /healthcheck")
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.HTTP.Authentication, "authentication", "", "usage: -authentication basic-auth-string")
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.HTTP.BodyRegex, "http-body-regex", `^[234][0-9][0-9]$`, `usage: -http-body-regex ^[234][0-9][0-9]$`)
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.RTSP.BodyRegex, "rtsp-body-regex", "", `usage: -rtsp-body-regex ""`)
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.RTSP.URIPath, "rtsp-path", "/", `usage: -rstp-path /`)
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.RTSP.StatusRegex, "rtsp-status-regex", `^[234][0-9][0-9]$`, `usage: -rtsp-status-regex ^[234][0-9][0-9]$`)
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.SCRIPT.Arguments, "script-arguments", "", `usage: -script-arguments ""`)
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.SCRIPT.Program, "script-program", "", `usage: -script-program ""`)
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.SIP.BodyRegex, "sip-body-regex", "", `usage: -sip-body-regex ""`)
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.SIP.StatusRegex, "sip-status-regex", `^[234][0-9][0-9]$`, `usage: -sip-status-regex ^[234][0-9][0-9]$`)
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.SIP.Transport, "sip-transport", "udp", `usage: -sip-transport "udp"`)
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.TCP.CloseString, "tcp-close-string", "", `usage: -tcp-close-string ""`)
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.TCP.WriteString, "tcp-write-string", "", `usage: -tcp-write-string ""`)
	createMonitorFlags.UintVar(&createMonitorStruct.Properties.TCP.MaxResponseLen, "tcp-max-response-len", 2048, `usage: -tcp-max-response-len 2048`)
	createMonitorFlags.StringVar(&createMonitorStruct.Properties.TCP.ResponseRegex, "tcp-response-regex", ".+", `usage: -tcp-response-regex ".+"`)
	RegisterCliCommand("monitor-create", createMonitorFlags, createMonitor)
}
