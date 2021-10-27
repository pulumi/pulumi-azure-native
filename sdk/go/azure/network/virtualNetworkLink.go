


package network

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
			Type: pulumi.String("azure-nextgen:network:VirtualNetworkLink"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180901:VirtualNetworkLink"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180901:VirtualNetworkLink"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200101:VirtualNetworkLink"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200101:VirtualNetworkLink"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualNetworkLink"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:VirtualNetworkLink"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkLink
	err := ctx.RegisterResource("azure-native:network:VirtualNetworkLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetworkLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkLinkState, opts ...pulumi.ResourceOption) (*VirtualNetworkLink, error) {
	var resource VirtualNetworkLink
	err := ctx.ReadResource("azure-native:network:VirtualNetworkLink", name, id, state, &resource, opts...)
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
	Etag                   *string           `pulumi:"etag"`
	Location               *string           `pulumi:"location"`
	PrivateZoneName        string            `pulumi:"privateZoneName"`
	RegistrationEnabled    *bool             `pulumi:"registrationEnabled"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	Tags                   map[string]string `pulumi:"tags"`
	VirtualNetwork         *SubResource      `pulumi:"virtualNetwork"`
	VirtualNetworkLinkName *string           `pulumi:"virtualNetworkLinkName"`
}


type VirtualNetworkLinkArgs struct {
	Etag                   pulumi.StringPtrInput
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
	return reflect.TypeOf((*VirtualNetworkLink)(nil))
}

func (i *VirtualNetworkLink) ToVirtualNetworkLinkOutput() VirtualNetworkLinkOutput {
	return i.ToVirtualNetworkLinkOutputWithContext(context.Background())
}

func (i *VirtualNetworkLink) ToVirtualNetworkLinkOutputWithContext(ctx context.Context) VirtualNetworkLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkLinkOutput)
}

type VirtualNetworkLinkOutput struct{ *pulumi.OutputState }

func (VirtualNetworkLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkLink)(nil))
}

func (o VirtualNetworkLinkOutput) ToVirtualNetworkLinkOutput() VirtualNetworkLinkOutput {
	return o
}

func (o VirtualNetworkLinkOutput) ToVirtualNetworkLinkOutputWithContext(ctx context.Context) VirtualNetworkLinkOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkLinkInput)(nil)).Elem(), &VirtualNetworkLink{})
	pulumi.RegisterOutputType(VirtualNetworkLinkOutput{})
}
