package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-pulse-vtm/api"
	"os"
	"time"
)

var apiVersion string

// ExecFunc executes the function for cli.
type ExecFunc func(client *api.Client, flagSet *flag.FlagSet)

// Command struct - defines a cli command with flags and exec
type Command struct {
	flagSet *flag.FlagSet
	exec    ExecFunc
}

var (
	pulseVTMServer   string
	debug            bool
	timeout          time.Duration
	pulseVTMUsername string
	pulseVTMPassword string
	pulseAPIVersion  string

	commandMap = make(map[string]Command, 0)
)

// RegisterCliCommand - allows additional cli commands to be registered.
func RegisterCliCommand(name string, flagSet *flag.FlagSet, exec ExecFunc) {
	commandMap[name] = Command{flagSet, exec}
}

// InitFlags - initiall cli flags.
func InitFlags() {
	flag.StringVar(&pulseVTMServer, "server", os.Getenv("PULSEVTM_SERVER"),
		"Pulse vTM API server hostname or address. (Env: PULSEVTM_SERVER)")
	flag.StringVar(&pulseVTMUsername, "username", os.Getenv("PULSEVTM_USERNAME"),
		"Pulse vTM authentication username (Env: PULSEVTM_USERNAME)")
	flag.StringVar(&pulseVTMPassword, "password", os.Getenv("PULSEVTM_PASSWORD"),
		"Pulse vTM authentication password (Env: PULSEVTM_PASSWORD)")
	flag.StringVar(&pulseAPIVersion, "api_version", "3.8",
		"Pulse vTM REST API version")
	flag.BoolVar(&debug, "debug", false, "Debug output. Default:false")
	flag.DurationVar(&timeout, "timeout", 0, "Client timeout value. Default: 0")

}

func usage() {
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "  Commands:\n")
	for name := range commandMap {
		fmt.Fprintf(os.Stderr, "    %s\n", name)
	}
}

func main() {
	InitFlags()
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
		os.Exit(2)
	}

	command := flag.Arg(0)
	cmd, inMap := commandMap[command]
	if !inMap {
		usage()
		os.Exit(2)
	}

	flagSet := cmd.flagSet
	if flag.NArg() > 1 {
		flagSet.Parse(flag.Args()[1:])
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	params := api.Params{
		APIVersion: pulseAPIVersion,
		Server:     pulseVTMServer,
		Username:   pulseVTMUsername,
		Password:   pulseVTMPassword,
		IgnoreSSL:  true,
		Debug:      debug,
	}

	client, err := api.Connect(params)
	if err != nil {
		fmt.Println("Error connecting to the PulsevTM server")
		os.Exit(1)
	}

	cmd.exec(client, flagSet)
}
