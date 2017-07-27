package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/rule"
	"github.com/sky-uk/go-rest-api"
	"os"
)

var readRuleName string

func showRule(client *rest.Client, flagSet *flag.FlagSet) {

	var trafficScriptRule rule.TrafficScriptRule
	headers := make(map[string]string)
	headers["Content-Type"] = "application/octet-stream"
	headers["Content-Transfer-Encoding"] = "BINARY"
	client.Headers = headers

	if readRuleName == "" {
		fmt.Printf("\nError: name argument required. Useage: -name rule-name")
		os.Exit(1)
	}

	readAPI := rule.NewGet(readRuleName)
	err := client.Do(readAPI)
	if err != nil {
		fmt.Printf("\nError retrieving rule %s from API. Error %+v", readRuleName, err)
		os.Exit(2)
	}

	trafficScriptRule.Name = readRuleName
	trafficScriptBytes := readAPI.ResponseObject().(*[]byte)
	trafficScriptRule.Script = string(*trafficScriptBytes)

	fmt.Printf("The traffic script for the %s rule is: \n", trafficScriptRule.Name)
	fmt.Printf("%s", trafficScriptRule.Script)

}

func init() {
	readRuleFlags := flag.NewFlagSet("rule-show", flag.ExitOnError)
	readRuleFlags.StringVar(&readRuleName, "name", "", "usage: -name rule-name")
	RegisterCliCommand("rule-show", readRuleFlags, showRule)
}
