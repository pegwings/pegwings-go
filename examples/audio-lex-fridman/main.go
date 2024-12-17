// Package main is an example of using pegwings-go to create a transcription
// using the whisper model.
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
	response, err := client.Transcribe(ctx, pegwings.AudioRequest{
		Model:    pegwings.ModelWhisperLargeV3,
		FilePath: "./The Roman Emperors who went insane Gregory Aldrete and Lex Fridman.mp3",
	})
	if err != nil {
		return err
	}
	fmt.Println(response.Text)
	return nil
}
