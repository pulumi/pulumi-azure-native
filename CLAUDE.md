# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is the Pulumi Azure Native provider, which generates a Pulumi provider from Azure's OpenAPI specifications. The provider works directly with the Azure Resource Manager (ARM) REST API rather than using a handwritten abstraction layer. It's a **code generation-focused repository** where the majority of the SDK and schema are generated from the `azure-rest-api-specs` submodule.

## Key Architecture Concepts

### Code Generation Pipeline

The entire provider is built around a multi-stage code generation pipeline:

1. **Specs Ingestion**: Azure OpenAPI specs are read from the `azure-rest-api-specs` git submodule
2. **Schema Generation**: The `pulumi-gen-azure-native` codegen tool (`bin/$(CODEGEN)`) processes specs and generates:
   - `bin/schema-full.json` - Complete schema with all API versions
   - `bin/schema-default-versions.json` - Schema with only default versions (used for v3+)
   - `bin/metadata-compact.json` - Resource metadata
3. **SDK Generation**: Pulumi's SDK generators create language-specific SDKs (Node.js, Python, .NET, Go, Java)
4. **Provider Binary**: The `pulumi-resource-azure-native` binary embeds the schema and serves as the gRPC provider

### Version Management System

This provider has a sophisticated version management system due to Azure's massive API surface:

- **API Versions**: Azure services are versioned per-service using ISO dates (e.g., `2020-01-01`, `2020-01-01-preview`)
- **Default Versions**: To simplify discovery, the provider exposes a single "default" version for each API that combines resources from multiple API versions at their latest stable versions
- **Version Configuration Files** (in `versions/`):
  - `v{N}-config.yaml` - Hand-edited configuration controlling version selection
  - `v{N}-spec.yaml` - Auto-generated spec file (can be manually edited)
  - `v{N}.yaml` - Calculated default versions file (generated, do not edit directly)
  - `v{N}-removed-resources.json` - Resources removed for size optimization
  - `v{N}-removed.json` - Removed API versions
- **Version Squeezing**: `make schema_squeeze` identifies redundant API versions that can be removed to reduce SDK size

### Sub-Resources Pattern

Some Azure resources can be defined inline on a parent resource OR as standalone resources (e.g., VirtualNetwork subnets). The provider has special CRUD lifecycle handling for this:

- **On Create**: Sets sub-resource properties to default (empty array) if not specified
- **On Update**: Retrieves existing sub-resources and round-trips them to avoid deletion
- **On Read**: Resets sub-resource properties to empty if they were defined standalone

See `provider/pkg/provider` for the implementation in `setUnsetSubresourcePropertiesToDefaults`, `maintainSubResourcePropertiesIfNotSet`, and `resetUnsetSubResourceProperties` methods.

### Go SDK Split Modules

The Go SDK is published to a separate repository (`github.com/pulumi/pulumi-azure-native-sdk`) with a Go module per Azure namespace rather than a single root module, due to the 512MB Go module size limit. The codegen automatically generates the required module structure and dependencies.

## Common Development Commands

### Building

```bash
# Full build (schema generation + SDK generation + provider binary)
make

# Build just the provider binary
make provider

# Build just the codegen tool
make codegen

# Build individual SDKs
make build_nodejs
make build_python
make build_dotnet
make build_go
make build_java
```

### Schema Generation

```bash
# Generate schema from Azure OpenAPI specs
make schema

# Alias for schema generation
make generate_schema

# Generate docs schema (with examples)
make generate_docs

# Calculate which API versions can be removed (for major version releases)
make schema_squeeze
```

### Testing

```bash
# Run provider unit tests
make test_provider

# Run provider unit tests only (default, fast)
make test_provider PROVIDER_TEST_TAGS=unit

# Run all provider tests including integration tests
make test_provider PROVIDER_TEST_TAGS=all

# Run SDK integration tests (requires Azure credentials)
make test           # All languages
make test_nodejs
make test_python
make test_dotnet
make test_go

# Run specific test by name
make test TEST_NAME=TestAccResourceGroup
```

### SDK Development

```bash
# Generate all SDKs
make generate

# Install SDKs locally for testing
make install_sdks           # Node.js + .NET
make install_nodejs_sdk
make install_dotnet_sdk

# Link Node.js SDK for local examples
cd examples/simple
yarn link @pulumi/azure-native
pulumi up
```

### Cleaning

```bash
# Clean all generated artifacts
make clean
```

### Debugging Code Generation

Use environment variables for faster iteration:

```bash
# Generate only specific namespace(s)
DEBUG_CODEGEN_NAMESPACES="compute" make schema

# Generate only specific API version(s)
DEBUG_CODEGEN_APIVERSIONS="2019-09-01" make schema

# Combine both for targeted generation
DEBUG_CODEGEN_NAMESPACES="compute" DEBUG_CODEGEN_APIVERSIONS="2019*" make schema

# Enable debug logging
DEBUG_CODEGEN=true make schema
```

## Project Structure

- `provider/` - Provider implementation in Go
  - `cmd/pulumi-gen-azure-native/` - Codegen tool that generates schemas from OpenAPI specs
  - `cmd/pulumi-resource-azure-native/` - Provider binary that serves resources via gRPC
  - `pkg/gen/` - Schema generation logic
  - `pkg/openapi/` - OpenAPI spec parsing and resolution
  - `pkg/versioning/` - Version selection and management logic
  - `pkg/resources/` - Resource CRUD logic and custom resource implementations
  - `pkg/convert/` - Conversion between SDK shapes and ARM API request/response formats
  - `pkg/tle/` - Template Language Expression parser (for ARM template expressions)
- `sdk/` - Generated SDK code (not committed except Go SDK structure)
  - `nodejs/`, `python/`, `dotnet/`, `java/` - Language-specific SDKs
  - `pulumi-azure-native-sdk/` - Split Go SDK (published to separate repo)
- `versions/` - Version configuration and metadata files
- `azure-rest-api-specs/` - Git submodule with Azure OpenAPI specifications
- `examples/` - Integration test examples in Go test framework
- `docs/resources/` - Additional hand-written resource documentation (markdown files)

## Version Variables

- `PROVIDER_VERSION` - Must be set in CI, defaults to `3.0.0-alpha.0+dev` locally
- Major version 3+ uses `schema-default-versions.json`; earlier versions use `schema-full.json`
- Version is embedded in binaries via `-ldflags` during build

## Important Implementation Details

### Custom Resources

Some resources require special handling beyond standard CRUD. Custom resource implementations are in `provider/pkg/resources/customresources/`. Examples include:
- Key Vault access policies
- Portal dashboards
- Recovery services
- PIM eligibility assignments

### Metadata Compact

The `metadata-compact.json` file contains resource metadata including:
- Resource IDs and ARM paths
- PUT/GET/DELETE method information
- API version mappings
This metadata is used by the provider at runtime for constructing API calls.

### Schema Files

- `bin/schema-full.json` - Full schema with all API versions (large, ~100MB+)
- `bin/schema-default-versions.json` - Default versions only (smaller, used for v3+)
- `provider/cmd/pulumi-resource-azure-native/schema.json` - Docs schema with examples (committed to repo)

## Testing Notes

- Examples are written as Go tests in `examples/`
- Tests are tagged: `unit`, `nodejs`, `python`, `dotnet`, `go`, `java`, `all`
- Tests require Azure credentials configured
- Use `TEST_NAME` to run specific tests
- Full test suite can take 2+ hours

## Submodule Management

```bash
# Update azure-rest-api-specs submodule
make update_submodules

# This also updates az-provider-list.json with latest API versions from Azure
```

## Adding Resource Documentation

1. Create a markdown file in `docs/resources/` named `[module]-[Resource].md` (e.g., `sql-Server.md`)
2. Run `make generate generate_docs` to regenerate SDKs with the new documentation
3. Commit changes and open a pull request
