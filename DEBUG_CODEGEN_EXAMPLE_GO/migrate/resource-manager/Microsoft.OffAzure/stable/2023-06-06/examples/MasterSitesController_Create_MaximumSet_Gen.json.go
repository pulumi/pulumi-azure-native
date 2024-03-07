package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewMasterSitesController(ctx, "masterSitesController", &offazure.MasterSitesControllerArgs{
			AllowMultipleSites:          pulumi.Bool(true),
			CustomerStorageAccountArmId: pulumi.String("cdxrihtiskkn"),
			Location:                    pulumi.String("plyak"),
			PublicNetworkAccess:         pulumi.String("NotSpecified"),
			ResourceGroupName:           pulumi.String("rgmigrate"),
			SiteName:                    pulumi.String("74c1J1m56557t52H-75"),
			Sites: pulumi.StringArray{
				pulumi.String("zxupfq"),
			},
			Tags: pulumi.StringMap{
				"key7125": pulumi.String("jbhnzfuxjovyamasouyfhhzhevagvw"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
