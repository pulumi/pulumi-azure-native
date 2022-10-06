


package v20170601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupNetworkInterface(ctx *pulumi.Context, args *LookupNetworkInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceResult, error) {
	var rv LookupNetworkInterfaceResult
	err := ctx.Invoke("azure-native:network/v20170601:getNetworkInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkInterfaceArgs struct {
	Expand               *string `pulumi:"expand"`
	NetworkInterfaceName string  `pulumi:"networkInterfaceName"`
	ResourceGroupName    string  `pulumi:"resourceGroupName"`
}


type LookupNetworkInterfaceResult struct {
	DnsSettings                 *NetworkInterfaceDnsSettingsResponse      `pulumi:"dnsSettings"`
	EnableAcceleratedNetworking *bool                                     `pulumi:"enableAcceleratedNetworking"`
	EnableIPForwarding          *bool                                     `pulumi:"enableIPForwarding"`
	Etag                        *string                                   `pulumi:"etag"`
	Id                          *string                                   `pulumi:"id"`
	IpConfigurations            []NetworkInterfaceIPConfigurationResponse `pulumi:"ipConfigurations"`
	Location                    *string                                   `pulumi:"location"`
	MacAddress                  *string                                   `pulumi:"macAddress"`
	Name                        string                                    `pulumi:"name"`
	NetworkSecurityGroup        *NetworkSecurityGroupResponse             `pulumi:"networkSecurityGroup"`
	Primary                     *bool                                     `pulumi:"primary"`
	ProvisioningState           *string                                   `pulumi:"provisioningState"`
	ResourceGuid                *string                                   `pulumi:"resourceGuid"`
	Tags                        map[string]string                         `pulumi:"tags"`
	Type                        string                                    `pulumi:"type"`
	VirtualMachine              *SubResourceResponse                      `pulumi:"virtualMachine"`
}

func LookupNetworkInterfaceOutput(ctx *pulumi.Context, args LookupNetworkInterfaceOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkInterfaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkInterfaceResult, error) {
			args := v.(LookupNetworkInterfaceArgs)
			r, err := LookupNetworkInterface(ctx, &args, opts...)
			var s LookupNetworkInterfaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkInterfaceResultOutput)
}

type LookupNetworkInterfaceOutputArgs struct {
	Expand               pulumi.StringPtrInput `pulumi:"expand"`
	NetworkInterfaceName pulumi.StringInput    `pulumi:"networkInterfaceName"`
	ResourceGroupName    pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupNetworkInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfaceArgs)(nil)).Elem()
}


type LookupNetworkInterfaceResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfaceResult)(nil)).Elem()
}

func (o LookupNetworkInterfaceResultOutput) ToLookupNetworkInterfaceResultOutput() LookupNetworkInterfaceResultOutput {
	return o
}

func (o LookupNetworkInterfaceResultOutput) ToLookupNetworkInterfaceResultOutputWithContext(ctx context.Context) LookupNetworkInterfaceResultOutput {
	return o
}

func (o LookupNetworkInterfaceResultOutput) DnsSettings() NetworkInterfaceDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *NetworkInterfaceDnsSettingsResponse { return v.DnsSettings }).(NetworkInterfaceDnsSettingsResponsePtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *bool { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) IpConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) []NetworkInterfaceIPConfigurationResponse {
		return v.IpConfigurations
	}).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

func (o LookupNetworkInterfaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkInterfaceResultOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *NetworkSecurityGroupResponse { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNetworkInterfaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNetworkInterfaceResultOutput) VirtualMachine() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *SubResourceResponse { return v.VirtualMachine }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkInterfaceResultOutput{})
}
