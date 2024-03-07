package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := app.NewDaprComponentResiliencyPolicy(ctx, "daprComponentResiliencyPolicy", &app.DaprComponentResiliencyPolicyArgs{
ComponentName: pulumi.String("mydaprcomponent"),
EnvironmentName: pulumi.String("myenvironment"),
InboundPolicy: app.DaprComponentResiliencyPolicyConfigurationResponse{
HttpRetryPolicy: interface{}{
MaxRetries: pulumi.Int(15),
RetryBackOff: &app.DaprComponentResiliencyPolicyHttpRetryBackOffConfigurationArgs{
InitialDelayInMilliseconds: pulumi.Int(2000),
MaxIntervalInMilliseconds: pulumi.Int(5500),
},
},
TimeoutPolicy: &app.DaprComponentResiliencyPolicyTimeoutPolicyConfigurationArgs{
ResponseTimeoutInSeconds: pulumi.Int(30),
},
},
Name: pulumi.String("myresiliencypolicy"),
OutboundPolicy: app.DaprComponentResiliencyPolicyConfigurationResponse{
HttpRetryPolicy: interface{}{
MaxRetries: pulumi.Int(5),
RetryBackOff: &app.DaprComponentResiliencyPolicyHttpRetryBackOffConfigurationArgs{
InitialDelayInMilliseconds: pulumi.Int(100),
MaxIntervalInMilliseconds: pulumi.Int(30000),
},
},
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
