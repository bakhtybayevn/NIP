# Slack Bot with Wit.ai and Wolfram Alpha Integration
This is a Go-based Slack bot that leverages Wit.ai and Wolfram Alpha to provide answers to questions asked by users in a Slack workspace. The bot listens for a specific command, processes the user's query using Wit.ai, and then queries Wolfram Alpha to generate and provide answers.
# Features
1. Question Answering: Users can ask questions to the Slack bot using the defined command, and the bot will attempt to provide answers based on the query.
2. Wit.ai Integration: The bot uses Wit.ai for natural language understanding to extract meaningful information from user queries.
3. Wolfram Alpha Integration: Queries made to Wolfram Alpha are used to find answers to the user's questions.
4. Error Handling: The bot includes error handling to provide appropriate responses in case of issues with Wit.ai or Wolfram Alpha queries.
# Prerequisites
Before you start using this application, ensure you have the following:
1. Go installed on your system.
2. Slack API tokens: You'll need a Slack bot token and a Slack app token.
3. Wit.ai API token: You'll need an API token for Wit.ai.
4. Wolfram Alpha app ID: You'll need an app ID from Wolfram Alpha.
# Installation
1. Clone the repository to your local machine:
git clone https://github.com/bakhtybayevn/NIP.git
cd NIP
2. Create a .env file in the project directory and add the following environment variables:
SLACK_BOT_TOKEN=<your_slack_bot_token>
SLACK_APP_TOKEN=<your_slack_app_token>
WIT_AI_TOKEN=<your_witai_token>
WOLFRAM_APP_ID=<your_wolfram_app_id>
3. Install the required Go packages:
go get github.com/joho/godotenv
go get github.com/krognol/go-wolfram
go get github.com/shomali11/slacker
go get github.com/tidwall/gjson
go get github.com/wit-ai/wit-go/v2
4. Build and run the bot:
go build
./slack-bot-witai-wolfram
The bot should now be up and running in your Slack workspace.
# Usage
In your Slack workspace, you can interact with the bot using the following command:
/ask <your_question>
Replace <your_question> with the question you want to ask the bot. The bot will process the question and provide an answer based on the information obtained from Wolfram Alpha.
# Contributing
Feel free to contribute to this project by opening issues, suggesting improvements, or submitting pull requests.
