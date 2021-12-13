


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

func init() {
	pulumi.RegisterOutputType(SignalROutput{})
}
