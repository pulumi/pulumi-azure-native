package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/costmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := costmanagement.NewTagInheritanceSetting(ctx, "tagInheritanceSetting", &costmanagement.TagInheritanceSettingArgs{
			Kind: pulumi.String("taginheritance"),
			Properties: &costmanagement.TagInheritancePropertiesArgs{
				PreferContainerTags: pulumi.Bool(false),
			},
			Scope: pulumi.String("subscriptions/00000000-0000-0000-0000-000000000000"),
			Type:  pulumi.String("taginheritance"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
