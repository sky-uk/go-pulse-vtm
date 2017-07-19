package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm"
	"github.com/sky-uk/go-brocade-vtm/api/rule"
	"net/http"
	"os"
)

var readRuleName string

func showRule(client *brocadevtm.VTMClient, flagSet *flag.FlagSet) {

	var trafficScriptRule rule.TrafficScriptRule

	if readRuleName == "" {
		fmt.Printf("\nError: name argument required. Useage: -name rule-name")
		os.Exit(1)
	}

	readAPI := rule.NewGetRule(readRuleName)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/octet-stream"
	headers["Content-Transfer-Encoding"] = "text"
	client.Headers = headers
	err := client.Do(readAPI)
	if err != nil {
		fmt.Printf("\nError retrieving rule %s from API. Error %+v", readRuleName, err)
		os.Exit(2)
	}
	httpResponseCode := readAPI.StatusCode()
	if httpResponseCode != http.StatusOK {
		fmt.Printf("\nError API returned invalid HTTP response code %d for rule %s", httpResponseCode, readRuleName)
		os.Exit(3)
	}

	trafficScriptRule.Name = readRuleName
	trafficScriptRule.Script = readAPI.GetResponse()

	fmt.Printf("The traffic script for the %s rule is: \n", trafficScriptRule.Name)
	fmt.Printf("%s", trafficScriptRule.Script)

}

func init() {
	readRuleFlags := flag.NewFlagSet("rule-show", flag.ExitOnError)
	readRuleFlags.StringVar(&readRuleName, "name", "", "usage: -name rule-name")
	RegisterCliCommand("rule-show", readRuleFlags, showRule)
}
