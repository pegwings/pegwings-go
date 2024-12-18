package pegwings

import "net/http"

type (
	// ChatCompletionResponse represents a response structure for chat
	// completion API.
	ChatCompletionResponse struct {
		// ID is the id of the response.
		ID string `json:"id"`
		// Object is the object of the response.
		Object string `json:"object"`
		// Created is the created time of the response.
		Created int64 `json:"created"`
		// Model is the model of the response.
		Model ChatModel `json:"model"`
		// Choices is the choices of the response.
		Choices []ChatCompletionChoice `json:"choices"`
		// Usage is the usage of the response.
		Usage Usage `json:"usage"`
		// SystemFingerprint is the system fingerprint of the response.
		SystemFingerprint string `json:"system_fingerprint"`
		// Header is the header of the response.
		header http.Header `json:"-"`
	}
	// ChatCompletionStreamResponse represents a response structure for chat
	// completion API.
	ChatCompletionStreamResponse struct {
		// ID is the identifier for the chat completion stream response.
		ID string `json:"id"`
		// Object is the object type of the chat completion stream
		// response.
		Object string `json:"object"`
		// Created is the creation time of the chat completion stream
		// response.
		Created int64 `json:"created"`
		// Model is the model used for the chat completion stream
		// response.
		Model ChatModel `json:"model"`
		// Choices is the choices for the chat completion stream
		// response.
		Choices []ChatCompletionStreamChoice `json:"choices"`
		// SystemFingerprint is the system fingerprint for the chat
		// completion stream response.
		SystemFingerprint string `json:"system_fingerprint"`
		// PromptAnnotations is the prompt annotations for the chat
		// completion stream response.
		PromptAnnotations []PromptAnnotation `json:"prompt_annotations,omitempty"`
		// PromptFilterResults is the prompt filter results for the chat
		// completion stream response.
		PromptFilterResults []struct {
			Index int `json:"index"`
		} `json:"prompt_filter_results,omitempty"`
		// Usage is an optional field that will only be present when you
		// set stream_options: {"include_usage": true} in your request.
		//
		// When present, it contains a null value except for the last
		// chunk which contains the token usage statistics for the
		// entire request.
		Usage *Usage `json:"usage,omitempty"`
	}
	// AudioResponse represents a response structure for audio API.
	AudioResponse struct {
		// Task is the task of the response.
		Task string `json:"task"`
		// Language is the language of the response.
		Language string `json:"language"`
		// Duration is the duration of the response.
		Duration float64 `json:"duration"`
		// Segments is the segments of the response.
		Segments Segments `json:"segments"`
		// Words is the words of the response.
		Words Words `json:"words"`
		// Text is the text of the response.
		Text string `json:"text"`

		header http.Header `json:"-"`
	}
	// audioTextResponse is the response structure for the audio API when the
	// response format is text.
	audioTextResponse struct {
		// Text is the text of the response.
		Text   string      `json:"text"`
		header http.Header `json:"-"`
	}
)
