package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/easm/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := easm.NewLabelByWorkspace(ctx, "labelByWorkspace", &easm.LabelByWorkspaceArgs{
			LabelName:         pulumi.String("ThisisaLabel"),
			ResourceGroupName: pulumi.String("dummyrg"),
			WorkspaceName:     pulumi.String("ThisisaWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
