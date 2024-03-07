package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kusto/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kusto.NewSandboxCustomImage(ctx, "sandboxCustomImage", &kusto.SandboxCustomImageArgs{
			ClusterName:             pulumi.String("kustoCluster"),
			Language:                pulumi.String("Python"),
			LanguageVersion:         pulumi.String("3.10.8"),
			RequirementsFileContent: pulumi.String("Requests"),
			ResourceGroupName:       pulumi.String("kustorptest"),
			SandboxCustomImageName:  pulumi.String("customImage8"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
