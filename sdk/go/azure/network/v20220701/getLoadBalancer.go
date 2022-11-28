


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("azure-native:network/v20220701:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoadBalancerArgs struct {
	Expand            *string `pulumi:"expand"`
	LoadBalancerName  string  `pulumi:"loadBalancerName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupLoadBalancerResult struct {
	BackendAddressPools      []BackendAddressPoolResponse      `pulumi:"backendAddressPools"`
	Etag                     string                            `pulumi:"etag"`
	ExtendedLocation         *ExtendedLocationResponse         `pulumi:"extendedLocation"`
	FrontendIPConfigurations []FrontendIPConfigurationResponse `pulumi:"frontendIPConfigurations"`
	Id                       *string                           `pulumi:"id"`
	InboundNatPools          []InboundNatPoolResponse          `pulumi:"inboundNatPools"`
	InboundNatRules          []InboundNatRuleResponse          `pulumi:"inboundNatRules"`
	LoadBalancingRules       []LoadBalancingRuleResponse       `pulumi:"loadBalancingRules"`
	Location                 *string                           `pulumi:"location"`
	Name                     string                            `pulumi:"name"`
	OutboundRules            []OutboundRuleResponse            `pulumi:"outboundRules"`
	Probes                   []ProbeResponse                   `pulumi:"probes"`
	ProvisioningState        string                            `pulumi:"provisioningState"`
	ResourceGuid             string                            `pulumi:"resourceGuid"`
	Sku                      *LoadBalancerSkuResponse          `pulumi:"sku"`
	Tags                     map[string]string                 `pulumi:"tags"`
	Type                     string                            `pulumi:"type"`
}

func LookupLoadBalancerOutput(ctx *pulumi.Context, args LookupLoadBalancerOutputArgs, opts ...pulumi.InvokeOption) LookupLoadBalancerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLoadBalancerResult, error) {
			args := v.(LookupLoadBalancerArgs)
			r, err := LookupLoadBalancer(ctx, &args, opts...)
			var s LookupLoadBalancerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLoadBalancerResultOutput)
}

type LookupLoadBalancerOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LoadBalancerName  pulumi.StringInput    `pulumi:"loadBalancerName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupLoadBalancerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerArgs)(nil)).Elem()
}


type LookupLoadBalancerResultOutput struct{ *pulumi.OutputState }

func (LookupLoadBalancerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerResult)(nil)).Elem()
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutput() LookupLoadBalancerResultOutput {
	return o
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutputWithContext(ctx context.Context) LookupLoadBalancerResultOutput {
	return o
}

func (o LookupLoadBalancerResultOutput) BackendAddressPools() BackendAddressPoolResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []BackendAddressPoolResponse { return v.BackendAddressPools }).(BackendAddressPoolResponseArrayOutput)
}

func (o LookupLoadBalancerResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupLoadBalancerResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupLoadBalancerResultOutput) FrontendIPConfigurations() FrontendIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []FrontendIPConfigurationResponse { return v.FrontendIPConfigurations }).(FrontendIPConfigurationResponseArrayOutput)
}

func (o LookupLoadBalancerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) InboundNatPools() InboundNatPoolResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []InboundNatPoolResponse { return v.InboundNatPools }).(InboundNatPoolResponseArrayOutput)
}

func (o LookupLoadBalancerResultOutput) InboundNatRules() InboundNatRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []InboundNatRuleResponse { return v.InboundNatRules }).(InboundNatRuleResponseArrayOutput)
}

func (o LookupLoadBalancerResultOutput) LoadBalancingRules() LoadBalancingRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []LoadBalancingRuleResponse { return v.LoadBalancingRules }).(LoadBalancingRuleResponseArrayOutput)
}

func (o LookupLoadBalancerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLoadBalancerResultOutput) OutboundRules() OutboundRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []OutboundRuleResponse { return v.OutboundRules }).(OutboundRuleResponseArrayOutput)
}

func (o LookupLoadBalancerResultOutput) Probes() ProbeResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []ProbeResponse { return v.Probes }).(ProbeResponseArrayOutput)
}

func (o LookupLoadBalancerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupLoadBalancerResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupLoadBalancerResultOutput) Sku() LoadBalancerSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *LoadBalancerSkuResponse { return v.Sku }).(LoadBalancerSkuResponsePtrOutput)
}

func (o LookupLoadBalancerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupLoadBalancerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoadBalancerResultOutput{})
}
