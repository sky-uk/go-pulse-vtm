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

var updateRuleName string
var updateTrafficScriptFile string

func updateRule(client *brocadevtm.VTMClient, flagSet *flag.FlagSet) {

	if updateRuleName == "" {
		fmt.Printf("\nName argument is required. Usage: -name vtm-rule-name\n")
		os.Exit(1)
	}

	if updateTrafficScriptFile == "" {
		fmt.Printf("\nRule argument is required. Usage: -script trafficScript\n")
		os.Exit(1)
	}

	updateTrafficScriptFile, fileErr := ioutil.ReadFile(updateTrafficScriptFile)

	if fileErr != nil {
		fmt.Printf("\nError reading file %s\n", updateTrafficScriptFile)
		os.Exit(2)
	}

	updateRuleAPI := rule.NewCreate(updateRuleName, updateTrafficScriptFile)
	err := client.Do(updateRuleAPI)
	if err != nil {
		fmt.Printf("\nError occurred while creating rule %s. Error: %+v\n", updateRuleName, err)
		os.Exit(3)
	}
	httpResponseCode := updateRuleAPI.StatusCode()
	if httpResponseCode == http.StatusCreated || httpResponseCode == http.StatusNoContent {
		fmt.Printf("Successfully updated rule %s\n", updateRuleName)
	} else {
		fmt.Printf("\nError occurred while updating rule %s. Received invalid http response code %d\n", updateRuleName, httpResponseCode)
		os.Exit(4)
	}
}

func init() {
	updateRuleFlags := flag.NewFlagSet("rule-update", flag.ExitOnError)
	updateRuleFlags.StringVar(&updateRuleName, "name", "", "usage: -name vtm-rule-name")
	updateRuleFlags.StringVar(&updateTrafficScriptFile, "script", "", "usage: -script location-of-traffic-script-file")
	RegisterCliCommand("rule-update", updateRuleFlags, updateRule)
}
