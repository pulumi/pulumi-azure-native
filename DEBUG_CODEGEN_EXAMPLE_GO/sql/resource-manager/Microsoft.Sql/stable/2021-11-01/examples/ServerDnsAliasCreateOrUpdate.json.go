package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewServerDnsAlias(ctx, "serverDnsAlias", &sql.ServerDnsAliasArgs{
			DnsAliasName:      pulumi.String("dns-alias-name-1"),
			ResourceGroupName: pulumi.String("Default"),
			ServerName:        pulumi.String("dns-alias-server"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
