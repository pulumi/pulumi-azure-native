package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/elasticsan/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticsan.NewVolumeSnapshot(ctx, "volumeSnapshot", &elasticsan.VolumeSnapshotArgs{
			CreationData: &elasticsan.SnapshotCreationDataArgs{
				SourceId: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"),
			},
			ElasticSanName:    pulumi.String("elasticsanname"),
			ResourceGroupName: pulumi.String("resourcegroupname"),
			SnapshotName:      pulumi.String("snapshotname"),
			VolumeGroupName:   pulumi.String("volumegroupname"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
