package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/krognol/go-wolfram"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
)

var (
	wolframClient *wolfram.Client
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		log.Printf("Command Event - Timestamp: %s, Command: %s, Parameters: %s, Event: %s\n", event.Timestamp, event.Command, event.Parameters, event.Event)
	}
}

func main() {
	// Load environment variables from a .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Create a Slack bot client and Wit.ai client
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
	wolframClient = &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}

	// Start a goroutine to print command events
	go printCommandEvents(bot.CommandEvents())

	// Define a Slack bot command
	bot.Command("ask <query>", &slacker.CommandDefinition{
		Description: "Ask the bot a question",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("query")
			log.Printf("Received query: %s\n", query)

			// Parse the user's query with Wit.ai
			msg, err := client.Parse(&witai.MessageRequest{
				Query: query,
			})
			if err != nil {
				log.Println("Wit.ai error:", err)
				response.Reply("I'm sorry, but I couldn't understand your question.")
				return
			}

			// Extract the Wolfram Alpha search query from Wit.ai response
			data, err := json.MarshalIndent(msg, "", "  ")
			if err != nil {
				log.Println("JSON marshaling error:", err)
				response.Reply("I encountered an error while processing your question.")
				return
			}
			rough := string(data)
			val := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")
			answer := val.String()

			if answer == "" {
				response.Reply("I couldn't find a specific question in your query.")
				return
			}

			// Query Wolfram Alpha and get the answer
			res, err := wolframClient.GetSpokentAnswerQuery(answer, wolfram.Metric, 1000)
			if err != nil {
				log.Println("Wolfram Alpha error:", err)
				response.Reply("I encountered an error while looking up the answer.")
				return
			}

			// Reply to the user with the Wolfram Alpha answer
			response.Reply(res)
		},
	})

	// Create a context for the bot and handle errors
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = bot.Listen(ctx)
	if err != nil {
		log.Fatal("Bot Listen error:", err)
	}
}
