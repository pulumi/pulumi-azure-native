package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventhub.NewApplicationGroup(ctx, "applicationGroup", &eventhub.ApplicationGroupArgs{
			ApplicationGroupName:     pulumi.String("appGroup1"),
			ClientAppGroupIdentifier: pulumi.String("SASKeyName=KeyName"),
			IsEnabled:                pulumi.Bool(true),
			NamespaceName:            pulumi.String("contoso-ua-test-eh-system-1"),
			Policies: []eventhub.ThrottlingPolicyArgs{
				{
					MetricId:           pulumi.String("IncomingMessages"),
					Name:               pulumi.String("ThrottlingPolicy1"),
					RateLimitThreshold: pulumi.Float64(7912),
					Type:               pulumi.String("ThrottlingPolicy"),
				},
				{
					MetricId:           pulumi.String("IncomingBytes"),
					Name:               pulumi.String("ThrottlingPolicy2"),
					RateLimitThreshold: pulumi.Float64(3951729),
					Type:               pulumi.String("ThrottlingPolicy"),
				},
				{
					MetricId:           pulumi.String("OutgoingBytes"),
					Name:               pulumi.String("ThrottlingPolicy3"),
					RateLimitThreshold: pulumi.Float64(245175),
					Type:               pulumi.String("ThrottlingPolicy"),
				},
			},
			ResourceGroupName: pulumi.String("contosotest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
