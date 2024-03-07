package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datafactory/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datafactory.NewTrigger(ctx, "trigger", &datafactory.TriggerArgs{
			FactoryName: pulumi.String("exampleFactoryName"),
			Properties: datafactory.ScheduleTrigger{
				Description: "Example description",
				Pipelines: []datafactory.TriggerPipelineReference{
					{
						Parameters: {
							"OutputBlobNameList": []string{
								"exampleoutput.csv",
							},
						},
						PipelineReference: {
							ReferenceName: "examplePipeline",
							Type:          "PipelineReference",
						},
					},
				},
				Recurrence: datafactory.ScheduleTriggerRecurrence{
					EndTime:   "2018-06-16T00:55:14.905167Z",
					Frequency: "Minute",
					Interval:  4,
					StartTime: "2018-06-16T00:39:14.905167Z",
					TimeZone:  "UTC",
				},
				Type: "ScheduleTrigger",
			},
			ResourceGroupName: pulumi.String("exampleResourceGroup"),
			TriggerName:       pulumi.String("exampleTrigger"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
