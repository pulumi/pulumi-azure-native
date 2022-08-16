


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupNetworkInterface(ctx *pulumi.Context, args *LookupNetworkInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceResult, error) {
	var rv LookupNetworkInterfaceResult
	err := ctx.Invoke("azure-native:network/v20150501preview:getNetworkInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkInterfaceArgs struct {
	NetworkInterfaceName string `pulumi:"networkInterfaceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupNetworkInterfaceResult struct {
	DnsSettings          *NetworkInterfaceDnsSettingsResponse      `pulumi:"dnsSettings"`
	EnableIPForwarding   *bool                                     `pulumi:"enableIPForwarding"`
	Etag                 *string                                   `pulumi:"etag"`
	Id                   string                                    `pulumi:"id"`
	IpConfigurations     []NetworkInterfaceIpConfigurationResponse `pulumi:"ipConfigurations"`
	Location             string                                    `pulumi:"location"`
	MacAddress           *string                                   `pulumi:"macAddress"`
	Name                 string                                    `pulumi:"name"`
	NetworkSecurityGroup *SubResourceResponse                      `pulumi:"networkSecurityGroup"`
	Primary              *bool                                     `pulumi:"primary"`
	ProvisioningState    *string                                   `pulumi:"provisioningState"`
	ResourceGuid         *string                                   `pulumi:"resourceGuid"`
	Tags                 map[string]string                         `pulumi:"tags"`
	Type                 string                                    `pulumi:"type"`
	VirtualMachine       *SubResourceResponse                      `pulumi:"virtualMachine"`
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
	NetworkInterfaceName pulumi.StringInput `pulumi:"networkInterfaceName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
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

func (o LookupNetworkInterfaceResultOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *bool { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkInterfaceResultOutput) IpConfigurations() NetworkInterfaceIpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) []NetworkInterfaceIpConfigurationResponse {
		return v.IpConfigurations
	}).(NetworkInterfaceIpConfigurationResponseArrayOutput)
}

func (o LookupNetworkInterfaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupNetworkInterfaceResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkInterfaceResultOutput) NetworkSecurityGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *SubResourceResponse { return v.NetworkSecurityGroup }).(SubResourceResponsePtrOutput)
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
