package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cognitiveservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cognitiveservices.NewRaiBlocklist(ctx, "raiBlocklist", &cognitiveservices.RaiBlocklistArgs{
			AccountName: pulumi.String("accountName"),
			Properties: &cognitiveservices.RaiBlocklistPropertiesArgs{
				Description: pulumi.String("Basic blocklist description"),
			},
			RaiBlocklistName:  pulumi.String("raiBlocklistName"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
