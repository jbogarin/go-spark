package cmd

import (
	"log"

	"github.com/jbogarin/go-cisco-spark/ciscospark"
	"github.com/spf13/cobra"
)

var organizationID string

// organizationsCmd represents the organizations command
var organizationsCmd = &cobra.Command{
	Use:   "organizations",
	Short: "A set of people in Cisco Spark.",
	Long:  `A set of people in Cisco Spark. Organizations may manage other organizations or be managed themselves. This organizations resource can be accessed only by an admin.`,
}

// organizationsListCmd represents the organizations GET command
var organizationsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List organizations",
	Long:  `List organizations in your organization.`,
	Run: func(cmd *cobra.Command, args []string) {
		queryParams := &ciscospark.GetOrganizationsQueryParams{
			Max: Max,
		}

		organizations, response, err := SparkClient.Organizations.Get(queryParams)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintResponseFormat(organizations)
	},
}

// organizationsGetCmd represents the organizations GET/<id> command
var organizationsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get organization details",
	Long: `Shows details for a organization, by ID.

Specify the organization ID with the -i/--id flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		organization, response, err := SparkClient.Organizations.GetOrganization(organizationID)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintResponseFormat(organization)
	},
}

func init() {
	RootCmd.AddCommand(organizationsCmd)
	organizationsCmd.AddCommand(organizationsListCmd)
	organizationsCmd.AddCommand(organizationsGetCmd)

	organizationsGetCmd.Flags().StringVarP(&organizationID, "id", "i", "", "The organization ID")
}
