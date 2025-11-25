# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is the Pulumi Azure Native provider, which generates a Pulumi provider from Azure's OpenAPI specifications. The provider works directly with the Azure Resource Manager (ARM) REST API rather than using a handwritten abstraction layer. It's a **code generation-focused repository** where the majority of the SDK and schema are generated from the `azure-rest-api-specs` submodule.

## v2 Branch - Maintenance Only

**You are currently on the v2 maintenance branch.** This branch is for v2.x patch releases only.

### Key Points

- **Purpose**: Bug fixes and small improvements for v2.x users
- **No new features**: All new development happens on `master` (v3.x)
- **Schema**: v2 uses `bin/schema-full.json` (complete schema with all API versions)
- **Version**: Set `PROVIDER_VERSION=2.91.0-alpha.0+dev` for local builds
- **First release**: Next v2 release will be v2.91.0 (minor version increment)

### Workflows

- **head.yml**: CI for v2 branch - runs prerequisites + provider unit tests only (no expensive SDK tests)
- **release.yml**: Placeholder for v2.x.x stable releases (no prereleases)
  - Only triggers on `v2.*.*` tags (no `-beta`, `-rc` suffixes)
  - When ready to release, will need to add build/publish steps with:
    - `make_latest: false` for GitHub releases
    - NPM tag "v2" (not "latest")
    - ESC secrets for publishing

### Testing on v2

```bash
# Build schema (should generate schema-full.json, not schema-default-versions.json)
make schema

# Run provider unit tests only (fast)
make test_provider PROVIDER_TEST_TAGS=unit

# Build provider
make provider
```

### Release Process

1. Create feature branch from `v2` (e.g., `feature/fix-xyz-v2`)
2. Implement fix/patch
3. Open PR to `v2` branch (not master)
4. After merge, tag with next minor version (v2.91.0, v2.92.0, etc.)
5. Release workflow publishes to registries

### Cherry-picking from Master

To backport a fix from master (v3):

```bash
git checkout v2
git cherry-pick <commit-hash>
# Resolve conflicts if version-specific code differs
```

**Note**: Some v3 fixes may not apply cleanly to v2 due to architectural differences (e.g., v3 has `azureApiVersion` output, v2 doesn't).

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

### Manual Testing for API Version Upgrades

When testing scenarios involving API version changes (e.g., state migration, refresh behavior), it's important to understand how `azureApiVersion` works:

**How `azureApiVersion` is Determined:**

1. **NOT from SDK**: The `azureApiVersion` field is computed by the provider, not supplied as an input from the SDK
2. **From Provider Metadata**: The API version comes from `metadata-compact.json` baked into the provider binary
3. **Lookup Process**:
   - SDK sends URN with type token (e.g., `azure-native:containerservice/v20241002preview:ManagedCluster`)
   - Provider looks up this token in its embedded metadata
   - Provider uses `res.APIVersion` from metadata for all operations
   - Provider sets `azureApiVersion` output after every CRUD operation

**API Version Format Conversion:**

The provider uses two different formats for API versions:

1. **SDK Version Format** (used in type tokens): `v20241002preview`
   - Example: `azure-native:containerservice/v20241002preview:ManagedCluster`
   - Format: `v` + `YYYYMMDD` + optional suffix (preview, beta, privatepreview)

2. **API Version Format** (used in Azure ARM API calls and state): `2024-10-02-preview`
   - Stored in `azureApiVersion` output property
   - Format: `YYYY-MM-DD` + optional `-suffix`

**Conversion Functions** (in `provider/pkg/openapi/versioner.go`):
- `openapi.ApiToSdkVersion()` - Converts `2024-10-02-preview` → `v20241002preview`
- `openapi.SdkToApiVersion()` - Converts `v20241002preview` → `2024-10-02-preview`

When working with API versions in provider code, always use these canonical conversion functions rather than reimplementing the logic.

**Testing Approaches for API Version Changes:**

**Option 1: Modify Version Configuration (Most Realistic)**
```bash
# 1. Modify default versions for a resource
vi versions/v3-config.yaml  # Change StorageAccount default version

# 2. Regenerate schema and metadata
make schema

# 3. Rebuild provider with new metadata
make provider

# 4. Deploy resources, then repeat steps 1-3 with different version
# 5. Run pulumi refresh to test migration behavior
```

**Option 2: State Manipulation (Quick & Effective)**
```bash
# 1. Deploy resources with current provider
pulumi up

# 2. Export and modify state to simulate old API version
pulumi stack export > state.json
jq '(.deployment.resources[] | select(.type == "azure-native:storage:StorageAccount")) |= (.outputs.azureApiVersion = "2023-01-01")' state.json > state-old.json

# 3. Import modified state
pulumi stack import --file state-old.json

# 4. Run refresh (provider reads with new API, compares with old)
pulumi refresh --yes

# 5. Verify no spurious changes detected
```

**Option 3: Two Provider Binaries (Most Thorough)**
```bash
# 1. Checkout and build provider at earlier version
git checkout v3.50.0
make provider
cp bin/pulumi-resource-azure-native ~/provider-old

# 2. Deploy resources using old provider
PATH="$HOME:$PATH" pulumi up

# 3. Checkout and build provider at newer version
git checkout v3.100.0
make provider
cp bin/pulumi-resource-azure-native ~/provider-new

# 4. Run refresh with new provider
PATH="$HOME:$PATH" pulumi refresh --yes

# 5. Verify no spurious changes or replacements
```

**Why This Matters:**

- The SDK version does NOT control API versions - the provider binary's metadata does
- When users upgrade provider versions, default API versions can change
- The provider must correctly handle state migration when `azureApiVersion` differs between old state and new metadata
- Test scenarios should validate that refresh operations don't report spurious property changes when only the API version changed

See issue #4400 for an example of a bug that required this type of testing to validate the fix.

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
