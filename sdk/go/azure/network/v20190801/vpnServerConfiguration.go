


package v20190801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpnServerConfiguration struct {
	pulumi.CustomResourceState

	AadAuthenticationParameters  AadAuthenticationParametersResponsePtrOutput                  `pulumi:"aadAuthenticationParameters"`
	Etag                         pulumi.StringOutput                                           `pulumi:"etag"`
	Location                     pulumi.StringPtrOutput                                        `pulumi:"location"`
	Name                         pulumi.StringOutput                                           `pulumi:"name"`
	P2SVpnGateways               P2SVpnGatewayResponseArrayOutput                              `pulumi:"p2SVpnGateways"`
	ProvisioningState            pulumi.StringOutput                                           `pulumi:"provisioningState"`
	RadiusClientRootCertificates VpnServerConfigRadiusClientRootCertificateResponseArrayOutput `pulumi:"radiusClientRootCertificates"`
	RadiusServerAddress          pulumi.StringPtrOutput                                        `pulumi:"radiusServerAddress"`
	RadiusServerRootCertificates VpnServerConfigRadiusServerRootCertificateResponseArrayOutput `pulumi:"radiusServerRootCertificates"`
	RadiusServerSecret           pulumi.StringPtrOutput                                        `pulumi:"radiusServerSecret"`
	Tags                         pulumi.StringMapOutput                                        `pulumi:"tags"`
	Type                         pulumi.StringOutput                                           `pulumi:"type"`
	VpnAuthenticationTypes       pulumi.StringArrayOutput                                      `pulumi:"vpnAuthenticationTypes"`
	VpnClientIpsecPolicies       IpsecPolicyResponseArrayOutput                                `pulumi:"vpnClientIpsecPolicies"`
	VpnClientRevokedCertificates VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates    VpnServerConfigVpnClientRootCertificateResponseArrayOutput    `pulumi:"vpnClientRootCertificates"`
	VpnProtocols                 pulumi.StringArrayOutput                                      `pulumi:"vpnProtocols"`
}


func NewVpnServerConfiguration(ctx *pulumi.Context,
	name string, args *VpnServerConfigurationArgs, opts ...pulumi.ResourceOption) (*VpnServerConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VpnServerConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource VpnServerConfiguration
	err := ctx.RegisterResource("azure-native:network/v20190801:VpnServerConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVpnServerConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnServerConfigurationState, opts ...pulumi.ResourceOption) (*VpnServerConfiguration, error) {
	var resource VpnServerConfiguration
	err := ctx.ReadResource("azure-native:network/v20190801:VpnServerConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type vpnServerConfigurationState struct {
}

type VpnServerConfigurationState struct {
}

func (VpnServerConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnServerConfigurationState)(nil)).Elem()
}

type vpnServerConfigurationArgs struct {
	AadAuthenticationParameters  *AadAuthenticationParameters                 `pulumi:"aadAuthenticationParameters"`
	Id                           *string                                      `pulumi:"id"`
	Location                     *string                                      `pulumi:"location"`
	Name                         *string                                      `pulumi:"name"`
	RadiusClientRootCertificates []VpnServerConfigRadiusClientRootCertificate `pulumi:"radiusClientRootCertificates"`
	RadiusServerAddress          *string                                      `pulumi:"radiusServerAddress"`
	RadiusServerRootCertificates []VpnServerConfigRadiusServerRootCertificate `pulumi:"radiusServerRootCertificates"`
	RadiusServerSecret           *string                                      `pulumi:"radiusServerSecret"`
	ResourceGroupName            string                                       `pulumi:"resourceGroupName"`
	Tags                         map[string]string                            `pulumi:"tags"`
	VpnAuthenticationTypes       []string                                     `pulumi:"vpnAuthenticationTypes"`
	VpnClientIpsecPolicies       []IpsecPolicy                                `pulumi:"vpnClientIpsecPolicies"`
	VpnClientRevokedCertificates []VpnServerConfigVpnClientRevokedCertificate `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates    []VpnServerConfigVpnClientRootCertificate    `pulumi:"vpnClientRootCertificates"`
	VpnProtocols                 []string                                     `pulumi:"vpnProtocols"`
	VpnServerConfigurationName   *string                                      `pulumi:"vpnServerConfigurationName"`
}


type VpnServerConfigurationArgs struct {
	AadAuthenticationParameters  AadAuthenticationParametersPtrInput
	Id                           pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	Name                         pulumi.StringPtrInput
	RadiusClientRootCertificates VpnServerConfigRadiusClientRootCertificateArrayInput
	RadiusServerAddress          pulumi.StringPtrInput
	RadiusServerRootCertificates VpnServerConfigRadiusServerRootCertificateArrayInput
	RadiusServerSecret           pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	Tags                         pulumi.StringMapInput
	VpnAuthenticationTypes       pulumi.StringArrayInput
	VpnClientIpsecPolicies       IpsecPolicyArrayInput
	VpnClientRevokedCertificates VpnServerConfigVpnClientRevokedCertificateArrayInput
	VpnClientRootCertificates    VpnServerConfigVpnClientRootCertificateArrayInput
	VpnProtocols                 pulumi.StringArrayInput
	VpnServerConfigurationName   pulumi.StringPtrInput
}

func (VpnServerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnServerConfigurationArgs)(nil)).Elem()
}

type VpnServerConfigurationInput interface {
	pulumi.Input

	ToVpnServerConfigurationOutput() VpnServerConfigurationOutput
	ToVpnServerConfigurationOutputWithContext(ctx context.Context) VpnServerConfigurationOutput
}

func (*VpnServerConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnServerConfiguration)(nil)).Elem()
}

func (i *VpnServerConfiguration) ToVpnServerConfigurationOutput() VpnServerConfigurationOutput {
	return i.ToVpnServerConfigurationOutputWithContext(context.Background())
}

func (i *VpnServerConfiguration) ToVpnServerConfigurationOutputWithContext(ctx context.Context) VpnServerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationOutput)
}

type VpnServerConfigurationOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnServerConfiguration)(nil)).Elem()
}

func (o VpnServerConfigurationOutput) ToVpnServerConfigurationOutput() VpnServerConfigurationOutput {
	return o
}

func (o VpnServerConfigurationOutput) ToVpnServerConfigurationOutputWithContext(ctx context.Context) VpnServerConfigurationOutput {
	return o
}

func (o VpnServerConfigurationOutput) AadAuthenticationParameters() AadAuthenticationParametersResponsePtrOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) AadAuthenticationParametersResponsePtrOutput {
		return v.AadAuthenticationParameters
	}).(AadAuthenticationParametersResponsePtrOutput)
}

func (o VpnServerConfigurationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnServerConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VpnServerConfigurationOutput) P2SVpnGateways() P2SVpnGatewayResponseArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) P2SVpnGatewayResponseArrayOutput { return v.P2SVpnGateways }).(P2SVpnGatewayResponseArrayOutput)
}

func (o VpnServerConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnServerConfigurationOutput) RadiusClientRootCertificates() VpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) VpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
		return v.RadiusClientRootCertificates
	}).(VpnServerConfigRadiusClientRootCertificateResponseArrayOutput)
}

func (o VpnServerConfigurationOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringPtrOutput { return v.RadiusServerAddress }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationOutput) RadiusServerRootCertificates() VpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) VpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
		return v.RadiusServerRootCertificates
	}).(VpnServerConfigRadiusServerRootCertificateResponseArrayOutput)
}

func (o VpnServerConfigurationOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringPtrOutput { return v.RadiusServerSecret }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VpnServerConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VpnServerConfigurationOutput) VpnAuthenticationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringArrayOutput { return v.VpnAuthenticationTypes }).(pulumi.StringArrayOutput)
}

func (o VpnServerConfigurationOutput) VpnClientIpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) IpsecPolicyResponseArrayOutput { return v.VpnClientIpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

func (o VpnServerConfigurationOutput) VpnClientRevokedCertificates() VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
		return v.VpnClientRevokedCertificates
	}).(VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput)
}

func (o VpnServerConfigurationOutput) VpnClientRootCertificates() VpnServerConfigVpnClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) VpnServerConfigVpnClientRootCertificateResponseArrayOutput {
		return v.VpnClientRootCertificates
	}).(VpnServerConfigVpnClientRootCertificateResponseArrayOutput)
}

func (o VpnServerConfigurationOutput) VpnProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnServerConfiguration) pulumi.StringArrayOutput { return v.VpnProtocols }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VpnServerConfigurationOutput{})
}
