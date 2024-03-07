package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := machinelearningservices.NewEnvironmentVersion(ctx, "environmentVersion", &machinelearningservices.EnvironmentVersionArgs{
EnvironmentVersionProperties: machinelearningservices.EnvironmentVersionResponse{
Build: &machinelearningservices.BuildContextArgs{
ContextUri: pulumi.String("https://storage-account.blob.core.windows.net/azureml/DockerBuildContext/95ddede6b9b8c4e90472db3acd0a8d28/"),
DockerfilePath: pulumi.String("prod/Dockerfile"),
},
CondaFile: pulumi.String("string"),
Description: pulumi.String("string"),
Image: pulumi.String("docker.io/tensorflow/serving:latest"),
InferenceConfig: interface{}{
LivenessRoute: &machinelearningservices.RouteArgs{
Path: pulumi.String("string"),
Port: pulumi.Int(1),
},
ReadinessRoute: &machinelearningservices.RouteArgs{
Path: pulumi.String("string"),
Port: pulumi.Int(1),
},
ScoringRoute: &machinelearningservices.RouteArgs{
Path: pulumi.String("string"),
Port: pulumi.Int(1),
},
},
IsAnonymous: pulumi.Bool(false),
Properties: pulumi.StringMap{
"string": pulumi.String("string"),
},
Tags: pulumi.StringMap{
"string": pulumi.String("string"),
},
},
Name: pulumi.String("string"),
ResourceGroupName: pulumi.String("test-rg"),
Version: pulumi.String("string"),
WorkspaceName: pulumi.String("my-aml-workspace"),
})
if err != nil {
return err
}
return nil
})
}
