module goaks

go 1.15

require (
	github.com/pulumi/pulumi-azure-nextgen/sdk v0.2.1
	github.com/pulumi/pulumi-azuread/sdk/v2 v2.4.0
	github.com/pulumi/pulumi-random/sdk/v2 v2.2.0
	github.com/pulumi/pulumi-tls/sdk/v2 v2.2.0
	github.com/pulumi/pulumi/sdk/v2 v2.6.0
)

replace github.com/pulumi/pulumi-azure-nextgen/sdk => ../../sdk
