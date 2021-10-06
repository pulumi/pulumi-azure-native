


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkGroup(ctx *pulumi.Context, args *LookupNetworkGroupArgs, opts ...pulumi.InvokeOption) (*LookupNetworkGroupResult, error) {
	var rv LookupNetworkGroupResult
	err := ctx.Invoke("azure-native:network/v20210201preview:getNetworkGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkGroupArgs struct {
	NetworkGroupName   string `pulumi:"networkGroupName"`
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupNetworkGroupResult struct {
	ConditionalMembership *string                    `pulumi:"conditionalMembership"`
	Description           *string                    `pulumi:"description"`
	DisplayName           *string                    `pulumi:"displayName"`
	Etag                  string                     `pulumi:"etag"`
	GroupMembers          []GroupMembersItemResponse `pulumi:"groupMembers"`
	Id                    string                     `pulumi:"id"`
	MemberType            *string                    `pulumi:"memberType"`
	Name                  string                     `pulumi:"name"`
	ProvisioningState     string                     `pulumi:"provisioningState"`
	SystemData            SystemDataResponse         `pulumi:"systemData"`
	Type                  string                     `pulumi:"type"`
}
