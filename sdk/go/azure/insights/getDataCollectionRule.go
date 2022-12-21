


package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataCollectionRule(ctx *pulumi.Context, args *LookupDataCollectionRuleArgs, opts ...pulumi.InvokeOption) (*LookupDataCollectionRuleResult, error) {
	var rv LookupDataCollectionRuleResult
	err := ctx.Invoke("azure-native:insights:getDataCollectionRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataCollectionRuleArgs struct {
	DataCollectionRuleName string `pulumi:"dataCollectionRuleName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupDataCollectionRuleResult struct {
	DataFlows         []DataFlowResponse                      `pulumi:"dataFlows"`
	DataSources       *DataCollectionRuleResponseDataSources  `pulumi:"dataSources"`
	Description       *string                                 `pulumi:"description"`
	Destinations      *DataCollectionRuleResponseDestinations `pulumi:"destinations"`
	Etag              string                                  `pulumi:"etag"`
	Id                string                                  `pulumi:"id"`
	ImmutableId       string                                  `pulumi:"immutableId"`
	Kind              *string                                 `pulumi:"kind"`
	Location          string                                  `pulumi:"location"`
	Name              string                                  `pulumi:"name"`
	ProvisioningState string                                  `pulumi:"provisioningState"`
	Tags              map[string]string                       `pulumi:"tags"`
	Type              string                                  `pulumi:"type"`
}

func LookupDataCollectionRuleOutput(ctx *pulumi.Context, args LookupDataCollectionRuleOutputArgs, opts ...pulumi.InvokeOption) LookupDataCollectionRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataCollectionRuleResult, error) {
			args := v.(LookupDataCollectionRuleArgs)
			r, err := LookupDataCollectionRule(ctx, &args, opts...)
			var s LookupDataCollectionRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataCollectionRuleResultOutput)
}

type LookupDataCollectionRuleOutputArgs struct {
	DataCollectionRuleName pulumi.StringInput `pulumi:"dataCollectionRuleName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataCollectionRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataCollectionRuleArgs)(nil)).Elem()
}


type LookupDataCollectionRuleResultOutput struct{ *pulumi.OutputState }

func (LookupDataCollectionRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataCollectionRuleResult)(nil)).Elem()
}

func (o LookupDataCollectionRuleResultOutput) ToLookupDataCollectionRuleResultOutput() LookupDataCollectionRuleResultOutput {
	return o
}

func (o LookupDataCollectionRuleResultOutput) ToLookupDataCollectionRuleResultOutputWithContext(ctx context.Context) LookupDataCollectionRuleResultOutput {
	return o
}

func (o LookupDataCollectionRuleResultOutput) DataFlows() DataFlowResponseArrayOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) []DataFlowResponse { return v.DataFlows }).(DataFlowResponseArrayOutput)
}

func (o LookupDataCollectionRuleResultOutput) DataSources() DataCollectionRuleResponseDataSourcesPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) *DataCollectionRuleResponseDataSources { return v.DataSources }).(DataCollectionRuleResponseDataSourcesPtrOutput)
}

func (o LookupDataCollectionRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupDataCollectionRuleResultOutput) Destinations() DataCollectionRuleResponseDestinationsPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) *DataCollectionRuleResponseDestinations { return v.Destinations }).(DataCollectionRuleResponseDestinationsPtrOutput)
}

func (o LookupDataCollectionRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDataCollectionRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataCollectionRuleResultOutput) ImmutableId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) string { return v.ImmutableId }).(pulumi.StringOutput)
}

func (o LookupDataCollectionRuleResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupDataCollectionRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDataCollectionRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataCollectionRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDataCollectionRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDataCollectionRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataCollectionRuleResultOutput{})
}
