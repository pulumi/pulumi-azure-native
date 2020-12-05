module gosimple

go 1.14

require (
	github.com/pulumi/pulumi-azure-nextgen/sdk v0.2.1
	github.com/pulumi/pulumi-random/sdk/v2 v2.3.0
	github.com/pulumi/pulumi/sdk/v2 v2.6.0
)

replace github.com/pulumi/pulumi-azure-nextgen/sdk => ../../sdk
