


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespaceNetworkRuleSet(ctx *pulumi.Context, args *LookupNamespaceNetworkRuleSetArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceNetworkRuleSetResult, error) {
	var rv LookupNamespaceNetworkRuleSetResult
	err := ctx.Invoke("azure-native:servicebus/v20221001preview:getNamespaceNetworkRuleSet", args, &rv, opts...)
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

func LookupNamespaceNetworkRuleSetOutput(ctx *pulumi.Context, args LookupNamespaceNetworkRuleSetOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceNetworkRuleSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamespaceNetworkRuleSetResult, error) {
			args := v.(LookupNamespaceNetworkRuleSetArgs)
			r, err := LookupNamespaceNetworkRuleSet(ctx, &args, opts...)
			var s LookupNamespaceNetworkRuleSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamespaceNetworkRuleSetResultOutput)
}

type LookupNamespaceNetworkRuleSetOutputArgs struct {
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNamespaceNetworkRuleSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceNetworkRuleSetArgs)(nil)).Elem()
}


type LookupNamespaceNetworkRuleSetResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceNetworkRuleSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceNetworkRuleSetResult)(nil)).Elem()
}

func (o LookupNamespaceNetworkRuleSetResultOutput) ToLookupNamespaceNetworkRuleSetResultOutput() LookupNamespaceNetworkRuleSetResultOutput {
	return o
}

func (o LookupNamespaceNetworkRuleSetResultOutput) ToLookupNamespaceNetworkRuleSetResultOutputWithContext(ctx context.Context) LookupNamespaceNetworkRuleSetResultOutput {
	return o
}

func (o LookupNamespaceNetworkRuleSetResultOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceNetworkRuleSetResult) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceNetworkRuleSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceNetworkRuleSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNamespaceNetworkRuleSetResultOutput) IpRules() NWRuleSetIpRulesResponseArrayOutput {
	return o.ApplyT(func(v LookupNamespaceNetworkRuleSetResult) []NWRuleSetIpRulesResponse { return v.IpRules }).(NWRuleSetIpRulesResponseArrayOutput)
}

func (o LookupNamespaceNetworkRuleSetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceNetworkRuleSetResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupNamespaceNetworkRuleSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceNetworkRuleSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNamespaceNetworkRuleSetResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceNetworkRuleSetResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceNetworkRuleSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNamespaceNetworkRuleSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupNamespaceNetworkRuleSetResultOutput) TrustedServiceAccessEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNamespaceNetworkRuleSetResult) *bool { return v.TrustedServiceAccessEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupNamespaceNetworkRuleSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceNetworkRuleSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNamespaceNetworkRuleSetResultOutput) VirtualNetworkRules() NWRuleSetVirtualNetworkRulesResponseArrayOutput {
	return o.ApplyT(func(v LookupNamespaceNetworkRuleSetResult) []NWRuleSetVirtualNetworkRulesResponse {
		return v.VirtualNetworkRules
	}).(NWRuleSetVirtualNetworkRulesResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceNetworkRuleSetResultOutput{})
}
