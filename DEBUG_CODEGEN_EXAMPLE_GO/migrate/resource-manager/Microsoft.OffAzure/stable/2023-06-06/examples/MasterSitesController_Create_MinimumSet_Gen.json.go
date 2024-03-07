package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewMasterSitesController(ctx, "masterSitesController", &offazure.MasterSitesControllerArgs{
			Location:          pulumi.String("plyak"),
			ResourceGroupName: pulumi.String("rgmigrate"),
			SiteName:          pulumi.String("-3A8SuY-jRr63J"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
