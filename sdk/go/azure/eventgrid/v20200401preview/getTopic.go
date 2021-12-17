


package v20200401preview

import (
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
