


package v20220601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PeerAsn struct {
	pulumi.CustomResourceState

	ErrorMessage      pulumi.StringOutput              `pulumi:"errorMessage"`
	Name              pulumi.StringOutput              `pulumi:"name"`
	PeerAsn           pulumi.IntPtrOutput              `pulumi:"peerAsn"`
	PeerContactDetail ContactDetailResponseArrayOutput `pulumi:"peerContactDetail"`
	PeerName          pulumi.StringPtrOutput           `pulumi:"peerName"`
	Type              pulumi.StringOutput              `pulumi:"type"`
	ValidationState   pulumi.StringOutput              `pulumi:"validationState"`
}


func NewPeerAsn(ctx *pulumi.Context,
	name string, args *PeerAsnArgs, opts ...pulumi.ResourceOption) (*PeerAsn, error) {
	if args == nil {
		args = &PeerAsnArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:peering:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20190801preview:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20190901preview:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200101preview:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200401:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20201001:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210101:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210601:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220101:PeerAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20221001:PeerAsn"),
		},
	})
	opts = append(opts, aliases)
	var resource PeerAsn
	err := ctx.RegisterResource("azure-native:peering/v20220601:PeerAsn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPeerAsn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeerAsnState, opts ...pulumi.ResourceOption) (*PeerAsn, error) {
	var resource PeerAsn
	err := ctx.ReadResource("azure-native:peering/v20220601:PeerAsn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type peerAsnState struct {
}

type PeerAsnState struct {
}

func (PeerAsnState) ElementType() reflect.Type {
	return reflect.TypeOf((*peerAsnState)(nil)).Elem()
}

type peerAsnArgs struct {
	PeerAsn           *int            `pulumi:"peerAsn"`
	PeerAsnName       *string         `pulumi:"peerAsnName"`
	PeerContactDetail []ContactDetail `pulumi:"peerContactDetail"`
	PeerName          *string         `pulumi:"peerName"`
}


type PeerAsnArgs struct {
	PeerAsn           pulumi.IntPtrInput
	PeerAsnName       pulumi.StringPtrInput
	PeerContactDetail ContactDetailArrayInput
	PeerName          pulumi.StringPtrInput
}

func (PeerAsnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peerAsnArgs)(nil)).Elem()
}

type PeerAsnInput interface {
	pulumi.Input

	ToPeerAsnOutput() PeerAsnOutput
	ToPeerAsnOutputWithContext(ctx context.Context) PeerAsnOutput
}

func (*PeerAsn) ElementType() reflect.Type {
	return reflect.TypeOf((**PeerAsn)(nil)).Elem()
}

func (i *PeerAsn) ToPeerAsnOutput() PeerAsnOutput {
	return i.ToPeerAsnOutputWithContext(context.Background())
}

func (i *PeerAsn) ToPeerAsnOutputWithContext(ctx context.Context) PeerAsnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeerAsnOutput)
}

type PeerAsnOutput struct{ *pulumi.OutputState }

func (PeerAsnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeerAsn)(nil)).Elem()
}

func (o PeerAsnOutput) ToPeerAsnOutput() PeerAsnOutput {
	return o
}

func (o PeerAsnOutput) ToPeerAsnOutputWithContext(ctx context.Context) PeerAsnOutput {
	return o
}

func (o PeerAsnOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerAsn) pulumi.StringOutput { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o PeerAsnOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerAsn) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PeerAsnOutput) PeerAsn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PeerAsn) pulumi.IntPtrOutput { return v.PeerAsn }).(pulumi.IntPtrOutput)
}

func (o PeerAsnOutput) PeerContactDetail() ContactDetailResponseArrayOutput {
	return o.ApplyT(func(v *PeerAsn) ContactDetailResponseArrayOutput { return v.PeerContactDetail }).(ContactDetailResponseArrayOutput)
}

func (o PeerAsnOutput) PeerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeerAsn) pulumi.StringPtrOutput { return v.PeerName }).(pulumi.StringPtrOutput)
}

func (o PeerAsnOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerAsn) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PeerAsnOutput) ValidationState() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerAsn) pulumi.StringOutput { return v.ValidationState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PeerAsnOutput{})
}
