package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewServiceEndpointPolicy(ctx, "serviceEndpointPolicy", &network.ServiceEndpointPolicyArgs{
			Location:                  pulumi.String("westus"),
			ResourceGroupName:         pulumi.String("rg1"),
			ServiceEndpointPolicyName: pulumi.String("testPolicy"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
