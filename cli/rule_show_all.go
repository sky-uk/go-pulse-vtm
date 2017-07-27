package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/rule"
	"github.com/sky-uk/go-rest-api"
	"os"
)

func showAllRules(client *rest.Client, flagSet *flag.FlagSet) {

	ruleShowAllAPI := rule.NewGetAll()
	err := client.Do(ruleShowAllAPI)
	if err != nil {
		fmt.Printf("\nError retrieving the rule list: %+v", err)
		os.Exit(1)
	}

	ruleList := ruleShowAllAPI.ResponseObject().(*rule.Rules)
	rows := []map[string]interface{}{}
	headers := []string{"Name", "HREF"}

	for _, rule := range ruleList.Children {
		row := map[string]interface{}{}
		row["Name"] = rule.Name
		row["HREF"] = rule.HRef
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)

}

func init() {
	showAllRulesFlags := flag.NewFlagSet("rule-show-all", flag.ExitOnError)
	RegisterCliCommand("rule-show-all", showAllRulesFlags, showAllRules)
}
