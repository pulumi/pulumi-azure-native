package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewJob(ctx, "job", &media.JobArgs{
			AccountName: pulumi.String("contosomedia"),
			CorrelationData: pulumi.StringMap{
				"Key 2": pulumi.String("Value 2"),
				"key1":  pulumi.String("value1"),
			},
			Input: media.JobInputAsset{
				AssetName: "job1-InputAsset",
				OdataType: "#Microsoft.Media.JobInputAsset",
			},
			JobName: pulumi.String("job1"),
			Outputs: []media.JobOutputAssetArgs{
				{
					AssetName: pulumi.String("job1-OutputAsset"),
					OdataType: pulumi.String("#Microsoft.Media.JobOutputAsset"),
				},
			},
			ResourceGroupName: pulumi.String("contosoresources"),
			TransformName:     pulumi.String("exampleTransform"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
