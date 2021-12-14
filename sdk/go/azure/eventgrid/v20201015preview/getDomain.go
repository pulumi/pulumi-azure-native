


package v20201015preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	var rv LookupDomainResult
	err := ctx.Invoke("azure-native:eventgrid/v20201015preview:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDomainArgs struct {
	DomainName        string `pulumi:"domainName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDomainResult struct {
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
	SystemData                 SystemDataResponse                  `pulumi:"systemData"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
}


func (val *LookupDomainResult) Defaults() *LookupDomainResult {
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
	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}
