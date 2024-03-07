package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewModule(ctx, "module", &automation.ModuleArgs{
			AutomationAccountName: pulumi.String("myAutomationAccount33"),
			ContentLink: &automation.ContentLinkArgs{
				ContentHash: &automation.ContentHashArgs{
					Algorithm: pulumi.String("sha265"),
					Value:     pulumi.String("07E108A962B81DD9C9BAA89BB47C0F6EE52B29E83758B07795E408D258B2B87A"),
				},
				Uri:     pulumi.String("https://teststorage.blob.core.windows.net/dsccomposite/OmsCompositeResources.zip"),
				Version: pulumi.String("1.0.0.0"),
			},
			ModuleName:        pulumi.String("OmsCompositeResources"),
			ResourceGroupName: pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
