package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewComponent(ctx, "component", &insights.ComponentArgs{
			Kind:              pulumi.String("web"),
			Location:          pulumi.String("South Central US"),
			ResourceGroupName: pulumi.String("my-resource-group"),
			ResourceName:      pulumi.String("my-component"),
			Tags: pulumi.StringMap{
				"ApplicationGatewayType": pulumi.String("Internal-Only"),
				"BillingEntity":          pulumi.String("Self"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
