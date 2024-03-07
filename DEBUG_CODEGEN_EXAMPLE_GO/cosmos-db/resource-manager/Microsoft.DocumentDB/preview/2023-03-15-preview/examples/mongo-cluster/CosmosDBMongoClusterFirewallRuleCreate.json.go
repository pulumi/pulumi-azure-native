package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewMongoClusterFirewallRule(ctx, "mongoClusterFirewallRule", &documentdb.MongoClusterFirewallRuleArgs{
			EndIpAddress:      pulumi.String("255.255.255.255"),
			FirewallRuleName:  pulumi.String("rule1"),
			MongoClusterName:  pulumi.String("myMongoCluster"),
			ResourceGroupName: pulumi.String("TestGroup"),
			StartIpAddress:    pulumi.String("0.0.0.0"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
