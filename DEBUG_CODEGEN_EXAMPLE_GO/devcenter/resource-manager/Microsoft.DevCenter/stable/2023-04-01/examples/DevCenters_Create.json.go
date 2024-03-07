package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devcenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devcenter.NewDevCenter(ctx, "devCenter", &devcenter.DevCenterArgs{
			DevCenterName:     pulumi.String("Contoso"),
			Location:          pulumi.String("centralus"),
			ResourceGroupName: pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"CostCode": pulumi.String("12345"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
