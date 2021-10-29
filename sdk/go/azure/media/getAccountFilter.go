


package media

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccountFilter(ctx *pulumi.Context, args *LookupAccountFilterArgs, opts ...pulumi.InvokeOption) (*LookupAccountFilterResult, error) {
	var rv LookupAccountFilterResult
	err := ctx.Invoke("azure-native:media:getAccountFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountFilterArgs struct {
	AccountName       string `pulumi:"accountName"`
	FilterName        string `pulumi:"filterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccountFilterResult struct {
	FirstQuality          *FirstQualityResponse          `pulumi:"firstQuality"`
	Id                    string                         `pulumi:"id"`
	Name                  string                         `pulumi:"name"`
	PresentationTimeRange *PresentationTimeRangeResponse `pulumi:"presentationTimeRange"`
	SystemData            SystemDataResponse             `pulumi:"systemData"`
	Tracks                []FilterTrackSelectionResponse `pulumi:"tracks"`
	Type                  string                         `pulumi:"type"`
}
