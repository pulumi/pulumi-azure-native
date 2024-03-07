package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/timeseriesinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := timeseriesinsights.NewGen2Environment(ctx, "gen2Environment", &timeseriesinsights.Gen2EnvironmentArgs{
			EnvironmentName:   pulumi.String("env1"),
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
