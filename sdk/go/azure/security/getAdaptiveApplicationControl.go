


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAdaptiveApplicationControl(ctx *pulumi.Context, args *LookupAdaptiveApplicationControlArgs, opts ...pulumi.InvokeOption) (*LookupAdaptiveApplicationControlResult, error) {
	var rv LookupAdaptiveApplicationControlResult
	err := ctx.Invoke("azure-native:security:getAdaptiveApplicationControl", args, &rv, opts...)
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
	ConfigurationStatus  string                                           `pulumi:"configurationStatus"`
	EnforcementMode      *string                                          `pulumi:"enforcementMode"`
	Id                   string                                           `pulumi:"id"`
	Issues               []AdaptiveApplicationControlIssueSummaryResponse `pulumi:"issues"`
	Location             string                                           `pulumi:"location"`
	Name                 string                                           `pulumi:"name"`
	PathRecommendations  []PathRecommendationResponse                     `pulumi:"pathRecommendations"`
	ProtectionMode       *ProtectionModeResponse                          `pulumi:"protectionMode"`
	RecommendationStatus string                                           `pulumi:"recommendationStatus"`
	SourceSystem         string                                           `pulumi:"sourceSystem"`
	Type                 string                                           `pulumi:"type"`
	VmRecommendations    []VmRecommendationResponse                       `pulumi:"vmRecommendations"`
}

func LookupAdaptiveApplicationControlOutput(ctx *pulumi.Context, args LookupAdaptiveApplicationControlOutputArgs, opts ...pulumi.InvokeOption) LookupAdaptiveApplicationControlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAdaptiveApplicationControlResult, error) {
			args := v.(LookupAdaptiveApplicationControlArgs)
			r, err := LookupAdaptiveApplicationControl(ctx, &args, opts...)
			var s LookupAdaptiveApplicationControlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAdaptiveApplicationControlResultOutput)
}

type LookupAdaptiveApplicationControlOutputArgs struct {
	AscLocation pulumi.StringInput `pulumi:"ascLocation"`
	GroupName   pulumi.StringInput `pulumi:"groupName"`
}

func (LookupAdaptiveApplicationControlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdaptiveApplicationControlArgs)(nil)).Elem()
}

type LookupAdaptiveApplicationControlResultOutput struct{ *pulumi.OutputState }

func (LookupAdaptiveApplicationControlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdaptiveApplicationControlResult)(nil)).Elem()
}

func (o LookupAdaptiveApplicationControlResultOutput) ToLookupAdaptiveApplicationControlResultOutput() LookupAdaptiveApplicationControlResultOutput {
	return o
}

func (o LookupAdaptiveApplicationControlResultOutput) ToLookupAdaptiveApplicationControlResultOutputWithContext(ctx context.Context) LookupAdaptiveApplicationControlResultOutput {
	return o
}

func (o LookupAdaptiveApplicationControlResultOutput) ConfigurationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdaptiveApplicationControlResult) string { return v.ConfigurationStatus }).(pulumi.StringOutput)
}

func (o LookupAdaptiveApplicationControlResultOutput) EnforcementMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAdaptiveApplicationControlResult) *string { return v.EnforcementMode }).(pulumi.StringPtrOutput)
}

func (o LookupAdaptiveApplicationControlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdaptiveApplicationControlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAdaptiveApplicationControlResultOutput) Issues() AdaptiveApplicationControlIssueSummaryResponseArrayOutput {
	return o.ApplyT(func(v LookupAdaptiveApplicationControlResult) []AdaptiveApplicationControlIssueSummaryResponse {
		return v.Issues
	}).(AdaptiveApplicationControlIssueSummaryResponseArrayOutput)
}

func (o LookupAdaptiveApplicationControlResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdaptiveApplicationControlResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAdaptiveApplicationControlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdaptiveApplicationControlResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAdaptiveApplicationControlResultOutput) PathRecommendations() PathRecommendationResponseArrayOutput {
	return o.ApplyT(func(v LookupAdaptiveApplicationControlResult) []PathRecommendationResponse {
		return v.PathRecommendations
	}).(PathRecommendationResponseArrayOutput)
}

func (o LookupAdaptiveApplicationControlResultOutput) ProtectionMode() ProtectionModeResponsePtrOutput {
	return o.ApplyT(func(v LookupAdaptiveApplicationControlResult) *ProtectionModeResponse { return v.ProtectionMode }).(ProtectionModeResponsePtrOutput)
}

func (o LookupAdaptiveApplicationControlResultOutput) RecommendationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdaptiveApplicationControlResult) string { return v.RecommendationStatus }).(pulumi.StringOutput)
}

func (o LookupAdaptiveApplicationControlResultOutput) SourceSystem() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdaptiveApplicationControlResult) string { return v.SourceSystem }).(pulumi.StringOutput)
}

func (o LookupAdaptiveApplicationControlResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdaptiveApplicationControlResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAdaptiveApplicationControlResultOutput) VmRecommendations() VmRecommendationResponseArrayOutput {
	return o.ApplyT(func(v LookupAdaptiveApplicationControlResult) []VmRecommendationResponse { return v.VmRecommendations }).(VmRecommendationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAdaptiveApplicationControlResultOutput{})
}
