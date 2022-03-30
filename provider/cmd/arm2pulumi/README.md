# Arm2Pulumi

A tool to convert Azure ARM templates to Pulumi code

The tool is currently hosted at https://www.pulumi.com/arm2pulumi/. It can also be used as a CLI tool locally.

## Building

### Dependencies

- Go 1.16+

Please refer to [Contributing to Pulumi](https://github.com/pulumi/pulumi/blob/master/CONTRIBUTING.md) for installation
guidance.

### Building locally

Run the following commands to install Go modules, generate all SDKs, and build the provider: 

```
$ git clone git@github.com:pulumi/pulumi-azure-native.git
$ cd pulumi-azure-native/provider
$ make ensure
$ make codegen generate_schema arm2pulumi
```

Add the `bin` folder to your `$PATH` or copy the `bin/arm2pulumi` file to another location in your `$PATH`.

### Performing a conversion

Suppose you have an ARM template - `storageAccount.json` which creates a storage account:
```
{
    "$schema": "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
    "contentVersion": "1.0.0.0",
    "parameters": {
        "storageAccountType": {
            "type": "string",
            "defaultValue": "Standard_LRS",
            "allowedValues": [
                "Standard_LRS",
                "Standard_GRS",
                "Standard_ZRS",
                "Premium_LRS"
            ],
            "metadata": {
                "description": "Storage Account type"
            }
        },
        "location": {
            "type": "string",
            "defaultValue": "[resourceGroup().location]",
            "metadata": {
                "description": "Location for all resources."
            }
        },
        "storageAccountName": {
            "type": "string"
        }
    },
    "resources": [
        {
            "type": "Microsoft.Storage/storageAccounts",
            "apiVersion": "2019-04-01",
            "name": "[parameters('storageAccountName')]",
            "location": "[parameters('location')]",
            "sku": {
                "name": "[parameters('storageAccountType')]"
            },
            "properties": {
                "sku": {
                    "name": "[parameters('storageAccountType')]"
                },
                "kind": "StorageV2"
            }
        }
    ],
    "outputs": {
        "storageAccountName": {
            "type": "string",
            "value": "[reference(parameters('storageAccountName')).name]"
        }
    }
}
```

To convert run the following command:
`$ arm2pulumi storageAccount.json <language>`
where language is one of `python`, `go`, `nodejs`, `dotnet (C#)`


This will print corresponding Pulumi code in the desired language.

For further information on using the Pulumi Azure Native SDK, please visit the [native Azure provider announcement](https://www.pulumi.com/blog/full-coverage-of-azure-resources-with-azure-native/) or for detailed reference documentation, please visit [the API docs](
https://www.pulumi.com/docs/reference/pkg/azure-native).

Issues with `arm2pulumi` can be reported to https://github.com/pulumi/arm2pulumi/issues.
