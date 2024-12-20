package pegwings_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/pegwings/pegwings-go"
	"github.com/pegwings/pegwings-go/pkg/builders"
	"github.com/pegwings/pegwings-go/pkg/test"
	"github.com/pegwings/pegwings-go/pkg/tools"
	"github.com/stretchr/testify/assert"
)

func TestChat(t *testing.T) {
	ctx := context.Background()
	a := assert.New(t)
	ts := test.NewTestServer()
	returnObj := pegwings.ChatCompletionResponse{
		ID:      "chatcmpl-123",
		Object:  "chat.completion.chunk",
		Created: 1693721698,
		Model:   "llama3-groq-70b-8192-tool-use-preview",
		Choices: []pegwings.ChatCompletionChoice{
			{
				Index: 0,
				Message: pegwings.ChatCompletionMessage{
					Role:    pegwings.RoleAssistant,
					Content: "Hello!",
				},
			},
		},
	}
	ts.RegisterHandler("/v1/chat/completions", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsval, err := json.Marshal(returnObj)
		a.NoError(err)
		_, err = w.Write(jsval)
		if err != nil {
			t.Fatal(err)
		}
	})
	testS := ts.GroqTestServer()
	testS.Start()
	client, err := pegwings.NewClient(
		test.GetTestToken(),
		pegwings.WithBaseURL(testS.URL+"/v1"),
	)
	a.NoError(err)
	resp, err := client.ChatCompletion(ctx, pegwings.ChatCompletionRequest{
		Model: pegwings.ModelLlama3Groq70B8192ToolUsePreview,
		Messages: []pegwings.ChatCompletionMessage{
			{
				Role:    pegwings.RoleUser,
				Content: "Hello!",
			},
		},
		MaxTokens: 2000,
		Tools:     []tools.Tool{},
	})
	a.NoError(err)
	a.NotEmpty(resp.Choices[0].Message.Content)
}

func TestAudioWithFailingFormBuilder(t *testing.T) {
	a := assert.New(t)
	dir, cleanup := test.CreateTestDirectory(t)
	defer cleanup()
	path := filepath.Join(dir, "fake.mp3")
	test.CreateTestFile(t, path)

	req := pegwings.AudioRequest{
		FilePath:    path,
		Prompt:      "test",
		Temperature: 0.5,
		Language:    "en",
		Format:      pegwings.FormatSRT,
	}

	mockFailedErr := fmt.Errorf("mock form builder fail")
	mockBuilder := &mockFormBuilder{}

	mockBuilder.mockCreateFormFile = func(string, *os.File) error {
		return mockFailedErr
	}
	err := pegwings.AudioMultipartForm(req, mockBuilder)
	a.ErrorIs(
		err,
		mockFailedErr,
		"audioMultipartForm should return error if form builder fails",
	)

	mockBuilder.mockCreateFormFile = func(string, *os.File) error {
		return nil
	}

	var failForField string
	mockBuilder.mockWriteField = func(fieldname, _ string) error {
		if fieldname == failForField {
			return mockFailedErr
		}
		return nil
	}

	failOn := []string{
		"model",
		"prompt",
		"temperature",
		"language",
		"response_format",
	}
	for _, failingField := range failOn {
		failForField = failingField
		mockFailedErr = fmt.Errorf(
			"mock form builder fail on field %s",
			failingField,
		)

		err = pegwings.AudioMultipartForm(req, mockBuilder)
		a.Error(
			err,
			mockFailedErr,
			"audioMultipartForm should return error if form builder fails",
		)
	}
}

func TestModeration(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	client, server, teardown := setupGroqTestServer()
	defer teardown()
	server.RegisterHandler("/v1/chat/completions", handleModerationEndpoint)
	mod, err := client.Moderate(ctx,
		[]pegwings.ChatCompletionMessage{
			{
				Role:    pegwings.RoleUser,
				Content: "I want to kill them.",
			},
		},
		pegwings.ModelLlamaGuard38B,
	)
	a.NoError(err)
	a.NotEmpty(mod)
}
func TestCreateFileField(t *testing.T) {
	a := assert.New(t)
	t.Run("createFileField failing file", func(t *testing.T) {
		t.Parallel()
		dir, cleanup := test.CreateTestDirectory(t)
		defer cleanup()
		path := filepath.Join(dir, "fake.mp3")
		test.CreateTestFile(t, path)
		req := pegwings.AudioRequest{
			FilePath: path,
		}
		mockFailedErr := fmt.Errorf("mock form builder fail")
		mockBuilder := &mockFormBuilder{
			mockCreateFormFile: func(string, *os.File) error {
				return mockFailedErr
			},
		}
		err := pegwings.AudioMultipartForm(req, mockBuilder)
		a.ErrorIs(
			err,
			mockFailedErr,
			"createFileField using a file should return error if form builder fails",
		)
	})

	t.Run("createFileField failing reader", func(t *testing.T) {
		t.Parallel()
		req := pegwings.AudioRequest{
			FilePath: "test.wav",
			Reader:   bytes.NewBuffer([]byte(`wav test contents`)),
		}

		mockFailedErr := fmt.Errorf("mock form builder fail")
		mockBuilder := &mockFormBuilder{
			mockCreateFormFileReader: func(string, io.Reader, string) error {
				return mockFailedErr
			},
		}

		err := pegwings.AudioMultipartForm(req, mockBuilder)
		a.ErrorIs(
			err,
			mockFailedErr,
			"createFileField using a reader should return error if form builder fails",
		)
	})

	t.Run("createFileField failing open", func(t *testing.T) {
		t.Parallel()
		req := pegwings.AudioRequest{
			FilePath: "non_existing_file.wav",
		}
		mockBuilder := builders.NewFormBuilder(&test.FailingErrorBuffer{})
		err := pegwings.AudioMultipartForm(req, mockBuilder)
		a.Error(
			err,
			"createFileField using file should return error when open file fails",
		)
	})
}

// mockFormBuilder is a mock form builder.
type mockFormBuilder struct {
	mockCreateFormFile       func(string, *os.File) error
	mockCreateFormFileReader func(string, io.Reader, string) error
	mockWriteField           func(string, string) error
	mockClose                func() error
}

// CreateFormFile is a mock form builder create form file method.
func (fb *mockFormBuilder) CreateFormFile(
	fieldname string,
	file *os.File,
) error {
	return fb.mockCreateFormFile(fieldname, file)
}

// CreateFormFileReader is a mock form builder create form file reader method
func (fb *mockFormBuilder) CreateFormFileReader(
	fieldname string,
	r io.Reader,
	filename string,
) error {
	return fb.mockCreateFormFileReader(fieldname, r, filename)
}

// WriteField is a mock form builder write field method.
func (fb *mockFormBuilder) WriteField(fieldname, value string) error {
	return fb.mockWriteField(fieldname, value)
}

// Close is a mock form builder close method.
func (fb *mockFormBuilder) Close() error {
	return fb.mockClose()
}

// FormDataContentType is a mock form builder.
func (fb *mockFormBuilder) FormDataContentType() string {
	return ""
}
