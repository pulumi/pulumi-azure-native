package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewScriptExecution(ctx, "scriptExecution", &avs.ScriptExecutionArgs{
			HiddenParameters: pulumi.Array{
				avs.ScriptSecureStringExecutionParameter{
					Name:        "Password",
					SecureValue: "PlaceholderPassword",
					Type:        "SecureValue",
				},
			},
			Parameters: pulumi.Array{
				avs.ScriptStringExecutionParameter{
					Name:  "DomainName",
					Type:  "Value",
					Value: "placeholderDomain.local",
				},
				avs.ScriptStringExecutionParameter{
					Name:  "BaseUserDN",
					Type:  "Value",
					Value: "DC=placeholder, DC=placeholder",
				},
			},
			PrivateCloudName:    pulumi.String("cloud1"),
			ResourceGroupName:   pulumi.String("group1"),
			Retention:           pulumi.String("P0Y0M60DT0H60M60S"),
			ScriptCmdletId:      pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptPackages/AVS.PowerCommands@1.0.0/scriptCmdlets/New-SsoExternalIdentitySource"),
			ScriptExecutionName: pulumi.String("addSsoServer"),
			Timeout:             pulumi.String("P0Y0M0DT0H60M60S"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
