package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := web.NewContainerApp(ctx, "containerApp", &web.ContainerAppArgs{
Configuration: web.ConfigurationResponse{
Ingress: &web.IngressArgs{
External: pulumi.Bool(true),
TargetPort: pulumi.Int(3000),
},
},
Kind: pulumi.String("containerApp"),
KubeEnvironmentId: pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/kubeEnvironments/demokube"),
Location: pulumi.String("East US"),
Name: pulumi.String("testcontainerApp0"),
ResourceGroupName: pulumi.String("rg"),
Template: web.TemplateResponse{
Containers: web.ContainerArray{
&web.ContainerArgs{
Image: pulumi.String("repo/testcontainerApp0:v1"),
Name: pulumi.String("testcontainerApp0"),
},
},
Dapr: &web.DaprArgs{
AppPort: pulumi.Int(3000),
Enabled: pulumi.Bool(true),
},
Scale: interface{}{
MaxReplicas: pulumi.Int(5),
MinReplicas: pulumi.Int(1),
Rules: web.ScaleRuleArray{
interface{}{
Custom: &web.CustomScaleRuleArgs{
Metadata: pulumi.StringMap{
"concurrentRequests": pulumi.String("50"),
},
Type: pulumi.String("http"),
},
Name: pulumi.String("httpscalingrule"),
},
},
},
},
})
if err != nil {
return err
}
return nil
})
}
