package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := app.NewDaprComponentResiliencyPolicy(ctx, "daprComponentResiliencyPolicy", &app.DaprComponentResiliencyPolicyArgs{
			ComponentName:   pulumi.String("mydaprcomponent"),
			EnvironmentName: pulumi.String("myenvironment"),
			InboundPolicy: &app.DaprComponentResiliencyPolicyConfigurationArgs{
				HttpRetryPolicy: &app.DaprComponentResiliencyPolicyHttpRetryPolicyConfigurationArgs{
					MaxRetries: pulumi.Int(5),
					RetryBackOff: &app.DaprComponentResiliencyPolicyHttpRetryBackOffConfigurationArgs{
						InitialDelayInMilliseconds: pulumi.Int(2000),
						MaxIntervalInMilliseconds:  pulumi.Int(5500),
					},
				},
			},
			Name: pulumi.String("myresiliencypolicy"),
			OutboundPolicy: &app.DaprComponentResiliencyPolicyConfigurationArgs{
				TimeoutPolicy: &app.DaprComponentResiliencyPolicyTimeoutPolicyConfigurationArgs{
					ResponseTimeoutInSeconds: pulumi.Int(12),
				},
			},
			ResourceGroupName: pulumi.String("examplerg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
