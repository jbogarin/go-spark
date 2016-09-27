package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jbogarin/go-cisco-spark/ciscospark"
	"github.com/spf13/cobra"
)

var name string

// peopleCmd represents the people command
var peopleCmd = &cobra.Command{
	Use:   "people",
	Short: "Gets people information",
	Long:  `Gets people information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())
	},
}

// meCmd represents the people GET me command
var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Get my information",
	Long:  `Gets my user information.`,
	Run: func(cmd *cobra.Command, args []string) {
		me, _, err := SparkClient.People.GetMe()
		if err != nil {
			log.Fatal(err)
		}
		meJSON, err := json.MarshalIndent(me, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(meJSON))

	},
}

// listCmd represents the people GET command
var peopleListCmd = &cobra.Command{
	Use:   "list",
	Short: "Gets a list of people",
	Long:  `Gets a list of people in my organization`,
	Run: func(cmd *cobra.Command, args []string) {
		queryParams := &ciscospark.GetPeopleQueryParams{
			DisplayName: name,
			Max:         Max,
		}

		people, _, err := SparkClient.People.Get(queryParams)
		if err != nil {
			log.Fatal(err)
		}

		peopleJSON, err := json.MarshalIndent(people, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(peopleJSON))
	},
}

func init() {
	RootCmd.AddCommand(peopleCmd)
	peopleCmd.AddCommand(meCmd)
	peopleCmd.AddCommand(peopleListCmd)

	peopleListCmd.Flags().StringVarP(&name, "name", "n", "", "displayName to search for")
}
