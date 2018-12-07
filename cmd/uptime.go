package cmd

import (
	"github.com/spf13/cobra"
	"github.com/silinternational/nodeping-go-client"
	"fmt"
	"os"
	"github.com/silinternational/nodeping-cli/lib"
	"sort"
	"time"
)

var contactGroupName string
var period string

var uptimeCmd = &cobra.Command{
	Use: "uptime",
	Short: "Get the uptime for checks",
	Long: "Get the uptime for all the checks for a certain Contact Group.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if contactGroupName == "" {
			fmt.Println("Error: The 'contact-group' flag is required (e.g. -g AppsDev).\n")
			os.Exit(1)
		}

		if period != "" {
			if ! lib.IsPeriodValid(period) {
				fmt.Printf(
					"Error: The period value is not valid. It must be one of these ...\n%v\n",
					lib.GetValidPeriods())
				os.Exit(1)
			}
		}
		runUptime()
	},
}

func init() {
	periods := lib.GetValidPeriods()

	rootCmd.AddCommand(uptimeCmd)
	uptimeCmd.Flags().StringVarP(
		&contactGroupName,
		"contact-group",
		"g",
		"",
		`Name of the Contact Group to retrieve uptime data for.`,
	)
	uptimeCmd.Flags().StringVarP(
		&period,
		"period",
		"p",
		"",
		fmt.Sprintf(`Name of the time period to get uptime values for ... %v`, periods),
	)
}

func runUptime() {
	npClient, err := nodeping.New(nodeping.ClientConfig{Token: nodepingToken})

	if err != nil {
		fmt.Printf("Error initializing cli: \n%s\n", err.Error())
		os.Exit(1)
	}

	contactGroups, err := npClient.ListContactGroups()

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
		fmt.Printf("Could not find contact group with name \"%s\"\n", contactGroupName)
		os.Exit(1)
	}

	fmt.Printf("Found contact group \"%s\" having ID \"%s\".\n", contactGroupName, cgID)

	checks, err := npClient.ListChecks()

	if err != nil {
		fmt.Printf("Error retrieving checks: \n%s\n", err.Error())
		os.Exit(1)
	}

	//fmt.Printf("First Check:\n%+v\n", checks[0])

	checkIDs := map[string]string{}
	checkLabels := []string{}

	for _, check := range checks {
		// Notifications is a list of maps with the contactGroup ID as keys
		for _, notfctn := range check.Notifications {
			foundOne := false
			for nKey := range notfctn {
				if nKey == cgID {
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

	uptimes := map[string]float32{}
	start := int64(0)
	end := int64(0)

	if period != "" {
		periodObject := lib.GetPeriodByName(period, 0)
		start = periodObject.From * 1000
		end = periodObject.To * 1000
	}

	fmt.Printf(
		"Period: %s. From: %v      To: %v\n\n",
		period,
		//time.Unix(start/1000, 0).Format(time.RFC822Z),
		time.Unix(start/1000, 0).UTC(),
		time.Unix(end/1000, 0).UTC(),
	)

	for _, checkID := range checkIDs {
		nextUptime, err := npClient.GetUptime(checkID, start, end)
		if err != nil {
			fmt.Printf("Error getting uptime for check ID %s.\n%s\n", checkID, err.Error())
			continue
		}
		uptimes[checkID] = nextUptime["total"].Uptime
		//fmt.Printf("Got Uptime Response ...\n   %+v\n", nextUptime)
	}

	sort.Strings(checkLabels)
	for _, label := range checkLabels {
		fmt.Printf("%s, %v\n", label, uptimes[checkIDs[label]])
	}
}