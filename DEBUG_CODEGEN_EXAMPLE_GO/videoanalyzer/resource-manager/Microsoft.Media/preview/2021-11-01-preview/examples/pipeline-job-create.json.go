package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/videoanalyzer/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := videoanalyzer.NewPipelineJob(ctx, "pipelineJob", &videoanalyzer.PipelineJobArgs{
			AccountName: pulumi.String("testaccount2"),
			Description: pulumi.String("Pipeline Job 1 Dsecription"),
			Parameters: []videoanalyzer.ParameterDefinitionArgs{
				{
					Name:  pulumi.String("timesequences"),
					Value: pulumi.String("[[\"2020-10-05T03:30:00Z\", \"2020-10-05T04:30:00Z\"]]"),
				},
				{
					Name:  pulumi.String("videoSourceName"),
					Value: pulumi.String("camera001"),
				},
			},
			PipelineJobName:   pulumi.String("pipelineJob1"),
			ResourceGroupName: pulumi.String("testrg"),
			TopologyName:      pulumi.String("pipelinetopology1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
