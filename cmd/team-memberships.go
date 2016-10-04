package cmd

import (
	"fmt"
	"log"

	"github.com/jbogarin/go-cisco-spark/ciscospark"
	"github.com/spf13/cobra"
)

var teamMembershipsID, teamMembershipsPersonID, teamMembershipsPersonEmail string
var teamMembershipsModerator bool

// teamMembershipsCmd represents the team-memberships command
var teamMembershipsCmd = &cobra.Command{
	Use:   "team-memberships",
	Short: "Team Memberships represent a person's relationship to a team.",
	Long:  `Team Memberships represent a person's relationship to a team. Use this API to list members of any team that you're in or create memberships to invite someone to a team. Team memberships can also be updated to make someome a moderator or deleted to remove them from the team.`,
}

// teamMembershipsListCmd represents the team-memberships GET command
var teamMembershipsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Team Memberships",
	Long: `Lists all team memberships. By default, lists memberships for teams to which the authenticated user belongs.

Use query parameters to filter the response.

Use teamId with -i/--id flag to list memberships for a team.`,
	Run: func(cmd *cobra.Command, args []string) {
		teamMembershipsQueryParams := &ciscospark.TeamMembershipQueryParams{
			Max:    Max,
			TeamID: teamMembershipsID,
		}

		teamMemberships, response, err := SparkClient.TeamMemberships.Get(teamMembershipsQueryParams)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintJSON(teamMemberships)

	},
}

// teamMembershipsCreateCmd represents the team-memberships POST command
var teamMembershipsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Team Membership.",
	Long:  `Add someone to a team by Person ID or email address; optionally making them a moderator.`,

	Run: func(cmd *cobra.Command, args []string) {
		teamMembershipRequest := &ciscospark.TeamMembershipRequest{
			TeamID:      teamMembershipsID,
			PersonEmail: teamMembershipsPersonEmail,
			IsModerator: teamMembershipsModerator,
		}

		newTeamMembership, response, err := SparkClient.TeamMemberships.Post(teamMembershipRequest)
		if verbose {
			PrintRequestWithBody(response.Request, teamMembershipRequest)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintJSON(newTeamMembership)
	},
}

// teamMembershipsGetCmd represents the team-memberships GET/<id> command
var teamMembershipsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details for a membership by ID.",
	Long: `Get details for a membership by ID.

Specify the membership ID using the -i/--id flag.`,
	Run: func(cmd *cobra.Command, args []string) {

		teamMembership, response, err := SparkClient.TeamMemberships.GetTeamMembership(teamMembershipsID)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintJSON(teamMembership)
	},
}

// teamMembershipsUpdateCmd represents the team-memberships PUT command
var teamMembershipsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a team membership",
	Long: `Updates properties for a membership by ID.

Specify the membership ID with the -i/--id flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateTeamMembershipRequest := &ciscospark.UpdateTeamMembershipRequest{
			IsModerator: teamMembershipsModerator,
		}

		updatedTeamMembership, response, err := SparkClient.TeamMemberships.UpdateTeamMembership(teamMembershipsID, updateTeamMembershipRequest)
		if verbose {
			PrintRequestWithBody(response.Request, updateTeamMembershipRequest)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintJSON(updatedTeamMembership)

	},
}

// teamMembershipsDeleteCmd represents the team-memberships DELETE command
var teamMembershipsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a team membership.",
	Long: `Deletes a membership by ID.

Specify the membership ID with the -i/--id flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := SparkClient.TeamMemberships.DeleteTeamMembership(teamMembershipsID)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response.StatusCode)
	},
}

func init() {
	RootCmd.AddCommand(teamMembershipsCmd)
	teamMembershipsCmd.AddCommand(teamMembershipsListCmd)
	teamMembershipsCmd.AddCommand(teamMembershipsCreateCmd)
	teamMembershipsCmd.AddCommand(teamMembershipsGetCmd)
	teamMembershipsCmd.AddCommand(teamMembershipsUpdateCmd)
	teamMembershipsCmd.AddCommand(teamMembershipsDeleteCmd)

	teamMembershipsListCmd.Flags().StringVarP(&teamMembershipsID, "", "i", "", "Limit results to a specific team, by ID.")

	teamMembershipsCreateCmd.Flags().StringVarP(&teamMembershipsID, "i", "i", "", "The team ID.")
	teamMembershipsCreateCmd.Flags().StringVarP(&teamMembershipsPersonID, "person-id", "p", "", "The person ID.")
	teamMembershipsCreateCmd.Flags().StringVarP(&teamMembershipsPersonEmail, "person-email", "e", "", "The email address of the person.")
	teamMembershipsCreateCmd.Flags().BoolVarP(&teamMembershipsModerator, "moderator", "M", false, "Set to true to make the person a room moderator")

	teamMembershipsGetCmd.Flags().StringVarP(&teamMembershipsID, "id", "i", "", "The membership ID.")

	teamMembershipsUpdateCmd.Flags().StringVarP(&teamMembershipsID, "id", "i", "", "The membership ID.")
	teamMembershipsUpdateCmd.Flags().BoolVarP(&teamMembershipsModerator, "moderator", "M", false, "Set to true to make the person a room moderator")

	teamMembershipsDeleteCmd.Flags().StringVarP(&teamMembershipsID, "id", "i", "", "The teamMemberships ID.")

}
