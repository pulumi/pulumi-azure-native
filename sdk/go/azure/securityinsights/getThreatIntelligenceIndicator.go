


package securityinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupThreatIntelligenceIndicator(ctx *pulumi.Context, args *LookupThreatIntelligenceIndicatorArgs, opts ...pulumi.InvokeOption) (*LookupThreatIntelligenceIndicatorResult, error) {
	var rv LookupThreatIntelligenceIndicatorResult
	err := ctx.Invoke("azure-native:securityinsights:getThreatIntelligenceIndicator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupThreatIntelligenceIndicatorArgs struct {
	Name                                string `pulumi:"name"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupThreatIntelligenceIndicatorResult struct {
	Etag *string `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Kind string  `pulumi:"kind"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}

func LookupThreatIntelligenceIndicatorOutput(ctx *pulumi.Context, args LookupThreatIntelligenceIndicatorOutputArgs, opts ...pulumi.InvokeOption) LookupThreatIntelligenceIndicatorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupThreatIntelligenceIndicatorResult, error) {
			args := v.(LookupThreatIntelligenceIndicatorArgs)
			r, err := LookupThreatIntelligenceIndicator(ctx, &args, opts...)
			var s LookupThreatIntelligenceIndicatorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupThreatIntelligenceIndicatorResultOutput)
}

type LookupThreatIntelligenceIndicatorOutputArgs struct {
	Name                                pulumi.StringInput `pulumi:"name"`
	OperationalInsightsResourceProvider pulumi.StringInput `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName                       pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupThreatIntelligenceIndicatorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThreatIntelligenceIndicatorArgs)(nil)).Elem()
}


type LookupThreatIntelligenceIndicatorResultOutput struct{ *pulumi.OutputState }

func (LookupThreatIntelligenceIndicatorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThreatIntelligenceIndicatorResult)(nil)).Elem()
}

func (o LookupThreatIntelligenceIndicatorResultOutput) ToLookupThreatIntelligenceIndicatorResultOutput() LookupThreatIntelligenceIndicatorResultOutput {
	return o
}

func (o LookupThreatIntelligenceIndicatorResultOutput) ToLookupThreatIntelligenceIndicatorResultOutputWithContext(ctx context.Context) LookupThreatIntelligenceIndicatorResultOutput {
	return o
}

func (o LookupThreatIntelligenceIndicatorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceIndicatorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupThreatIntelligenceIndicatorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceIndicatorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupThreatIntelligenceIndicatorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceIndicatorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupThreatIntelligenceIndicatorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceIndicatorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupThreatIntelligenceIndicatorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceIndicatorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupThreatIntelligenceIndicatorResultOutput{})
}
