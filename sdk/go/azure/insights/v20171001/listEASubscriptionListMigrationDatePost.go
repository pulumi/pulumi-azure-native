


package v20171001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEASubscriptionListMigrationDatePost(ctx *pulumi.Context, args *ListEASubscriptionListMigrationDatePostArgs, opts ...pulumi.InvokeOption) (*ListEASubscriptionListMigrationDatePostResult, error) {
	var rv ListEASubscriptionListMigrationDatePostResult
	err := ctx.Invoke("azure-native:insights/v20171001:listEASubscriptionListMigrationDatePost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEASubscriptionListMigrationDatePostArgs struct {
}


type ListEASubscriptionListMigrationDatePostResult struct {
	IsGrandFatherableSubscription *bool   `pulumi:"isGrandFatherableSubscription"`
	OptedInDate                   *string `pulumi:"optedInDate"`
}
