


package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetworkLink struct {
	pulumi.CustomResourceState

	Etag                    pulumi.StringPtrOutput       `pulumi:"etag"`
	Location                pulumi.StringPtrOutput       `pulumi:"location"`
	Name                    pulumi.StringOutput          `pulumi:"name"`
	ProvisioningState       pulumi.StringOutput          `pulumi:"provisioningState"`
	RegistrationEnabled     pulumi.BoolPtrOutput         `pulumi:"registrationEnabled"`
	Tags                    pulumi.StringMapOutput       `pulumi:"tags"`
	Type                    pulumi.StringOutput          `pulumi:"type"`
	VirtualNetwork          SubResourceResponsePtrOutput `pulumi:"virtualNetwork"`
	VirtualNetworkLinkState pulumi.StringOutput          `pulumi:"virtualNetworkLinkState"`
}


func NewVirtualNetworkLink(ctx *pulumi.Context,
	name string, args *VirtualNetworkLinkArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateZoneName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateZoneName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualNetworkLink"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180901:VirtualNetworkLink"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200101:VirtualNetworkLink"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkLink
	err := ctx.RegisterResource("azure-native:network/v20200601:VirtualNetworkLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetworkLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkLinkState, opts ...pulumi.ResourceOption) (*VirtualNetworkLink, error) {
	var resource VirtualNetworkLink
	err := ctx.ReadResource("azure-native:network/v20200601:VirtualNetworkLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualNetworkLinkState struct {
}

type VirtualNetworkLinkState struct {
}

func (VirtualNetworkLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkLinkState)(nil)).Elem()
}

type virtualNetworkLinkArgs struct {
	Location               *string           `pulumi:"location"`
	PrivateZoneName        string            `pulumi:"privateZoneName"`
	RegistrationEnabled    *bool             `pulumi:"registrationEnabled"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	Tags                   map[string]string `pulumi:"tags"`
	VirtualNetwork         *SubResource      `pulumi:"virtualNetwork"`
	VirtualNetworkLinkName *string           `pulumi:"virtualNetworkLinkName"`
}


type VirtualNetworkLinkArgs struct {
	Location               pulumi.StringPtrInput
	PrivateZoneName        pulumi.StringInput
	RegistrationEnabled    pulumi.BoolPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
	VirtualNetwork         SubResourcePtrInput
	VirtualNetworkLinkName pulumi.StringPtrInput
}

func (VirtualNetworkLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkLinkArgs)(nil)).Elem()
}

type VirtualNetworkLinkInput interface {
	pulumi.Input

	ToVirtualNetworkLinkOutput() VirtualNetworkLinkOutput
	ToVirtualNetworkLinkOutputWithContext(ctx context.Context) VirtualNetworkLinkOutput
}

func (*VirtualNetworkLink) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkLink)(nil)).Elem()
}

func (i *VirtualNetworkLink) ToVirtualNetworkLinkOutput() VirtualNetworkLinkOutput {
	return i.ToVirtualNetworkLinkOutputWithContext(context.Background())
}

func (i *VirtualNetworkLink) ToVirtualNetworkLinkOutputWithContext(ctx context.Context) VirtualNetworkLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkLinkOutput)
}

type VirtualNetworkLinkOutput struct{ *pulumi.OutputState }

func (VirtualNetworkLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkLink)(nil)).Elem()
}

func (o VirtualNetworkLinkOutput) ToVirtualNetworkLinkOutput() VirtualNetworkLinkOutput {
	return o
}

func (o VirtualNetworkLinkOutput) ToVirtualNetworkLinkOutputWithContext(ctx context.Context) VirtualNetworkLinkOutput {
	return o
}

func (o VirtualNetworkLinkOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkLinkOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkLinkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkLinkOutput) RegistrationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) pulumi.BoolPtrOutput { return v.RegistrationEnabled }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkLinkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualNetworkLinkOutput) VirtualNetwork() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) SubResourceResponsePtrOutput { return v.VirtualNetwork }).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkLinkOutput) VirtualNetworkLinkState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) pulumi.StringOutput { return v.VirtualNetworkLinkState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkLinkOutput{})
}
