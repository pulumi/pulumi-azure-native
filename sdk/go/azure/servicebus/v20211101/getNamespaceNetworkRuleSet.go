


package v20211101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespaceNetworkRuleSet(ctx *pulumi.Context, args *LookupNamespaceNetworkRuleSetArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceNetworkRuleSetResult, error) {
	var rv LookupNamespaceNetworkRuleSetResult
	err := ctx.Invoke("azure-native:servicebus/v20211101:getNamespaceNetworkRuleSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNamespaceNetworkRuleSetArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNamespaceNetworkRuleSetResult struct {
	DefaultAction               *string                                `pulumi:"defaultAction"`
	Id                          string                                 `pulumi:"id"`
	IpRules                     []NWRuleSetIpRulesResponse             `pulumi:"ipRules"`
	Location                    string                                 `pulumi:"location"`
	Name                        string                                 `pulumi:"name"`
	PublicNetworkAccess         *string                                `pulumi:"publicNetworkAccess"`
	SystemData                  SystemDataResponse                     `pulumi:"systemData"`
	TrustedServiceAccessEnabled *bool                                  `pulumi:"trustedServiceAccessEnabled"`
	Type                        string                                 `pulumi:"type"`
	VirtualNetworkRules         []NWRuleSetVirtualNetworkRulesResponse `pulumi:"virtualNetworkRules"`
}


func (val *LookupNamespaceNetworkRuleSetResult) Defaults() *LookupNamespaceNetworkRuleSetResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	return &tmp
}
