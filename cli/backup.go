package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"os"
	"time"
)

func getBackup(client *api.Client, flagSet *flag.FlagSet) {
	trafficManager := flagSet.Lookup("tm").Value.String()
	backupName := flagSet.Lookup("backup-name").Value.String()

	if backupName == "" {
		fmt.Println("backup-name must be set - usage: -backup-name <a valid backup name>")
		os.Exit(2)
	}

	res, err := client.GetBackup(trafficManager, backupName)
	if err != nil {
		fmt.Println("Error getting all configuration resources: ", err)
		os.Exit(1)
	}

	var backupMap, backupInfo map[string]interface{}
	backupMap = res["properties"].(map[string]interface{})
	backupInfo = backupMap["backup"].(map[string]interface{})
	backupInfo["time_stamp"] = time.Unix(int64(backupInfo["time_stamp"].(float64)), 0)
	PrettyPrintSingle(backupInfo)
}

func getAllBackups(client *api.Client, flagSet *flag.FlagSet) {
	res, err := client.GetAllBackups(flagSet.Lookup("tm").Value.String())
	if err != nil {
		fmt.Println("Error getting backups: ", err)
		os.Exit(1)
	}
	PrettyPrintMany([]string{"name", "href"}, res)
}

func createBackup(client *api.Client, flagSet *flag.FlagSet) {
	backupName := flagSet.Lookup("backup-name").Value.String()

	if backupName == "" {
		fmt.Println("backup-name must be set - usage: -backup-name <a valid backup name>")
		os.Exit(2)
	}
	res, err := client.CreateBackup(flagSet.Lookup("tm").Value.String(), backupName, flagSet.Lookup("description").Value.String())
	if err != nil {
		fmt.Println("Error creating backup: ", err)
		os.Exit(1)
	}
	var backupMap, backupInfo map[string]interface{}
	backupMap = res["properties"].(map[string]interface{})
	backupInfo = backupMap["backup"].(map[string]interface{})
	backupInfo["name"] = backupName
	backupInfo["time_stamp"] = time.Unix(int64(backupInfo["time_stamp"].(float64)), 0)
	PrettyPrintSingle(backupInfo)
}

func deleteBackup(client *api.Client, flagSet *flag.FlagSet) {
	backupName := flagSet.Lookup("backup-name").Value.String()
	if backupName == "" {
		fmt.Println("backup-name must be set - usage: -backup-name <a valid backup name>")
		os.Exit(2)
	}
	err := client.DeleteBackup(flagSet.Lookup("tm").Value.String(), backupName)

	if err != nil {
		fmt.Println("Error deleting backup: ", err)
		os.Exit(1)
	}
	fmt.Println("Backup: " + backupName + " deleted")
}

func restoreBackup(client *api.Client, flagSet *flag.FlagSet) {
	backupName := flagSet.Lookup("backup-name").Value.String()
	if backupName == "" {
		fmt.Println("backup-name must be set - usage: -backup-name <a valid backup name>")
		os.Exit(2)
	}
	res, err := client.RestoreBackup(flagSet.Lookup("tm").Value.String(), backupName)
	if err != nil {
		fmt.Println("Error restoring backup: ", err)
		os.Exit(1)
	}
	var backupMap, backupInfo map[string]interface{}
	backupMap = res["properties"].(map[string]interface{})
	backupInfo = backupMap["backup"].(map[string]interface{})
	backupInfo["name"] = backupName
	backupInfo["time_stamp"] = time.Unix(int64(backupInfo["time_stamp"].(float64)), 0)
	PrettyPrintSingle(backupInfo)
	fmt.Println("Backup: " + backupName + " restored")
}

func downloadBackup(client *api.Client, flagSet *flag.FlagSet) {
	backupName := flagSet.Lookup("backup-name").Value.String()
	if backupName == "" {
		fmt.Println("backup-name must be set - usage: -backup-name <a valid backup name>")
		os.Exit(2)
	}

	err := client.DownloadBackup(flagSet.Lookup("tm").Value.String(), backupName, "backups/")

	if err != nil {
		fmt.Println("Error restoring backup: ", err)
		os.Exit(1)
	}
}

func uploadBackup(client *api.Client, flagSet *flag.FlagSet) {
	filePath := flagSet.Lookup("backuppath").Value.String()

	if filePath == "" {
		fmt.Println("backuppath must be set - usage: -backuppath <a valid file location>")
		os.Exit(2)
	}

	err := client.UploadBackup(flagSet.Lookup("tm").Value.String(), flagSet.Lookup("backuppath").Value.String())

	if err != nil {
		fmt.Println("Error restoring backup: ", err)
		os.Exit(1)
	}
}

func init() {
	backupFlags := flag.NewFlagSet("get-backup", flag.ExitOnError)
	backupFlags.String("backup-name", "", "usage: -backup-name <backup name>")
	backupFlags.String("tm", "local_tm", "usage: -tm <traffic manager name>")
	backupFlags.String("description", "", "usage: -description <backup description>")
	backupFlags.String("backuppath", "", "usage: -backuppath <path to backup>")
	RegisterCliCommand("get-backup", backupFlags, getBackup)
	RegisterCliCommand("get-all-backups", backupFlags, getAllBackups)
	RegisterCliCommand("create-backup", backupFlags, createBackup)
	RegisterCliCommand("delete-backup", backupFlags, deleteBackup)
	RegisterCliCommand("restore-backup", backupFlags, restoreBackup)
	RegisterCliCommand("upload-backup", backupFlags, uploadBackup)
}
