package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewGroupUser(ctx, "groupUser", &apimanagement.GroupUserArgs{
			GroupId:           pulumi.String("tempgroup"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			UserId:            pulumi.String("59307d350af58404d8a26300"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
