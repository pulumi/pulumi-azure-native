


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualHubBgpConnection struct {
	pulumi.CustomResourceState

	ConnectionState             pulumi.StringOutput          `pulumi:"connectionState"`
	Etag                        pulumi.StringOutput          `pulumi:"etag"`
	HubVirtualNetworkConnection SubResourceResponsePtrOutput `pulumi:"hubVirtualNetworkConnection"`
	Name                        pulumi.StringPtrOutput       `pulumi:"name"`
	PeerAsn                     pulumi.Float64PtrOutput      `pulumi:"peerAsn"`
	PeerIp                      pulumi.StringPtrOutput       `pulumi:"peerIp"`
	ProvisioningState           pulumi.StringOutput          `pulumi:"provisioningState"`
	Type                        pulumi.StringOutput          `pulumi:"type"`
}


func NewVirtualHubBgpConnection(ctx *pulumi.Context,
	name string, args *VirtualHubBgpConnectionArgs, opts ...pulumi.ResourceOption) (*VirtualHubBgpConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualHubName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualHubBgpConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualHubBgpConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualHubBgpConnection
	err := ctx.RegisterResource("azure-native:network/v20210801:VirtualHubBgpConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualHubBgpConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubBgpConnectionState, opts ...pulumi.ResourceOption) (*VirtualHubBgpConnection, error) {
	var resource VirtualHubBgpConnection
	err := ctx.ReadResource("azure-native:network/v20210801:VirtualHubBgpConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualHubBgpConnectionState struct {
}

type VirtualHubBgpConnectionState struct {
}

func (VirtualHubBgpConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubBgpConnectionState)(nil)).Elem()
}

type virtualHubBgpConnectionArgs struct {
	ConnectionName              *string      `pulumi:"connectionName"`
	HubVirtualNetworkConnection *SubResource `pulumi:"hubVirtualNetworkConnection"`
	Id                          *string      `pulumi:"id"`
	Name                        *string      `pulumi:"name"`
	PeerAsn                     *float64     `pulumi:"peerAsn"`
	PeerIp                      *string      `pulumi:"peerIp"`
	ResourceGroupName           string       `pulumi:"resourceGroupName"`
	VirtualHubName              string       `pulumi:"virtualHubName"`
}


type VirtualHubBgpConnectionArgs struct {
	ConnectionName              pulumi.StringPtrInput
	HubVirtualNetworkConnection SubResourcePtrInput
	Id                          pulumi.StringPtrInput
	Name                        pulumi.StringPtrInput
	PeerAsn                     pulumi.Float64PtrInput
	PeerIp                      pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	VirtualHubName              pulumi.StringInput
}

func (VirtualHubBgpConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubBgpConnectionArgs)(nil)).Elem()
}

type VirtualHubBgpConnectionInput interface {
	pulumi.Input

	ToVirtualHubBgpConnectionOutput() VirtualHubBgpConnectionOutput
	ToVirtualHubBgpConnectionOutputWithContext(ctx context.Context) VirtualHubBgpConnectionOutput
}

func (*VirtualHubBgpConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubBgpConnection)(nil)).Elem()
}

func (i *VirtualHubBgpConnection) ToVirtualHubBgpConnectionOutput() VirtualHubBgpConnectionOutput {
	return i.ToVirtualHubBgpConnectionOutputWithContext(context.Background())
}

func (i *VirtualHubBgpConnection) ToVirtualHubBgpConnectionOutputWithContext(ctx context.Context) VirtualHubBgpConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubBgpConnectionOutput)
}

type VirtualHubBgpConnectionOutput struct{ *pulumi.OutputState }

func (VirtualHubBgpConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubBgpConnection)(nil)).Elem()
}

func (o VirtualHubBgpConnectionOutput) ToVirtualHubBgpConnectionOutput() VirtualHubBgpConnectionOutput {
	return o
}

func (o VirtualHubBgpConnectionOutput) ToVirtualHubBgpConnectionOutputWithContext(ctx context.Context) VirtualHubBgpConnectionOutput {
	return o
}

func (o VirtualHubBgpConnectionOutput) ConnectionState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) pulumi.StringOutput { return v.ConnectionState }).(pulumi.StringOutput)
}

func (o VirtualHubBgpConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualHubBgpConnectionOutput) HubVirtualNetworkConnection() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) SubResourceResponsePtrOutput { return v.HubVirtualNetworkConnection }).(SubResourceResponsePtrOutput)
}

func (o VirtualHubBgpConnectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualHubBgpConnectionOutput) PeerAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) pulumi.Float64PtrOutput { return v.PeerAsn }).(pulumi.Float64PtrOutput)
}

func (o VirtualHubBgpConnectionOutput) PeerIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) pulumi.StringPtrOutput { return v.PeerIp }).(pulumi.StringPtrOutput)
}

func (o VirtualHubBgpConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualHubBgpConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHubBgpConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualHubBgpConnectionOutput{})
}
