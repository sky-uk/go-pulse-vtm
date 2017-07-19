package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm"
	"github.com/sky-uk/go-brocade-vtm/api/rule"
	"io/ioutil"
	"net/http"
	"os"
)

var ruleName string
var trafficScriptFile string

func createRule(client *brocadevtm.VTMClient, flagSet *flag.FlagSet) {

	if ruleName == "" {
		fmt.Printf("\nName argument is required. Usage: -name vtm-rule-name\n")
		os.Exit(1)
	}

	if trafficScriptFile == "" {
		fmt.Printf("\nRule argument is required. Usage: -script trafficScript\n")
		os.Exit(1)
	}

	trafficScript, fileErr := ioutil.ReadFile(trafficScriptFile)

	if fileErr != nil {
		fmt.Printf("\nError reading file %s\n", trafficScriptFile)
		os.Exit(2)
	}

	createRuleAPI := rule.NewCreate(ruleName, trafficScript)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/octet-stream"
	headers["Content-Transfer-Encoding"] = "text"
	client.Headers = headers
	err := client.Do(createRuleAPI)
	if err != nil {
		fmt.Printf("\nError occurred while creating rule %s. Error: %+v\n", ruleName, err)
		os.Exit(3)
	}
	httpResponseCode := createRuleAPI.StatusCode()
	if httpResponseCode == http.StatusCreated || httpResponseCode == http.StatusNoContent {
		fmt.Printf("Successfully created new rule %s\n", ruleName)
	} else {
		fmt.Printf("\nError occurred while creating rule %s. Received invalid http response code %d\n", ruleName, httpResponseCode)
		os.Exit(4)
	}
}

func init() {
	createRuleFlags := flag.NewFlagSet("rule-create", flag.ExitOnError)
	createRuleFlags.StringVar(&ruleName, "name", "", "usage: -name vtm-rule-name")
	createRuleFlags.StringVar(&trafficScriptFile, "script", "", "usage: -script location-of-traffic-script-file")
	RegisterCliCommand("rule-create", createRuleFlags, createRule)
}
