package toolhouse_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/pegwings/pegwings-go/extensions/toolhouse"
	"github.com/pegwings/pegwings-go/pkg/test"
	"github.com/pegwings/pegwings-go/pkg/tools"
	"github.com/stretchr/testify/assert"
)

func TestGetTools(t *testing.T) {
	a := assert.New(t)
	ctx := context.Background()
	ts := test.NewTestServer()
	ts.RegisterHandler("/get_tools", func(w http.ResponseWriter, r *http.Request) {
		var ts []tools.Tool
		ts = append(ts, tools.Tool{
			Function: tools.FunctionDefinition{
				Name:        "tool",
				Description: "tool",
				Parameters:  tools.FunctionParameters{},
			},
			Type: tools.ToolTypeFunction,
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonBytes, err := json.Marshal(ts)
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
	tools, err := client.GetTools(ctx)
	a.NoError(err)
	a.NotEmpty(tools)
}
