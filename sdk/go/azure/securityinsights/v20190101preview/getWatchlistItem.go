


package v20190101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWatchlistItem(ctx *pulumi.Context, args *LookupWatchlistItemArgs, opts ...pulumi.InvokeOption) (*LookupWatchlistItemResult, error) {
	var rv LookupWatchlistItemResult
	err := ctx.Invoke("azure-native:securityinsights/v20190101preview:getWatchlistItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWatchlistItemArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WatchlistAlias                      string `pulumi:"watchlistAlias"`
	WatchlistItemId                     string `pulumi:"watchlistItemId"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupWatchlistItemResult struct {
	Created           *string                    `pulumi:"created"`
	CreatedBy         *WatchlistUserInfoResponse `pulumi:"createdBy"`
	EntityMapping     interface{}                `pulumi:"entityMapping"`
	Etag              *string                    `pulumi:"etag"`
	Id                string                     `pulumi:"id"`
	IsDeleted         *bool                      `pulumi:"isDeleted"`
	ItemsKeyValue     interface{}                `pulumi:"itemsKeyValue"`
	Name              string                     `pulumi:"name"`
	TenantId          *string                    `pulumi:"tenantId"`
	Type              string                     `pulumi:"type"`
	Updated           *string                    `pulumi:"updated"`
	UpdatedBy         *WatchlistUserInfoResponse `pulumi:"updatedBy"`
	WatchlistItemId   *string                    `pulumi:"watchlistItemId"`
	WatchlistItemType *string                    `pulumi:"watchlistItemType"`
}
