package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-rest-api"
	"os"
	"time"
)

// ExecFunc executes the function for cli.
type ExecFunc func(client *rest.Client, flagSet *flag.FlagSet)

// Command struct - defines a cli command with flags and exec
type Command struct {
	flagSet *flag.FlagSet
	exec    ExecFunc
}

var (
	/*
	 * InfoBlox API server
	 */
	brocadeVTMServer string
	brocadeVTMPort   int
	debug            bool
	timeout          time.Duration

	/*
	 * Authentication
	 */
	brocadeVTMUsername string
	brocadeVTMPassword string

	commandMap = make(map[string]Command, 0)
)

// RegisterCliCommand - allows additional cli commands to be registered.
func RegisterCliCommand(name string, flagSet *flag.FlagSet, exec ExecFunc) {
	commandMap[name] = Command{flagSet, exec}
}

// InitFlags - initiall cli flags.
func InitFlags() {
	flag.StringVar(&brocadeVTMServer, "server", "https://"+os.Getenv("BROCADEVTM_SERVER"),
		"Brocade vTM API server hostname or address. (Env: BROCADEVTM_SERVER)")
	flag.IntVar(&brocadeVTMPort, "port", 443,
		"Brocade vTM API server port. Default:443")
	flag.StringVar(&brocadeVTMUsername, "username", os.Getenv("BROCADEVTM_USERNAME"),
		"Brocade vTM authentication username (Env: BROCADEVTM_USERNAME)")
	flag.StringVar(&brocadeVTMPassword, "password", os.Getenv("BROCADEVTM_PASSWORD"),
		"Brocade vTM authentication password (Env: BROCADEVTM_PASSWORD)")
	flag.BoolVar(&debug, "debug", false, "Debug output. Default:false")
	flag.DurationVar(&timeout, "timeout", 300, "Client timeout value. Default: 300")

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
	flagSet.Parse(flag.Args()[1:])

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	client := rest.Client{brocadeVTMServer, brocadeVTMUsername, brocadeVTMPassword, true, debug, headers, timeout}

	cmd.exec(&client, flagSet)
}
