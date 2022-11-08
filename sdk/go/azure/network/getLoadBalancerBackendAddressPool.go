


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLoadBalancerBackendAddressPool(ctx *pulumi.Context, args *LookupLoadBalancerBackendAddressPoolArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerBackendAddressPoolResult, error) {
	var rv LookupLoadBalancerBackendAddressPoolResult
	err := ctx.Invoke("azure-native:network:getLoadBalancerBackendAddressPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoadBalancerBackendAddressPoolArgs struct {
	BackendAddressPoolName string `pulumi:"backendAddressPoolName"`
	LoadBalancerName       string `pulumi:"loadBalancerName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupLoadBalancerBackendAddressPoolResult struct {
	BackendIPConfigurations      []NetworkInterfaceIPConfigurationResponse `pulumi:"backendIPConfigurations"`
	Etag                         string                                    `pulumi:"etag"`
	Id                           *string                                   `pulumi:"id"`
	LoadBalancerBackendAddresses []LoadBalancerBackendAddressResponse      `pulumi:"loadBalancerBackendAddresses"`
	LoadBalancingRules           []SubResourceResponse                     `pulumi:"loadBalancingRules"`
	Name                         *string                                   `pulumi:"name"`
	OutboundRule                 SubResourceResponse                       `pulumi:"outboundRule"`
	OutboundRules                []SubResourceResponse                     `pulumi:"outboundRules"`
	ProvisioningState            string                                    `pulumi:"provisioningState"`
	Type                         string                                    `pulumi:"type"`
}

func LookupLoadBalancerBackendAddressPoolOutput(ctx *pulumi.Context, args LookupLoadBalancerBackendAddressPoolOutputArgs, opts ...pulumi.InvokeOption) LookupLoadBalancerBackendAddressPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLoadBalancerBackendAddressPoolResult, error) {
			args := v.(LookupLoadBalancerBackendAddressPoolArgs)
			r, err := LookupLoadBalancerBackendAddressPool(ctx, &args, opts...)
			var s LookupLoadBalancerBackendAddressPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLoadBalancerBackendAddressPoolResultOutput)
}

type LookupLoadBalancerBackendAddressPoolOutputArgs struct {
	BackendAddressPoolName pulumi.StringInput `pulumi:"backendAddressPoolName"`
	LoadBalancerName       pulumi.StringInput `pulumi:"loadBalancerName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLoadBalancerBackendAddressPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerBackendAddressPoolArgs)(nil)).Elem()
}


type LookupLoadBalancerBackendAddressPoolResultOutput struct{ *pulumi.OutputState }

func (LookupLoadBalancerBackendAddressPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerBackendAddressPoolResult)(nil)).Elem()
}

func (o LookupLoadBalancerBackendAddressPoolResultOutput) ToLookupLoadBalancerBackendAddressPoolResultOutput() LookupLoadBalancerBackendAddressPoolResultOutput {
	return o
}

func (o LookupLoadBalancerBackendAddressPoolResultOutput) ToLookupLoadBalancerBackendAddressPoolResultOutputWithContext(ctx context.Context) LookupLoadBalancerBackendAddressPoolResultOutput {
	return o
}

func (o LookupLoadBalancerBackendAddressPoolResultOutput) BackendIPConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) []NetworkInterfaceIPConfigurationResponse {
		return v.BackendIPConfigurations
	}).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

func (o LookupLoadBalancerBackendAddressPoolResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupLoadBalancerBackendAddressPoolResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerBackendAddressPoolResultOutput) LoadBalancerBackendAddresses() LoadBalancerBackendAddressResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) []LoadBalancerBackendAddressResponse {
		return v.LoadBalancerBackendAddresses
	}).(LoadBalancerBackendAddressResponseArrayOutput)
}

func (o LookupLoadBalancerBackendAddressPoolResultOutput) LoadBalancingRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) []SubResourceResponse { return v.LoadBalancingRules }).(SubResourceResponseArrayOutput)
}

func (o LookupLoadBalancerBackendAddressPoolResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupLoadBalancerBackendAddressPoolResultOutput) OutboundRule() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) SubResourceResponse { return v.OutboundRule }).(SubResourceResponseOutput)
}

func (o LookupLoadBalancerBackendAddressPoolResultOutput) OutboundRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) []SubResourceResponse { return v.OutboundRules }).(SubResourceResponseArrayOutput)
}

func (o LookupLoadBalancerBackendAddressPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupLoadBalancerBackendAddressPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerBackendAddressPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoadBalancerBackendAddressPoolResultOutput{})
}
