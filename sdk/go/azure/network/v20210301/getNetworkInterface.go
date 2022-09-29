


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkInterface(ctx *pulumi.Context, args *LookupNetworkInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceResult, error) {
	var rv LookupNetworkInterfaceResult
	err := ctx.Invoke("azure-native:network/v20210301:getNetworkInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNetworkInterfaceArgs struct {
	Expand               *string `pulumi:"expand"`
	NetworkInterfaceName string  `pulumi:"networkInterfaceName"`
	ResourceGroupName    string  `pulumi:"resourceGroupName"`
}


type LookupNetworkInterfaceResult struct {
	DnsSettings                 *NetworkInterfaceDnsSettingsResponse       `pulumi:"dnsSettings"`
	DscpConfiguration           SubResourceResponse                        `pulumi:"dscpConfiguration"`
	EnableAcceleratedNetworking *bool                                      `pulumi:"enableAcceleratedNetworking"`
	EnableIPForwarding          *bool                                      `pulumi:"enableIPForwarding"`
	Etag                        string                                     `pulumi:"etag"`
	ExtendedLocation            *ExtendedLocationResponse                  `pulumi:"extendedLocation"`
	HostedWorkloads             []string                                   `pulumi:"hostedWorkloads"`
	Id                          *string                                    `pulumi:"id"`
	IpConfigurations            []NetworkInterfaceIPConfigurationResponse  `pulumi:"ipConfigurations"`
	Location                    *string                                    `pulumi:"location"`
	MacAddress                  string                                     `pulumi:"macAddress"`
	MigrationPhase              *string                                    `pulumi:"migrationPhase"`
	Name                        string                                     `pulumi:"name"`
	NetworkSecurityGroup        *NetworkSecurityGroupResponse              `pulumi:"networkSecurityGroup"`
	NicType                     *string                                    `pulumi:"nicType"`
	Primary                     bool                                       `pulumi:"primary"`
	PrivateEndpoint             PrivateEndpointResponse                    `pulumi:"privateEndpoint"`
	PrivateLinkService          *PrivateLinkServiceResponse                `pulumi:"privateLinkService"`
	ProvisioningState           string                                     `pulumi:"provisioningState"`
	ResourceGuid                string                                     `pulumi:"resourceGuid"`
	Tags                        map[string]string                          `pulumi:"tags"`
	TapConfigurations           []NetworkInterfaceTapConfigurationResponse `pulumi:"tapConfigurations"`
	Type                        string                                     `pulumi:"type"`
	VirtualMachine              SubResourceResponse                        `pulumi:"virtualMachine"`
	WorkloadType                *string                                    `pulumi:"workloadType"`
}


func (val *LookupNetworkInterfaceResult) Defaults() *LookupNetworkInterfaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.PrivateEndpoint = *tmp.PrivateEndpoint.Defaults()

	return &tmp
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

func (o LookupNetworkInterfaceResultOutput) DscpConfiguration() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) SubResourceResponse { return v.DscpConfiguration }).(SubResourceResponseOutput)
}

func (o LookupNetworkInterfaceResultOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *bool { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupNetworkInterfaceResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) HostedWorkloads() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) []string { return v.HostedWorkloads }).(pulumi.StringArrayOutput)
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

func (o LookupNetworkInterfaceResultOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.MacAddress }).(pulumi.StringOutput)
}

func (o LookupNetworkInterfaceResultOutput) MigrationPhase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.MigrationPhase }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkInterfaceResultOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *NetworkSecurityGroupResponse { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) NicType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.NicType }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) Primary() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) bool { return v.Primary }).(pulumi.BoolOutput)
}

func (o LookupNetworkInterfaceResultOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponseOutput)
}

func (o LookupNetworkInterfaceResultOutput) PrivateLinkService() PrivateLinkServiceResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *PrivateLinkServiceResponse { return v.PrivateLinkService }).(PrivateLinkServiceResponsePtrOutput)
}

func (o LookupNetworkInterfaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNetworkInterfaceResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupNetworkInterfaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNetworkInterfaceResultOutput) TapConfigurations() NetworkInterfaceTapConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) []NetworkInterfaceTapConfigurationResponse {
		return v.TapConfigurations
	}).(NetworkInterfaceTapConfigurationResponseArrayOutput)
}

func (o LookupNetworkInterfaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNetworkInterfaceResultOutput) VirtualMachine() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) SubResourceResponse { return v.VirtualMachine }).(SubResourceResponseOutput)
}

func (o LookupNetworkInterfaceResultOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkInterfaceResultOutput{})
}
