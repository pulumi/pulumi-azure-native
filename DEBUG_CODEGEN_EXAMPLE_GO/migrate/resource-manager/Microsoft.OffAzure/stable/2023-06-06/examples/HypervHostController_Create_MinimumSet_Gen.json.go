package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewHypervHostController(ctx, "hypervHostController", &offazure.HypervHostControllerArgs{
			HostName:          pulumi.String("IIfm6suXoKL"),
			ResourceGroupName: pulumi.String("rgmigrate"),
			SiteName:          pulumi.String("-25mH3"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
