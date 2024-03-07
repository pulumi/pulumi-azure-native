package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/windowsiot/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := windowsiot.NewService(ctx, "service", &windowsiot.ServiceArgs{
			AdminDomainName:   pulumi.String("d.e.f"),
			BillingDomainName: pulumi.String("a.b.c"),
			DeviceName:        pulumi.String("service4445"),
			Location:          pulumi.String("East US"),
			Notes:             pulumi.String("blah"),
			Quantity:          pulumi.Float64(1000000),
			ResourceGroupName: pulumi.String("res9101"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
