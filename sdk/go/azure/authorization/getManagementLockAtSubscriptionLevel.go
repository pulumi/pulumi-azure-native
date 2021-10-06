


package authorization

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementLockAtSubscriptionLevel(ctx *pulumi.Context, args *LookupManagementLockAtSubscriptionLevelArgs, opts ...pulumi.InvokeOption) (*LookupManagementLockAtSubscriptionLevelResult, error) {
	var rv LookupManagementLockAtSubscriptionLevelResult
	err := ctx.Invoke("azure-native:authorization:getManagementLockAtSubscriptionLevel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementLockAtSubscriptionLevelArgs struct {
	LockName string `pulumi:"lockName"`
}


type LookupManagementLockAtSubscriptionLevelResult struct {
	Id     string                        `pulumi:"id"`
	Level  string                        `pulumi:"level"`
	Name   string                        `pulumi:"name"`
	Notes  *string                       `pulumi:"notes"`
	Owners []ManagementLockOwnerResponse `pulumi:"owners"`
	Type   string                        `pulumi:"type"`
}
