package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/virtualmachineimages/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := virtualmachineimages.NewTrigger(ctx, "trigger", &virtualmachineimages.TriggerArgs{
			ImageTemplateName: pulumi.String("myImageTemplate"),
			Kind:              pulumi.String("SourceImage"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			TriggerName:       pulumi.String("source"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
