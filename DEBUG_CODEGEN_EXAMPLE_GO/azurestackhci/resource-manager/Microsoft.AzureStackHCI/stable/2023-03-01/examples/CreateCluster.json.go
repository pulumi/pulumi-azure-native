package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestackhci.NewCluster(ctx, "cluster", &azurestackhci.ClusterArgs{
			AadClientId:             pulumi.String("24a6e53d-04e5-44d2-b7cc-1b732a847dfc"),
			AadTenantId:             pulumi.String("7e589cc1-a8b6-4dff-91bd-5ec0fa18db94"),
			CloudManagementEndpoint: pulumi.String("https://98294836-31be-4668-aeae-698667faf99b.waconazure.com"),
			ClusterName:             pulumi.String("myCluster"),
			Location:                pulumi.String("East US"),
			ResourceGroupName:       pulumi.String("test-rg"),
			Type:                    pulumi.String("SystemAssigned"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
