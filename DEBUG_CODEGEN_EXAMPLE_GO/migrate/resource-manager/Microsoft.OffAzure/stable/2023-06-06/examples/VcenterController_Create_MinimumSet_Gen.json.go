package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewVcenterController(ctx, "vcenterController", &offazure.VcenterControllerArgs{
			ResourceGroupName: pulumi.String("rgmigrate"),
			SiteName:          pulumi.String("6-qSc554IYc4U08"),
			VcenterName:       pulumi.String("R5I8Xj8--zsS6JYI-0FNhe"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
