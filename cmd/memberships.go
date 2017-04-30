package cmd

import (
	"fmt"
	"log"

	"github.com/jbogarin/go-cisco-spark/ciscospark"
	"github.com/spf13/cobra"
)

var membershipRoomID, membershipPersonID, membershipPersonEmail, membershipID string
var membershipModerator bool

// membershipsCmd represents the memberships command
var membershipsCmd = &cobra.Command{
	Use:   "memberships",
	Short: "Memberships represent a person's relationship to a room.",
	Long:  `Memberships represent a person's relationship to a room. Use this API to list members of any room that you're in or create memberships to invite someone to a room. Memberships can also be updated to make someome a moderator or deleted to remove them from the room.`,
}

// membershipsListCmd represents the memberships GET command
var membershipsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all room memberships. By default, lists memberships for rooms to which the authenticated user belongs.",
	Long: `Lists all room memberships. By default, lists memberships for rooms to which the authenticated user belongs.

Use query parameters to filter the response.

Use -r/--room to list memberships for a room, by ID.

Use either -p/--person-id or -e/--person-email to filter the results.`,
	Run: func(cmd *cobra.Command, args []string) {
		membershipQueryParams := &ciscospark.MembershipQueryParams{
			Max: Max,
		}

		if membershipRoomID != "" {
			membershipQueryParams.RoomID = membershipRoomID
		}

		if membershipPersonID != "" {
			membershipQueryParams.PersonID = membershipPersonID
		}

		if membershipPersonEmail != "" {
			membershipQueryParams.PersonEmail = membershipPersonEmail
		}

		memberships, response, err := SparkClient.Memberships.Get(membershipQueryParams)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintResponseFormat(memberships)

	},
}

// membershipsCreateCmd represents the memberships POST command
var membershipsCreateCmd = &cobra.Command{
	Use:   "add",
	Short: "Add someone to a room by Person ID or email address; optionally making them a moderator.",
	Long: `Add someone to a room by Person ID or email address; optionally making them a moderator.

Use -r/-room to define the room

Use -p/--person-id to define the person id.

Use -e/--person-email to define the person email.

Use -M/--moderator to define the person as moderator`,
	Run: func(cmd *cobra.Command, args []string) {
		membershipRequest := &ciscospark.MembershipRequest{
			RoomID: membershipRoomID,
		}

		if membershipPersonID != "" {
			membershipRequest.PersonID = membershipPersonID
		}

		if membershipPersonEmail != "" {
			membershipRequest.PersonEmail = membershipPersonEmail
		}

		if membershipModerator {
			membershipRequest.IsModerator = membershipModerator
		}

		membership, response, err := SparkClient.Memberships.Post(membershipRequest)
		if verbose {
			PrintRequestWithBody(response.Request, membershipRequest)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintResponseFormat(membership)

	},
}

// membershipsGetCmd represents the memberships GET/<id> command
var membershipsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details for a membership by ID.",
	Long: `Get details for a membership by ID.

Specify the membership ID with the -i/--id flag.

Use -r/-room to list memberships for a room, by ID.

Use -p/--person-id to filter by  person id.

Use -e/--person-email to filter by person email.`,
	Run: func(cmd *cobra.Command, args []string) {

		membership, response, err := SparkClient.Memberships.GetMembership(membershipID)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintResponseFormat(membership)

	},
}

// membershipsUpdateCmd represents the memberships PUT command
var membershipsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates properties for a membership by ID.",
	Long: `Updates properties for a membership by ID.

Specify the membership ID with the -i/--id flag.`,

	Run: func(cmd *cobra.Command, args []string) {
		updateMembershipRequest := &ciscospark.UpdateMembershipRequest{
			IsModerator: membershipModerator,
		}

		membership, response, err := SparkClient.Memberships.UpdateMembership(membershipID, updateMembershipRequest)

		if verbose {
			PrintRequestWithBody(response.Request, updateMembershipRequest)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintResponseFormat(membership)

	},
}

// membershipsDeleteCmd represents the memberships DELETE command
var membershipsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a membership by ID.",
	Long: `Deletes a membership by ID.

Specify the membership ID with the -i/--id flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := SparkClient.Memberships.DeleteMembership(membershipID)
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
	RootCmd.AddCommand(membershipsCmd)
	membershipsCmd.AddCommand(membershipsListCmd)
	membershipsCmd.AddCommand(membershipsCreateCmd)
	membershipsCmd.AddCommand(membershipsGetCmd)
	membershipsCmd.AddCommand(membershipsUpdateCmd)
	membershipsCmd.AddCommand(membershipsDeleteCmd)

	membershipsListCmd.Flags().StringVarP(&membershipRoomID, "room", "r", "", "Limit results to a specific room, by ID.")
	membershipsListCmd.Flags().StringVarP(&membershipPersonID, "person-id", "p", "", "Limit results to a specific person, by ID.")
	membershipsListCmd.Flags().StringVarP(&membershipPersonEmail, "person-email", "e", "", "Limit results to a specific person, by email address.")

	membershipsCreateCmd.Flags().StringVarP(&membershipRoomID, "room", "r", "", "The room ID.")
	membershipsCreateCmd.Flags().StringVarP(&membershipPersonID, "person-id", "p", "", "The person ID.")
	membershipsCreateCmd.Flags().StringVarP(&membershipPersonEmail, "person-email", "e", "", "The email address of the person.")
	membershipsCreateCmd.Flags().BoolVarP(&membershipModerator, "moderator", "M", false, "Set to true to make the person a room moderator")

	membershipsGetCmd.Flags().StringVarP(&membershipID, "id", "i", "", "The membership ID.")
	membershipsGetCmd.Flags().StringVarP(&membershipRoomID, "room", "r", "", "The room ID.")
	membershipsGetCmd.Flags().StringVarP(&membershipPersonID, "person-id", "p", "", "The person ID.")
	membershipsGetCmd.Flags().StringVarP(&membershipPersonEmail, "person-email", "e", "", "The email address of the person.")

	membershipsUpdateCmd.Flags().StringVarP(&membershipID, "id", "i", "", "The membership ID.")
	membershipsUpdateCmd.Flags().BoolVarP(&membershipModerator, "moderator", "M", false, "Set to true to make the person a room moderator")

	membershipsDeleteCmd.Flags().StringVarP(&membershipID, "id", "i", "", "The membership ID.")

}
