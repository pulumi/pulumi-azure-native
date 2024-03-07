package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicefabricmesh/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicefabricmesh.NewVolume(ctx, "volume", &servicefabricmesh.VolumeArgs{
			AzureFileParameters: &servicefabricmesh.VolumeProviderParametersAzureFileArgs{
				AccountKey:  pulumi.String("provide-account-key-here"),
				AccountName: pulumi.String("sbzdemoaccount"),
				ShareName:   pulumi.String("sharel"),
			},
			Description:        pulumi.String("Service Fabric Mesh sample volume."),
			Location:           pulumi.String("EastUS"),
			Provider:           pulumi.String("SFAzureFile"),
			ResourceGroupName:  pulumi.String("sbz_demo"),
			Tags:               nil,
			VolumeResourceName: pulumi.String("sampleVolume"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
