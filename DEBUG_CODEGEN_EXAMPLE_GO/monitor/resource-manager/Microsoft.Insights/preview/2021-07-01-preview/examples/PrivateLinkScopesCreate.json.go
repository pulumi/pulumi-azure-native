package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewPrivateLinkScope(ctx, "privateLinkScope", &insights.PrivateLinkScopeArgs{
			AccessModeSettings: &insights.AccessModeSettingsArgs{
				Exclusions:          insights.AccessModeSettingsExclusionArray{},
				IngestionAccessMode: pulumi.String("Open"),
				QueryAccessMode:     pulumi.String("Open"),
			},
			Location:          pulumi.String("Global"),
			ResourceGroupName: pulumi.String("my-resource-group"),
			ScopeName:         pulumi.String("my-privatelinkscope"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
