


package media

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssetFilter(ctx *pulumi.Context, args *LookupAssetFilterArgs, opts ...pulumi.InvokeOption) (*LookupAssetFilterResult, error) {
	var rv LookupAssetFilterResult
	err := ctx.Invoke("azure-native:media:getAssetFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssetFilterArgs struct {
	AccountName       string `pulumi:"accountName"`
	AssetName         string `pulumi:"assetName"`
	FilterName        string `pulumi:"filterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAssetFilterResult struct {
	FirstQuality          *FirstQualityResponse          `pulumi:"firstQuality"`
	Id                    string                         `pulumi:"id"`
	Name                  string                         `pulumi:"name"`
	PresentationTimeRange *PresentationTimeRangeResponse `pulumi:"presentationTimeRange"`
	SystemData            SystemDataResponse             `pulumi:"systemData"`
	Tracks                []FilterTrackSelectionResponse `pulumi:"tracks"`
	Type                  string                         `pulumi:"type"`
}
