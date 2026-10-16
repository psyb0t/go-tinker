# Testing

Refines `~/.claude/rule-details/go/writing-tests.md`.

- **TST1** Tests are black-box, in package `tinker_test`, and use only the exported API.
- **TST2** Endpoint tests run the generated client against `httptest.NewServer` and assert the HTTP method, the path, the `X-API-Key` header and the decoded response. They never call the live API.
- **TST3** Every test and subtest calls `t.Parallel()`.
- **TST4** `example_test.go` holds the README usage example verbatim in shape. It needs real credentials, so it only compiles (`//nolint:testableexamples` with the reason). Change both together.
- **TST5** `make test` runs `go test -race` and must pass. Test values are obvious placeholders, never real keys or model names from a live account.
