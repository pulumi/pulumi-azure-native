package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewInternetGatewayRule(ctx, "internetGatewayRule", &managednetworkfabric.InternetGatewayRuleArgs{
			Annotation:              pulumi.String("annotationValue"),
			InternetGatewayRuleName: pulumi.String("example-internetGatewayRule"),
			Location:                pulumi.String("eastus"),
			ResourceGroupName:       pulumi.String("example-rg"),
			RuleProperties: &managednetworkfabric.RulePropertiesArgs{
				Action: pulumi.String("Allow"),
				AddressList: pulumi.StringArray{
					pulumi.String("10.10.10.10"),
				},
			},
			Tags: pulumi.StringMap{
				"keyID": pulumi.String("keyValue"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
