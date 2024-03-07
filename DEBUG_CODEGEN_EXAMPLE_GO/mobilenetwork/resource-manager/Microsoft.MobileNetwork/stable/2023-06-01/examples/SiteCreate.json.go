package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/mobilenetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mobilenetwork.NewSite(ctx, "site", &mobilenetwork.SiteArgs{
			Location:          pulumi.String("testLocation"),
			MobileNetworkName: pulumi.String("testMobileNetwork"),
			ResourceGroupName: pulumi.String("rg1"),
			SiteName:          pulumi.String("testSite"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
