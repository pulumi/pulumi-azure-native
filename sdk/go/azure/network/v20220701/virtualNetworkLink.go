


package v20220701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetworkLink struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput       `pulumi:"etag"`
	Metadata          pulumi.StringMapOutput    `pulumi:"metadata"`
	Name              pulumi.StringOutput       `pulumi:"name"`
	ProvisioningState pulumi.StringOutput       `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput  `pulumi:"systemData"`
	Type              pulumi.StringOutput       `pulumi:"type"`
	VirtualNetwork    SubResourceResponseOutput `pulumi:"virtualNetwork"`
}


func NewVirtualNetworkLink(ctx *pulumi.Context,
	name string, args *VirtualNetworkLinkArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DnsForwardingRulesetName == nil {
		return nil, errors.New("invalid value for required argument 'DnsForwardingRulesetName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualNetwork == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetwork'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20200401preview:VirtualNetworkLink"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkLink
	err := ctx.RegisterResource("azure-native:network/v20220701:VirtualNetworkLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetworkLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkLinkState, opts ...pulumi.ResourceOption) (*VirtualNetworkLink, error) {
	var resource VirtualNetworkLink
	err := ctx.ReadResource("azure-native:network/v20220701:VirtualNetworkLink", name, id, state, &resource, opts...)
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
	DnsForwardingRulesetName string            `pulumi:"dnsForwardingRulesetName"`
	Metadata                 map[string]string `pulumi:"metadata"`
	ResourceGroupName        string            `pulumi:"resourceGroupName"`
	VirtualNetwork           SubResource       `pulumi:"virtualNetwork"`
	VirtualNetworkLinkName   *string           `pulumi:"virtualNetworkLinkName"`
}


type VirtualNetworkLinkArgs struct {
	DnsForwardingRulesetName pulumi.StringInput
	Metadata                 pulumi.StringMapInput
	ResourceGroupName        pulumi.StringInput
	VirtualNetwork           SubResourceInput
	VirtualNetworkLinkName   pulumi.StringPtrInput
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

func (o VirtualNetworkLinkOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualNetworkLinkOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkLinkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkLinkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VirtualNetworkLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualNetworkLinkOutput) VirtualNetwork() SubResourceResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkLink) SubResourceResponseOutput { return v.VirtualNetwork }).(SubResourceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkLinkOutput{})
}
