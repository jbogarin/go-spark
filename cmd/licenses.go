package cmd

import (
	"log"

	"github.com/jbogarin/go-cisco-spark/ciscospark"
	"github.com/spf13/cobra"
)

var licenseID, licenseOrgID string

// licensesCmd represents the licenses command
var licensesCmd = &cobra.Command{
	Use:   "licenses",
	Short: "A set of people in Cisco Spark.",
	Long:  `A set of people in Cisco Spark. Licenses may manage other licenses or be managed themselves. This licenses resource can be accessed only by an admin.`,
}

// licensesListCmd represents the licenses GET command
var licensesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List licenses",
	Long:  `List licenses in your license.`,
	Run: func(cmd *cobra.Command, args []string) {
		queryParams := &ciscospark.GetLicensesQueryParams{
			Max:   Max,
			OrgID: licenseOrgID,
		}

		licenses, response, err := SparkClient.Licenses.Get(queryParams)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintJSON(licenses)
	},
}

// licensesGetCmd represents the licenses GET/<id> command
var licensesGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get license details",
	Long: `Shows details for a license, by ID.

Specify the license ID with the -i/--id flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		license, response, err := SparkClient.Licenses.GetLicense(licenseID)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintJSON(license)
	},
}

func init() {
	RootCmd.AddCommand(licensesCmd)
	licensesCmd.AddCommand(licensesListCmd)
	licensesCmd.AddCommand(licensesGetCmd)

	licensesGetCmd.Flags().StringVarP(&licenseID, "id", "i", "", "The license ID")
	licensesListCmd.Flags().StringVarP(&licenseOrgID, "orgId", "o", "", "Specify the organization")

}
