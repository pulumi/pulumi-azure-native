package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewIpFirewallRule(ctx, "ipFirewallRule", &synapse.IpFirewallRuleArgs{
			EndIpAddress:      pulumi.String("10.0.0.254"),
			ResourceGroupName: pulumi.String("ExampleResourceGroup"),
			RuleName:          pulumi.String("ExampleIpFirewallRule"),
			StartIpAddress:    pulumi.String("10.0.0.0"),
			WorkspaceName:     pulumi.String("ExampleWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
