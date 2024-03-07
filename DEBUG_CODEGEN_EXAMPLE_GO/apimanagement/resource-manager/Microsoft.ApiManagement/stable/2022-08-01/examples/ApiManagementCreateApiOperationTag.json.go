package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewTagByOperation(ctx, "tagByOperation", &apimanagement.TagByOperationArgs{
			ApiId:             pulumi.String("5931a75ae4bbd512a88c680b"),
			OperationId:       pulumi.String("5931a75ae4bbd512a88c680a"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			TagId:             pulumi.String("tagId1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
