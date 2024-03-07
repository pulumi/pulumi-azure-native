package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/portal/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := portal.NewConsoleWithLocation(ctx, "consoleWithLocation", &portal.ConsoleWithLocationArgs{
			ConsoleName: pulumi.String("default"),
			Location:    pulumi.String("eastus"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
