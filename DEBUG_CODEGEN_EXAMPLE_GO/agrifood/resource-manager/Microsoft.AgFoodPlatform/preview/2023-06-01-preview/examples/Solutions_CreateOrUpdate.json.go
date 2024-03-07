package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/agfoodplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := agfoodplatform.NewSolution(ctx, "solution", &agfoodplatform.SolutionArgs{
			DataManagerForAgricultureResourceName: pulumi.String("examples-farmbeatsResourceName"),
			ResourceGroupName:                     pulumi.String("examples-rg"),
			SolutionId:                            pulumi.String("abc.partner"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
