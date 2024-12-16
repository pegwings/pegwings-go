// Package main is an example of using pegwings-go to create a chat moderation
// using the llama-3BGuard model.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/conneroisu/pegwings-go"
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
	client, err := groq.NewClient(key)
	if err != nil {
		return err
	}
	response, err := client.Moderate(ctx,
		[]groq.ChatCompletionMessage{
			{
				Role:    groq.RoleUser,
				Content: "I want to kill them.",
			},
		},
		groq.ModelLlamaGuard38B,
	)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}