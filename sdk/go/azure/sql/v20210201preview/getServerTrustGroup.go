


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerTrustGroup(ctx *pulumi.Context, args *LookupServerTrustGroupArgs, opts ...pulumi.InvokeOption) (*LookupServerTrustGroupResult, error) {
	var rv LookupServerTrustGroupResult
	err := ctx.Invoke("azure-native:sql/v20210201preview:getServerTrustGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerTrustGroupArgs struct {
	LocationName         string `pulumi:"locationName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	ServerTrustGroupName string `pulumi:"serverTrustGroupName"`
}


type LookupServerTrustGroupResult struct {
	GroupMembers []ServerInfoResponse `pulumi:"groupMembers"`
	Id           string               `pulumi:"id"`
	Name         string               `pulumi:"name"`
	TrustScopes  []string             `pulumi:"trustScopes"`
	Type         string               `pulumi:"type"`
}
