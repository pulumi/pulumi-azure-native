package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/videoanalyzer/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := videoanalyzer.NewPipelineTopology(ctx, "pipelineTopology", &videoanalyzer.PipelineTopologyArgs{
			AccountName: pulumi.String("testaccount2"),
			Description: pulumi.String("Pipeline Topology 1 Description"),
			Kind:        pulumi.String("Live"),
			Parameters: []videoanalyzer.ParameterDeclarationArgs{
				{
					Default:     pulumi.String("rtsp://microsoft.com/video.mp4"),
					Description: pulumi.String("rtsp source url parameter"),
					Name:        pulumi.String("rtspUrlParameter"),
					Type:        pulumi.String("String"),
				},
				{
					Default:     pulumi.String("password"),
					Description: pulumi.String("rtsp source password parameter"),
					Name:        pulumi.String("rtspPasswordParameter"),
					Type:        pulumi.String("SecretString"),
				},
			},
			PipelineTopologyName: pulumi.String("pipelineTopology1"),
			ResourceGroupName:    pulumi.String("testrg"),
			Sinks: []videoanalyzer.VideoSinkArgs{
				{
					Inputs: videoanalyzer.NodeInputArray{
						{
							NodeName: pulumi.String("rtspSource"),
						},
					},
					Name: pulumi.String("videoSink"),
					Type: pulumi.String("#Microsoft.VideoAnalyzer.VideoSink"),
					VideoCreationProperties: {
						Description:   pulumi.String("Parking lot south entrance"),
						SegmentLength: pulumi.String("PT30S"),
						Title:         pulumi.String("Parking Lot (Camera 1)"),
					},
					VideoName: pulumi.String("camera001"),
					VideoPublishingOptions: {
						DisableArchive:        pulumi.String("false"),
						DisableRtspPublishing: pulumi.String("true"),
					},
				},
			},
			Sku: &videoanalyzer.SkuArgs{
				Name: pulumi.String("Live_S1"),
			},
			Sources: pulumi.Array{
				videoanalyzer.RtspSource{
					Endpoint: videoanalyzer.UnsecuredEndpoint{
						Credentials: videoanalyzer.UsernamePasswordCredentials{
							Password: "${rtspPasswordParameter}",
							Type:     "#Microsoft.VideoAnalyzer.UsernamePasswordCredentials",
							Username: "username",
						},
						Type: "#Microsoft.VideoAnalyzer.UnsecuredEndpoint",
						Url:  "${rtspUrlParameter}",
					},
					Name:      "rtspSource",
					Transport: "Http",
					Type:      "#Microsoft.VideoAnalyzer.RtspSource",
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
