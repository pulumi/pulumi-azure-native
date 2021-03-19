[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fazure-native.svg)](https://npmjs.com/package/@pulumi/azure-native)
[![Python version](https://badge.fury.io/py/pulumi-azure-native.svg)](https://pypi.org/project/pulumi-azure-native)
[![NuGet version](https://badge.fury.io/nu/pulumi.azurenative.svg)](https://badge.fury.io/nu/pulumi.azurenative)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-azure-native/sdk/go)](https://pkg.go.dev/github.com/pulumi/pulumi-azure-native/sdk/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fazure-native.svg)](https://github.com/pulumi/pulumi-azure-native/blob/master/LICENSE)

# Native Azure Pulumi Provider (preview)

The native Azure Provider for Pulumi lets you use Azure resources in your cloud programs.
This provider uses the Azure Resource Manager REST API directly and therefore provides full access to ARM API.

The provider is currently in public preview, and is the recommended provider for new projects targeting Azure. We
will attempt to minimize any breaking changes, but it is possible that there will be a few between now and GA (expected
this spring).

To use this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/azure-native

or `yarn`:

    $ yarn add @pulumi/azure-native

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_azure_native

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-azure-native/sdk

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.AzureNative

## Concepts

The `@pulumi/azure-native` package provides a strongly-typed means to build cloud applications that create
and interact closely with Azure resources.  Resources are exposed for the entire Azure surface area,
including (but not limited to) 'compute', 'keyvault', 'network', 'storage', and more.

The Azure Native provider works directly with the Azure Resource Manager (ARM) platform instead of depending on a
handwritten layer as with the [previous provider](https://github.com/pulumi/pulumi-azure). This approach ensures higher
quality and higher fidelity with the Azure platform.

The Azure Native provider is currently in beta. While we try to minimize impact on users, breaking changes may be shipped in minor version updates.

## Configuring credentials

Credentials configuration is compatible with the existing Terraform-based Azure provider.

Please refer to [this quickstart guide](
https://www.pulumi.com/docs/intro/cloud-providers/azure/setup/) for possible configuration options.

## Building

### Dependencies

- Go 1.15
- NodeJS 10.X.X or later
- Python 3.6 or later
- .NET Core 3.1

Please refer to [Contributing to Pulumi](https://github.com/pulumi/pulumi/blob/master/CONTRIBUTING.md) for installation
guidance.

### Building locally

Run the following commands to install Go modules, generate all SDKs, and build the provider: 

```
$ make ensure
$ make build
```

Add the `bin` folder to your `$PATH` or copy the `bin/pulumi-resource-azurerm` file to another location in your `$PATH`.

### Running an example

Navigate to one of the `examples` and run Pulumi:

```
$ cd ./exampes/simple
$ yarn link @pulumi/azure-native
$ pulumi up
``` 

## Reference

For further information, please visit the [Next Generation Azure provider announcement](https://www.pulumi.com/blog/announcing-nextgen-azure-provider/) or for detailed reference documentation, please visit [the API docs](
https://www.pulumi.com/docs/reference/pkg/azure-native).
