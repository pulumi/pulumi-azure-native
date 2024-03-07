package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsdataprocessor/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotoperationsdataprocessor.NewPipeline(ctx, "pipeline", &iotoperationsdataprocessor.PipelineArgs{
			Description: pulumi.String("vayzklhg"),
			Enabled:     pulumi.Bool(true),
			ExtendedLocation: &iotoperationsdataprocessor.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/e0aaa3df-e9a4-456a-9824-3c3b5c438110/resourceGroups/IoTOperationsDataProcessor-rg/providers/Microsoft.ExtendedLocation/customLocations/dev-space"),
				Type: pulumi.String("CustomLocation"),
			},
			Input: &iotoperationsdataprocessor.PipelineInputTypeArgs{
				Next: pulumi.StringArray{
					pulumi.String("umnuwjk"),
				},
				Type: pulumi.String("xrnubjkvzajxjzb"),
			},
			InstanceName:      pulumi.String("056k5pl8t7761-2--ej25u2l28ttb-22mh79-75-2ch-t8"),
			Location:          pulumi.String("westus"),
			PipelineName:      pulumi.String("j8-8--3"),
			ResourceGroupName: pulumi.String("rgopenapi"),
			Stages: iotoperationsdataprocessor.PipelineStageMap{
				"stageId": &iotoperationsdataprocessor.PipelineStageArgs{
					Next: pulumi.StringArray{
						pulumi.String("gxqgqh"),
					},
					Type: pulumi.String("cxqgblrzqniowabexbztempdpkuib"),
				},
			},
			Tags: nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
