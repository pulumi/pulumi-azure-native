package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/customerinsights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := customerinsights.NewView(ctx, "view", &customerinsights.ViewArgs{
			Definition: pulumi.String("{\\\"isProfileType\\\":false,\\\"profileTypes\\\":[],\\\"widgets\\\":[],\\\"style\\\":[]}"),
			DisplayName: pulumi.StringMap{
				"en": pulumi.String("some name"),
			},
			HubName:           pulumi.String("sdkTestHub"),
			ResourceGroupName: pulumi.String("TestHubRG"),
			UserId:            pulumi.String("testUser"),
			ViewName:          pulumi.String("testView"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
