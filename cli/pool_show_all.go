package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/pool"
	"github.com/sky-uk/go-rest-api"
	"os"
)

func showAllPools(client *rest.Client, flagSet *flag.FlagSet) {

	if apiVersion != "" {
		pool.PoolEndpoint = "/api/tm/" + apiVersion + "/config/active/pools/"
	}

	getAllPoolsAPI := pool.NewGetAll()
	err := client.Do(getAllPoolsAPI)
	if err != nil {
		fmt.Printf("\nError retreiving the list of pools\n")
		os.Exit(1)
	}

	poolList := getAllPoolsAPI.ResponseObject().(*pool.LBPoolList)
	rows := []map[string]interface{}{}
	headers := []string{"Name", "HRef"}

	for _, poolItem := range poolList.ChildPools {
		row := map[string]interface{}{}
		row["Name"] = poolItem.Name
		row["HRef"] = poolItem.Href
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)
}

func init() {
	showAllPoolFlags := flag.NewFlagSet("pool-show-all", flag.ExitOnError)
	showAllPoolFlags.StringVar(&apiVersion, "apiversion", "", "usage: -apiversion 3.8")
	RegisterCliCommand("pool-show-all", showAllPoolFlags, showAllPools)
}
