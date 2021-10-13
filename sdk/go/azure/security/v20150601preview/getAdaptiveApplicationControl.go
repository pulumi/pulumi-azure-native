


package v20150601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAdaptiveApplicationControl(ctx *pulumi.Context, args *LookupAdaptiveApplicationControlArgs, opts ...pulumi.InvokeOption) (*LookupAdaptiveApplicationControlResult, error) {
	var rv LookupAdaptiveApplicationControlResult
	err := ctx.Invoke("azure-native:security/v20150601preview:getAdaptiveApplicationControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAdaptiveApplicationControlArgs struct {
	AscLocation string `pulumi:"ascLocation"`
	GroupName   string `pulumi:"groupName"`
}

type LookupAdaptiveApplicationControlResult struct {
	ConfigurationStatus  *string                               `pulumi:"configurationStatus"`
	EnforcementMode      *string                               `pulumi:"enforcementMode"`
	Id                   string                                `pulumi:"id"`
	Issues               []AppWhitelistingIssueSummaryResponse `pulumi:"issues"`
	Location             string                                `pulumi:"location"`
	Name                 string                                `pulumi:"name"`
	PathRecommendations  []PathRecommendationResponse          `pulumi:"pathRecommendations"`
	ProtectionMode       *ProtectionModeResponse               `pulumi:"protectionMode"`
	RecommendationStatus *string                               `pulumi:"recommendationStatus"`
	SourceSystem         *string                               `pulumi:"sourceSystem"`
	Type                 string                                `pulumi:"type"`
	VmRecommendations    []VmRecommendationResponse            `pulumi:"vmRecommendations"`
}
