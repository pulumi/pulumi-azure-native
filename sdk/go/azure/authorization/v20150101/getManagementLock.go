


package v20150101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementLock(ctx *pulumi.Context, args *LookupManagementLockArgs, opts ...pulumi.InvokeOption) (*LookupManagementLockResult, error) {
	var rv LookupManagementLockResult
	err := ctx.Invoke("azure-native:authorization/v20150101:getManagementLock", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementLockArgs struct {
	LockName string `pulumi:"lockName"`
}


type LookupManagementLockResult struct {
	Id    string  `pulumi:"id"`
	Level *string `pulumi:"level"`
	Name  *string `pulumi:"name"`
	Notes *string `pulumi:"notes"`
	Type  string  `pulumi:"type"`
}
