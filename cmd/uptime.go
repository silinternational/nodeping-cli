package cmd

import (
	"github.com/spf13/cobra"
	"github.com/silinternational/nodeping-go-client"
	"fmt"
	"os"
	"github.com/silinternational/nodeping-cli/lib"
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
			fmt.Println("Error: The 'contact-group' flag is required (e.g. -g AppsDev).")
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

	cgID, err := lib.GetContactGroupIDFromName(contactGroupName, npClient)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	checkLabels, checkIDs, err := lib.GetCheckIDsAndLabels(cgID, npClient)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

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

	uptimes := lib.GetUptimes(checkIDs, start, end, npClient)

	for _, label := range checkLabels {
		fmt.Printf("%s, %v\n", label, uptimes[checkIDs[label]])
	}
}

