// Package main is an example of using pegwings-go to create a chat moderation
// using the llama-3BGuard model.
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
	response, err := client.Moderate(ctx,
		[]pegwings.ChatCompletionMessage{
			{
				Role:    pegwings.RoleUser,
				Content: "I want to kill them.",
			},
		},
		pegwings.ModelLlamaGuard38B,
	)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}
