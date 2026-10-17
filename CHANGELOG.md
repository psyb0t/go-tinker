# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.2, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.1, 2026-10-17

- Documentation only. The README and `AGENTS.md` describe the daily regeneration workflow: what it checks, how it fixes failing checks, how it picks the next version and which paths a fix may not touch.

## v0.1.0, 2026-10-16

- Added `tinker`, a REST client generated with oapi-codegen from the OpenAPI document published with Thinking Machines' Tinker SDK (Tinker API, spec version 0.1.0, OpenAPI 3.0.3). It covers all 15 paths: model creation and unloading, forward and backward passes, optimizer steps, sampling, weight saving and loading, futures, telemetry and health checks.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- Response wrappers use the `HTTPResponse` suffix, because the spec's schema names already end in `Response`.
- `make generate` reads the current spec URL from the `.stats.yml` of thinking-machines-lab/tinker, downloads the spec, rewrites its 3.1 style nullable type lists into 3.0 form, and regenerates the client. The spec is not committed.
- Added the `go-tinker` agent skill: `claude plugin install go-tinker@psyb0t` for Claude Code, `codex plugin add go-tinker@psyb0t` for Codex, `openclaw skills install @psyb0t/go-tinker` from ClawHub.
