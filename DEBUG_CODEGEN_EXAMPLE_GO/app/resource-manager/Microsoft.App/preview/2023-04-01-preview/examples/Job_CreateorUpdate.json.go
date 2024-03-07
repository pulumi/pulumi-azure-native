package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := app.NewJob(ctx, "job", &app.JobArgs{
Configuration: app.JobConfigurationResponse{
ManualTriggerConfig: &app.JobConfigurationManualTriggerConfigArgs{
Parallelism: pulumi.Int(4),
ReplicaCompletionCount: pulumi.Int(1),
},
ReplicaRetryLimit: pulumi.Int(10),
ReplicaTimeout: pulumi.Int(10),
TriggerType: pulumi.String("Manual"),
},
EnvironmentId: pulumi.String("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
JobName: pulumi.String("testcontainerappsjob0"),
Location: pulumi.String("East US"),
ResourceGroupName: pulumi.String("rg"),
Template: app.JobTemplateResponse{
Containers: app.ContainerArray{
interface{}{
Image: pulumi.String("repo/testcontainerappsjob0:v1"),
Name: pulumi.String("testcontainerappsjob0"),
Probes: app.ContainerAppProbeArray{
interface{}{
HttpGet: interface{}{
HttpHeaders: app.ContainerAppProbeHttpHeadersArray{
&app.ContainerAppProbeHttpHeadersArgs{
Name: pulumi.String("Custom-Header"),
Value: pulumi.String("Awesome"),
},
},
Path: pulumi.String("/health"),
Port: pulumi.Int(8080),
},
InitialDelaySeconds: pulumi.Int(5),
PeriodSeconds: pulumi.Int(3),
Type: pulumi.String("Liveness"),
},
},
},
},
InitContainers: app.InitContainerArray{
interface{}{
Args: pulumi.StringArray{
pulumi.String("-c"),
pulumi.String("while true; do echo hello; sleep 10;done"),
},
Command: pulumi.StringArray{
pulumi.String("/bin/sh"),
},
Image: pulumi.String("repo/testcontainerappsjob0:v4"),
Name: pulumi.String("testinitcontainerAppsJob0"),
Resources: &app.ContainerResourcesArgs{
Cpu: pulumi.Float64(0.5),
Memory: pulumi.String("1Gi"),
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
