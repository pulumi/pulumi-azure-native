


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetEntitiesGetTimeline(ctx *pulumi.Context, args *GetEntitiesGetTimelineArgs, opts ...pulumi.InvokeOption) (*GetEntitiesGetTimelineResult, error) {
	var rv GetEntitiesGetTimelineResult
	err := ctx.Invoke("azure-native:securityinsights/v20221101preview:getEntitiesGetTimeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetEntitiesGetTimelineArgs struct {
	EndTime           string   `pulumi:"endTime"`
	EntityId          string   `pulumi:"entityId"`
	Kinds             []string `pulumi:"kinds"`
	NumberOfBucket    *int     `pulumi:"numberOfBucket"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	StartTime         string   `pulumi:"startTime"`
	WorkspaceName     string   `pulumi:"workspaceName"`
}


type GetEntitiesGetTimelineResult struct {
	MetaData *TimelineResultsMetadataResponse `pulumi:"metaData"`
	Value    []interface{}                    `pulumi:"value"`
}

func GetEntitiesGetTimelineOutput(ctx *pulumi.Context, args GetEntitiesGetTimelineOutputArgs, opts ...pulumi.InvokeOption) GetEntitiesGetTimelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEntitiesGetTimelineResult, error) {
			args := v.(GetEntitiesGetTimelineArgs)
			r, err := GetEntitiesGetTimeline(ctx, &args, opts...)
			var s GetEntitiesGetTimelineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEntitiesGetTimelineResultOutput)
}

type GetEntitiesGetTimelineOutputArgs struct {
	EndTime           pulumi.StringInput      `pulumi:"endTime"`
	EntityId          pulumi.StringInput      `pulumi:"entityId"`
	Kinds             pulumi.StringArrayInput `pulumi:"kinds"`
	NumberOfBucket    pulumi.IntPtrInput      `pulumi:"numberOfBucket"`
	ResourceGroupName pulumi.StringInput      `pulumi:"resourceGroupName"`
	StartTime         pulumi.StringInput      `pulumi:"startTime"`
	WorkspaceName     pulumi.StringInput      `pulumi:"workspaceName"`
}

func (GetEntitiesGetTimelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEntitiesGetTimelineArgs)(nil)).Elem()
}


type GetEntitiesGetTimelineResultOutput struct{ *pulumi.OutputState }

func (GetEntitiesGetTimelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEntitiesGetTimelineResult)(nil)).Elem()
}

func (o GetEntitiesGetTimelineResultOutput) ToGetEntitiesGetTimelineResultOutput() GetEntitiesGetTimelineResultOutput {
	return o
}

func (o GetEntitiesGetTimelineResultOutput) ToGetEntitiesGetTimelineResultOutputWithContext(ctx context.Context) GetEntitiesGetTimelineResultOutput {
	return o
}

func (o GetEntitiesGetTimelineResultOutput) MetaData() TimelineResultsMetadataResponsePtrOutput {
	return o.ApplyT(func(v GetEntitiesGetTimelineResult) *TimelineResultsMetadataResponse { return v.MetaData }).(TimelineResultsMetadataResponsePtrOutput)
}

func (o GetEntitiesGetTimelineResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetEntitiesGetTimelineResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEntitiesGetTimelineResultOutput{})
}
