


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualnetworkRetrieve struct {
	pulumi.CustomResourceState

	ExtendedLocation  ExtendedLocationResponsePtrOutput                   `pulumi:"extendedLocation"`
	Location          pulumi.StringOutput                                 `pulumi:"location"`
	Name              pulumi.StringOutput                                 `pulumi:"name"`
	NetworkType       pulumi.StringPtrOutput                              `pulumi:"networkType"`
	ProvisioningState pulumi.StringOutput                                 `pulumi:"provisioningState"`
	ResourceName      pulumi.StringPtrOutput                              `pulumi:"resourceName"`
	Status            VirtualNetworkStatusResponseOutput                  `pulumi:"status"`
	Subnets           VirtualnetworksPropertiesResponseSubnetsArrayOutput `pulumi:"subnets"`
	SystemData        SystemDataResponseOutput                            `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                              `pulumi:"tags"`
	Type              pulumi.StringOutput                                 `pulumi:"type"`
}


func NewVirtualnetworkRetrieve(ctx *pulumi.Context,
	name string, args *VirtualnetworkRetrieveArgs, opts ...pulumi.ResourceOption) (*VirtualnetworkRetrieve, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210701preview:virtualnetworkRetrieve"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualnetworkRetrieve
	err := ctx.RegisterResource("azure-native:azurestackhci/v20210901preview:virtualnetworkRetrieve", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualnetworkRetrieve(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualnetworkRetrieveState, opts ...pulumi.ResourceOption) (*VirtualnetworkRetrieve, error) {
	var resource VirtualnetworkRetrieve
	err := ctx.ReadResource("azure-native:azurestackhci/v20210901preview:virtualnetworkRetrieve", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualnetworkRetrieveState struct {
}

type VirtualnetworkRetrieveState struct {
}

func (VirtualnetworkRetrieveState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualnetworkRetrieveState)(nil)).Elem()
}

type virtualnetworkRetrieveArgs struct {
	ExtendedLocation    *ExtendedLocation                  `pulumi:"extendedLocation"`
	Location            *string                            `pulumi:"location"`
	NetworkType         *string                            `pulumi:"networkType"`
	ResourceGroupName   string                             `pulumi:"resourceGroupName"`
	ResourceName        *string                            `pulumi:"resourceName"`
	Subnets             []VirtualnetworksPropertiesSubnets `pulumi:"subnets"`
	Tags                map[string]string                  `pulumi:"tags"`
	VirtualnetworksName *string                            `pulumi:"virtualnetworksName"`
}


type VirtualnetworkRetrieveArgs struct {
	ExtendedLocation    ExtendedLocationPtrInput
	Location            pulumi.StringPtrInput
	NetworkType         pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	ResourceName        pulumi.StringPtrInput
	Subnets             VirtualnetworksPropertiesSubnetsArrayInput
	Tags                pulumi.StringMapInput
	VirtualnetworksName pulumi.StringPtrInput
}

func (VirtualnetworkRetrieveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualnetworkRetrieveArgs)(nil)).Elem()
}

type VirtualnetworkRetrieveInput interface {
	pulumi.Input

	ToVirtualnetworkRetrieveOutput() VirtualnetworkRetrieveOutput
	ToVirtualnetworkRetrieveOutputWithContext(ctx context.Context) VirtualnetworkRetrieveOutput
}

func (*VirtualnetworkRetrieve) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualnetworkRetrieve)(nil)).Elem()
}

func (i *VirtualnetworkRetrieve) ToVirtualnetworkRetrieveOutput() VirtualnetworkRetrieveOutput {
	return i.ToVirtualnetworkRetrieveOutputWithContext(context.Background())
}

func (i *VirtualnetworkRetrieve) ToVirtualnetworkRetrieveOutputWithContext(ctx context.Context) VirtualnetworkRetrieveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworkRetrieveOutput)
}

type VirtualnetworkRetrieveOutput struct{ *pulumi.OutputState }

func (VirtualnetworkRetrieveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualnetworkRetrieve)(nil)).Elem()
}

func (o VirtualnetworkRetrieveOutput) ToVirtualnetworkRetrieveOutput() VirtualnetworkRetrieveOutput {
	return o
}

func (o VirtualnetworkRetrieveOutput) ToVirtualnetworkRetrieveOutputWithContext(ctx context.Context) VirtualnetworkRetrieveOutput {
	return o
}

func (o VirtualnetworkRetrieveOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualnetworkRetrieve) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o VirtualnetworkRetrieveOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualnetworkRetrieve) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualnetworkRetrieveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualnetworkRetrieve) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualnetworkRetrieveOutput) NetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualnetworkRetrieve) pulumi.StringPtrOutput { return v.NetworkType }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworkRetrieveOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualnetworkRetrieve) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualnetworkRetrieveOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualnetworkRetrieve) pulumi.StringPtrOutput { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworkRetrieveOutput) Status() VirtualNetworkStatusResponseOutput {
	return o.ApplyT(func(v *VirtualnetworkRetrieve) VirtualNetworkStatusResponseOutput { return v.Status }).(VirtualNetworkStatusResponseOutput)
}

func (o VirtualnetworkRetrieveOutput) Subnets() VirtualnetworksPropertiesResponseSubnetsArrayOutput {
	return o.ApplyT(func(v *VirtualnetworkRetrieve) VirtualnetworksPropertiesResponseSubnetsArrayOutput { return v.Subnets }).(VirtualnetworksPropertiesResponseSubnetsArrayOutput)
}

func (o VirtualnetworkRetrieveOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualnetworkRetrieve) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VirtualnetworkRetrieveOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualnetworkRetrieve) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualnetworkRetrieveOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualnetworkRetrieve) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualnetworkRetrieveOutput{})
}
