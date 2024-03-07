package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/integrationspaces/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := integrationspaces.NewInfrastructureResource(ctx, "infrastructureResource", &integrationspaces.InfrastructureResourceArgs{
			InfrastructureResourceName: pulumi.String("InfrastructureResource1"),
			ResourceGroupName:          pulumi.String("testrg"),
			ResourceId:                 pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.ApiManagement/service/APIM1"),
			ResourceType:               pulumi.String("Microsoft.ApiManagement/service"),
			SpaceName:                  pulumi.String("Space1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
