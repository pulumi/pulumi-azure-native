package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/portal/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := portal.NewConsole(ctx, "console", &portal.ConsoleArgs{
			ConsoleName: pulumi.String("default"),
			Properties: &portal.ConsoleCreatePropertiesArgs{
				OsType: pulumi.String("Linux"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
