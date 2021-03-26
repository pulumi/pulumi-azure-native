module github.com/pulumi/pulumi-azure-native/examples

go 1.16

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/pulumi/pulumi/pkg/v3 v3.0.0-20210324220902-b543e235f01d
	github.com/pulumi/pulumi/sdk/v3 v3.0.0-20210324220902-b543e235f01d // indirect
	github.com/stretchr/testify v1.6.1
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
)
