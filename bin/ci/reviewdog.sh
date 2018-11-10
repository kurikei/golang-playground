#!/usr/bin/env bash
set -ex

REVIEWDOG_ARG="-reporter='github-pr-review'"
if [ "$CI_PULL_REQUEST" = "" ]; then
  REVIEWDOG_ARG="-diff='git diff master'"
fi

golint ./cmd/... ./test/... | eval reviewdog -f=golint $REVIEWDOG_ARG

golangci-lint run \
    --issues-exit-code 0 \
    --out-format checkstyle \
    --disable-all \
    -E govet \
    -E errcheck \
    -E ineffassign \
    -E interfacer \
    -E unconvert \
    -E misspell \
    -E unparam \
    -E nakedret \
    -E prealloc \
    ./cmd/... ./test/... \
    | eval reviewdog -name=golangci-lint -f=checkstyle $REVIEWDOG_ARG
