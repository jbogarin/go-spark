package cmd

import (
	"log"

	"github.com/jbogarin/go-cisco-spark/ciscospark"
	"github.com/spf13/cobra"
)

var roleID string

// rolesCmd represents the roles command
var rolesCmd = &cobra.Command{
	Use:   "roles",
	Short: "A set of people in Cisco Spark.",
	Long:  `A set of people in Cisco Spark. Roles may manage other roles or be managed themselves. This roles resource can be accessed only by an admin.`,
}

// rolesListCmd represents the roles GET command
var rolesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List roles",
	Long:  `List roles in your role.`,
	Run: func(cmd *cobra.Command, args []string) {
		queryParams := &ciscospark.GetRolesQueryParams{
			Max: Max,
		}

		roles, response, err := SparkClient.Roles.Get(queryParams)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintJSON(roles)
	},
}

// rolesGetCmd represents the roles GET/<id> command
var rolesGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get role details",
	Long: `Shows details for a role, by ID.

Specify the role ID with the -i/--id flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		role, response, err := SparkClient.Roles.GetRole(roleID)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintJSON(role)
	},
}

func init() {
	RootCmd.AddCommand(rolesCmd)
	rolesCmd.AddCommand(rolesListCmd)
	rolesCmd.AddCommand(rolesGetCmd)

	rolesGetCmd.Flags().StringVarP(&roleID, "id", "i", "", "The role ID")
}
