package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewApplicationLiveView(ctx, "applicationLiveView", &appplatform.ApplicationLiveViewArgs{
			ApplicationLiveViewName: pulumi.String("default"),
			ResourceGroupName:       pulumi.String("myResourceGroup"),
			ServiceName:             pulumi.String("myservice"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
