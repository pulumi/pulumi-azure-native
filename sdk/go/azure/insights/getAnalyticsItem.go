


package insights

import (
	"context"
	"reflect"

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

func LookupAnalyticsItemOutput(ctx *pulumi.Context, args LookupAnalyticsItemOutputArgs, opts ...pulumi.InvokeOption) LookupAnalyticsItemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAnalyticsItemResult, error) {
			args := v.(LookupAnalyticsItemArgs)
			r, err := LookupAnalyticsItem(ctx, &args, opts...)
			var s LookupAnalyticsItemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAnalyticsItemResultOutput)
}

type LookupAnalyticsItemOutputArgs struct {
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput    `pulumi:"resourceName"`
	ScopePath         pulumi.StringInput    `pulumi:"scopePath"`
}

func (LookupAnalyticsItemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnalyticsItemArgs)(nil)).Elem()
}


type LookupAnalyticsItemResultOutput struct{ *pulumi.OutputState }

func (LookupAnalyticsItemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnalyticsItemResult)(nil)).Elem()
}

func (o LookupAnalyticsItemResultOutput) ToLookupAnalyticsItemResultOutput() LookupAnalyticsItemResultOutput {
	return o
}

func (o LookupAnalyticsItemResultOutput) ToLookupAnalyticsItemResultOutputWithContext(ctx context.Context) LookupAnalyticsItemResultOutput {
	return o
}

func (o LookupAnalyticsItemResultOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalyticsItemResult) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o LookupAnalyticsItemResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalyticsItemResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupAnalyticsItemResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalyticsItemResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupAnalyticsItemResultOutput) Properties() ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput {
	return o.ApplyT(func(v LookupAnalyticsItemResult) ApplicationInsightsComponentAnalyticsItemPropertiesResponse {
		return v.Properties
	}).(ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput)
}

func (o LookupAnalyticsItemResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalyticsItemResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupAnalyticsItemResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnalyticsItemResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o LookupAnalyticsItemResultOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnalyticsItemResult) string { return v.TimeModified }).(pulumi.StringOutput)
}

func (o LookupAnalyticsItemResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnalyticsItemResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupAnalyticsItemResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnalyticsItemResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAnalyticsItemResultOutput{})
}
