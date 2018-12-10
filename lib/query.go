package lib

import (
	"fmt"
	"github.com/silinternational/nodeping-go-client"
	"os"
	"sort"
)

func GetContactGroupIDFromName(contactGroupName string, npClient *nodeping.NodePingClient) (string, error) {

	contactGroups, err := npClient.ListContactGroups()

	if err != nil {
		return "", err
	}

	if err != nil {
		fmt.Printf("Error retrieving contact groups: \n%s\n", err.Error())
		os.Exit(1)
	}

	cgID := ""
	for cgKey, cg := range contactGroups{
		if cg.Name == contactGroupName {
			cgID = cgKey
			break
		}
	}

	if cgID == "" {
		return "", fmt.Errorf("Could not find contact group with name \"%s\"\n", contactGroupName)
	}

	return cgID, nil
}

func GetCheckIDsAndLabels(
	contactGroupID string,
	npClient *nodeping.NodePingClient,
) ([]string, map[string]string, error) {

	checkIDs := map[string]string{}
	checkLabels := []string{}

	checks, err := npClient.ListChecks()

	if err != nil {
		return checkLabels, checkIDs, err
	}

	//fmt.Printf("First Check:\n%+v\n", checks[0])


	for _, check := range checks {
		// Notifications is a list of maps with the contactGroup ID as keys
		for _, notfctn := range check.Notifications {
			foundOne := false
			for nKey := range notfctn {
				if nKey == contactGroupID {
					checkIDs[check.Label] = check.ID
					checkLabels = append(checkLabels, check.Label)
					foundOne = true
					break
				}
			}
			if foundOne {
				break
			}
		}
	}

	sort.Strings(checkLabels)
	return checkLabels, checkIDs, nil
}

func GetUptimes(
	checkIDs map[string]string,
	start, end int64,
	npClient *nodeping.NodePingClient,
) map[string]float32 {

	uptimes := map[string]float32{}

	for _, checkID := range checkIDs {
		nextUptime, err := npClient.GetUptime(checkID, start, end)
		if err != nil {
			fmt.Printf("Error getting uptime for check ID %s.\n%s\n", checkID, err.Error())
			continue
		}
		uptimes[checkID] = nextUptime["total"].Uptime
		//fmt.Printf("Got Uptime Response ...\n   %+v\n", nextUptime)
	}

	return uptimes
}