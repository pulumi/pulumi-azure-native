package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/desktopvirtualization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := desktopvirtualization.NewApplication(ctx, "application", &desktopvirtualization.ApplicationArgs{
			ApplicationGroupName: pulumi.String("applicationGroup1"),
			ApplicationName:      pulumi.String("application1"),
			CommandLineArguments: pulumi.String("arguments"),
			CommandLineSetting:   pulumi.String("Allow"),
			Description:          pulumi.String("des1"),
			FilePath:             pulumi.String("path"),
			FriendlyName:         pulumi.String("friendly"),
			IconIndex:            pulumi.Int(1),
			IconPath:             pulumi.String("icon"),
			ResourceGroupName:    pulumi.String("resourceGroup1"),
			ShowInPortal:         pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
