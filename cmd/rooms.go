package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jbogarin/go-cisco-spark/ciscospark"
	"github.com/spf13/cobra"
)

var roomType string
var roomName string
var roomTeamID string

func filterRooms(s []*ciscospark.Room, fn func(*ciscospark.Room) bool) []*ciscospark.Room {
	var p []*ciscospark.Room // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func checkTitle(room *ciscospark.Room) bool {
	return strings.Contains(room.Title, roomName)
}

func checkTeamID(room *ciscospark.Room) bool {
	return room.TeamID == roomTeamID
}

// roomsCmd represents the rooms command
var roomsCmd = &cobra.Command{
	Use:   "rooms",
	Short: "Rooms are virtual meeting places where people post messages and collaborate to get work done.",
	Long:  `Rooms are virtual meeting places where people post messages and collaborate to get work done. This API is used to manage the rooms themselves. Rooms are create and deleted with this API. You can also update a room to change its title, for example..`,
}

// roomsListCmd represents the rooms GET command
var roomsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List rooms",
	Long: `List rooms.

By default, lists rooms to which the authenticated user belongs.

Use -r/--room-type to define the room type`,
	Run: func(cmd *cobra.Command, args []string) {
		roomsQueryParams := &ciscospark.RoomQueryParams{
			Max:  Max,
			Type: roomType,
		}

		rooms, response, err := SparkClient.Rooms.Get(roomsQueryParams)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		var myRooms []*ciscospark.Room
		if roomName != "" {
			myRooms = filterRooms(rooms, checkTitle)
		} else if roomTeamID != "" {
			myRooms = filterRooms(rooms, checkTeamID)
		} else {
			myRooms = rooms
		}

		PrintJSON(myRooms)
	},
}

// roomsCreateCmd represents the rooms POST command
var roomsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a room",
	Long:  `Creates a room. The authenticated user is automatically added as a member of the room.`,
	Run: func(cmd *cobra.Command, args []string) {
		roomRequest := &ciscospark.RoomRequest{
			Title: roomName,
		}

		if roomTeamID != "" {
			roomRequest.TeamID = roomTeamID
		}

		newRoom, response, err := SparkClient.Rooms.Post(roomRequest)
		if verbose {
			PrintRequestWithBody(response.Request, roomRequest)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintJSON(newRoom)

	},
}

// roomsGetCmd represents the rooms GET/<id> command
var roomsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get room details",
	Long: `Shows details for a room, by ID.

Specify the room ID with the -i/--id flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		room, response, err := SparkClient.Rooms.GetRoom(roomID)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintJSON(room)

	},
}

// roomsUpdateCmd represents the rooms PUT command
var roomsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a room",
	Long: `Updates details for a room, by ID.

Specify the room ID with the -i/--id flag.`,
	Run: func(cmd *cobra.Command, args []string) {

		updateRoomRequest := &ciscospark.UpdateRoomRequest{
			Title: roomName,
		}

		updatedRoom, response, err := SparkClient.Rooms.UpdateRoom(roomID, updateRoomRequest)
		if verbose {
			PrintRequestWithBody(response.Request, updateRoomRequest)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintJSON(updatedRoom)

	},
}

// roomsCreateCmd represents the rooms DELETE command
var roomsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a room",
	Long: `Deletes a room, by ID.

Specify the room ID with the -i/--id flag`,
	Run: func(cmd *cobra.Command, args []string) {

		if roomID == "" {
			fmt.Println(cmd.Help())
		} else {
			response, err := SparkClient.Rooms.DeleteRoom(roomID)
			if verbose {
				PrintRequestWithoutBody(response.Request)
			}
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(response.StatusCode)
		}
	},
}

func init() {
	RootCmd.AddCommand(roomsCmd)
	roomsCmd.AddCommand(roomsListCmd)
	roomsCmd.AddCommand(roomsCreateCmd)
	roomsCmd.AddCommand(roomsUpdateCmd)
	roomsCmd.AddCommand(roomsDeleteCmd)
	roomsCmd.AddCommand(roomsGetCmd)

	roomsListCmd.Flags().StringVarP(&roomType, "type", "r", "", "Available values: direct and group. direct returns all 1-to-1 rooms. group returns all group rooms. If not specified or values not matched, will return all room types.")
	roomsListCmd.Flags().StringVarP(&roomName, "name", "n", "", "Filter by room name")
	roomsListCmd.Flags().StringVarP(&roomTeamID, "team", "T", "", "Limit the rooms to those associatedwith a team, by ID.")

	roomsCreateCmd.Flags().StringVarP(&roomName, "name", "n", "", "A user-friendly name for the room.")
	roomsCreateCmd.Flags().StringVarP(&roomTeamID, "team", "T", "", "The ID for the team with which this room is associated.")

	roomsUpdateCmd.Flags().StringVarP(&roomID, "id", "i", "", "The Room ID")

	roomsGetCmd.Flags().StringVarP(&roomID, "id", "i", "", "The Room ID")

	roomsDeleteCmd.Flags().StringVarP(&roomID, "id", "i", "", "The Room ID")

}
