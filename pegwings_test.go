package pegwings_test

import (
	"log/slog"
	"net/http"
	"testing"

	"github.com/pegwings/pegwings-go"
	"github.com/stretchr/testify/assert"
)

// TestClient tests the creation of a new client.
func TestClient(t *testing.T) {
	a := assert.New(t)
	client, err := pegwings.NewClient(
		"test",
		pegwings.WithBaseURL("http://localhost/v1"),
		pegwings.WithClient(http.DefaultClient),
		pegwings.WithLogger(slog.Default()),
	)
	a.NoError(err)
	a.NotNil(client)
}
