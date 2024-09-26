[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fazure-native.svg)](https://npmjs.com/package/@pulumi/azure-native)
[![Python version](https://badge.fury.io/py/pulumi-azure-native.svg)](https://pypi.org/project/pulumi-azure-native)
[![NuGet version](https://badge.fury.io/nu/pulumi.azurenative.svg)](https://badge.fury.io/nu/pulumi.azurenative)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-azure-native/sdk/go)](https://pkg.go.dev/github.com/pulumi/pulumi-azure-native/sdk/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fazure-native.svg)](https://github.com/pulumi/pulumi-azure-native/blob/master/LICENSE)

# Native Azure Pulumi Provider

The [Azure Native](https://www.pulumi.com/docs/intro/cloud-providers/azure/) provider for Pulumi lets you use Azure resources in your cloud programs.
This provider uses the Azure Resource Manager REST API directly and therefore provides full access to the ARM API.

The Azure Native provider is the recommended provider for projects targeting Azure.

To use this package, [install the Pulumi CLI](https://www.pulumi.com/docs/get-started/install/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    npm install @pulumi/azure-native

or `yarn`:

    yarn add @pulumi/azure-native

### Python

To use from Python, install using `pip`:

    pip install pulumi_azure_native

### Go

To use from Go, use `go get` to grab the latest version of the library

    go get github.com/pulumi/pulumi-azure-native/sdk

### .NET

To use from .NET, install using `dotnet add package`:

    dotnet add package Pulumi.AzureNative

## Concepts

The `@pulumi/azure-native` package provides a strongly-typed means to build cloud applications that create
and interact closely with Azure resources.  Resources are exposed for the entire Azure surface area,
including (but not limited to) 'compute', 'keyvault', 'network', 'storage', and more.

The Azure Native provider works directly with the Azure Resource Manager (ARM) platform instead of depending on a
handwritten layer as with the [classic provider](https://github.com/pulumi/pulumi-azure). This approach ensures higher
quality and higher fidelity with the Azure platform.

## Configuring credentials

To learn how to configure credentials refer to the [Azure configuration options](https://www.pulumi.com/registry/packages/azure-native/installation-configuration/#configuration-options).

## Other Configuration

In addition to the configuration options in the official documentation linked above, the following environment variables can be used to tweak lower-level behavior of the provider:

  - `PULUMI_FORCE_NEW_FROM_SUBTYPES`: if, after a change in your program, the provider does not replace a resource that should be replaced because it cannot be updated, setting this variable to `true` might allow the provider to infer the correct behavior. For more details please see PR #2970. We're planning to make this behavior the default in the future.

## Building

See [contributing](CONTRIBUTING.md) for details on how to build and contribute to this provider.

## Reference

For further information, visit [Azure Native in the Pulumi Registry](https://www.pulumi.com/registry/packages/azure-native/)
or for detailed API reference documentation, visit [Azure Native API Docs in the Pulumi Registry](https://www.pulumi.com/registry/packages/azure-native/api-docs/).  
