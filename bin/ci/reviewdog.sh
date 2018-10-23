#!/usr/bin/env bash
set -ex

REVIEWDOG_ARG="-reporter='github-pr-review'"
if [ "$CI_PULL_REQUEST" = "" ]; then
  REVIEWDOG_ARG="-diff='git diff master'"
fi

golint $(go list ./${SRC_PATH}/...) | eval reviewdog -f=golint $REVIEWDOG_ARG
