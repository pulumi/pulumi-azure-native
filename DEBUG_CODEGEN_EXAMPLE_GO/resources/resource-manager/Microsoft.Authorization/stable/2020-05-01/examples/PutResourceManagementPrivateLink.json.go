package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewResourceManagementPrivateLink(ctx, "resourceManagementPrivateLink", &authorization.ResourceManagementPrivateLinkArgs{
			Location:          pulumi.String("eastus"),
			ResourceGroupName: pulumi.String("my-resource-group"),
			RmplName:          pulumi.String("my-rmplName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
