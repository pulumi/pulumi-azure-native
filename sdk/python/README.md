# Next Generation Azure Pulumi Provider (preview)

The Next Generation Azure Provider for Pulumi lets you use Azure resources in your cloud programs.
This provider uses the Azure Resource Manager REST API directly and therefore provides full access to ARM API.

The provider is currently in private preview:

- Breaking changes are regularly introduced, including changes in SDKs, provider plugin, and state format
- No artifacts are published in public package managers or Pulumi provider registry. You'll have to build them
yourself in order to use the provider.

To use this package, please [install the Pulumi CLI first](https://pulumi.io/).

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
$ yarn link @pulumi/azure-nextgen
$ pulumi up
``` 
