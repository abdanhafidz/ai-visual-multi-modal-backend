package config

import (
	"github.com/sashabaranov/go-openai"
)

var OpenAIClient *openai.Client

func init() {
	OpenAIClient = openai.NewClient(OPEN_AI_API_KEY)
}
