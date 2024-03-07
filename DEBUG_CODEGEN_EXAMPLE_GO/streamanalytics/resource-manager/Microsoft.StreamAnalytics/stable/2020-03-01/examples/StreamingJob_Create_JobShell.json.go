package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/streamanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := streamanalytics.NewStreamingJob(ctx, "streamingJob", &streamanalytics.StreamingJobArgs{
			CompatibilityLevel:                 pulumi.String("1.0"),
			DataLocale:                         pulumi.String("en-US"),
			EventsLateArrivalMaxDelayInSeconds: pulumi.Int(16),
			EventsOutOfOrderMaxDelayInSeconds:  pulumi.Int(5),
			EventsOutOfOrderPolicy:             pulumi.String("Drop"),
			Functions:                          streamanalytics.FunctionTypeArray{},
			Inputs:                             streamanalytics.InputTypeArray{},
			JobName:                            pulumi.String("sj59"),
			Location:                           pulumi.String("West US"),
			OutputErrorPolicy:                  pulumi.String("Drop"),
			Outputs:                            streamanalytics.OutputTypeArray{},
			ResourceGroupName:                  pulumi.String("sjrg6936"),
			Sku: &streamanalytics.SkuArgs{
				Name: pulumi.String("Standard"),
			},
			Tags: pulumi.StringMap{
				"key1":      pulumi.String("value1"),
				"key3":      pulumi.String("value3"),
				"randomKey": pulumi.String("randomValue"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
