package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/testbase/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := testbase.NewTestBaseAccount(ctx, "testBaseAccount", &testbase.TestBaseAccountArgs{
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("contoso-rg1"),
			Sku: &testbase.TestBaseAccountSKUArgs{
				Name: pulumi.String("S0"),
				Tier: pulumi.String("Standard"),
			},
			TestBaseAccountName: pulumi.String("contoso-testBaseAccount1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
