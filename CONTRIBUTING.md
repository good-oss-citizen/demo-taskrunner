# Contributing to taskrunner

Welcome! We're glad you want to contribute.

## First-Time Contributors

Look for issues labeled `good first issue` — these are specifically chosen to be
approachable for newcomers. They are scoped to be learning opportunities: you should
be able to understand the problem, read the surrounding code, and write a solution
that teaches you how the project works.

## Process

1. Comment on the issue to claim it
2. Fork and create a branch
3. Write tests for your change
4. Run `make test` and `make lint`
5. Open a PR referencing the issue

## Code Style

We use `gofmt` and `golangci-lint`. See `.golangci.yml` for the lint configuration.

## Commit Messages

Use the format: `type(scope): description`
Examples: `fix(parser): validate config keys`, `feat(scheduler): add cron support`
