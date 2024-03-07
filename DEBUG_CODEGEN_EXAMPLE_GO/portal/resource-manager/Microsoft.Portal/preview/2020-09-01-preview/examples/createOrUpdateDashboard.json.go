package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/portal/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := portal.NewDashboard(ctx, "dashboard", &portal.DashboardArgs{
			DashboardName: pulumi.String("testDashboard"),
			Lenses: portal.DashboardLensArray{
				&portal.DashboardLensArgs{
					Order: pulumi.Int(1),
					Parts: []portal.DashboardPartsArgs{
						{
							Position: {
								ColSpan: pulumi.Int(3),
								RowSpan: pulumi.Int(4),
								X:       pulumi.Int(1),
								Y:       pulumi.Int(2),
							},
						},
						{
							Position: {
								ColSpan: pulumi.Int(6),
								RowSpan: pulumi.Int(6),
								X:       pulumi.Int(5),
								Y:       pulumi.Int(5),
							},
						},
					},
				},
				&portal.DashboardLensArgs{
					Order: pulumi.Int(2),
					Parts: portal.DashboardPartsArray{},
				},
			},
			Location: pulumi.String("eastus"),
			Metadata: pulumi.Map{
				"metadata": pulumi.Any{
					ColSpan: 2,
					RowSpan: 1,
					X:       4,
					Y:       3,
				},
			},
			ResourceGroupName: pulumi.String("testRG"),
			Tags: pulumi.StringMap{
				"aKey":       pulumi.String("aValue"),
				"anotherKey": pulumi.String("anotherValue"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
