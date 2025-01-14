#!/usr/bin/env bash

set -e

make schema_squeeze
make generate_schema
make generate_docs
git add .
git commit -m "Generate squeezed schema and docs"
rm -rf sdk/nodejs
git add sdk/nodejs
make generate_nodejs
git add sdk/nodejs
git commit -m "Generate Node.js SDK"
rm -rf sdk/python
git add sdk/python
make generate_python
git add sdk/python
git commit -m "Generate Python SDK"
rm -rf sdk/dotnet
git add sdk/dotnet
make generate_dotnet
git add sdk/dotnet
git commit -m "Generate .NET SDK"
