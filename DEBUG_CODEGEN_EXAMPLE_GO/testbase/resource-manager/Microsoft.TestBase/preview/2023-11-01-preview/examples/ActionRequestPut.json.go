package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/testbase/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := testbase.NewActionRequest(ctx, "actionRequest", &testbase.ActionRequestArgs{
			ActionRequestName:   pulumi.String("167184141414254"),
			ResourceGroupName:   pulumi.String("contoso-rg"),
			TestBaseAccountName: pulumi.String("contoso-testBaseAccount"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
