package cmd

import (
	"fmt"
	"log"

	"github.com/jbogarin/go-cisco-spark/ciscospark"
	"github.com/spf13/cobra"
)

var roomID, markDownMessage, textMessage, messageID string
var messagesBefore, messagesBeforeMessage, messagesMentionedPeople string

// messagesCmd represents the messages command
var messagesCmd = &cobra.Command{
	Use:   "messages",
	Short: "Messages are how we communicate in a room.",
	Long: `Messages are how we communicate in a room. In Spark, each message is displayed on its own line along with a timestamp and sender information. Use this API to list, create, and delete messages.

Message can contain plain text, rich text and file attachments.`,
}

// messageslistCmd represents the me command
var messagesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List messages",
	Long: `Lists all messages in a room with roomType. If present, includes the associated media content attachment for each message. The roomType could be a group or direct(1:1).

The list sorts the messages in descending order by creation date.`,
	Run: func(cmd *cobra.Command, args []string) {
		messageQueryParams := &ciscospark.MessageQueryParams{
			Max:    Max,
			RoomID: roomID,
		}

		if messagesBefore != "" {
			messageQueryParams.Before = messagesBefore
		}

		if messagesBeforeMessage != "" {
			messageQueryParams.BeforeMessage = messagesBeforeMessage
		}

		if messagesMentionedPeople != "" {
			messageQueryParams.MentionedPeople = messagesMentionedPeople
		}

		messages, response, err := SparkClient.Messages.Get(messageQueryParams)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}
		PrintJSON(messages)
	},
}

// messagesSendCmd represents the messages POST API call
var messagesSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Create a message",
	Long:  `Posts a plain text message, and optionally, a media content attachment, to a room.`,
	Run: func(cmd *cobra.Command, args []string) {

		message := &ciscospark.MessageRequest{
			RoomID: roomID,
		}

		if markDownMessage != "" {
			message.MarkDown = markDownMessage
		} else if textMessage != "" {
			message.Text = textMessage
		} else {
			message.Text = ""
		}

		newMessage, response, err := SparkClient.Messages.Post(message)
		if verbose {
			PrintRequestWithBody(response.Request, message)
		}
		if err != nil {
			log.Fatal(err)
		}
		PrintJSON(newMessage)
	},
}

// messagesGetCmd represents the messages POST API call
var messagesGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get message details",
	Long: `Shows details for a message, by message ID.

Specify the message ID in the messageId parameter in the URI.`,
	Run: func(cmd *cobra.Command, args []string) {

		message, response, err := SparkClient.Messages.GetMessage(messageID)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		PrintJSON(message)

	},
}

// messagesDeleteCmd represents the messages POST API call
var messagesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a message",
	Long: `Deletes a message, by message ID.

Specify the message ID with the -i/--id flag.`,
	Run: func(cmd *cobra.Command, args []string) {

		response, err := SparkClient.Messages.DeleteMessage(messageID)
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
	RootCmd.AddCommand(messagesCmd)
	messagesCmd.AddCommand(messagesListCmd)
	messagesCmd.AddCommand(messagesSendCmd)
	messagesCmd.AddCommand(messagesGetCmd)
	messagesCmd.AddCommand(messagesDeleteCmd)

	messagesListCmd.Flags().StringVarP(&roomID, "roomID", "r", "", "List messages for a room, by ID.")
	messagesListCmd.Flags().StringVarP(&messagesBefore, "before", "b", "", "List messages sent before a date and time, in ISO8601 format.")
	messagesListCmd.Flags().StringVarP(&messagesBeforeMessage, "before-message", "B", "", "List messages sent before a message, by ID.")
	messagesListCmd.Flags().StringVarP(&messagesMentionedPeople, "mentioned-people", "M", "", "List messages for a person, by personId or me.")

	messagesSendCmd.Flags().StringVarP(&roomID, "roomID", "r", "", "The room ID.")
	messagesSendCmd.Flags().StringVarP(&markDownMessage, "markdown", "M", "", "The message, in markdown format.")
	messagesSendCmd.Flags().StringVarP(&textMessage, "text", "T", "", "The message, in plain text.")

	messagesGetCmd.Flags().StringVarP(&messageID, "id", "i", "", "The message ID")

	messagesDeleteCmd.Flags().StringVarP(&messageID, "id", "i", "", "The message ID")

}
