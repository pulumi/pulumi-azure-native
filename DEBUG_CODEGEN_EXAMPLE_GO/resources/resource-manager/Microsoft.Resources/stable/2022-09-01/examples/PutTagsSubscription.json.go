package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resources.NewTagAtScope(ctx, "tagAtScope", &resources.TagAtScopeArgs{
			Properties: &resources.TagsArgs{
				Tags: pulumi.StringMap{
					"tagKey1": pulumi.String("tag-value-1"),
					"tagKey2": pulumi.String("tag-value-2"),
				},
			},
			Scope: pulumi.String("subscriptions/00000000-0000-0000-0000-000000000000"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
