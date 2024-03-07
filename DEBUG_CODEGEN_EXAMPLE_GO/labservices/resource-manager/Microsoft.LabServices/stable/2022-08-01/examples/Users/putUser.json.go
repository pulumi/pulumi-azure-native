package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/labservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := labservices.NewUser(ctx, "user", &labservices.UserArgs{
			AdditionalUsageQuota: pulumi.String("PT10H"),
			Email:                pulumi.String("testuser@contoso.com"),
			LabName:              pulumi.String("testlab"),
			ResourceGroupName:    pulumi.String("testrg123"),
			UserName:             pulumi.String("testuser"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
