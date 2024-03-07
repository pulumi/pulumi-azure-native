package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewDscConfiguration(ctx, "dscConfiguration", &automation.DscConfigurationArgs{
			AutomationAccountName: pulumi.String("myAutomationAccount18"),
			ConfigurationName:     pulumi.String("SetupServer"),
			Description:           pulumi.String("sample configuration"),
			Location:              pulumi.String("East US 2"),
			Name:                  pulumi.String("SetupServer"),
			ResourceGroupName:     pulumi.String("rg"),
			Source: automation.ContentSourceResponse{
				Hash: &automation.ContentHashArgs{
					Algorithm: pulumi.String("sha256"),
					Value:     pulumi.String("A9E5DB56BA21513F61E0B3868816FDC6D4DF5131F5617D7FF0D769674BD5072F"),
				},
				Type: pulumi.String("embeddedContent"),
				Value: pulumi.String(`Configuration SetupServer {
    Node localhost {
                               WindowsFeature IIS {
                               Name = "Web-Server";
            Ensure = "Present"
        }
    }
}`),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
