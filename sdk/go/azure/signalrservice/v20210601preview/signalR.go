


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SignalR struct {
	pulumi.CustomResourceState

	Cors                       SignalRCorsSettingsResponsePtrOutput         `pulumi:"cors"`
	DisableAadAuth             pulumi.BoolPtrOutput                         `pulumi:"disableAadAuth"`
	DisableLocalAuth           pulumi.BoolPtrOutput                         `pulumi:"disableLocalAuth"`
	ExternalIP                 pulumi.StringOutput                          `pulumi:"externalIP"`
	Features                   SignalRFeatureResponseArrayOutput            `pulumi:"features"`
	HostName                   pulumi.StringOutput                          `pulumi:"hostName"`
	Identity                   ManagedIdentityResponsePtrOutput             `pulumi:"identity"`
	Kind                       pulumi.StringPtrOutput                       `pulumi:"kind"`
	Location                   pulumi.StringPtrOutput                       `pulumi:"location"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	NetworkACLs                SignalRNetworkACLsResponsePtrOutput          `pulumi:"networkACLs"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                          `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrOutput                       `pulumi:"publicNetworkAccess"`
	PublicPort                 pulumi.IntOutput                             `pulumi:"publicPort"`
	ServerPort                 pulumi.IntOutput                             `pulumi:"serverPort"`
	SharedPrivateLinkResources SharedPrivateLinkResourceResponseArrayOutput `pulumi:"sharedPrivateLinkResources"`
	Sku                        ResourceSkuResponsePtrOutput                 `pulumi:"sku"`
	SystemData                 SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	Tls                        SignalRTlsSettingsResponsePtrOutput          `pulumi:"tls"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
	Upstream                   ServerlessUpstreamSettingsResponsePtrOutput  `pulumi:"upstream"`
	Version                    pulumi.StringOutput                          `pulumi:"version"`
}


func NewSignalR(ctx *pulumi.Context,
	name string, args *SignalRArgs, opts ...pulumi.ResourceOption) (*SignalR, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.DisableAadAuth) {
		args.DisableAadAuth = pulumi.BoolPtr(false)
	}
	if isZero(args.DisableLocalAuth) {
		args.DisableLocalAuth = pulumi.BoolPtr(false)
	}
	if args.NetworkACLs != nil {
		args.NetworkACLs = args.NetworkACLs.ToSignalRNetworkACLsPtrOutput().ApplyT(func(v *SignalRNetworkACLs) *SignalRNetworkACLs { return v.Defaults() }).(SignalRNetworkACLsPtrOutput)
	}
	if isZero(args.PublicNetworkAccess) {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	if args.Tls != nil {
		args.Tls = args.Tls.ToSignalRTlsSettingsPtrOutput().ApplyT(func(v *SignalRTlsSettings) *SignalRTlsSettings { return v.Defaults() }).(SignalRTlsSettingsPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:signalrservice:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20180301preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20181001:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20200501:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20200701preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210401preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210901preview:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20211001:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220201:SignalR"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220801preview:SignalR"),
		},
	})
	opts = append(opts, aliases)
	var resource SignalR
	err := ctx.RegisterResource("azure-native:signalrservice/v20210601preview:SignalR", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSignalR(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignalRState, opts ...pulumi.ResourceOption) (*SignalR, error) {
	var resource SignalR
	err := ctx.ReadResource("azure-native:signalrservice/v20210601preview:SignalR", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type signalRState struct {
}

type SignalRState struct {
}

func (SignalRState) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRState)(nil)).Elem()
}

type signalRArgs struct {
	Cors                *SignalRCorsSettings        `pulumi:"cors"`
	DisableAadAuth      *bool                       `pulumi:"disableAadAuth"`
	DisableLocalAuth    *bool                       `pulumi:"disableLocalAuth"`
	Features            []SignalRFeature            `pulumi:"features"`
	Identity            *ManagedIdentity            `pulumi:"identity"`
	Kind                *string                     `pulumi:"kind"`
	Location            *string                     `pulumi:"location"`
	NetworkACLs         *SignalRNetworkACLs         `pulumi:"networkACLs"`
	PublicNetworkAccess *string                     `pulumi:"publicNetworkAccess"`
	ResourceGroupName   string                      `pulumi:"resourceGroupName"`
	ResourceName        *string                     `pulumi:"resourceName"`
	Sku                 *ResourceSku                `pulumi:"sku"`
	Tags                map[string]string           `pulumi:"tags"`
	Tls                 *SignalRTlsSettings         `pulumi:"tls"`
	Upstream            *ServerlessUpstreamSettings `pulumi:"upstream"`
}


type SignalRArgs struct {
	Cors                SignalRCorsSettingsPtrInput
	DisableAadAuth      pulumi.BoolPtrInput
	DisableLocalAuth    pulumi.BoolPtrInput
	Features            SignalRFeatureArrayInput
	Identity            ManagedIdentityPtrInput
	Kind                pulumi.StringPtrInput
	Location            pulumi.StringPtrInput
	NetworkACLs         SignalRNetworkACLsPtrInput
	PublicNetworkAccess pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	ResourceName        pulumi.StringPtrInput
	Sku                 ResourceSkuPtrInput
	Tags                pulumi.StringMapInput
	Tls                 SignalRTlsSettingsPtrInput
	Upstream            ServerlessUpstreamSettingsPtrInput
}

func (SignalRArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRArgs)(nil)).Elem()
}

type SignalRInput interface {
	pulumi.Input

	ToSignalROutput() SignalROutput
	ToSignalROutputWithContext(ctx context.Context) SignalROutput
}

func (*SignalR) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalR)(nil)).Elem()
}

func (i *SignalR) ToSignalROutput() SignalROutput {
	return i.ToSignalROutputWithContext(context.Background())
}

func (i *SignalR) ToSignalROutputWithContext(ctx context.Context) SignalROutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalROutput)
}

type SignalROutput struct{ *pulumi.OutputState }

func (SignalROutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalR)(nil)).Elem()
}

func (o SignalROutput) ToSignalROutput() SignalROutput {
	return o
}

func (o SignalROutput) ToSignalROutputWithContext(ctx context.Context) SignalROutput {
	return o
}

func (o SignalROutput) Cors() SignalRCorsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) SignalRCorsSettingsResponsePtrOutput { return v.Cors }).(SignalRCorsSettingsResponsePtrOutput)
}

func (o SignalROutput) DisableAadAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.BoolPtrOutput { return v.DisableAadAuth }).(pulumi.BoolPtrOutput)
}

func (o SignalROutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.BoolPtrOutput { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o SignalROutput) ExternalIP() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.ExternalIP }).(pulumi.StringOutput)
}

func (o SignalROutput) Features() SignalRFeatureResponseArrayOutput {
	return o.ApplyT(func(v *SignalR) SignalRFeatureResponseArrayOutput { return v.Features }).(SignalRFeatureResponseArrayOutput)
}

func (o SignalROutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

func (o SignalROutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) ManagedIdentityResponsePtrOutput { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

func (o SignalROutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SignalROutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o SignalROutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SignalROutput) NetworkACLs() SignalRNetworkACLsResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) SignalRNetworkACLsResponsePtrOutput { return v.NetworkACLs }).(SignalRNetworkACLsResponsePtrOutput)
}

func (o SignalROutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *SignalR) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o SignalROutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SignalROutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o SignalROutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SignalR) pulumi.IntOutput { return v.PublicPort }).(pulumi.IntOutput)
}

func (o SignalROutput) ServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SignalR) pulumi.IntOutput { return v.ServerPort }).(pulumi.IntOutput)
}

func (o SignalROutput) SharedPrivateLinkResources() SharedPrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v *SignalR) SharedPrivateLinkResourceResponseArrayOutput { return v.SharedPrivateLinkResources }).(SharedPrivateLinkResourceResponseArrayOutput)
}

func (o SignalROutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) ResourceSkuResponsePtrOutput { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

func (o SignalROutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SignalR) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SignalROutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SignalROutput) Tls() SignalRTlsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) SignalRTlsSettingsResponsePtrOutput { return v.Tls }).(SignalRTlsSettingsResponsePtrOutput)
}

func (o SignalROutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SignalROutput) Upstream() ServerlessUpstreamSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SignalR) ServerlessUpstreamSettingsResponsePtrOutput { return v.Upstream }).(ServerlessUpstreamSettingsResponsePtrOutput)
}

func (o SignalROutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalR) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SignalROutput{})
}
