package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewAssessment(ctx, "assessment", &security.AssessmentArgs{
			AssessmentName: pulumi.String("8bb8be0a-6010-4789-812f-e4d661c4ed0e"),
			ResourceDetails: security.AzureResourceDetails{
				Source: "Azure",
			},
			ResourceId: pulumi.String("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2"),
			Status: &security.AssessmentStatusArgs{
				Code: pulumi.String("Healthy"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
