package cmd

import (
	"log"

	"github.com/jbogarin/go-cisco-spark/ciscospark"
	"github.com/spf13/cobra"
)

// webhooksCmd represents the webhooks command
var webhooksCmd = &cobra.Command{
	Use:   "webhooks",
	Short: "Events trigger in near real-time allowing your app and backend IT systems to stay in sync with new content and room activity.",
	Long: `Events trigger in near real-time allowing your app and backend IT systems to stay in sync with new content and room activity.
	
	Webhooks created via this API will not appear in a room's 'Integrations' list within the Spark client.`,
}

// webhooksListCmd represents the webhooks GET command
var webhooksListCmd = &cobra.Command{
	Use:   "list",
	Short: "List webhooks",
	Long: `List webhooks.

By default, lists webhooks to which the authenticated user belongs.

Use -r/--webhook-type to define the webhook type`,
	Run: func(cmd *cobra.Command, args []string) {
		webhooksQueryParams := &ciscospark.WebhookQueryParams{
			Max: 10,
		}

		webhooks, response, err := SparkClient.Webhooks.Get(webhooksQueryParams)
		if verbose {
			PrintRequestWithoutBody(response.Request)
		}
		if err != nil {
			log.Fatal(err)
		}

		// var myWebhooks []*ciscospark.Webhook
		// if webhookName != "" {
		// 	myWebhooks = filterWebhooks(webhooks, checkTitle)
		// } else if webhookTeamID != "" {
		// 	myWebhooks = filterWebhooks(webhooks, checkTeamID)
		// } else {
		// 	myWebhooks = webhooks
		// }

		PrintResponseFormat(webhooks)
	},
}

// // webhooksCreateCmd represents the webhooks POST command
// var webhooksCreateCmd = &cobra.Command{
// 	Use:   "create",
// 	Short: "Create a webhook",
// 	Long:  `Creates a webhook. The authenticated user is automatically added as a member of the webhook.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		webhookRequest := &ciscospark.WebhookRequest{
// 			Title: webhookName,
// 		}

// 		if webhookTeamID != "" {
// 			webhookRequest.TeamID = webhookTeamID
// 		}

// 		newWebhook, response, err := SparkClient.Webhooks.Post(webhookRequest)
// 		if verbose {
// 			PrintRequestWithBody(response.Request, webhookRequest)
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		PrintResponseFormat(newWebhook)

// 	},
// }

// // webhooksGetCmd represents the webhooks GET/<id> command
// var webhooksGetCmd = &cobra.Command{
// 	Use:   "get",
// 	Short: "Get webhook details",
// 	Long: `Shows details for a webhook, by ID.

// Specify the webhook ID with the -i/--id flag.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		webhook, response, err := SparkClient.Webhooks.GetWebhook(webhookID)
// 		if verbose {
// 			PrintRequestWithoutBody(response.Request)
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		PrintResponseFormat(webhook)

// 	},
// }

// // webhooksUpdateCmd represents the webhooks PUT command
// var webhooksUpdateCmd = &cobra.Command{
// 	Use:   "update",
// 	Short: "Update a webhook",
// 	Long: `Updates details for a webhook, by ID.

// Specify the webhook ID with the -i/--id flag.`,
// 	Run: func(cmd *cobra.Command, args []string) {

// 		updateWebhookRequest := &ciscospark.UpdateWebhookRequest{
// 			Title: webhookName,
// 		}

// 		updatedWebhook, response, err := SparkClient.Webhooks.UpdateWebhook(webhookID, updateWebhookRequest)
// 		if verbose {
// 			PrintRequestWithBody(response.Request, updateWebhookRequest)
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		PrintResponseFormat(updatedWebhook)

// 	},
// }

// // webhooksCreateCmd represents the webhooks DELETE command
// var webhooksDeleteCmd = &cobra.Command{
// 	Use:   "delete",
// 	Short: "Delete a webhook",
// 	Long: `Deletes a webhook, by ID.

// Specify the webhook ID with the -i/--id flag`,
// 	Run: func(cmd *cobra.Command, args []string) {

// 		if webhookID == "" {
// 			fmt.Println(cmd.Help())
// 		} else {
// 			response, err := SparkClient.Webhooks.DeleteWebhook(webhookID)
// 			if verbose {
// 				PrintRequestWithoutBody(response.Request)
// 			}
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			fmt.Println(response.StatusCode)
// 		}
// 	},
// }

func init() {
	RootCmd.AddCommand(webhooksCmd)
	webhooksCmd.AddCommand(webhooksListCmd)
	// webhooksCmd.AddCommand(webhooksCreateCmd)
	// webhooksCmd.AddCommand(webhooksUpdateCmd)
	// webhooksCmd.AddCommand(webhooksDeleteCmd)
	// webhooksCmd.AddCommand(webhooksGetCmd)

	// webhooksCreateCmd.Flags().StringVarP(&webhookName, "name", "n", "", "A user-friendly name for the webhook.")
	// webhooksCreateCmd.Flags().StringVarP(&webhookTeamID, "team", "T", "", "The ID for the team with which this webhook is associated.")

	// webhooksUpdateCmd.Flags().StringVarP(&webhookID, "id", "i", "", "The Webhook ID")

	// webhooksGetCmd.Flags().StringVarP(&webhookID, "id", "i", "", "The Webhook ID")

	// webhooksDeleteCmd.Flags().StringVarP(&webhookID, "id", "i", "", "The Webhook ID")

}
