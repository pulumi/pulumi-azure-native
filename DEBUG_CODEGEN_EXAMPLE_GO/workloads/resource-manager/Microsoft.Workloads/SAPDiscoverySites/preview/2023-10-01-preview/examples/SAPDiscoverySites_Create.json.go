package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workloads.NewSapDiscoverySite(ctx, "sapDiscoverySite", &workloads.SapDiscoverySiteArgs{
			Location:             pulumi.String("eastus"),
			MasterSiteId:         pulumi.String("MasterSiteIdResourceId"),
			MigrateProjectId:     pulumi.String("MigrateProjectId"),
			ResourceGroupName:    pulumi.String("test-rg"),
			SapDiscoverySiteName: pulumi.String("SampleSite"),
			Tags: pulumi.StringMap{
				"property1": pulumi.String("value1"),
				"property2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
