package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbformysql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbformysql.NewConfiguration(ctx, "configuration", &dbformysql.ConfigurationArgs{
			ConfigurationName: pulumi.String("event_scheduler"),
			ResourceGroupName: pulumi.String("TestGroup"),
			ServerName:        pulumi.String("testserver"),
			Source:            pulumi.String("user-override"),
			Value:             pulumi.String("off"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
