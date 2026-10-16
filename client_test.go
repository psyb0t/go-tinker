package tinker_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/psyb0t/go-tinker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testAPIKey    = "test-key"
	testModelID   = "123e4567-e89b-12d3-a456-426614174000"
	testRequestID = "223e4567-e89b-12d3-a456-426614174001"
)

func newTestClient(
	t *testing.T,
	handler http.HandlerFunc,
) *tinker.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	auth := tinker.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("X-API-Key", testAPIKey)

			return nil
		},
	)

	client, err := tinker.NewClientWithResponses(srv.URL, auth)
	require.NoError(t, err)

	return client
}

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/v1/healthz", r.URL.Path)
		assert.Equal(t, testAPIKey, r.Header.Get("X-API-Key"))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	res, err := client.HealthCheckWithResponse(context.Background())
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)

	assert.Equal(t, tinker.HealthResponseStatus("ok"), res.JSON200.Status)
}

func TestCreateModel(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/v1/create_model", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, testAPIKey, r.Header.Get("X-API-Key"))

		var body tinker.CreateModelRequest

		assert.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "base-model-x", body.BaseModel)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"request_id":"` + testRequestID + `",` +
				`"model_id":"` + testModelID + `"}`,
		))
	})

	res, err := client.CreateModelWithResponse(
		context.Background(),
		tinker.CreateModelRequest{BaseModel: "base-model-x"},
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)
	require.NotNil(t, res.JSON200.ModelId)

	assert.Equal(t, testRequestID, res.JSON200.RequestId.String())
	assert.Equal(t, testModelID, res.JSON200.ModelId.String())
}
