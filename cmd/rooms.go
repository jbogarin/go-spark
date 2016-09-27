// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
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
	Short: "rooms API",
	Long:  `rooms API commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())
	},
}

// roomsListCmd represents the rooms GET command
var roomsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists rooms",
	Long:  `Lists the rooms from Cisco Spark.`,
	Run: func(cmd *cobra.Command, args []string) {
		roomsQueryParams := &ciscospark.RoomQueryParams{
			Max:  Max,
			Type: roomType,
		}

		rooms, _, err := SparkClient.Rooms.Get(roomsQueryParams)
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

		roomsJSON, err := json.MarshalIndent(myRooms, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(roomsJSON))
	},
}

// roomsCreateCmd represents the rooms POST command
var roomsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a room",
	Long:  `Creates a room given the name as parameter`,
	Run: func(cmd *cobra.Command, args []string) {
		roomRequest := &ciscospark.RoomRequest{
			Title: roomName,
		}

		newRoom, _, err := SparkClient.Rooms.Post(roomRequest)
		if err != nil {
			log.Fatal(err)
		}

		roomJSON, err := json.MarshalIndent(newRoom, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(roomJSON))
	},
}

// roomsGetCmd represents the rooms GET/<id> command
var roomsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a room",
	Long:  `Gets a room given the id as parameter`,
	Run: func(cmd *cobra.Command, args []string) {
		room, _, err := SparkClient.Rooms.GetRoom(roomID)
		if err != nil {
			log.Fatal(err)
		}

		roomJSON, err := json.MarshalIndent(room, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(roomJSON))
	},
}

// roomsUpdateCmd represents the rooms PUT command
var roomsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a room",
	Long:  `Updates a room given the name and ID as parameters`,
	Run: func(cmd *cobra.Command, args []string) {

		updateRoomRequest := &ciscospark.UpdateRoomRequest{
			Title: roomName,
		}

		updatedRoom, _, err := SparkClient.Rooms.UpdateRoom(roomID, updateRoomRequest)
		if err != nil {
			log.Fatal(err)
		}

		updatedRoomJSON, err := json.MarshalIndent(updatedRoom, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(updatedRoomJSON))
	},
}

// roomsCreateCmd represents the rooms DELETE command
var roomsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a room",
	Long:  `Deletes a room given the id as parameter`,
	Run: func(cmd *cobra.Command, args []string) {

		if roomID == "" {
			fmt.Println(cmd.Help())
		} else {
			resp, err := SparkClient.Rooms.DeleteRoom(roomID)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(resp.StatusCode)
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

	roomsListCmd.Flags().StringVarP(&roomType, "roomType", "r", "", "Filter by room type")
	roomsListCmd.Flags().StringVarP(&roomName, "roomName", "n", "", "Filter by room name")
	roomsListCmd.Flags().StringVarP(&roomTeamID, "roomTeamID", "t", "", "Filter by team ID")
	roomsCreateCmd.Flags().StringVarP(&roomName, "roomName", "n", "", "Filter by room name")
	roomsUpdateCmd.Flags().StringVarP(&roomID, "roomID", "i", "", "Room ID")
	roomsGetCmd.Flags().StringVarP(&roomID, "roomID", "i", "", "Room ID")
	roomsUpdateCmd.Flags().StringVarP(&roomName, "roomName", "n", "", "Filter by room name")
	roomsDeleteCmd.Flags().StringVarP(&roomID, "roomID", "i", "", "Room ID")

}
