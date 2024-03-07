package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewServiceEndpointPolicy(ctx, "serviceEndpointPolicy", &network.ServiceEndpointPolicyArgs{
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceEndpointPolicyDefinitions: []network.ServiceEndpointPolicyDefinitionTypeArgs{
				{
					Description: pulumi.String("Storage Service EndpointPolicy Definition"),
					Name:        pulumi.String("StorageServiceEndpointPolicyDefinition"),
					Service:     pulumi.String("Microsoft.Storage"),
					ServiceResources: pulumi.StringArray{
						pulumi.String("/subscriptions/subid1"),
						pulumi.String("/subscriptions/subid1/resourceGroups/storageRg"),
						pulumi.String("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount"),
					},
				},
			},
			ServiceEndpointPolicyName: pulumi.String("testPolicy"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
