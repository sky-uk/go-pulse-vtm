package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"os"
	"time"
)

// ExecFunc executes the function for cli.
type ExecFunc func(client *api.Client, flagSet *flag.FlagSet)

// Command struct - defines a cli command with flags and exec
type Command struct {
	flagSet *flag.FlagSet
	exec    ExecFunc
}

var (
	brocadeVTMServer   string
	debug              bool
	timeout            time.Duration
	brocadeVTMUsername string
	brocadeVTMPassword string
	brocadeAPIVersion  string

	commandMap = make(map[string]Command, 0)
)

// RegisterCliCommand - allows additional cli commands to be registered.
func RegisterCliCommand(name string, flagSet *flag.FlagSet, exec ExecFunc) {
	commandMap[name] = Command{flagSet, exec}
}

// InitFlags - initiall cli flags.
func InitFlags() {
	flag.StringVar(&brocadeVTMServer, "server", os.Getenv("BROCADEVTM_SERVER"),
		"Brocade vTM API server hostname or address. (Env: BROCADEVTM_SERVER)")
	flag.StringVar(&brocadeVTMUsername, "username", os.Getenv("BROCADEVTM_USERNAME"),
		"Brocade vTM authentication username (Env: BROCADEVTM_USERNAME)")
	flag.StringVar(&brocadeVTMPassword, "password", os.Getenv("BROCADEVTM_PASSWORD"),
		"Brocade vTM authentication password (Env: BROCADEVTM_PASSWORD)")
	flag.StringVar(&brocadeAPIVersion, "api_version", "5.1",
		"Brocade vTM REST API version")
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
	headers["Content-Type"] = contentType(command)

	if isFileDownload(command) {
		headers["Accept"] = "application/x-tar"
	}

	params := api.Params{
		APIVersion: brocadeAPIVersion,
		Server:     brocadeVTMServer,
		Username:   brocadeVTMUsername,
		Password:   brocadeVTMPassword,
		IgnoreSSL:  true,
		Debug:      debug,
		Headers:    headers,
	}

	client, err := api.Connect(params)
	if err != nil {
		fmt.Println("Error connecting to the BrocadevTM server")
		os.Exit(1)
	}

	cmd.exec(client, flagSet)
}

func contentType(command string) string {
	switch command {
	case "upload-backup", "download-backup":
		return "application/x-tar"
	default:
		return "application/json"
	}
}

func isFileDownload(command string) bool {
	switch command {
	case "upload-backup":
		return true
	default:
		return false
	}
}
