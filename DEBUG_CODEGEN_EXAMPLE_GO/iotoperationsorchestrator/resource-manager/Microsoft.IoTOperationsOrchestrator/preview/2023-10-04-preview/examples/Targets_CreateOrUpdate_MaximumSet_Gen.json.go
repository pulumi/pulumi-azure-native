package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsorchestrator/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotoperationsorchestrator.NewTarget(ctx, "target", &iotoperationsorchestrator.TargetArgs{
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
			Location: pulumi.String("pjjkifnrwvzcyohz"),
			Name:     pulumi.String("7---s--1-hl-fl-3f0-wfy34e08-4"),
			ReconciliationPolicy: &iotoperationsorchestrator.ReconciliationPolicyArgs{
				Interval: pulumi.String("wyrrzydmfgzymvzbppscxyfobku"),
				Type:     pulumi.String("periodic"),
			},
			ResourceGroupName: pulumi.String("rgopenapi"),
			Scope:             pulumi.String("lm"),
			Tags:              nil,
			Topologies: iotoperationsorchestrator.TopologiesPropertiesArray{
				&iotoperationsorchestrator.TopologiesPropertiesArgs{
					Bindings: iotoperationsorchestrator.BindingPropertiesArray{
						&iotoperationsorchestrator.BindingPropertiesArgs{
							Config:   nil,
							Provider: pulumi.String("qpwesjlyyggcbehwigbobqum"),
							Role:     pulumi.String("role"),
						},
					},
				},
			},
			Version: pulumi.String("prbigsnjltnzqliu"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
