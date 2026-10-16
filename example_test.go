package tinker_test

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-tinker"
)

// This is the README usage example. It talks to the live API, so it has no
// Output comment and only has to compile.
//
//nolint:testableexamples // needs real credentials, compile check only
func Example() {
	auth := tinker.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("X-API-Key", os.Getenv("TINKER_API_KEY"))

			return nil
		},
	)

	client, err := tinker.NewClientWithResponses(
		"https://tinker.thinkingmachines.dev/services/tinker-prod",
		auth,
	)
	if err != nil {
		panic(err)
	}

	res, err := client.GetServerCapabilitiesWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
