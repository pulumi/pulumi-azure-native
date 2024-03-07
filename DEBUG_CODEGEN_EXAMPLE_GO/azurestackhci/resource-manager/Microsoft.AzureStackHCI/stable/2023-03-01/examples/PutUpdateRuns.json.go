package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestackhci.NewUpdateRun(ctx, "updateRun", &azurestackhci.UpdateRunArgs{
			ClusterName:        pulumi.String("testcluster"),
			Description:        pulumi.String("Update Azure Stack."),
			EndTimeUtc:         pulumi.String("2022-04-06T13:58:42.969006+00:00"),
			ErrorMessage:       pulumi.String(""),
			LastUpdatedTimeUtc: pulumi.String("2022-04-06T13:58:42.969006+00:00"),
			Name:               pulumi.String("Unnamed step"),
			ResourceGroupName:  pulumi.String("testrg"),
			StartTimeUtc:       pulumi.String("2022-04-06T01:36:33.3876751+00:00"),
			Status:             pulumi.String("Success"),
			Steps: azurestackhci.StepArray{
				&azurestackhci.StepArgs{
					Description:        pulumi.String("Prepare for SSU update"),
					EndTimeUtc:         pulumi.String("2022-04-06T01:37:16.8728314+00:00"),
					ErrorMessage:       pulumi.String(""),
					LastUpdatedTimeUtc: pulumi.String("2022-04-06T01:37:16.8728314+00:00"),
					Name:               pulumi.String("PreUpdate Cloud"),
					StartTimeUtc:       pulumi.String("2022-04-06T01:36:33.3876751+00:00"),
					Status:             pulumi.String("Success"),
					Steps:              azurestackhci.StepArray{},
				},
			},
			UpdateName:    pulumi.String("Microsoft4.2203.2.32"),
			UpdateRunName: pulumi.String("23b779ba-0d52-4a80-8571-45ca74664ec3"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
