package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/testbase/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := testbase.NewDraftPackage(ctx, "draftPackage", &testbase.DraftPackageArgs{
			AppFileName:         pulumi.String("TestBaseM365DigitalClock.msi"),
			ApplicationName:     pulumi.String("contoso-package"),
			DraftPackageName:    pulumi.String("61d99543-14ff-47ae-bf03-8a8b8445502e"),
			ResourceGroupName:   pulumi.String("contoso-rg1"),
			SourceType:          pulumi.String("Native"),
			TestBaseAccountName: pulumi.String("contoso-testBaseAccount1"),
			UseSample:           pulumi.Bool(false),
			Version:             pulumi.String("1.0"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
