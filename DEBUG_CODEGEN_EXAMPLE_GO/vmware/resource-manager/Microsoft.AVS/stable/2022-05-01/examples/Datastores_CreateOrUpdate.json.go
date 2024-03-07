package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewDatastore(ctx, "datastore", &avs.DatastoreArgs{
			ClusterName:   pulumi.String("cluster1"),
			DatastoreName: pulumi.String("datastore1"),
			NetAppVolume: &avs.NetAppVolumeArgs{
				Id: pulumi.String("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/ResourceGroup1/providers/Microsoft.NetApp/netAppAccounts/NetAppAccount1/capacityPools/CapacityPool1/volumes/NFSVol1"),
			},
			PrivateCloudName:  pulumi.String("cloud1"),
			ResourceGroupName: pulumi.String("group1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
