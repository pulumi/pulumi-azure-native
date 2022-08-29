


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTopic(ctx *pulumi.Context, args *LookupTopicArgs, opts ...pulumi.InvokeOption) (*LookupTopicResult, error) {
	var rv LookupTopicResult
	err := ctx.Invoke("azure-native:eventgrid/v20200401preview:getTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupTopicArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TopicName         string `pulumi:"topicName"`
}


type LookupTopicResult struct {
	Endpoint                   string                              `pulumi:"endpoint"`
	Id                         string                              `pulumi:"id"`
	Identity                   *IdentityInfoResponse               `pulumi:"identity"`
	InboundIpRules             []InboundIpRuleResponse             `pulumi:"inboundIpRules"`
	InputSchema                *string                             `pulumi:"inputSchema"`
	InputSchemaMapping         *JsonInputSchemaMappingResponse     `pulumi:"inputSchemaMapping"`
	Location                   string                              `pulumi:"location"`
	MetricResourceId           string                              `pulumi:"metricResourceId"`
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                             `pulumi:"publicNetworkAccess"`
	Sku                        *ResourceSkuResponse                `pulumi:"sku"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
}


func (val *LookupTopicResult) Defaults() *LookupTopicResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.InputSchema) {
		inputSchema_ := "EventGridSchema"
		tmp.InputSchema = &inputSchema_
	}
	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	return &tmp
}

func LookupTopicOutput(ctx *pulumi.Context, args LookupTopicOutputArgs, opts ...pulumi.InvokeOption) LookupTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTopicResult, error) {
			args := v.(LookupTopicArgs)
			r, err := LookupTopic(ctx, &args, opts...)
			var s LookupTopicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTopicResultOutput)
}

type LookupTopicOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TopicName         pulumi.StringInput `pulumi:"topicName"`
}

func (LookupTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicArgs)(nil)).Elem()
}


type LookupTopicResultOutput struct{ *pulumi.OutputState }

func (LookupTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicResult)(nil)).Elem()
}

func (o LookupTopicResultOutput) ToLookupTopicResultOutput() LookupTopicResultOutput {
	return o
}

func (o LookupTopicResultOutput) ToLookupTopicResultOutputWithContext(ctx context.Context) LookupTopicResultOutput {
	return o
}

func (o LookupTopicResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o LookupTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTopicResultOutput) Identity() IdentityInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *IdentityInfoResponse { return v.Identity }).(IdentityInfoResponsePtrOutput)
}

func (o LookupTopicResultOutput) InboundIpRules() InboundIpRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupTopicResult) []InboundIpRuleResponse { return v.InboundIpRules }).(InboundIpRuleResponseArrayOutput)
}

func (o LookupTopicResultOutput) InputSchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.InputSchema }).(pulumi.StringPtrOutput)
}

func (o LookupTopicResultOutput) InputSchemaMapping() JsonInputSchemaMappingResponsePtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *JsonInputSchemaMappingResponse { return v.InputSchemaMapping }).(JsonInputSchemaMappingResponsePtrOutput)
}

func (o LookupTopicResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupTopicResultOutput) MetricResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.MetricResourceId }).(pulumi.StringOutput)
}

func (o LookupTopicResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTopicResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupTopicResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupTopicResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupTopicResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupTopicResultOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *ResourceSkuResponse { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

func (o LookupTopicResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTopicResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupTopicResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTopicResultOutput{})
}
