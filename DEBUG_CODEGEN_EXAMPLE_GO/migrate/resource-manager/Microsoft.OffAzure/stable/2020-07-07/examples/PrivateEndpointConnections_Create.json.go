package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &offazure.PrivateEndpointConnectionArgs{
			PeConnectionName:  pulumi.String("privateendpt1938mastersit9007pe.4f2f2970-0bfa-45d4-9ee1-d9f79502fc6f"),
			ResourceGroupName: pulumi.String("ayagrawrg"),
			SiteName:          pulumi.String("privateendpt1938mastersite"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
