


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualHubIpConfiguration(ctx *pulumi.Context, args *LookupVirtualHubIpConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHubIpConfigurationResult, error) {
	var rv LookupVirtualHubIpConfigurationResult
	err := ctx.Invoke("azure-native:network:getVirtualHubIpConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualHubIpConfigurationArgs struct {
	IpConfigName      string `pulumi:"ipConfigName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupVirtualHubIpConfigurationResult struct {
	Etag                      string                   `pulumi:"etag"`
	Id                        *string                  `pulumi:"id"`
	Name                      *string                  `pulumi:"name"`
	PrivateIPAddress          *string                  `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string                  `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         string                   `pulumi:"provisioningState"`
	PublicIPAddress           *PublicIPAddressResponse `pulumi:"publicIPAddress"`
	Subnet                    *SubnetResponse          `pulumi:"subnet"`
	Type                      string                   `pulumi:"type"`
}


func (val *LookupVirtualHubIpConfigurationResult) Defaults() *LookupVirtualHubIpConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.PublicIPAddress = tmp.PublicIPAddress.Defaults()

	tmp.Subnet = tmp.Subnet.Defaults()

	return &tmp
}

func LookupVirtualHubIpConfigurationOutput(ctx *pulumi.Context, args LookupVirtualHubIpConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualHubIpConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualHubIpConfigurationResult, error) {
			args := v.(LookupVirtualHubIpConfigurationArgs)
			r, err := LookupVirtualHubIpConfiguration(ctx, &args, opts...)
			var s LookupVirtualHubIpConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualHubIpConfigurationResultOutput)
}

type LookupVirtualHubIpConfigurationOutputArgs struct {
	IpConfigName      pulumi.StringInput `pulumi:"ipConfigName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualHubName    pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupVirtualHubIpConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubIpConfigurationArgs)(nil)).Elem()
}


type LookupVirtualHubIpConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualHubIpConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubIpConfigurationResult)(nil)).Elem()
}

func (o LookupVirtualHubIpConfigurationResultOutput) ToLookupVirtualHubIpConfigurationResultOutput() LookupVirtualHubIpConfigurationResultOutput {
	return o
}

func (o LookupVirtualHubIpConfigurationResultOutput) ToLookupVirtualHubIpConfigurationResultOutputWithContext(ctx context.Context) LookupVirtualHubIpConfigurationResultOutput {
	return o
}

func (o LookupVirtualHubIpConfigurationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVirtualHubIpConfigurationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubIpConfigurationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubIpConfigurationResultOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubIpConfigurationResultOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubIpConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualHubIpConfigurationResultOutput) PublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) *PublicIPAddressResponse { return v.PublicIPAddress }).(PublicIPAddressResponsePtrOutput)
}

func (o LookupVirtualHubIpConfigurationResultOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

func (o LookupVirtualHubIpConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubIpConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualHubIpConfigurationResultOutput{})
}
