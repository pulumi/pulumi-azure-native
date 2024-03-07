package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/videoanalyzer/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := videoanalyzer.NewEdgeModule(ctx, "edgeModule", &videoanalyzer.EdgeModuleArgs{
			AccountName:       pulumi.String("testaccount2"),
			EdgeModuleName:    pulumi.String("edgeModule1"),
			ResourceGroupName: pulumi.String("testrg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
