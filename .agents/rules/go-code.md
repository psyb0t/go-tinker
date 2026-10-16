# Go code

Refines `~/.claude/rule-details/working-with-generated-code.md` and `~/.claude/rule-details/coding-style.md`.

- **GO1** `client.gen.go` is oapi-codegen output. Never edit it by hand. Change `config.yaml` or a script under `scripts/`, then run `make generate` and commit the result.
- **GO2** A spec that oapi-codegen cannot read gets fixed by a patch step in `scripts/`, run from a `//go:generate` line in `gen.go` between `fetch-spec.sh` and oapi-codegen. Patch scripts read `api/openapi.yaml`, write `api/openapi.fixed.yaml`, and copy anything they do not recognize through untouched so oapi-codegen fails loudly.
- **GO3** The spec is never committed. `api/` is gitignored; `scripts/fetch-spec.sh` resolves the current URL from the `.stats.yml` of thinking-machines-lab/tinker on every run.
- **GO4** `config.yaml` keeps `number` mapped to `float64` (including `format: float`) and the `HTTPResponse` response suffix. Dropping the suffix brings back the `OptimStepResponse2` name collision.
- **GO5** Hand-written Go is limited to `gen.go` and the tests. The package never logs and never prints.
- **GO6** `make lint` passes with zero issues. Shell scripts pass shellcheck and shfmt.
