


package authorization

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementLockAtResourceLevel(ctx *pulumi.Context, args *LookupManagementLockAtResourceLevelArgs, opts ...pulumi.InvokeOption) (*LookupManagementLockAtResourceLevelResult, error) {
	var rv LookupManagementLockAtResourceLevelResult
	err := ctx.Invoke("azure-native:authorization:getManagementLockAtResourceLevel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementLockAtResourceLevelArgs struct {
	LockName                  string `pulumi:"lockName"`
	ParentResourcePath        string `pulumi:"parentResourcePath"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	ResourceName              string `pulumi:"resourceName"`
	ResourceProviderNamespace string `pulumi:"resourceProviderNamespace"`
	ResourceType              string `pulumi:"resourceType"`
}


type LookupManagementLockAtResourceLevelResult struct {
	Id     string                        `pulumi:"id"`
	Level  string                        `pulumi:"level"`
	Name   string                        `pulumi:"name"`
	Notes  *string                       `pulumi:"notes"`
	Owners []ManagementLockOwnerResponse `pulumi:"owners"`
	Type   string                        `pulumi:"type"`
}
