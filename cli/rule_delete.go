package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/rule"
	"github.com/sky-uk/go-rest-api"
	"net/http"
	"os"
)

var deleteRuleName string

func deleteRule(client *rest.Client, flagSet *flag.FlagSet) {

	if deleteRuleName == "" {
		fmt.Printf("\nError: name argument is required. Usage: -name rule-name")
		os.Exit(1)
	}

	deleteRuleAPI := rule.NewDelete(deleteRuleName)
	err := client.Do(deleteRuleAPI)
	if err != nil {
		fmt.Printf("\nError while deleting rule %s. Error: %+v", deleteRuleName, err)
		os.Exit(2)
	}
	httpResponseCode := deleteRuleAPI.StatusCode()
	if httpResponseCode != http.StatusNoContent {
		fmt.Printf("\nError while deleting rule %s. Received invalid http response code %d", deleteRuleName, httpResponseCode)
		os.Exit(3)
	}
	fmt.Printf("Successfully deleted rule %s\n", deleteRuleName)
}

func init() {
	deleteRuleFlags := flag.NewFlagSet("rule-delete", flag.ExitOnError)
	deleteRuleFlags.StringVar(&deleteRuleName, "name", "", "usage: -name rule-name")
	RegisterCliCommand("rule-delete", deleteRuleFlags, deleteRule)
}
