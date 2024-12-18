package pegwings

import (
	"io"
	"time"

	"github.com/pegwings/pegwings-go/pkg/tools"
)

type (
	// # [Chat](https://doc.pegwings.com/api-reference/chat)

	// ChatCompletionRequest represents a request structure for the chat
	// completion API.
	ChatCompletionRequest struct {
		// Model is the used during the chat completion request.
		Model ChatModel `json:"model"`
		// Messages is the messages of the chat completion request.
		//
		// These act as the prompt for the model.
		Messages []ChatCompletionMessage `json:"messages"`
		// MaxTokens is the max tokens of the chat completion request.
		MaxTokens int `json:"max_tokens,omitempty"`
		// Temperature is the temperature of the chat completion
		// request.
		Temperature float32 `json:"temperature,omitempty"`
		// TopP is the top p of the chat completion request.
		TopP float32 `json:"top_p,omitempty"`
		// N is the n of the chat completion request.
		N int `json:"n,omitempty"`
		// Stream is the stream of the chat completion request.
		Stream bool `json:"stream,omitempty"`
		// Stop is the stop of the chat completion request.
		Stop []string `json:"stop,omitempty"`
		// PresencePenalty is the presence penalty of the chat
		// completion request.
		PresencePenalty float32 `json:"presence_penalty,omitempty"`
		// ResponseFormat is the response format of the chat completion
		// request.
		ResponseFormat *ChatCompletionResponseFormat `json:"response_format,omitempty"`
		// Seed is the seed of the chat completion request.
		Seed *int `json:"seed,omitempty"`
		// FrequencyPenalty is the frequency penalty of the chat
		// completion request.
		FrequencyPenalty float32 `json:"frequency_penalty,omitempty"`
		// LogitBias is must be a token id string (specified by their
		// token ID in the tokenizer), not a word string.
		// incorrect: `"logit_bias":{ "You": 6}`, correct: `"logit_bias":{"1639": 6}`
		// refs: https://platform.openai.com/docs/api-reference/chat/create#chat/create-logit_bias
		LogitBias map[string]int `json:"logit_bias,omitempty"`
		// LogProbs indicates whether to return log probabilities of the
		// output tokens or not. If true, returns the log probabilities
		// of each output token returned in the content of message.
		//
		// This option is currently not available on the
		// gpt-4-vision-preview model.
		LogProbs bool `json:"logprobs,omitempty"`
		// TopLogProbs is an integer between 0 and 5 specifying the
		// number of most likely tokens to return at each token
		// position, each with an associated log probability. Logprobs
		// must be set to true if this parameter is used.
		TopLogProbs int `json:"top_logprobs,omitempty"`
		// User is the user of the chat completion request.
		User string `json:"user,omitempty"`
		// Tools is the tools of the chat completion request.
		Tools []tools.Tool `json:"tools,omitempty"`
		// This can be either a string or an ToolChoice object.
		ToolChoice any `json:"tool_choice,omitempty"`
		// Options for streaming response. Only set this when you set stream: true.
		StreamOptions *StreamOptions `json:"stream_options,omitempty"`
		// Disable the default behavior of parallel tool calls by setting it: false.
		ParallelToolCalls any `json:"parallel_tool_calls,omitempty"`
		// RetryDelay is the delay between retries.
		RetryDelay time.Duration `json:"-"`
	}

	// EmbeddingRequest represents a request structure for embedding API.
	// TODO: update reference to the official API.
	EmbeddingRequest struct {
		// Input text to embed, encoded as a string or array of tokens. To embed multiple
		// inputs in a single request, pass an array of strings or array of token arrays.
		// The input must not exceed the max input tokens for the model (8192 tokens for
		// `text-embedding-ada-002`), cannot be an empty string, and any array must be 2048
		// dimensions or less.
		// [Example Python code](https://cookbook.openai.com/examples/how_to_count_tokens_with_tiktoken)
		// for counting tokens.
		Input string `json:"input,required"`
		// ID of the model to use. You can use the
		// [List models](https://platform.openai.com/docs/api-reference/models/list) API to
		// see all of your available models, or see our
		// [Model overview](https://platform.openai.com/docs/models) for descriptions of
		// them.
		Model EmbeddingModel `json:"model,required"`
		// The format to return the embeddings in. Can be either `float` or
		// [`base64`](https://pypi.org/project/pybase64/).
		EncodingFormat EmbeddingsFormat `json:"encoding_format"`
		// The number of dimensions the resulting output embeddings should have. Only
		// supported in `text-embedding-3` and later models.
		Dimensions int64 `json:"dimensions"`
		// A unique identifier representing your end-user, which can help OpenAI to monitor
		// and detect abuse.
		// [Learn more](https://platform.openai.com/docs/guides/safety-best-practices#end-user-ids).
		User string `json:"user"`
	}

	// AudioRequest represents a request structure for audio API.
	AudioRequest struct {
		// Model is the model to use for the transcription.
		Model AudioModel
		// FilePath is either an existing file in your filesystem or a
		// filename representing the contents of Reader.
		FilePath string
		// Reader is an optional io.Reader when you do not want to use
		// an existing file.
		Reader io.Reader
		// Prompt is the prompt for the transcription.
		Prompt string
		// Temperature is the temperature for the transcription.
		Temperature float32
		// Language is the language for the transcription. Only for
		// transcription.
		Language string
		// Format is the format for the transcription response.
		//
		// Valid formats are: FormatJSON, FormatSRT, FormatVTT,
		// FormatVerboseJSON, FormatText
		Format Format
	}

	// SpeechRequest represents a request structure for speech API.
	//
	// See [Speech](https://doc.pegwings.com/api-reference/speech) for more.
	SpeechRequest struct {
		// Model is the model to use to generate the speech.
		Model SpeechModel `json:"model"`
		// Input is the input to generate audio for.
		//
		// Max length is 4096 characters.
		Input string `json:"input"`
		// Voice is the voice to use for the speech.
		Voice Voice `json:"voice"`
		// ResponseFormat is the response format for the speech.
		ResponseFormat *AudioFormat `json:"response_format,omitempty"`
		// Speed is the speed of the generated speech.
		//
		// Default is 1.0.
		Speed float32 `json:"speed,omitempty"`
	}
)
