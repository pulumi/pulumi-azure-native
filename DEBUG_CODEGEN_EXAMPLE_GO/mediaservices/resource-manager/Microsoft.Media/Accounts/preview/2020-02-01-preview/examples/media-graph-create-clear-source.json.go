package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewMediaGraph(ctx, "mediaGraph", &media.MediaGraphArgs{
			AccountName:       pulumi.String("contosomedia"),
			Description:       pulumi.String("updated description"),
			MediaGraphName:    pulumi.String("SampleMediaGraph"),
			ResourceGroupName: pulumi.String("contoso"),
			Sinks: []media.MediaGraphAssetSinkArgs{
				{
					AssetName: pulumi.String("SampleAsset"),
					Inputs: pulumi.StringArray{
						pulumi.String("rtspSource"),
					},
					Name:      pulumi.String("AssetSink"),
					OdataType: pulumi.String("#Microsoft.Media.MediaGraphAssetSink"),
				},
			},
			Sources: []media.MediaGraphRtspSourceArgs{
				{
					Endpoint: {
						Credentials: {
							OdataType: "#Microsoft.Media.MediaGraphUsernamePasswordCredentials",
							Password:  "examplepassword",
							Username:  "exampleusername",
						},
						OdataType: "#Microsoft.Media.MediaGraphClearEndpoint",
						Url:       "rtsp://contoso.com:554/stream1",
					},
					Name:      pulumi.String("rtspSource"),
					OdataType: pulumi.String("#Microsoft.Media.MediaGraphRtspSource"),
					Transport: pulumi.String("Http"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
