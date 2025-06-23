package config

import (
	"github.com/sashabaranov/go-openai"
)

var OpenAIClient *openai.Client

func InitializeOpenAIClient() {
	OpenAIClient = openai.NewClient(OPEN_AI_API_KEY)
}
