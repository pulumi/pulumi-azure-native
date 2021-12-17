


package v20170401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementLockByScope(ctx *pulumi.Context, args *LookupManagementLockByScopeArgs, opts ...pulumi.InvokeOption) (*LookupManagementLockByScopeResult, error) {
	var rv LookupManagementLockByScopeResult
	err := ctx.Invoke("azure-native:authorization/v20170401:getManagementLockByScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementLockByScopeArgs struct {
	LockName string `pulumi:"lockName"`
	Scope    string `pulumi:"scope"`
}


type LookupManagementLockByScopeResult struct {
	Id     string                        `pulumi:"id"`
	Level  string                        `pulumi:"level"`
	Name   string                        `pulumi:"name"`
	Notes  *string                       `pulumi:"notes"`
	Owners []ManagementLockOwnerResponse `pulumi:"owners"`
	Type   string                        `pulumi:"type"`
}
