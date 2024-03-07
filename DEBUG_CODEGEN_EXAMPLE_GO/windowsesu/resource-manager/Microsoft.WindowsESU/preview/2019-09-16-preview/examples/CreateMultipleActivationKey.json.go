package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/windowsesu/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := windowsesu.NewMultipleActivationKey(ctx, "multipleActivationKey", &windowsesu.MultipleActivationKeyArgs{
			AgreementNumber:           pulumi.String("1a2b45ag"),
			InstalledServerNumber:     pulumi.Int(100),
			IsEligible:                pulumi.Bool(true),
			Location:                  pulumi.String("East US"),
			MultipleActivationKeyName: pulumi.String("server08-key-2019"),
			OsType:                    pulumi.String("WindowsServer2008"),
			ResourceGroupName:         pulumi.String("testgr1"),
			SupportType:               pulumi.String("SupplementalServicing"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
