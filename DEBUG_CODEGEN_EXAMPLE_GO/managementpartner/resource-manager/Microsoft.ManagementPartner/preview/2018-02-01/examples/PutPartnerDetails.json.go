package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managementpartner/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managementpartner.NewPartner(ctx, "partner", &managementpartner.PartnerArgs{
			PartnerId: pulumi.String("123456"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
