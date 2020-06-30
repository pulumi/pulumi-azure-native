module gosimple

go 1.14

require (
	github.com/pulumi/pulumi-azurerm/sdk v0.1.0
	github.com/pulumi/pulumi/sdk/v2 v2.0.0
)

replace github.com/pulumi/pulumi-azurerm/sdk => ../../sdk
