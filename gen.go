// Package tinker is a Go client for the Thinking Machines Tinker API,
// generated with oapi-codegen from the OpenAPI document published with
// Thinking Machines' Tinker SDK.
package tinker

//go:generate bash scripts/fetch-spec.sh
//go:generate bash scripts/fix-nullable-types.sh
//go:generate go tool oapi-codegen --config=config.yaml api/openapi.fixed.yaml
