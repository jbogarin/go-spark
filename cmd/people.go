package cmd

import (
	"log"

	"github.com/jbogarin/go-cisco-spark/ciscospark"
	"github.com/spf13/cobra"
)

var peopleName, peopleEmail, personID string

// peopleCmd represents the people command
var peopleCmd = &cobra.Command{
	Use:   "people",
	Short: "People are registered users of the Spark application.",
	Long:  `People are registered users of the Spark application. Currently, people can only be searched with this API.`,
}

// peopleMeCmd represents the people GET me command
var peopleMeCmd = &cobra.Command{
	Use:   "me",
	Short: "Get my details",
	Long:  `Show the profile for the authenticated user.`,
	Run: func(cmd *cobra.Command, args []string) {
		me, response, err := SparkClient.People.GetMe()
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}
		PrintResponseFormat(me)
	},
}

// peopleListCmd represents the people GET command
var peopleListCmd = &cobra.Command{
	Use:   "list",
	Short: "List people",
	Long:  `List people in your organization.`,
	Run: func(cmd *cobra.Command, args []string) {
		queryParams := &ciscospark.GetPeopleQueryParams{
			Max: Max,
		}

		if peopleName != "" {
			queryParams.DisplayName = peopleName
		}

		if peopleEmail != "" {
			queryParams.Email = peopleEmail
		}

		people, response, err := SparkClient.People.Get(queryParams)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintResponseFormat(people)
	},
}

// peopleGetCmd represents the people GET/<id> command
var peopleGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get person details",
	Long: `Shows details for a person, by ID.

Specify the person ID with the -i/--id flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		person, response, err := SparkClient.People.GetPerson(personID)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintResponseFormat(person)
	},
}

func init() {
	RootCmd.AddCommand(peopleCmd)
	peopleCmd.AddCommand(peopleMeCmd)
	peopleCmd.AddCommand(peopleListCmd)
	peopleCmd.AddCommand(peopleGetCmd)

	peopleListCmd.Flags().StringVarP(&peopleName, "name", "n", "", "List people whose name starts with this string.")
	peopleListCmd.Flags().StringVarP(&peopleEmail, "email", "e", "", "List people with this email address.")

	peopleGetCmd.Flags().StringVarP(&personID, "id", "i", "", "The person ID")
}
