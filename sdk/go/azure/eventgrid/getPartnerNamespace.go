


package eventgrid

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPartnerNamespace(ctx *pulumi.Context, args *LookupPartnerNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupPartnerNamespaceResult, error) {
	var rv LookupPartnerNamespaceResult
	err := ctx.Invoke("azure-native:eventgrid:getPartnerNamespace", args, &rv, opts...)
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
	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	return &tmp
}
