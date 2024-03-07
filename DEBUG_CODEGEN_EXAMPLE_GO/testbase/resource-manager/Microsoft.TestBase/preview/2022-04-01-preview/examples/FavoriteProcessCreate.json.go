package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/testbase/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := testbase.NewFavoriteProcess(ctx, "favoriteProcess", &testbase.FavoriteProcessArgs{
			ActualProcessName:           pulumi.String("testApp&.exe"),
			FavoriteProcessResourceName: pulumi.String("testAppProcess"),
			PackageName:                 pulumi.String("contoso-package2"),
			ResourceGroupName:           pulumi.String("contoso-rg1"),
			TestBaseAccountName:         pulumi.String("contoso-testBaseAccount1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
