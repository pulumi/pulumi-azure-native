package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cache/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cache.NewAccessPolicyAssignment(ctx, "accessPolicyAssignment", &cache.AccessPolicyAssignmentArgs{
			AccessPolicyAssignmentName: pulumi.String("accessPolicyAssignmentName1"),
			AccessPolicyName:           pulumi.String("accessPolicy1"),
			CacheName:                  pulumi.String("cache1"),
			ObjectId:                   pulumi.String("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
			ObjectIdAlias:              pulumi.String("TestAADAppRedis"),
			ResourceGroupName:          pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
