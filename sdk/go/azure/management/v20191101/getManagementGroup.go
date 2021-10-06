


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementGroup(ctx *pulumi.Context, args *LookupManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupManagementGroupResult, error) {
	var rv LookupManagementGroupResult
	err := ctx.Invoke("azure-native:management/v20191101:getManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementGroupArgs struct {
	Expand  *string `pulumi:"expand"`
	Filter  *string `pulumi:"filter"`
	GroupId string  `pulumi:"groupId"`
	Recurse *bool   `pulumi:"recurse"`
}


type LookupManagementGroupResult struct {
	Children    []ManagementGroupChildInfoResponse   `pulumi:"children"`
	Details     *ManagementGroupDetailsResponse      `pulumi:"details"`
	DisplayName *string                              `pulumi:"displayName"`
	Id          string                               `pulumi:"id"`
	Name        string                               `pulumi:"name"`
	Path        []ManagementGroupPathElementResponse `pulumi:"path"`
	Roles       []string                             `pulumi:"roles"`
	TenantId    *string                              `pulumi:"tenantId"`
	Type        string                               `pulumi:"type"`
}
