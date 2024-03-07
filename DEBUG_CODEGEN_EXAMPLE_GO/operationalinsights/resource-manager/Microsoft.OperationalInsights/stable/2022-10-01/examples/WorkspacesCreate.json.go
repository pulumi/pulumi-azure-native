package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/operationalinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := operationalinsights.NewWorkspace(ctx, "workspace", &operationalinsights.WorkspaceArgs{
			Location:          pulumi.String("australiasoutheast"),
			ResourceGroupName: pulumi.String("oiautorest6685"),
			RetentionInDays:   pulumi.Int(30),
			Sku: &operationalinsights.WorkspaceSkuArgs{
				Name: pulumi.String("PerGB2018"),
			},
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("val1"),
			},
			WorkspaceName: pulumi.String("oiautorest6685"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
