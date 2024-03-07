package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/logic/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := logic.NewIntegrationServiceEnvironmentManagedApi(ctx, "integrationServiceEnvironmentManagedApi", &logic.IntegrationServiceEnvironmentManagedApiArgs{
			ApiName:                           pulumi.String("servicebus"),
			IntegrationServiceEnvironmentName: pulumi.String("testIntegrationServiceEnvironment"),
			Location:                          pulumi.String("brazilsouth"),
			ResourceGroup:                     pulumi.String("testResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
