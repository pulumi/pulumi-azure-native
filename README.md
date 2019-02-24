# Azure Resource Manager Provider

An implementation of a Pulumi Resource Provider directly targeting the ARM REST
APIs and Swagger specs.

### Building and Testing

Clone this repository into `$GOPATH/src/github.com/pulumi/pulumi-azure` - note
that the name of the repository currently differs from the import path.

```
$ go install ./cmd/pulumi-resource-azurerm/
$ cd ./exampes/simple
$ pulumi up
```
