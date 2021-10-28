


package v20190801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualRouterPeering struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput     `pulumi:"etag"`
	Name              pulumi.StringPtrOutput  `pulumi:"name"`
	PeerAsn           pulumi.Float64PtrOutput `pulumi:"peerAsn"`
	PeerIp            pulumi.StringPtrOutput  `pulumi:"peerIp"`
	ProvisioningState pulumi.StringOutput     `pulumi:"provisioningState"`
	Type              pulumi.StringOutput     `pulumi:"type"`
}


func NewVirtualRouterPeering(ctx *pulumi.Context,
	name string, args *VirtualRouterPeeringArgs, opts ...pulumi.ResourceOption) (*VirtualRouterPeering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualRouterName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualRouterName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualRouterPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210501:VirtualRouterPeering"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualRouterPeering
	err := ctx.RegisterResource("azure-native:network/v20190801:VirtualRouterPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualRouterPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualRouterPeeringState, opts ...pulumi.ResourceOption) (*VirtualRouterPeering, error) {
	var resource VirtualRouterPeering
	err := ctx.ReadResource("azure-native:network/v20190801:VirtualRouterPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualRouterPeeringState struct {
}

type VirtualRouterPeeringState struct {
}

func (VirtualRouterPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRouterPeeringState)(nil)).Elem()
}

type virtualRouterPeeringArgs struct {
	Id                *string  `pulumi:"id"`
	Name              *string  `pulumi:"name"`
	PeerAsn           *float64 `pulumi:"peerAsn"`
	PeerIp            *string  `pulumi:"peerIp"`
	PeeringName       *string  `pulumi:"peeringName"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	VirtualRouterName string   `pulumi:"virtualRouterName"`
}


type VirtualRouterPeeringArgs struct {
	Id                pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	PeerAsn           pulumi.Float64PtrInput
	PeerIp            pulumi.StringPtrInput
	PeeringName       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	VirtualRouterName pulumi.StringInput
}

func (VirtualRouterPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRouterPeeringArgs)(nil)).Elem()
}

type VirtualRouterPeeringInput interface {
	pulumi.Input

	ToVirtualRouterPeeringOutput() VirtualRouterPeeringOutput
	ToVirtualRouterPeeringOutputWithContext(ctx context.Context) VirtualRouterPeeringOutput
}

func (*VirtualRouterPeering) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualRouterPeering)(nil))
}

func (i *VirtualRouterPeering) ToVirtualRouterPeeringOutput() VirtualRouterPeeringOutput {
	return i.ToVirtualRouterPeeringOutputWithContext(context.Background())
}

func (i *VirtualRouterPeering) ToVirtualRouterPeeringOutputWithContext(ctx context.Context) VirtualRouterPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualRouterPeeringOutput)
}

type VirtualRouterPeeringOutput struct{ *pulumi.OutputState }

func (VirtualRouterPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualRouterPeering)(nil))
}

func (o VirtualRouterPeeringOutput) ToVirtualRouterPeeringOutput() VirtualRouterPeeringOutput {
	return o
}

func (o VirtualRouterPeeringOutput) ToVirtualRouterPeeringOutputWithContext(ctx context.Context) VirtualRouterPeeringOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualRouterPeeringOutput{})
}
