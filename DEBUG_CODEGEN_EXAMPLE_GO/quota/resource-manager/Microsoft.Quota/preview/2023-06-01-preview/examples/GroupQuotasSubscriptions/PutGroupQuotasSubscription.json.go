package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/quota/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quota.NewGroupQuotaSubscription(ctx, "groupQuotaSubscription", &quota.GroupQuotaSubscriptionArgs{
			GroupQuotaName: pulumi.String("groupquota1"),
			MgId:           pulumi.String("E7EC67B3-7657-4966-BFFC-41EFD36BAA09"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
