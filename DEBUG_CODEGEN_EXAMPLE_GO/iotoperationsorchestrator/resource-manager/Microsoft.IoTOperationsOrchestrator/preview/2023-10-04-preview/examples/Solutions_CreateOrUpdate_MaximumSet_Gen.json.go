package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsorchestrator/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotoperationsorchestrator.NewSolution(ctx, "solution", &iotoperationsorchestrator.SolutionArgs{
			Components: iotoperationsorchestrator.ComponentPropertiesArray{
				&iotoperationsorchestrator.ComponentPropertiesArgs{
					Dependencies: pulumi.StringArray{
						pulumi.String("x"),
					},
					Name:       pulumi.String("yhnelpxsobdyurwvhkq"),
					Properties: nil,
					Type:       pulumi.String("wiabwsfqhhxru"),
				},
			},
			ExtendedLocation: &iotoperationsorchestrator.ExtendedLocationArgs{
				Name: pulumi.String("bjjhfqsplgzdlbdlddleetyg"),
				Type: pulumi.String("sosibrbmmrfbbyp"),
			},
			Location:          pulumi.String("svzwmojzvarczmvgfhjk"),
			Name:              pulumi.String("49-gj2-mwgt--1m611----35u"),
			ResourceGroupName: pulumi.String("rgopenapi"),
			Tags:              nil,
			Version:           pulumi.String("jwxk"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
