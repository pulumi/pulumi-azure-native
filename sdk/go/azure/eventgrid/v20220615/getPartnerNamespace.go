


package v20220615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPartnerNamespace(ctx *pulumi.Context, args *LookupPartnerNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupPartnerNamespaceResult, error) {
	var rv LookupPartnerNamespaceResult
	err := ctx.Invoke("azure-native:eventgrid/v20220615:getPartnerNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPartnerNamespaceArgs struct {
	PartnerNamespaceName string `pulumi:"partnerNamespaceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupPartnerNamespaceResult struct {
	DisableLocalAuth                    *bool                               `pulumi:"disableLocalAuth"`
	Endpoint                            string                              `pulumi:"endpoint"`
	Id                                  string                              `pulumi:"id"`
	InboundIpRules                      []InboundIpRuleResponse             `pulumi:"inboundIpRules"`
	Location                            string                              `pulumi:"location"`
	Name                                string                              `pulumi:"name"`
	PartnerRegistrationFullyQualifiedId *string                             `pulumi:"partnerRegistrationFullyQualifiedId"`
	PartnerTopicRoutingMode             *string                             `pulumi:"partnerTopicRoutingMode"`
	PrivateEndpointConnections          []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState                   string                              `pulumi:"provisioningState"`
	PublicNetworkAccess                 *string                             `pulumi:"publicNetworkAccess"`
	SystemData                          SystemDataResponse                  `pulumi:"systemData"`
	Tags                                map[string]string                   `pulumi:"tags"`
	Type                                string                              `pulumi:"type"`
}


func (val *LookupPartnerNamespaceResult) Defaults() *LookupPartnerNamespaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DisableLocalAuth) {
		disableLocalAuth_ := false
		tmp.DisableLocalAuth = &disableLocalAuth_
	}
	if isZero(tmp.PartnerTopicRoutingMode) {
		partnerTopicRoutingMode_ := "SourceEventAttribute"
		tmp.PartnerTopicRoutingMode = &partnerTopicRoutingMode_
	}
	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	return &tmp
}

func LookupPartnerNamespaceOutput(ctx *pulumi.Context, args LookupPartnerNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupPartnerNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPartnerNamespaceResult, error) {
			args := v.(LookupPartnerNamespaceArgs)
			r, err := LookupPartnerNamespace(ctx, &args, opts...)
			var s LookupPartnerNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPartnerNamespaceResultOutput)
}

type LookupPartnerNamespaceOutputArgs struct {
	PartnerNamespaceName pulumi.StringInput `pulumi:"partnerNamespaceName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPartnerNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerNamespaceArgs)(nil)).Elem()
}


type LookupPartnerNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupPartnerNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerNamespaceResult)(nil)).Elem()
}

func (o LookupPartnerNamespaceResultOutput) ToLookupPartnerNamespaceResultOutput() LookupPartnerNamespaceResultOutput {
	return o
}

func (o LookupPartnerNamespaceResultOutput) ToLookupPartnerNamespaceResultOutputWithContext(ctx context.Context) LookupPartnerNamespaceResultOutput {
	return o
}

func (o LookupPartnerNamespaceResultOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupPartnerNamespaceResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o LookupPartnerNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPartnerNamespaceResultOutput) InboundIpRules() InboundIpRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) []InboundIpRuleResponse { return v.InboundIpRules }).(InboundIpRuleResponseArrayOutput)
}

func (o LookupPartnerNamespaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPartnerNamespaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPartnerNamespaceResultOutput) PartnerRegistrationFullyQualifiedId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) *string { return v.PartnerRegistrationFullyQualifiedId }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerNamespaceResultOutput) PartnerTopicRoutingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) *string { return v.PartnerTopicRoutingMode }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerNamespaceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupPartnerNamespaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPartnerNamespaceResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerNamespaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPartnerNamespaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPartnerNamespaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerNamespaceResultOutput{})
}
