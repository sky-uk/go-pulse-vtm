package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/pool"
	"github.com/sky-uk/go-rest-api"
	"net/http"
	"os"
)

func deletePool(client *rest.Client, flagSet *flag.FlagSet) {

	poolName := flagSet.Lookup("name").Value.String()
	if poolName == "" {
		fmt.Printf("\nError: -name argument required\n")
		os.Exit(1)
	}
	deletePoolAPI := pool.NewDelete(poolName)
	err := client.Do(deletePoolAPI)
	if err != nil && deletePoolAPI.StatusCode() != http.StatusNotFound {
		fmt.Printf("\nError deleting pool %s\n", poolName)
		os.Exit(2)
	}
	fmt.Printf("\nSuccessfully deleted pool %s\n", poolName)
}

func init() {
	deletePoolFlags := flag.NewFlagSet("pool-delete", flag.ExitOnError)
	deletePoolFlags.String("name", "", "usage: -name pool-name")
	RegisterCliCommand("pool-delete", deletePoolFlags, deletePool)
}
