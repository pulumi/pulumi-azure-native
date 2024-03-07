package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotoperationsorchestrator/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotoperationsorchestrator.NewInstance(ctx, "instance", &iotoperationsorchestrator.InstanceArgs{
			ExtendedLocation: &iotoperationsorchestrator.ExtendedLocationArgs{
				Name: pulumi.String("bjjhfqsplgzdlbdlddleetyg"),
				Type: pulumi.String("sosibrbmmrfbbyp"),
			},
			Location: pulumi.String("uzehbktba"),
			Name:     pulumi.String("j6-r05-43h--55-q97-9ig--6w44a-2--3g-k53-13263ncl6-2q-h0-q5"),
			ReconciliationPolicy: &iotoperationsorchestrator.ReconciliationPolicyArgs{
				Interval: pulumi.String("wyrrzydmfgzymvzbppscxyfobku"),
				Type:     pulumi.String("periodic"),
			},
			ResourceGroupName: pulumi.String("rgopenapi"),
			Scope:             pulumi.String("rkargnjeljnivwjly"),
			Solution:          pulumi.String("p"),
			Tags:              nil,
			Target: &iotoperationsorchestrator.TargetSelectorPropertiesArgs{
				Name: pulumi.String("mqxcv"),
			},
			Version: pulumi.String("nf"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
