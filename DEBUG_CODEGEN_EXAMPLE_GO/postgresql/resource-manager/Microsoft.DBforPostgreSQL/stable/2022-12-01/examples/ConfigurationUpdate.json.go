package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbforpostgresql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbforpostgresql.NewConfiguration(ctx, "configuration", &dbforpostgresql.ConfigurationArgs{
			ConfigurationName: pulumi.String("event_scheduler"),
			ResourceGroupName: pulumi.String("testrg"),
			ServerName:        pulumi.String("testserver"),
			Source:            pulumi.String("user-override"),
			Value:             pulumi.String("on"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
