


package securityinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetEntitiesGetTimeline(ctx *pulumi.Context, args *GetEntitiesGetTimelineArgs, opts ...pulumi.InvokeOption) (*GetEntitiesGetTimelineResult, error) {
	var rv GetEntitiesGetTimelineResult
	err := ctx.Invoke("azure-native:securityinsights:getEntitiesGetTimeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetEntitiesGetTimelineArgs struct {
	EndTime                             string   `pulumi:"endTime"`
	EntityId                            string   `pulumi:"entityId"`
	Kinds                               []string `pulumi:"kinds"`
	NumberOfBucket                      *int     `pulumi:"numberOfBucket"`
	OperationalInsightsResourceProvider string   `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string   `pulumi:"resourceGroupName"`
	StartTime                           string   `pulumi:"startTime"`
	WorkspaceName                       string   `pulumi:"workspaceName"`
}


type GetEntitiesGetTimelineResult struct {
	MetaData *TimelineResultsMetadataResponse `pulumi:"metaData"`
	Value    []interface{}                    `pulumi:"value"`
}
