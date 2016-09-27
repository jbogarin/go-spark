package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jbogarin/go-cisco-spark/ciscospark"
	"github.com/spf13/cobra"
)

var roomID string
var htmlMessage string
var markDownMessage string
var textMessage string

// messagesCmd represents the messages command
var messagesCmd = &cobra.Command{
	Use:   "messages",
	Short: "messages API",
	Long:  `messages APIs commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())
	},
}

// messageslistCmd represents the me command
var messagesListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists messages",
	Long:  `Lists the messages from a Cisco Spark room.`,
	Run: func(cmd *cobra.Command, args []string) {
		messageQueryParams := &ciscospark.MessageQueryParams{
			Max:    Max,
			RoomID: roomID,
		}

		messages, _, err := SparkClient.Messages.Get(messageQueryParams)
		if err != nil {
			log.Fatal(err)
		}
		messagesJSON, err := json.MarshalIndent(messages, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(messagesJSON))
	},
}

// messagesSendCmd represents the messages POST API call
var messagesSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Sends a message",
	Long:  `Sends a html, markdown or text messages to the room specified by -r.`,
	Run: func(cmd *cobra.Command, args []string) {

		message := &ciscospark.MessageRequest{
			RoomID: roomID,
		}

		if htmlMessage != "" {
			message.HTML = htmlMessage
		} else if markDownMessage != "" {
			message.MarkDown = markDownMessage
		} else if textMessage != "" {
			message.Text = textMessage
		} else {
			message.Text = ""
		}

		newMessage, _, err := SparkClient.Messages.Post(message)
		if err != nil {
			log.Fatal(err)
		}
		messagesJSON, err := json.MarshalIndent(newMessage, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(messagesJSON))
	},
}

func init() {
	RootCmd.AddCommand(messagesCmd)
	messagesCmd.AddCommand(messagesListCmd)
	messagesCmd.AddCommand(messagesSendCmd)

	messagesListCmd.Flags().StringVarP(&roomID, "roomID", "r", "", "roomID from where to get the messages")
	messagesSendCmd.Flags().StringVarP(&roomID, "roomID", "r", "", "roomID from where to get the messages")
	messagesSendCmd.Flags().StringVarP(&htmlMessage, "html", "H", "", "HTML message to send")
	messagesSendCmd.Flags().StringVarP(&markDownMessage, "markdown", "M", "", "Markdown message to send")
	messagesSendCmd.Flags().StringVarP(&textMessage, "text", "T", "", "Text message to send")
}
