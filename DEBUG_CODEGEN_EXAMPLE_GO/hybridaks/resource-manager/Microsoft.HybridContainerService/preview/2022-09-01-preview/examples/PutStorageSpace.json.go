package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridcontainerservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridcontainerservice.NewStorageSpaceRetrieve(ctx, "storageSpaceRetrieve", &hybridcontainerservice.StorageSpaceRetrieveArgs{
			ExtendedLocation: &hybridcontainerservice.StorageSpacesExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"),
				Type: pulumi.String("CustomLocation"),
			},
			Location: pulumi.String("westus"),
			Properties: hybridcontainerservice.StorageSpacesPropertiesResponse{
				HciStorageProfile: &hybridcontainerservice.StorageSpacesPropertiesHciStorageProfileArgs{
					MocGroup:            pulumi.String("target-group"),
					MocLocation:         pulumi.String("MocLocation"),
					MocStorageContainer: pulumi.String("WssdStorageContainer"),
				},
			},
			ResourceGroupName: pulumi.String("test-arcappliance-resgrp"),
			StorageSpacesName: pulumi.String("test-storage"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
