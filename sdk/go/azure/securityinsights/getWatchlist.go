


package securityinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWatchlist(ctx *pulumi.Context, args *LookupWatchlistArgs, opts ...pulumi.InvokeOption) (*LookupWatchlistResult, error) {
	var rv LookupWatchlistResult
	err := ctx.Invoke("azure-native:securityinsights:getWatchlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWatchlistArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WatchlistAlias                      string `pulumi:"watchlistAlias"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupWatchlistResult struct {
	ContentType         *string                    `pulumi:"contentType"`
	Created             *string                    `pulumi:"created"`
	CreatedBy           *WatchlistUserInfoResponse `pulumi:"createdBy"`
	DefaultDuration     *string                    `pulumi:"defaultDuration"`
	Description         *string                    `pulumi:"description"`
	DisplayName         string                     `pulumi:"displayName"`
	Etag                *string                    `pulumi:"etag"`
	Id                  string                     `pulumi:"id"`
	IsDeleted           *bool                      `pulumi:"isDeleted"`
	ItemsSearchKey      string                     `pulumi:"itemsSearchKey"`
	Labels              []string                   `pulumi:"labels"`
	Name                string                     `pulumi:"name"`
	NumberOfLinesToSkip *int                       `pulumi:"numberOfLinesToSkip"`
	Provider            string                     `pulumi:"provider"`
	RawContent          *string                    `pulumi:"rawContent"`
	Source              string                     `pulumi:"source"`
	SystemData          SystemDataResponse         `pulumi:"systemData"`
	TenantId            *string                    `pulumi:"tenantId"`
	Type                string                     `pulumi:"type"`
	Updated             *string                    `pulumi:"updated"`
	UpdatedBy           *WatchlistUserInfoResponse `pulumi:"updatedBy"`
	UploadStatus        *string                    `pulumi:"uploadStatus"`
	WatchlistAlias      *string                    `pulumi:"watchlistAlias"`
	WatchlistId         *string                    `pulumi:"watchlistId"`
	WatchlistItemsCount *int                       `pulumi:"watchlistItemsCount"`
	WatchlistType       *string                    `pulumi:"watchlistType"`
}
