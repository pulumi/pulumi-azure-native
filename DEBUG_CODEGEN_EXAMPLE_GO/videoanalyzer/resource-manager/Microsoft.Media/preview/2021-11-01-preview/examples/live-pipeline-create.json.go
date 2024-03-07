package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/videoanalyzer/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := videoanalyzer.NewLivePipeline(ctx, "livePipeline", &videoanalyzer.LivePipelineArgs{
			AccountName:      pulumi.String("testaccount2"),
			BitrateKbps:      pulumi.Int(500),
			Description:      pulumi.String("Live Pipeline 1 Description"),
			LivePipelineName: pulumi.String("livePipeline1"),
			Parameters: []videoanalyzer.ParameterDefinitionArgs{
				{
					Name:  pulumi.String("rtspUrlParameter"),
					Value: pulumi.String("rtsp://contoso.com/stream"),
				},
			},
			ResourceGroupName: pulumi.String("testrg"),
			TopologyName:      pulumi.String("pipelinetopology1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
