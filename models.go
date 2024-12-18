package pegwings

type (
	// Model is a ai model accessible through the groq api.
	Model string

	// ChatModel is the type for chat models present on the groq api.
	ChatModel Model

	// ModerationModel is the type for moderation models present on the groq api.
	ModerationModel Model

	// AudioModel is the type for audio models present on the groq api.
	AudioModel Model

	// EmbeddingModel is the type for embedding models present on the groq api.
	EmbeddingModel Model

	// SpeechModel is the type for speech models present on the groq api.
	SpeechModel Model
)

var (
	// ModelGemma29BIt is an AI text chat model.
	//
	// It is created/provided by Google.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- ChatCompletion
	// 	- ChatCompletionStream
	// 	- ChatCompletionJSON
	ModelGemma29BIt ChatModel = "gemma2-9b-it"
	// ModelGemma7BIt is an AI text chat model.
	//
	// It is created/provided by Google.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- ChatCompletion
	// 	- ChatCompletionStream
	// 	- ChatCompletionJSON
	ModelGemma7BIt ChatModel = "gemma-7b-it"
	// ModelLlama3170BVersatile is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 32768 context window.
	//
	// It can be used with the following client methods:
	//	- ChatCompletion
	// 	- ChatCompletionStream
	// 	- ChatCompletionJSON
	ModelLlama3170BVersatile ChatModel = "llama-3.1-70b-versatile"
	// ModelLlama318BInstant is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 131072 context window.
	//
	// It can be used with the following client methods:
	//	- ChatCompletion
	// 	- ChatCompletionStream
	// 	- ChatCompletionJSON
	ModelLlama318BInstant ChatModel = "llama-3.1-8b-instant"
	// ModelLlama3211BVisionPreview is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- ChatCompletion
	// 	- ChatCompletionStream
	// 	- ChatCompletionJSON
	ModelLlama3211BVisionPreview ChatModel = "llama-3.2-11b-vision-preview"
	// ModelLlama321BPreview is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- ChatCompletion
	// 	- ChatCompletionStream
	// 	- ChatCompletionJSON
	ModelLlama321BPreview ChatModel = "llama-3.2-1b-preview"
	// ModelLlama323BPreview is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- ChatCompletion
	// 	- ChatCompletionStream
	// 	- ChatCompletionJSON
	ModelLlama323BPreview ChatModel = "llama-3.2-3b-preview"
	// ModelLlama3290BVisionPreview is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- ChatCompletion
	// 	- ChatCompletionStream
	// 	- ChatCompletionJSON
	ModelLlama3290BVisionPreview ChatModel = "llama-3.2-90b-vision-preview"
	// ModelLlama3370BSpecdec is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- ChatCompletion
	// 	- ChatCompletionStream
	// 	- ChatCompletionJSON
	ModelLlama3370BSpecdec ChatModel = "llama-3.3-70b-specdec"
	// ModelLlama3370BVersatile is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 32768 context window.
	//
	// It can be used with the following client methods:
	//	- ChatCompletion
	// 	- ChatCompletionStream
	// 	- ChatCompletionJSON
	ModelLlama3370BVersatile ChatModel = "llama-3.3-70b-versatile"
	// ModelLlama370B8192 is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- ChatCompletion
	// 	- ChatCompletionStream
	// 	- ChatCompletionJSON
	ModelLlama370B8192 ChatModel = "llama3-70b-8192"
	// ModelLlama38B8192 is an AI text chat model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- ChatCompletion
	// 	- ChatCompletionStream
	// 	- ChatCompletionJSON
	ModelLlama38B8192 ChatModel = "llama3-8b-8192"
	// ModelLlama3Groq70B8192ToolUsePreview is an AI text chat model.
	//
	// It is created/provided by pegwings.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- ChatCompletion
	// 	- ChatCompletionStream
	// 	- ChatCompletionJSON
	ModelLlama3Groq70B8192ToolUsePreview ChatModel = "llama3-groq-70b-8192-tool-use-preview"
	// ModelLlama3Groq8B8192ToolUsePreview is an AI text chat model.
	//
	// It is created/provided by pegwings.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- ChatCompletion
	// 	- ChatCompletionStream
	// 	- ChatCompletionJSON
	ModelLlama3Groq8B8192ToolUsePreview ChatModel = "llama3-groq-8b-8192-tool-use-preview"
	// ModelMixtral8X7B32768 is an AI text chat model.
	//
	// It is created/provided by Mistral AI.
	//
	// It has 32768 context window.
	//
	// It can be used with the following client methods:
	//	- ChatCompletion
	// 	- ChatCompletionStream
	// 	- ChatCompletionJSON
	ModelMixtral8X7B32768 ChatModel = "mixtral-8x7b-32768"
	// ModelDistilWhisperLargeV3En is an AI audio transcription model.
	//
	// It is created/provided by Hugging Face.
	//
	// It has 448 context window.
	//
	// It can be used with the following client methods:
	//	- CreateTranscription
	// 	- CreateTranslation
	ModelDistilWhisperLargeV3En AudioModel = "distil-whisper-large-v3-en"
	// ModelWhisperLargeV3 is an AI audio transcription model.
	//
	// It is created/provided by OpenAI.
	//
	// It has 448 context window.
	//
	// It can be used with the following client methods:
	//	- CreateTranscription
	// 	- CreateTranslation
	ModelWhisperLargeV3 AudioModel = "whisper-large-v3"
	// ModelWhisperLargeV3Turbo is an AI audio transcription model.
	//
	// It is created/provided by OpenAI.
	//
	// It has 448 context window.
	//
	// It can be used with the following client methods:
	//	- CreateTranscription
	// 	- CreateTranslation
	ModelWhisperLargeV3Turbo AudioModel = "whisper-large-v3-turbo"
	// ModelLlamaGuard38B is an AI moderation model.
	//
	// It is created/provided by Meta.
	//
	// It has 8192 context window.
	//
	// It can be used with the following client methods:
	//	- Moderate
	ModelLlamaGuard38B ModerationModel = "llama-guard-3-8b"
)
