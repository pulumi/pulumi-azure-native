package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewPolicy(ctx, "policy", &devtestlab.PolicyArgs{
			Description:       pulumi.String("{policyDescription}"),
			EvaluatorType:     pulumi.String("{policyEvaluatorType}"),
			FactData:          pulumi.String("{policyFactData}"),
			FactName:          pulumi.String("{policyFactName}"),
			LabName:           pulumi.String("{labName}"),
			Location:          pulumi.String("{location}"),
			Name:              pulumi.String("{policyName}"),
			PolicySetName:     pulumi.String("{policySetName}"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			Status:            pulumi.String("{policyStatus}"),
			Tags: pulumi.StringMap{
				"tagName1": pulumi.String("tagValue1"),
			},
			Threshold: pulumi.String("{policyThreshold}"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
