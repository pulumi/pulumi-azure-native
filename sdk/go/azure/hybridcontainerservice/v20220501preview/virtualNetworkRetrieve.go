


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetworkRetrieve struct {
	pulumi.CustomResourceState

	ExtendedLocation VirtualNetworksResponseExtendedLocationPtrOutput `pulumi:"extendedLocation"`
	Location         pulumi.StringOutput                              `pulumi:"location"`
	Name             pulumi.StringOutput                              `pulumi:"name"`
	Properties       VirtualNetworksPropertiesResponseOutput          `pulumi:"properties"`
	SystemData       SystemDataResponseOutput                         `pulumi:"systemData"`
	Tags             pulumi.StringMapOutput                           `pulumi:"tags"`
	Type             pulumi.StringOutput                              `pulumi:"type"`
}


func NewVirtualNetworkRetrieve(ctx *pulumi.Context,
	name string, args *VirtualNetworkRetrieveArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkRetrieve, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcontainerservice:virtualNetworkRetrieve"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkRetrieve
	err := ctx.RegisterResource("azure-native:hybridcontainerservice/v20220501preview:virtualNetworkRetrieve", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetworkRetrieve(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkRetrieveState, opts ...pulumi.ResourceOption) (*VirtualNetworkRetrieve, error) {
	var resource VirtualNetworkRetrieve
	err := ctx.ReadResource("azure-native:hybridcontainerservice/v20220501preview:virtualNetworkRetrieve", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualNetworkRetrieveState struct {
}

type VirtualNetworkRetrieveState struct {
}

func (VirtualNetworkRetrieveState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRetrieveState)(nil)).Elem()
}

type virtualNetworkRetrieveArgs struct {
	ExtendedLocation    *VirtualNetworksExtendedLocation `pulumi:"extendedLocation"`
	Location            *string                          `pulumi:"location"`
	Properties          *VirtualNetworksProperties       `pulumi:"properties"`
	ResourceGroupName   string                           `pulumi:"resourceGroupName"`
	Tags                map[string]string                `pulumi:"tags"`
	VirtualNetworksName *string                          `pulumi:"virtualNetworksName"`
}


type VirtualNetworkRetrieveArgs struct {
	ExtendedLocation    VirtualNetworksExtendedLocationPtrInput
	Location            pulumi.StringPtrInput
	Properties          VirtualNetworksPropertiesPtrInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
	VirtualNetworksName pulumi.StringPtrInput
}

func (VirtualNetworkRetrieveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkRetrieveArgs)(nil)).Elem()
}

type VirtualNetworkRetrieveInput interface {
	pulumi.Input

	ToVirtualNetworkRetrieveOutput() VirtualNetworkRetrieveOutput
	ToVirtualNetworkRetrieveOutputWithContext(ctx context.Context) VirtualNetworkRetrieveOutput
}

func (*VirtualNetworkRetrieve) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkRetrieve)(nil)).Elem()
}

func (i *VirtualNetworkRetrieve) ToVirtualNetworkRetrieveOutput() VirtualNetworkRetrieveOutput {
	return i.ToVirtualNetworkRetrieveOutputWithContext(context.Background())
}

func (i *VirtualNetworkRetrieve) ToVirtualNetworkRetrieveOutputWithContext(ctx context.Context) VirtualNetworkRetrieveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRetrieveOutput)
}

type VirtualNetworkRetrieveOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRetrieveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkRetrieve)(nil)).Elem()
}

func (o VirtualNetworkRetrieveOutput) ToVirtualNetworkRetrieveOutput() VirtualNetworkRetrieveOutput {
	return o
}

func (o VirtualNetworkRetrieveOutput) ToVirtualNetworkRetrieveOutputWithContext(ctx context.Context) VirtualNetworkRetrieveOutput {
	return o
}

func (o VirtualNetworkRetrieveOutput) ExtendedLocation() VirtualNetworksResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkRetrieve) VirtualNetworksResponseExtendedLocationPtrOutput {
		return v.ExtendedLocation
	}).(VirtualNetworksResponseExtendedLocationPtrOutput)
}

func (o VirtualNetworkRetrieveOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkRetrieve) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualNetworkRetrieveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkRetrieve) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkRetrieveOutput) Properties() VirtualNetworksPropertiesResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkRetrieve) VirtualNetworksPropertiesResponseOutput { return v.Properties }).(VirtualNetworksPropertiesResponseOutput)
}

func (o VirtualNetworkRetrieveOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkRetrieve) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VirtualNetworkRetrieveOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkRetrieve) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkRetrieveOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkRetrieve) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkRetrieveOutput{})
}
