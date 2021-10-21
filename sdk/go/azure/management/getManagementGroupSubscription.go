


package management

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementGroupSubscription(ctx *pulumi.Context, args *LookupManagementGroupSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupManagementGroupSubscriptionResult, error) {
	var rv LookupManagementGroupSubscriptionResult
	err := ctx.Invoke("azure-native:management:getManagementGroupSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementGroupSubscriptionArgs struct {
	GroupId        string  `pulumi:"groupId"`
	SubscriptionId *string `pulumi:"subscriptionId"`
}


type LookupManagementGroupSubscriptionResult struct {
	DisplayName *string                            `pulumi:"displayName"`
	Id          string                             `pulumi:"id"`
	Name        string                             `pulumi:"name"`
	Parent      *DescendantParentGroupInfoResponse `pulumi:"parent"`
	State       *string                            `pulumi:"state"`
	Tenant      *string                            `pulumi:"tenant"`
	Type        string                             `pulumi:"type"`
}
