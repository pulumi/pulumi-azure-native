package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := machinelearningservices.NewComponentVersion(ctx, "componentVersion", &machinelearningservices.ComponentVersionArgs{
ComponentVersionProperties: &machinelearningservices.ComponentVersionTypeArgs{
ComponentSpec: pulumi.Any{
8ced901bD826477dBfef329da9672513: nil,
},
Description: pulumi.String("string"),
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
