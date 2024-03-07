package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewManagedServerDnsAlias(ctx, "managedServerDnsAlias", &sql.ManagedServerDnsAliasArgs{
			DnsAliasName:        pulumi.String("dns-alias-mi"),
			ManagedInstanceName: pulumi.String("dns-mi"),
			ResourceGroupName:   pulumi.String("Default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
