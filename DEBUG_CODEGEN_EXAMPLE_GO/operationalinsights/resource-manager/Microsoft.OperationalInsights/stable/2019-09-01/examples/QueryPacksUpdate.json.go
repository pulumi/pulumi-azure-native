package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/operationalinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := operationalinsights.NewQueryPack(ctx, "queryPack", &operationalinsights.QueryPackArgs{
			Location:          pulumi.String("South Central US"),
			QueryPackName:     pulumi.String("my-querypack"),
			ResourceGroupName: pulumi.String("my-resource-group"),
			Tags: pulumi.StringMap{
				"Tag1": pulumi.String("Value1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
