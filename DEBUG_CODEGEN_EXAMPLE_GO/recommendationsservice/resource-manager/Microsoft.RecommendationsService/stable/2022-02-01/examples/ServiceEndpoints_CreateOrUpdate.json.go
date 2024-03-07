package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recommendationsservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recommendationsservice.NewServiceEndpoint(ctx, "serviceEndpoint", &recommendationsservice.ServiceEndpointArgs{
			AccountName: pulumi.String("sampleAccount"),
			Location:    pulumi.String("West US"),
			Properties: &recommendationsservice.ServiceEndpointResourcePropertiesArgs{
				PreAllocatedCapacity: pulumi.Int(100),
			},
			ResourceGroupName:   pulumi.String("rg"),
			ServiceEndpointName: pulumi.String("s1"),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Prod"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
