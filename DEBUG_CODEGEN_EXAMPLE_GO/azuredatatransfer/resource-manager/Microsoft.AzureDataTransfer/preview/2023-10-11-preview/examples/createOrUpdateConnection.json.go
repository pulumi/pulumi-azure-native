package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azuredatatransfer/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azuredatatransfer.NewConnection(ctx, "connection", &azuredatatransfer.ConnectionArgs{
			ConnectionName: pulumi.String("testConnection"),
			Location:       pulumi.String("East US"),
			Properties: &azuredatatransfer.ConnectionPropertiesArgs{
				Justification: pulumi.String("justification"),
				Pipeline:      pulumi.String("testdc"),
				RequirementId: pulumi.String("id"),
			},
			ResourceGroupName: pulumi.String("testRG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
