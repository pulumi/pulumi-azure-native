


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInboundNatRule(ctx *pulumi.Context, args *LookupInboundNatRuleArgs, opts ...pulumi.InvokeOption) (*LookupInboundNatRuleResult, error) {
	var rv LookupInboundNatRuleResult
	err := ctx.Invoke("azure-native:network/v20210301:getInboundNatRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupInboundNatRuleArgs struct {
	Expand             *string `pulumi:"expand"`
	InboundNatRuleName string  `pulumi:"inboundNatRuleName"`
	LoadBalancerName   string  `pulumi:"loadBalancerName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type LookupInboundNatRuleResult struct {
	BackendAddressPool      *SubResourceResponse                    `pulumi:"backendAddressPool"`
	BackendIPConfiguration  NetworkInterfaceIPConfigurationResponse `pulumi:"backendIPConfiguration"`
	BackendPort             *int                                    `pulumi:"backendPort"`
	EnableFloatingIP        *bool                                   `pulumi:"enableFloatingIP"`
	EnableTcpReset          *bool                                   `pulumi:"enableTcpReset"`
	Etag                    string                                  `pulumi:"etag"`
	FrontendIPConfiguration *SubResourceResponse                    `pulumi:"frontendIPConfiguration"`
	FrontendPort            *int                                    `pulumi:"frontendPort"`
	FrontendPortRangeEnd    *int                                    `pulumi:"frontendPortRangeEnd"`
	FrontendPortRangeStart  *int                                    `pulumi:"frontendPortRangeStart"`
	Id                      *string                                 `pulumi:"id"`
	IdleTimeoutInMinutes    *int                                    `pulumi:"idleTimeoutInMinutes"`
	Name                    *string                                 `pulumi:"name"`
	Protocol                *string                                 `pulumi:"protocol"`
	ProvisioningState       string                                  `pulumi:"provisioningState"`
	Type                    string                                  `pulumi:"type"`
}


func (val *LookupInboundNatRuleResult) Defaults() *LookupInboundNatRuleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.BackendIPConfiguration = *tmp.BackendIPConfiguration.Defaults()

	return &tmp
}

func LookupInboundNatRuleOutput(ctx *pulumi.Context, args LookupInboundNatRuleOutputArgs, opts ...pulumi.InvokeOption) LookupInboundNatRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInboundNatRuleResult, error) {
			args := v.(LookupInboundNatRuleArgs)
			r, err := LookupInboundNatRule(ctx, &args, opts...)
			var s LookupInboundNatRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInboundNatRuleResultOutput)
}

type LookupInboundNatRuleOutputArgs struct {
	Expand             pulumi.StringPtrInput `pulumi:"expand"`
	InboundNatRuleName pulumi.StringInput    `pulumi:"inboundNatRuleName"`
	LoadBalancerName   pulumi.StringInput    `pulumi:"loadBalancerName"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupInboundNatRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInboundNatRuleArgs)(nil)).Elem()
}


type LookupInboundNatRuleResultOutput struct{ *pulumi.OutputState }

func (LookupInboundNatRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInboundNatRuleResult)(nil)).Elem()
}

func (o LookupInboundNatRuleResultOutput) ToLookupInboundNatRuleResultOutput() LookupInboundNatRuleResultOutput {
	return o
}

func (o LookupInboundNatRuleResultOutput) ToLookupInboundNatRuleResultOutputWithContext(ctx context.Context) LookupInboundNatRuleResultOutput {
	return o
}

func (o LookupInboundNatRuleResultOutput) BackendAddressPool() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *SubResourceResponse { return v.BackendAddressPool }).(SubResourceResponsePtrOutput)
}

func (o LookupInboundNatRuleResultOutput) BackendIPConfiguration() NetworkInterfaceIPConfigurationResponseOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) NetworkInterfaceIPConfigurationResponse {
		return v.BackendIPConfiguration
	}).(NetworkInterfaceIPConfigurationResponseOutput)
}

func (o LookupInboundNatRuleResultOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o LookupInboundNatRuleResultOutput) EnableFloatingIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *bool { return v.EnableFloatingIP }).(pulumi.BoolPtrOutput)
}

func (o LookupInboundNatRuleResultOutput) EnableTcpReset() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *bool { return v.EnableTcpReset }).(pulumi.BoolPtrOutput)
}

func (o LookupInboundNatRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupInboundNatRuleResultOutput) FrontendIPConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *SubResourceResponse { return v.FrontendIPConfiguration }).(SubResourceResponsePtrOutput)
}

func (o LookupInboundNatRuleResultOutput) FrontendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *int { return v.FrontendPort }).(pulumi.IntPtrOutput)
}

func (o LookupInboundNatRuleResultOutput) FrontendPortRangeEnd() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *int { return v.FrontendPortRangeEnd }).(pulumi.IntPtrOutput)
}

func (o LookupInboundNatRuleResultOutput) FrontendPortRangeStart() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *int { return v.FrontendPortRangeStart }).(pulumi.IntPtrOutput)
}

func (o LookupInboundNatRuleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupInboundNatRuleResultOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o LookupInboundNatRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupInboundNatRuleResultOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o LookupInboundNatRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupInboundNatRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInboundNatRuleResultOutput{})
}
