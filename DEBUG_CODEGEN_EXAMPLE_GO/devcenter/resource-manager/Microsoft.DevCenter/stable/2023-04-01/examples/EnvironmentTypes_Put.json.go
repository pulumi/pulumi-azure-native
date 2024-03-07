package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devcenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devcenter.NewEnvironmentType(ctx, "environmentType", &devcenter.EnvironmentTypeArgs{
			DevCenterName:       pulumi.String("Contoso"),
			EnvironmentTypeName: pulumi.String("DevTest"),
			ResourceGroupName:   pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"Owner": pulumi.String("superuser"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
