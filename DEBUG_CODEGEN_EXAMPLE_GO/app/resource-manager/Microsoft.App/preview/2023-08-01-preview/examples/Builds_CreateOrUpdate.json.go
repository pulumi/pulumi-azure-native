package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/app/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := app.NewBuild(ctx, "build", &app.BuildArgs{
BuildName: pulumi.String("testBuild-123456789az"),
BuilderName: pulumi.String("testBuilder"),
Configuration: app.BuildConfigurationResponse{
BaseOs: pulumi.String("DebianBullseye"),
EnvironmentVariables: app.EnvironmentVariableArray{
&app.EnvironmentVariableArgs{
Name: pulumi.String("foo1"),
Value: pulumi.String("bar1"),
},
&app.EnvironmentVariableArgs{
Name: pulumi.String("foo2"),
Value: pulumi.String("bar2"),
},
},
Platform: pulumi.String("dotnetcore"),
PlatformVersion: pulumi.String("7.0"),
PreBuildSteps: app.PreBuildStepArray{
interface{}{
Description: pulumi.String("First pre build step."),
HttpGet: &app.HttpGetArgs{
FileName: pulumi.String("output.txt"),
Headers: pulumi.StringArray{
pulumi.String("foo"),
pulumi.String("bar"),
},
Url: pulumi.String("https://microsoft.com"),
},
Scripts: pulumi.StringArray{
pulumi.String("echo 'hello'"),
pulumi.String("echo 'world'"),
},
},
interface{}{
Description: pulumi.String("Second pre build step."),
HttpGet: &app.HttpGetArgs{
FileName: pulumi.String("output.txt"),
Headers: pulumi.StringArray{
pulumi.String("foo"),
},
Url: pulumi.String("https://microsoft.com"),
},
Scripts: pulumi.StringArray{
pulumi.String("echo 'hello'"),
pulumi.String("echo 'again'"),
},
},
},
},
DestinationContainerRegistry: &app.ContainerRegistryWithCustomImageArgs{
Image: pulumi.String("test.azurecr.io/repo:tag"),
Server: pulumi.String("test.azurecr.io"),
},
ResourceGroupName: pulumi.String("rg"),
})
if err != nil {
return err
}
return nil
})
}
