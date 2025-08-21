# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This repository contains the Pulumi Azure Native provider, which enables Pulumi users to manage Azure resources using TypeScript, Python, Go, .NET, and Java. The provider directly calls the Azure Resource Manager (ARM) REST API, providing comprehensive coverage of Azure services.

## Build and Development Commands

### Setup and Dependencies

- **Dependencies**:
  - Go 1.21+
  - NodeJS 18.X.X+
  - Python 3.10+
  - .NET 6+
  - Gradle 8+

### Core Build Commands

```bash
# Full build (generate schema, SDKs, and build the provider)
make

# Generate schema only
make schema

# Regenerate schema and update docs
make generate_schema generate_docs

# Generate SDKs for all languages
make generate

# Build the provider binary
make provider

# Install the provider locally (used for testing)
make install_provider

# Clean the build
make clean
```

### SDK Generation and Building

```bash
# Generate SDKs for specific languages
make generate_java
make generate_nodejs
make generate_python
make generate_dotnet
make generate_go

# Build SDKs for specific languages
make build_nodejs
make build_python
make build_dotnet
make build_java
make build_go
```

### Testing

```bash
# Run all tests
make test

# Run tests for specific languages
make test_dotnet
make test_python
make test_go
make test_nodejs

# Run provider tests
make test_provider

# Run specific test by name
make test TEST_NAME=TestXXX

# Run specific test category
make test TEST_TAGS=nodejs
```

### Schema Management

```bash
# Generate schema squeeze (for major version releases)
make schema_squeeze
```

## Code Architecture

### Key Components

1. **Azure Specifications**: The repository uses Azure's REST API specifications from the `azure-rest-api-specs` submodule.

2. **Provider**: The core provider implementation is in `provider/pkg`, which handles:
   - Authentication with Azure
   - CRUD operations for resources
   - Type conversions between Pulumi and Azure
   - Handling of special cases and customizations

3. **Code Generator**: The `pulumi-gen-azure-native` tool generates schemas and SDKs based on Azure's OpenAPI specifications.

4. **SDKs**: Generated SDKs for multiple languages in the `sdk/` directory:
   - TypeScript/JavaScript (`sdk/nodejs`)
   - Python (`sdk/python`) 
   - .NET (`sdk/dotnet`)
   - Go (`sdk/go`)
   - Java (`sdk/java`)

5. **Version Management**: The provider supports multiple API versions through configuration in the `versions/` directory:
   - Each major version has its own configuration
   - API versions are selectively included based on compatibility and features

### Important Subsystems

#### Sub-Resources

The provider handles "sub-resource properties" specially. These are resources that can be defined inline with a parent resource or as standalone resources. The provider ensures that:
- On Create: Default values are set for unset sub-resource properties
- On Update: Existing sub-resources are preserved during updates
- On Read: Unset sub-resource properties are reset to empty values

#### Custom Resources

Some Azure resources require special handling beyond what can be generated from the OpenAPI specs. These custom implementations are in `provider/pkg/resources/customresources/`.

## Development Workflow

1. **Understanding API Versions**: Azure's REST API has many versions with ISO date format (e.g., `2020-01-01`).
   - Not all versions contain all resources
   - Preview versions (`-preview` suffix) are sometimes stable for everyday use
   - The provider selectively includes versions to maintain compatibility

2. **Making Changes**:
   - For schema changes, modify files in `versions/` directory
   - For provider behavior changes, modify files in `provider/pkg/`
   - For custom resource implementations, modify files in `provider/pkg/resources/customresources/`

3. **Testing Changes**:
   - Run unit tests with `make test_provider`
   - Test examples with `make test` or language-specific tests

4. **Adding Resource Documentation**:
   - Add documentation in `docs/resources/` folder
   - File naming follows the pattern: `[module]-[Resource].md` (e.g., `sql-Server.md`)
   - Regenerate docs with `make generate generate_docs`

## Common Issues and Solutions

1. **Schema Generation**: If schema generation fails, check that the Azure REST API specs submodule is properly initialized:
   ```bash
   git submodule update --init --recursive
   ```

2. **Sub-resource Handling**: When troubleshooting issues with resources that contain sub-resources, check the implementation in `provider/pkg/provider/`.

3. **Version Compatibility**: When adding support for new API versions, check for compatibility with existing versions in the `versions/` directory.