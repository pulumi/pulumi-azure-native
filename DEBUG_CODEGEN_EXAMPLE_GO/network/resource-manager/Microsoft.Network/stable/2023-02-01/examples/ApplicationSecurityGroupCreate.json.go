package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewApplicationSecurityGroup(ctx, "applicationSecurityGroup", &network.ApplicationSecurityGroupArgs{
			ApplicationSecurityGroupName: pulumi.String("test-asg"),
			Location:                     pulumi.String("westus"),
			ResourceGroupName:            pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
