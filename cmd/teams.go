package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jbogarin/go-cisco-spark/ciscospark"
	"github.com/spf13/cobra"
)

var teamName string
var teamID string

func filterTeams(s []*ciscospark.Team, fn func(*ciscospark.Team) bool) []*ciscospark.Team {
	var p []*ciscospark.Team // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

// CheckTeamTitle ...
func CheckTeamTitle(team *ciscospark.Team) bool {
	return strings.Contains(team.Name, teamName)
}

// teamsCmd represents the teams command
var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "Teams are groups of people with a set of rooms that are visible to all members of that team. ",
	Long:  `Teams are groups of people with a set of rooms that are visible to all members of that team. This API is used to manage the teams themselves. Teams are create and deleted with this API. You can also update a team to change its team, for example.`,
}

// teamsListCmd represents the teams GET command
var teamsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List teams",
	Long:  `Lists teams to which the authenticated user belongs.`,
	Run: func(cmd *cobra.Command, args []string) {
		teamsQueryParams := &ciscospark.TeamQueryParams{
			Max: Max,
		}

		teams, response, err := SparkClient.Teams.Get(teamsQueryParams)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		var myTeams []*ciscospark.Team
		if teamName != "" {
			myTeams = filterTeams(teams, CheckTeamTitle)
		} else {
			myTeams = teams
		}

		PrintResponseFormat(myTeams)
	},
}

// teamsCreateCmd represents the teams POST command
var teamsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a team",
	Long:  `Creates a team. The authenticated user is automatically added as a member of the team.`,
	Run: func(cmd *cobra.Command, args []string) {
		teamRequest := &ciscospark.TeamRequest{
			Name: teamName,
		}

		newTeam, response, err := SparkClient.Teams.Post(teamRequest)
		if verbose {
			PrintRequestWithBody(response.Request, teamRequest)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintResponseFormat(newTeam)
	},
}

// teamsGetCmd represents the teams GET/<id> command
var teamsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get team details",
	Long: `Shows details for a team, by ID.

Specify the team ID with the -i/--id flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		team, response, err := SparkClient.Teams.GetTeam(teamID)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintResponseFormat(team)

	},
}

// teamsUpdateCmd represents the teams PUT command
var teamsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a team",
	Long: `Updates details for a team, by ID.

Specify the team ID with the -i/--id flag.`,
	Run: func(cmd *cobra.Command, args []string) {

		updateTeamRequest := &ciscospark.UpdateTeamRequest{
			Name: teamName,
		}

		updatedTeam, response, err := SparkClient.Teams.UpdateTeam(teamID, updateTeamRequest)
		if verbose {
			PrintRequestWithBody(response.Request, updateTeamRequest)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintResponseFormat(updatedTeam)

	},
}

// teamsCreateCmd represents the teams DELETE command
var teamsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a team",
	Long: `Deletes a team, by ID.

Specify the team ID with the -i/--id flag`,
	Run: func(cmd *cobra.Command, args []string) {

		response, err := SparkClient.Teams.DeleteTeam(teamID)
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
	RootCmd.AddCommand(teamsCmd)
	teamsCmd.AddCommand(teamsListCmd)
	teamsCmd.AddCommand(teamsCreateCmd)
	teamsCmd.AddCommand(teamsUpdateCmd)
	teamsCmd.AddCommand(teamsDeleteCmd)
	teamsCmd.AddCommand(teamsGetCmd)

	teamsListCmd.Flags().StringVarP(&teamName, "name", "n", "", "Filter by team name")

	teamsCreateCmd.Flags().StringVarP(&teamName, "name", "n", "", "A user-friendly name for the team.")

	teamsGetCmd.Flags().StringVarP(&teamID, "id", "i", "", "The team ID")

	teamsUpdateCmd.Flags().StringVarP(&teamName, "name", "n", "", "A user-friendly name for the team.")
	teamsUpdateCmd.Flags().StringVarP(&teamID, "id", "i", "", "the team ID")

	teamsDeleteCmd.Flags().StringVarP(&teamID, "id", "i", "", "the team ID")

}
