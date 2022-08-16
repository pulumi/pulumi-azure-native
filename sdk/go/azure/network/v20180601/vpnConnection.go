


package v20180601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpnConnection struct {
	pulumi.CustomResourceState

	ConnectionBandwidthInMbps pulumi.IntOutput               `pulumi:"connectionBandwidthInMbps"`
	ConnectionStatus          pulumi.StringOutput            `pulumi:"connectionStatus"`
	EgressBytesTransferred    pulumi.Float64Output           `pulumi:"egressBytesTransferred"`
	EnableBgp                 pulumi.BoolPtrOutput           `pulumi:"enableBgp"`
	Etag                      pulumi.StringOutput            `pulumi:"etag"`
	IngressBytesTransferred   pulumi.Float64Output           `pulumi:"ingressBytesTransferred"`
	IpsecPolicies             IpsecPolicyResponseArrayOutput `pulumi:"ipsecPolicies"`
	Name                      pulumi.StringPtrOutput         `pulumi:"name"`
	ProvisioningState         pulumi.StringOutput            `pulumi:"provisioningState"`
	RemoteVpnSite             SubResourceResponsePtrOutput   `pulumi:"remoteVpnSite"`
	RoutingWeight             pulumi.IntPtrOutput            `pulumi:"routingWeight"`
	SharedKey                 pulumi.StringPtrOutput         `pulumi:"sharedKey"`
}


func NewVpnConnection(ctx *pulumi.Context,
	name string, args *VpnConnectionArgs, opts ...pulumi.ResourceOption) (*VpnConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayName == nil {
		return nil, errors.New("invalid value for required argument 'GatewayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VpnConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VpnConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource VpnConnection
	err := ctx.RegisterResource("azure-native:network/v20180601:VpnConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVpnConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnConnectionState, opts ...pulumi.ResourceOption) (*VpnConnection, error) {
	var resource VpnConnection
	err := ctx.ReadResource("azure-native:network/v20180601:VpnConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type vpnConnectionState struct {
}

type VpnConnectionState struct {
}

func (VpnConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnConnectionState)(nil)).Elem()
}

type vpnConnectionArgs struct {
	ConnectionName    *string       `pulumi:"connectionName"`
	EnableBgp         *bool         `pulumi:"enableBgp"`
	GatewayName       string        `pulumi:"gatewayName"`
	Id                *string       `pulumi:"id"`
	IpsecPolicies     []IpsecPolicy `pulumi:"ipsecPolicies"`
	Name              *string       `pulumi:"name"`
	RemoteVpnSite     *SubResource  `pulumi:"remoteVpnSite"`
	ResourceGroupName string        `pulumi:"resourceGroupName"`
	RoutingWeight     *int          `pulumi:"routingWeight"`
	SharedKey         *string       `pulumi:"sharedKey"`
}


type VpnConnectionArgs struct {
	ConnectionName    pulumi.StringPtrInput
	EnableBgp         pulumi.BoolPtrInput
	GatewayName       pulumi.StringInput
	Id                pulumi.StringPtrInput
	IpsecPolicies     IpsecPolicyArrayInput
	Name              pulumi.StringPtrInput
	RemoteVpnSite     SubResourcePtrInput
	ResourceGroupName pulumi.StringInput
	RoutingWeight     pulumi.IntPtrInput
	SharedKey         pulumi.StringPtrInput
}

func (VpnConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnConnectionArgs)(nil)).Elem()
}

type VpnConnectionInput interface {
	pulumi.Input

	ToVpnConnectionOutput() VpnConnectionOutput
	ToVpnConnectionOutputWithContext(ctx context.Context) VpnConnectionOutput
}

func (*VpnConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnConnection)(nil)).Elem()
}

func (i *VpnConnection) ToVpnConnectionOutput() VpnConnectionOutput {
	return i.ToVpnConnectionOutputWithContext(context.Background())
}

func (i *VpnConnection) ToVpnConnectionOutputWithContext(ctx context.Context) VpnConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnConnectionOutput)
}

type VpnConnectionOutput struct{ *pulumi.OutputState }

func (VpnConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnConnection)(nil)).Elem()
}

func (o VpnConnectionOutput) ToVpnConnectionOutput() VpnConnectionOutput {
	return o
}

func (o VpnConnectionOutput) ToVpnConnectionOutputWithContext(ctx context.Context) VpnConnectionOutput {
	return o
}

func (o VpnConnectionOutput) ConnectionBandwidthInMbps() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnConnection) pulumi.IntOutput { return v.ConnectionBandwidthInMbps }).(pulumi.IntOutput)
}

func (o VpnConnectionOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnConnection) pulumi.StringOutput { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o VpnConnectionOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v *VpnConnection) pulumi.Float64Output { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnConnectionOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpnConnection) pulumi.BoolPtrOutput { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnConnection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnConnectionOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v *VpnConnection) pulumi.Float64Output { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnConnectionOutput) IpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v *VpnConnection) IpsecPolicyResponseArrayOutput { return v.IpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

func (o VpnConnectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnConnection) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnConnectionOutput) RemoteVpnSite() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VpnConnection) SubResourceResponsePtrOutput { return v.RemoteVpnSite }).(SubResourceResponsePtrOutput)
}

func (o VpnConnectionOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpnConnection) pulumi.IntPtrOutput { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func (o VpnConnectionOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnConnection) pulumi.StringPtrOutput { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VpnConnectionOutput{})
}
