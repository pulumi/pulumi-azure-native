module github.com/pulumi/pulumi-azure-native/examples

go 1.16

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/pulumi/pulumi/pkg/v2 v2.21.2
	github.com/pulumi/pulumi/sdk/v2 v2.21.2 // indirect
	github.com/stretchr/testify v1.6.1
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
)
