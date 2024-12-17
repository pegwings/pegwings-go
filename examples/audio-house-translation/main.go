// Package main is an example of using the pegwings-go library to create a
// transcription/translation using the whisper-large-v3 model.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pegwings/pegwings-go"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(
	ctx context.Context,
) error {
	client, err := pegwings.NewClient(os.Getenv("GROQ_KEY"))
	if err != nil {
		return err
	}
	response, err := client.Translate(ctx, pegwings.AudioRequest{
		Model:    pegwings.ModelWhisperLargeV3,
		FilePath: "./house-speaks-mandarin.mp3",
		Prompt:   "english and mandarin",
	})
	if err != nil {
		return err
	}
	fmt.Println(response.Text)
	return nil
}
