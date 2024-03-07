package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/quota/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quota.NewGroupQuota(ctx, "groupQuota", &quota.GroupQuotaArgs{
			GroupQuotaName: pulumi.String("groupquota1"),
			MgId:           pulumi.String("E7EC67B3-7657-4966-BFFC-41EFD36BAA09"),
			Properties: &quota.GroupQuotasEntityBaseArgs{
				AdditionalAttributes: &quota.AdditionalAttributesArgs{
					Environment: pulumi.Any("Production"),
					GroupId: &quota.GroupingIdArgs{
						GroupingIdType: pulumi.String("ServiceTreeId"),
						Value:          pulumi.String("yourServiceTreeIdHere"),
					},
				},
				DisplayName: pulumi.String("GroupQuota1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
