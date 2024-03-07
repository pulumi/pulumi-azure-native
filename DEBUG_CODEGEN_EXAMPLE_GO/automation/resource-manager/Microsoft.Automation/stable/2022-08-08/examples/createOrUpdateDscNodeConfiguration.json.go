package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewDscNodeConfiguration(ctx, "dscNodeConfiguration", &automation.DscNodeConfigurationArgs{
			AutomationAccountName: pulumi.String("myAutomationAccount20"),
			Configuration: &automation.DscConfigurationAssociationPropertyArgs{
				Name: pulumi.String("configName"),
			},
			IncrementNodeConfigurationBuild: pulumi.Bool(true),
			Name:                            pulumi.String("configName.nodeConfigName"),
			NodeConfigurationName:           pulumi.String("configName.nodeConfigName"),
			ResourceGroupName:               pulumi.String("rg"),
			Source: &automation.ContentSourceArgs{
				Hash: &automation.ContentHashArgs{
					Algorithm: pulumi.String("sha256"),
					Value:     pulumi.String("6DE256A57F01BFA29B88696D5E77A383D6E61484C7686E8DB955FA10ACE9FFE5"),
				},
				Type: pulumi.String("embeddedContent"),
				Value: pulumi.String(`
instance of MSFT_RoleResource as $MSFT_RoleResource1ref
{
ResourceID = "[WindowsFeature]IIS";
 Ensure = "Present";
 SourceInfo = "::3::32::WindowsFeature";
 Name = "Web-Server";
 ModuleName = "PsDesiredStateConfiguration";

ModuleVersion = "1.0";
 ConfigurationName = "configName";
};
instance of OMI_ConfigurationDocument

                    {
 Version="2.0.0";
 
                        MinimumCompatibleVersion = "1.0.0";
 
                        CompatibleVersionAdditionalProperties= {"Omi_BaseResource:ConfigurationName"};
 
                        Author="weijiel";
 
                        GenerationDate="03/30/2017 13:40:25";
 
                        GenerationHost="TEST-BACKEND";
 
                        Name="configName";

                    };
`),
				Version: pulumi.String("1.0"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
