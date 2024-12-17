// Package main demonstrates an example application of pegwings-go.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pegwings/pegwings-go"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(
	ctx context.Context,
) error {
	key := os.Getenv("GROQ_KEY")
	client, err := pegwings.NewClient(key)
	if err != nil {
		return err
	}
	response, err := client.ChatCompletion(
		ctx,
		pegwings.ChatCompletionRequest{
			Model: pegwings.ModelLlama3211BVisionPreview,
			Messages: []pegwings.ChatCompletionMessage{
				{
					Role: pegwings.RoleUser,
					MultiContent: []pegwings.ChatMessagePart{
						{
							Type: pegwings.ChatMessagePartTypeText,
							Text: "What is the contents of the image?",
						},
						{
							Type: pegwings.ChatMessagePartTypeImageURL,
							ImageURL: &pegwings.ChatMessageImageURL{
								URL:    "https://cdnimg.webstaurantstore.com/images/products/large/87539/251494.jpg",
								Detail: "auto",
							},
						}},
				},
			},
			MaxTokens: 2000,
		},
	)
	if err != nil {
		return err
	}
	fmt.Println(response.Choices[0].Message.Content)
	return nil
}
