#!/bin/bash

test "${1}" = "" && echo "Usage: $0 <version>" && exit 1

# needs to be exported so the subshells of the find command can see it
export VERSION=${1}

# Tag root module
git tag "v${VERSION}"

# This script tags all go modules in the repository with the current version
find . -name "go.mod" -type f -mindepth 2 -maxdepth 2 -exec sh -cx '
  DIR=$(dirname ${1})
  BASE=$(basename $DIR)
  TAG="${BASE}/v${VERSION}"
  git tag "$TAG"
  ' shell {} \;
