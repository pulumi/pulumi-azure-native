package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewAppServicePlan(ctx, "appServicePlan", &web.AppServicePlanArgs{
			Kind:              pulumi.String("app"),
			Location:          pulumi.String("East US"),
			Name:              pulumi.String("testsf6141"),
			ResourceGroupName: pulumi.String("testrg123"),
			Sku: &web.SkuDescriptionArgs{
				Capacity: pulumi.Int(1),
				Family:   pulumi.String("P"),
				Name:     pulumi.String("P1"),
				Size:     pulumi.String("P1"),
				Tier:     pulumi.String("Premium"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
