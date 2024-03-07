package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestackhci.NewUpdateSummary(ctx, "updateSummary", &azurestackhci.UpdateSummaryArgs{
			ClusterName:       pulumi.String("testcluster"),
			CurrentVersion:    pulumi.String("4.2203.2.32"),
			HardwareModel:     pulumi.String("PowerEdge R730xd"),
			LastChecked:       pulumi.String("2022-04-07T18:04:07Z"),
			LastUpdated:       pulumi.String("2022-04-06T14:08:18.254Z"),
			OemFamily:         pulumi.String("DellEMC"),
			ResourceGroupName: pulumi.String("testrg"),
			State:             pulumi.String("AppliedSuccessfully"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
