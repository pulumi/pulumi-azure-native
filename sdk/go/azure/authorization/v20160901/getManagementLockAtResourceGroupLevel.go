


package v20160901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementLockAtResourceGroupLevel(ctx *pulumi.Context, args *LookupManagementLockAtResourceGroupLevelArgs, opts ...pulumi.InvokeOption) (*LookupManagementLockAtResourceGroupLevelResult, error) {
	var rv LookupManagementLockAtResourceGroupLevelResult
	err := ctx.Invoke("azure-native:authorization/v20160901:getManagementLockAtResourceGroupLevel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementLockAtResourceGroupLevelArgs struct {
	LockName          string `pulumi:"lockName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupManagementLockAtResourceGroupLevelResult struct {
	Id     string                        `pulumi:"id"`
	Level  string                        `pulumi:"level"`
	Name   string                        `pulumi:"name"`
	Notes  *string                       `pulumi:"notes"`
	Owners []ManagementLockOwnerResponse `pulumi:"owners"`
	Type   string                        `pulumi:"type"`
}
