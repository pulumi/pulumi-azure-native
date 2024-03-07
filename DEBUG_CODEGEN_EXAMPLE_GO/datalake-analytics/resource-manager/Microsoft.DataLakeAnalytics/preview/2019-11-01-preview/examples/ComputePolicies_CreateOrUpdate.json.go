package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datalakeanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datalakeanalytics.NewComputePolicy(ctx, "computePolicy", &datalakeanalytics.ComputePolicyArgs{
			AccountName:                  pulumi.String("contosoadla"),
			ComputePolicyName:            pulumi.String("test_policy"),
			MaxDegreeOfParallelismPerJob: pulumi.Int(10),
			MinPriorityPerJob:            pulumi.Int(30),
			ObjectId:                     pulumi.String("776b9091-8916-4638-87f7-9c989a38da98"),
			ObjectType:                   pulumi.String("User"),
			ResourceGroupName:            pulumi.String("contosorg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
