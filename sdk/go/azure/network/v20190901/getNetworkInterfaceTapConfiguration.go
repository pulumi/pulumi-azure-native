


package v20190901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkInterfaceTapConfiguration(ctx *pulumi.Context, args *LookupNetworkInterfaceTapConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceTapConfigurationResult, error) {
	var rv LookupNetworkInterfaceTapConfigurationResult
	err := ctx.Invoke("azure-native:network/v20190901:getNetworkInterfaceTapConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkInterfaceTapConfigurationArgs struct {
	NetworkInterfaceName string `pulumi:"networkInterfaceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	TapConfigurationName string `pulumi:"tapConfigurationName"`
}


type LookupNetworkInterfaceTapConfigurationResult struct {
	Etag              string                     `pulumi:"etag"`
	Id                *string                    `pulumi:"id"`
	Name              *string                    `pulumi:"name"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	Type              string                     `pulumi:"type"`
	VirtualNetworkTap *VirtualNetworkTapResponse `pulumi:"virtualNetworkTap"`
}

func LookupNetworkInterfaceTapConfigurationOutput(ctx *pulumi.Context, args LookupNetworkInterfaceTapConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkInterfaceTapConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkInterfaceTapConfigurationResult, error) {
			args := v.(LookupNetworkInterfaceTapConfigurationArgs)
			r, err := LookupNetworkInterfaceTapConfiguration(ctx, &args, opts...)
			var s LookupNetworkInterfaceTapConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkInterfaceTapConfigurationResultOutput)
}

type LookupNetworkInterfaceTapConfigurationOutputArgs struct {
	NetworkInterfaceName pulumi.StringInput `pulumi:"networkInterfaceName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	TapConfigurationName pulumi.StringInput `pulumi:"tapConfigurationName"`
}

func (LookupNetworkInterfaceTapConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfaceTapConfigurationArgs)(nil)).Elem()
}


type LookupNetworkInterfaceTapConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkInterfaceTapConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfaceTapConfigurationResult)(nil)).Elem()
}

func (o LookupNetworkInterfaceTapConfigurationResultOutput) ToLookupNetworkInterfaceTapConfigurationResultOutput() LookupNetworkInterfaceTapConfigurationResultOutput {
	return o
}

func (o LookupNetworkInterfaceTapConfigurationResultOutput) ToLookupNetworkInterfaceTapConfigurationResultOutputWithContext(ctx context.Context) LookupNetworkInterfaceTapConfigurationResultOutput {
	return o
}

func (o LookupNetworkInterfaceTapConfigurationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceTapConfigurationResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupNetworkInterfaceTapConfigurationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceTapConfigurationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInterfaceTapConfigurationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceTapConfigurationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInterfaceTapConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceTapConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNetworkInterfaceTapConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceTapConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNetworkInterfaceTapConfigurationResultOutput) VirtualNetworkTap() VirtualNetworkTapResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceTapConfigurationResult) *VirtualNetworkTapResponse {
		return v.VirtualNetworkTap
	}).(VirtualNetworkTapResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkInterfaceTapConfigurationResultOutput{})
}
