module scaleset

go 1.14

require (
	github.com/pulumi/pulumi-azure/sdk/v3 v3.0.0
	github.com/pulumi/pulumi-azurerm/sdk v0.1.0
	github.com/pulumi/pulumi-random/sdk/v2 v2.2.0
	github.com/pulumi/pulumi/sdk/v2 v2.0.0
)

replace github.com/pulumi/pulumi-azurerm/sdk => ../../sdk
