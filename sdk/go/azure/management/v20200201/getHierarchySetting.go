


package v20200201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHierarchySetting(ctx *pulumi.Context, args *LookupHierarchySettingArgs, opts ...pulumi.InvokeOption) (*LookupHierarchySettingResult, error) {
	var rv LookupHierarchySettingResult
	err := ctx.Invoke("azure-native:management/v20200201:getHierarchySetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHierarchySettingArgs struct {
	GroupId string `pulumi:"groupId"`
}


type LookupHierarchySettingResult struct {
	DefaultManagementGroup               *string `pulumi:"defaultManagementGroup"`
	Id                                   string  `pulumi:"id"`
	Name                                 string  `pulumi:"name"`
	RequireAuthorizationForGroupCreation *bool   `pulumi:"requireAuthorizationForGroupCreation"`
	TenantId                             *string `pulumi:"tenantId"`
	Type                                 string  `pulumi:"type"`
}
