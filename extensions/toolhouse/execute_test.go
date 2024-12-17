package toolhouse_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/pegwings/pegwings-go"
	"github.com/pegwings/pegwings-go/extensions/toolhouse"
	"github.com/pegwings/pegwings-go/pkg/test"
	"github.com/pegwings/pegwings-go/pkg/tools"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	ts := test.NewTestServer()
	ts.RegisterHandler("/run_tools", func(w http.ResponseWriter, r *http.Request) {
		var runResp struct {
			Provider string `json:"provider"`
			Content  struct {
				Role       string `json:"role"`
				ToolCallID string `json:"tool_call_id"`
				Name       string `json:"name"`
				Content    string `json:"content"`
			} `json:"content"`
		}
		runResp.Content.Content = "response1"
		runResp.Content.Name = "tool"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonBytes, err := json.Marshal(runResp)
		a.NoError(err)
		_, err = w.Write(jsonBytes)
		a.NoError(err)
	})
	testS := ts.ToolhouseTestServer()
	testS.Start()
	client, err := toolhouse.NewExtension(
		test.GetTestToken(),
		toolhouse.WithBaseURL(testS.URL),
		toolhouse.WithClient(testS.Client()),
		toolhouse.WithLogger(test.DefaultLogger),
		toolhouse.WithMetadata(map[string]any{
			"id":       "conner",
			"timezone": 5,
		}),
	)
	a.NoError(err)
	history := []pegwings.ChatCompletionMessage{
		{
			Role:    pegwings.RoleUser,
			Content: "",
			ToolCalls: []tools.ToolCall{
				{
					Function: tools.FunctionCall{
						Name: "tool",
					},
				},
			},
		},
	}
	resp, err := client.Run(ctx, pegwings.ChatCompletionResponse{
		Choices: []pegwings.ChatCompletionChoice{
			{
				Message:      history[0],
				FinishReason: pegwings.ReasonFunctionCall,
			},
		},
	})
	a.NoError(err)
	assert.Equal(t, "response1", resp[0].Content)
}
