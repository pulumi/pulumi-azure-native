package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cognitiveservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cognitiveservices.NewRaiBlocklistItem(ctx, "raiBlocklistItem", &cognitiveservices.RaiBlocklistItemArgs{
			AccountName: pulumi.String("accountName"),
			Properties: &cognitiveservices.RaiBlocklistItemPropertiesArgs{
				IsRegex: pulumi.Bool(false),
				Pattern: pulumi.String("Pattern To Block"),
			},
			RaiBlocklistItemName: pulumi.String("raiBlocklistItemName"),
			RaiBlocklistName:     pulumi.String("raiBlocklistName"),
			ResourceGroupName:    pulumi.String("resourceGroupName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
