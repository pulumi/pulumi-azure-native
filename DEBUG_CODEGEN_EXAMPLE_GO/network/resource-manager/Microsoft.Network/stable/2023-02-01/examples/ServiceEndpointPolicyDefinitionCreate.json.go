package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewServiceEndpointPolicyDefinition(ctx, "serviceEndpointPolicyDefinition", &network.ServiceEndpointPolicyDefinitionArgs{
			Description:                         pulumi.String("Storage Service EndpointPolicy Definition"),
			ResourceGroupName:                   pulumi.String("rg1"),
			Service:                             pulumi.String("Microsoft.Storage"),
			ServiceEndpointPolicyDefinitionName: pulumi.String("testDefinition"),
			ServiceEndpointPolicyName:           pulumi.String("testPolicy"),
			ServiceResources: pulumi.StringArray{
				pulumi.String("/subscriptions/subid1"),
				pulumi.String("/subscriptions/subid1/resourceGroups/storageRg"),
				pulumi.String("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
