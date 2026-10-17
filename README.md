# go-tinker

[![Go Reference](https://pkg.go.dev/badge/github.com/psyb0t/go-tinker.svg)](https://pkg.go.dev/github.com/psyb0t/go-tinker)
[![CI](https://github.com/psyb0t/go-tinker/actions/workflows/pipeline.yml/badge.svg?branch=main)](https://github.com/psyb0t/go-tinker/actions/workflows/pipeline.yml)
[![version](https://raw.githubusercontent.com/psyb0t/go-tinker/badges/version.svg)](https://github.com/psyb0t/go-tinker/tags)
[![license](https://raw.githubusercontent.com/psyb0t/go-tinker/badges/license.svg)](LICENSE)
[![imported by](https://raw.githubusercontent.com/psyb0t/go-tinker/badges/importers.svg)](https://github.com/psyb0t/go-tinker/blob/badges/importers.md)

Unofficial Go client for the [Thinking Machines Tinker API](https://thinkingmachines.ai/tinker/), generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from the OpenAPI document that ships with Thinking Machines' [Tinker SDK](https://github.com/thinking-machines-lab/tinker). I am not affiliated with Thinking Machines.

The package is `tinker`, in the repository root. Everything in `client.gen.go` is generated. Don't edit it by hand, regenerate it.

## Install

```sh
go get github.com/psyb0t/go-tinker
```

## Usage

Every endpoint in the spec has a typed method on `ClientWithResponses`. Responses decode into `JSON200` and the other status fields.

```go
package main

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-tinker"
)

func main() {
	auth := tinker.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		req.Header.Set("X-API-Key", os.Getenv("TINKER_API_KEY"))

		return nil
	})

	client, err := tinker.NewClientWithResponses("https://tinker.thinkingmachines.dev/services/tinker-prod", auth)
	if err != nil {
		panic(err)
	}

	res, err := client.GetServerCapabilitiesWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
```

The same code is in `example_test.go`, so `make test` checks that it compiles.

Most training and sampling calls (`Forward`, `ForwardBackward`, `OptimStep`, `Sample`, `SaveWeights` and the rest) return an `UntypedAPIFuture` with a `request_id`. Pass that ID to `RetrieveFuture` to get the result once the server has it.

Response wrappers end in `HTTPResponse` (`CreateModelHTTPResponse`, `OptimStepHTTPResponse`), because the spec already uses names like `CreateModelResponse` for its own schemas.

## Auth

The spec declares one scheme: an API key in the `X-API-Key` header. Set it in a request editor as shown above. The base URL in the example is the default the official Python SDK uses. The spec's own `servers` entry points at a different host it labels internal.

## Agent integrations

The [skill](.agents/skills/go-tinker) teaches a coding agent how to use go-tinker when it writes Go. It works in any agent that reads `.agents/skills/`, and installs natively in the clients below.

### Claude Code

```bash
claude plugin marketplace add psyb0t/agents
claude plugin install go-tinker@psyb0t
```

### Codex

```bash
codex plugin marketplace add psyb0t/agents
codex plugin add go-tinker@psyb0t
```

Codex also picks the skill up with no install in any checkout of this repo.

### OpenClaw

The skill is published to ClawHub on every release:

```bash
openclaw skills install @psyb0t/go-tinker
```

## Regenerating

The spec is not committed. `make generate` downloads the current OpenAPI document into the gitignored `api/` directory and regenerates the client:

```sh
make generate
```

`gen.go` declares the steps, so `make generate` is just `go generate ./...`:

- `scripts/fetch-spec.sh` reads `openapi_spec_url` from the `.stats.yml` of [thinking-machines-lab/tinker](https://github.com/thinking-machines-lab/tinker) and downloads that URL to `api/openapi.yaml`. Stainless hosts the spec under a URL with a content hash, which changes on every spec update, so the script looks it up each time.
- `scripts/fix-nullable-types.sh` writes `api/openapi.fixed.yaml`. The spec declares OpenAPI 3.0.3 but uses the 3.1 type list form for two nullable array items (`type: [number, 'null']`), which oapi-codegen rejects. The script rewrites each one to `type: number` plus `nullable: true` and drops the `nullable: true` the spec already has next to them.
- oapi-codegen runs with `config.yaml`. The config maps `number` to `float64`, because oapi-codegen would otherwise use `float32`, and sets the response wrapper suffix to `HTTPResponse` so the wrappers don't collide with the spec's schema names.

`.github/workflows/regen.yml` does the same every day in CI. When the new spec changes the client, it runs `make lint`, `make test` and `make sec`. If a check fails, Claude edits the code and the checks run again, up to three times, and a review step reads the result. A passing change is committed with a CHANGELOG entry and tagged as the next release: a patch version for fixes and additions, a minor version for breaking changes.

## Development

Every command runs in the `Dockerfile.dev` image. The host needs only Docker and Make.

```sh
make help      # list targets
make dep       # go mod tidy + vendor
make lint      # shellcheck, shfmt, go fix, golangci-lint
make test      # go test -race
make sec       # gitleaks, govulncheck, semgrep
make generate  # download the spec and regenerate
```

## License

[MIT](LICENSE). The generated code comes from the OpenAPI specification Thinking Machines publishes with its Apache-2.0 licensed Tinker SDK. The specification itself is not redistributed here.
