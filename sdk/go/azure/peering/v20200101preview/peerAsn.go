


package v20200101preview

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
	ValidationState   pulumi.StringPtrOutput           `pulumi:"validationState"`
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
	})
	opts = append(opts, aliases)
	var resource PeerAsn
	err := ctx.RegisterResource("azure-native:peering/v20200101preview:PeerAsn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPeerAsn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeerAsnState, opts ...pulumi.ResourceOption) (*PeerAsn, error) {
	var resource PeerAsn
	err := ctx.ReadResource("azure-native:peering/v20200101preview:PeerAsn", name, id, state, &resource, opts...)
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
	ValidationState   *string         `pulumi:"validationState"`
}


type PeerAsnArgs struct {
	PeerAsn           pulumi.IntPtrInput
	PeerAsnName       pulumi.StringPtrInput
	PeerContactDetail ContactDetailArrayInput
	PeerName          pulumi.StringPtrInput
	ValidationState   pulumi.StringPtrInput
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
	return reflect.TypeOf((*PeerAsn)(nil))
}

func (i *PeerAsn) ToPeerAsnOutput() PeerAsnOutput {
	return i.ToPeerAsnOutputWithContext(context.Background())
}

func (i *PeerAsn) ToPeerAsnOutputWithContext(ctx context.Context) PeerAsnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeerAsnOutput)
}

type PeerAsnOutput struct{ *pulumi.OutputState }

func (PeerAsnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeerAsn)(nil))
}

func (o PeerAsnOutput) ToPeerAsnOutput() PeerAsnOutput {
	return o
}

func (o PeerAsnOutput) ToPeerAsnOutputWithContext(ctx context.Context) PeerAsnOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PeerAsnOutput{})
}
