


package insights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAnalyticsItem(ctx *pulumi.Context, args *LookupAnalyticsItemArgs, opts ...pulumi.InvokeOption) (*LookupAnalyticsItemResult, error) {
	var rv LookupAnalyticsItemResult
	err := ctx.Invoke("azure-native:insights:getAnalyticsItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAnalyticsItemArgs struct {
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceName      string  `pulumi:"resourceName"`
	ScopePath         string  `pulumi:"scopePath"`
}


type LookupAnalyticsItemResult struct {
	Content      *string                                                     `pulumi:"content"`
	Id           *string                                                     `pulumi:"id"`
	Name         *string                                                     `pulumi:"name"`
	Properties   ApplicationInsightsComponentAnalyticsItemPropertiesResponse `pulumi:"properties"`
	Scope        *string                                                     `pulumi:"scope"`
	TimeCreated  string                                                      `pulumi:"timeCreated"`
	TimeModified string                                                      `pulumi:"timeModified"`
	Type         *string                                                     `pulumi:"type"`
	Version      string                                                      `pulumi:"version"`
}
