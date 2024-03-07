package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestackhci.NewUpdate(ctx, "update", &azurestackhci.UpdateArgs{
			AdditionalProperties: pulumi.String("additional properties"),
			AvailabilityType:     pulumi.String("Local"),
			ClusterName:          pulumi.String("testcluster"),
			Description:          pulumi.String("AzS Update 4.2203.2.32"),
			DisplayName:          pulumi.String("AzS Update - 4.2203.2.32"),
			InstalledDate:        pulumi.String("2022-04-06T14:08:18.254Z"),
			NotifyMessage:        pulumi.String("Brief message with instructions for updates of AvailabilityType Notify"),
			PackagePath:          pulumi.String("\\\\SU1FileServer\\SU1_Infrastructure_2\\Updates\\Packages\\Microsoft4.2203.2.32"),
			PackageSizeInMb:      pulumi.Float64(18858),
			PackageType:          pulumi.String("Infrastructure"),
			Prerequisites: azurestackhci.UpdatePrerequisiteArray{
				&azurestackhci.UpdatePrerequisiteArgs{
					PackageName: pulumi.String("update package name"),
					UpdateType:  pulumi.String("update type"),
					Version:     pulumi.String("prerequisite version"),
				},
			},
			ProgressPercentage: pulumi.Float64(0),
			Publisher:          pulumi.String("Microsoft"),
			ReleaseLink:        pulumi.String("https://docs.microsoft.com/azure-stack/operator/release-notes?view=azs-2203"),
			ResourceGroupName:  pulumi.String("testrg"),
			State:              pulumi.String("Installed"),
			UpdateName:         pulumi.String("Microsoft4.2203.2.32"),
			Version:            pulumi.String("4.2203.2.32"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
