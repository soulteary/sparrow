package OpenaiAPI

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	openai "github.com/sashabaranov/go-openai"
	"github.com/soulteary/sparrow/internal/define"
)

func GetClient() *openai.Client {
	config := openai.DefaultConfig(define.OPENAI_API_KEY)
	if define.ENABLE_OPENAI_API_PROXY {
		proxyUrl, err := url.Parse(define.OPENAI_API_PROXY_ADDR)
		if err != nil {
			panic(err)
		}
		transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		config.HTTPClient = &http.Client{Transport: transport}
	}
	return openai.NewClientWithConfig(config)
}

func Get(prompt string) string {
	c := GetClient()
	resp, err := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}},
		},
	)
	if err != nil {
		return fmt.Sprintf("OpenAI API, Chat Completion error: %v\n", err)
	}
	return resp.Choices[0].Message.Content
}
