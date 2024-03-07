package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/timeseriesinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := timeseriesinsights.NewReferenceDataSet(ctx, "referenceDataSet", &timeseriesinsights.ReferenceDataSetArgs{
			EnvironmentName: pulumi.String("env1"),
			KeyProperties: timeseriesinsights.ReferenceDataSetKeyPropertyArray{
				&timeseriesinsights.ReferenceDataSetKeyPropertyArgs{
					Name: pulumi.String("DeviceId1"),
					Type: pulumi.String("String"),
				},
				&timeseriesinsights.ReferenceDataSetKeyPropertyArgs{
					Name: pulumi.String("DeviceFloor"),
					Type: pulumi.String("Double"),
				},
			},
			Location:             pulumi.String("West US"),
			ReferenceDataSetName: pulumi.String("rds1"),
			ResourceGroupName:    pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
