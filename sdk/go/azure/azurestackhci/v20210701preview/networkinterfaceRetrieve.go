


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkinterfaceRetrieve struct {
	pulumi.CustomResourceState

	DnsSettings       InterfaceDNSSettingsResponsePtrOutput `pulumi:"dnsSettings"`
	ExtendedLocation  ExtendedLocationResponsePtrOutput     `pulumi:"extendedLocation"`
	IpConfigurations  IpConfigurationResponseArrayOutput    `pulumi:"ipConfigurations"`
	Location          pulumi.StringOutput                   `pulumi:"location"`
	MacAddress        pulumi.StringPtrOutput                `pulumi:"macAddress"`
	Name              pulumi.StringOutput                   `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                   `pulumi:"provisioningState"`
	ResourceName      pulumi.StringPtrOutput                `pulumi:"resourceName"`
	Status            NetworkInterfaceStatusResponseOutput  `pulumi:"status"`
	SystemData        SystemDataResponseOutput              `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                `pulumi:"tags"`
	Type              pulumi.StringOutput                   `pulumi:"type"`
}


func NewNetworkinterfaceRetrieve(ctx *pulumi.Context,
	name string, args *NetworkinterfaceRetrieveArgs, opts ...pulumi.ResourceOption) (*NetworkinterfaceRetrieve, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource NetworkinterfaceRetrieve
	err := ctx.RegisterResource("azure-native:azurestackhci/v20210701preview:networkinterfaceRetrieve", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkinterfaceRetrieve(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkinterfaceRetrieveState, opts ...pulumi.ResourceOption) (*NetworkinterfaceRetrieve, error) {
	var resource NetworkinterfaceRetrieve
	err := ctx.ReadResource("azure-native:azurestackhci/v20210701preview:networkinterfaceRetrieve", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkinterfaceRetrieveState struct {
}

type NetworkinterfaceRetrieveState struct {
}

func (NetworkinterfaceRetrieveState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkinterfaceRetrieveState)(nil)).Elem()
}

type networkinterfaceRetrieveArgs struct {
	DnsSettings           *InterfaceDNSSettings `pulumi:"dnsSettings"`
	ExtendedLocation      *ExtendedLocation     `pulumi:"extendedLocation"`
	IpConfigurations      []IpConfiguration     `pulumi:"ipConfigurations"`
	Location              *string               `pulumi:"location"`
	MacAddress            *string               `pulumi:"macAddress"`
	NetworkinterfacesName *string               `pulumi:"networkinterfacesName"`
	ResourceGroupName     string                `pulumi:"resourceGroupName"`
	ResourceName          *string               `pulumi:"resourceName"`
	Tags                  map[string]string     `pulumi:"tags"`
}


type NetworkinterfaceRetrieveArgs struct {
	DnsSettings           InterfaceDNSSettingsPtrInput
	ExtendedLocation      ExtendedLocationPtrInput
	IpConfigurations      IpConfigurationArrayInput
	Location              pulumi.StringPtrInput
	MacAddress            pulumi.StringPtrInput
	NetworkinterfacesName pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	ResourceName          pulumi.StringPtrInput
	Tags                  pulumi.StringMapInput
}

func (NetworkinterfaceRetrieveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkinterfaceRetrieveArgs)(nil)).Elem()
}

type NetworkinterfaceRetrieveInput interface {
	pulumi.Input

	ToNetworkinterfaceRetrieveOutput() NetworkinterfaceRetrieveOutput
	ToNetworkinterfaceRetrieveOutputWithContext(ctx context.Context) NetworkinterfaceRetrieveOutput
}

func (*NetworkinterfaceRetrieve) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkinterfaceRetrieve)(nil)).Elem()
}

func (i *NetworkinterfaceRetrieve) ToNetworkinterfaceRetrieveOutput() NetworkinterfaceRetrieveOutput {
	return i.ToNetworkinterfaceRetrieveOutputWithContext(context.Background())
}

func (i *NetworkinterfaceRetrieve) ToNetworkinterfaceRetrieveOutputWithContext(ctx context.Context) NetworkinterfaceRetrieveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkinterfaceRetrieveOutput)
}

type NetworkinterfaceRetrieveOutput struct{ *pulumi.OutputState }

func (NetworkinterfaceRetrieveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkinterfaceRetrieve)(nil)).Elem()
}

func (o NetworkinterfaceRetrieveOutput) ToNetworkinterfaceRetrieveOutput() NetworkinterfaceRetrieveOutput {
	return o
}

func (o NetworkinterfaceRetrieveOutput) ToNetworkinterfaceRetrieveOutputWithContext(ctx context.Context) NetworkinterfaceRetrieveOutput {
	return o
}

func (o NetworkinterfaceRetrieveOutput) DnsSettings() InterfaceDNSSettingsResponsePtrOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) InterfaceDNSSettingsResponsePtrOutput { return v.DnsSettings }).(InterfaceDNSSettingsResponsePtrOutput)
}

func (o NetworkinterfaceRetrieveOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o NetworkinterfaceRetrieveOutput) IpConfigurations() IpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) IpConfigurationResponseArrayOutput { return v.IpConfigurations }).(IpConfigurationResponseArrayOutput)
}

func (o NetworkinterfaceRetrieveOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o NetworkinterfaceRetrieveOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) pulumi.StringPtrOutput { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkinterfaceRetrieveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkinterfaceRetrieveOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NetworkinterfaceRetrieveOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) pulumi.StringPtrOutput { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o NetworkinterfaceRetrieveOutput) Status() NetworkInterfaceStatusResponseOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) NetworkInterfaceStatusResponseOutput { return v.Status }).(NetworkInterfaceStatusResponseOutput)
}

func (o NetworkinterfaceRetrieveOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o NetworkinterfaceRetrieveOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkinterfaceRetrieveOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkinterfaceRetrieve) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkinterfaceRetrieveOutput{})
}
