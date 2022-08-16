


package v20181101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type P2sVpnServerConfiguration struct {
	pulumi.CustomResourceState

	Etag                                           pulumi.StringOutput                                              `pulumi:"etag"`
	Name                                           pulumi.StringPtrOutput                                           `pulumi:"name"`
	P2SVpnGateways                                 SubResourceResponseArrayOutput                                   `pulumi:"p2SVpnGateways"`
	P2SVpnServerConfigRadiusClientRootCertificates P2SVpnServerConfigRadiusClientRootCertificateResponseArrayOutput `pulumi:"p2SVpnServerConfigRadiusClientRootCertificates"`
	P2SVpnServerConfigRadiusServerRootCertificates P2SVpnServerConfigRadiusServerRootCertificateResponseArrayOutput `pulumi:"p2SVpnServerConfigRadiusServerRootCertificates"`
	P2SVpnServerConfigVpnClientRevokedCertificates P2SVpnServerConfigVpnClientRevokedCertificateResponseArrayOutput `pulumi:"p2SVpnServerConfigVpnClientRevokedCertificates"`
	P2SVpnServerConfigVpnClientRootCertificates    P2SVpnServerConfigVpnClientRootCertificateResponseArrayOutput    `pulumi:"p2SVpnServerConfigVpnClientRootCertificates"`
	ProvisioningState                              pulumi.StringOutput                                              `pulumi:"provisioningState"`
	RadiusServerAddress                            pulumi.StringPtrOutput                                           `pulumi:"radiusServerAddress"`
	RadiusServerSecret                             pulumi.StringPtrOutput                                           `pulumi:"radiusServerSecret"`
	VpnClientIpsecPolicies                         IpsecPolicyResponseArrayOutput                                   `pulumi:"vpnClientIpsecPolicies"`
	VpnProtocols                                   pulumi.StringArrayOutput                                         `pulumi:"vpnProtocols"`
}


func NewP2sVpnServerConfiguration(ctx *pulumi.Context,
	name string, args *P2sVpnServerConfigurationArgs, opts ...pulumi.ResourceOption) (*P2sVpnServerConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualWanName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualWanName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:P2sVpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:P2sVpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:P2sVpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:P2sVpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:P2sVpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:P2sVpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:P2sVpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:P2sVpnServerConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource P2sVpnServerConfiguration
	err := ctx.RegisterResource("azure-native:network/v20181101:P2sVpnServerConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetP2sVpnServerConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *P2sVpnServerConfigurationState, opts ...pulumi.ResourceOption) (*P2sVpnServerConfiguration, error) {
	var resource P2sVpnServerConfiguration
	err := ctx.ReadResource("azure-native:network/v20181101:P2sVpnServerConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type p2sVpnServerConfigurationState struct {
}

type P2sVpnServerConfigurationState struct {
}

func (P2sVpnServerConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*p2sVpnServerConfigurationState)(nil)).Elem()
}

type p2sVpnServerConfigurationArgs struct {
	Id                                             *string                                         `pulumi:"id"`
	Name                                           *string                                         `pulumi:"name"`
	P2SVpnServerConfigRadiusClientRootCertificates []P2SVpnServerConfigRadiusClientRootCertificate `pulumi:"p2SVpnServerConfigRadiusClientRootCertificates"`
	P2SVpnServerConfigRadiusServerRootCertificates []P2SVpnServerConfigRadiusServerRootCertificate `pulumi:"p2SVpnServerConfigRadiusServerRootCertificates"`
	P2SVpnServerConfigVpnClientRevokedCertificates []P2SVpnServerConfigVpnClientRevokedCertificate `pulumi:"p2SVpnServerConfigVpnClientRevokedCertificates"`
	P2SVpnServerConfigVpnClientRootCertificates    []P2SVpnServerConfigVpnClientRootCertificate    `pulumi:"p2SVpnServerConfigVpnClientRootCertificates"`
	P2SVpnServerConfigurationName                  *string                                         `pulumi:"p2SVpnServerConfigurationName"`
	RadiusServerAddress                            *string                                         `pulumi:"radiusServerAddress"`
	RadiusServerSecret                             *string                                         `pulumi:"radiusServerSecret"`
	ResourceGroupName                              string                                          `pulumi:"resourceGroupName"`
	VirtualWanName                                 string                                          `pulumi:"virtualWanName"`
	VpnClientIpsecPolicies                         []IpsecPolicy                                   `pulumi:"vpnClientIpsecPolicies"`
	VpnProtocols                                   []string                                        `pulumi:"vpnProtocols"`
}


type P2sVpnServerConfigurationArgs struct {
	Id                                             pulumi.StringPtrInput
	Name                                           pulumi.StringPtrInput
	P2SVpnServerConfigRadiusClientRootCertificates P2SVpnServerConfigRadiusClientRootCertificateArrayInput
	P2SVpnServerConfigRadiusServerRootCertificates P2SVpnServerConfigRadiusServerRootCertificateArrayInput
	P2SVpnServerConfigVpnClientRevokedCertificates P2SVpnServerConfigVpnClientRevokedCertificateArrayInput
	P2SVpnServerConfigVpnClientRootCertificates    P2SVpnServerConfigVpnClientRootCertificateArrayInput
	P2SVpnServerConfigurationName                  pulumi.StringPtrInput
	RadiusServerAddress                            pulumi.StringPtrInput
	RadiusServerSecret                             pulumi.StringPtrInput
	ResourceGroupName                              pulumi.StringInput
	VirtualWanName                                 pulumi.StringInput
	VpnClientIpsecPolicies                         IpsecPolicyArrayInput
	VpnProtocols                                   pulumi.StringArrayInput
}

func (P2sVpnServerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*p2sVpnServerConfigurationArgs)(nil)).Elem()
}

type P2sVpnServerConfigurationInput interface {
	pulumi.Input

	ToP2sVpnServerConfigurationOutput() P2sVpnServerConfigurationOutput
	ToP2sVpnServerConfigurationOutputWithContext(ctx context.Context) P2sVpnServerConfigurationOutput
}

func (*P2sVpnServerConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**P2sVpnServerConfiguration)(nil)).Elem()
}

func (i *P2sVpnServerConfiguration) ToP2sVpnServerConfigurationOutput() P2sVpnServerConfigurationOutput {
	return i.ToP2sVpnServerConfigurationOutputWithContext(context.Background())
}

func (i *P2sVpnServerConfiguration) ToP2sVpnServerConfigurationOutputWithContext(ctx context.Context) P2sVpnServerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(P2sVpnServerConfigurationOutput)
}

type P2sVpnServerConfigurationOutput struct{ *pulumi.OutputState }

func (P2sVpnServerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**P2sVpnServerConfiguration)(nil)).Elem()
}

func (o P2sVpnServerConfigurationOutput) ToP2sVpnServerConfigurationOutput() P2sVpnServerConfigurationOutput {
	return o
}

func (o P2sVpnServerConfigurationOutput) ToP2sVpnServerConfigurationOutputWithContext(ctx context.Context) P2sVpnServerConfigurationOutput {
	return o
}

func (o P2sVpnServerConfigurationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o P2sVpnServerConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o P2sVpnServerConfigurationOutput) P2SVpnGateways() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) SubResourceResponseArrayOutput { return v.P2SVpnGateways }).(SubResourceResponseArrayOutput)
}

func (o P2sVpnServerConfigurationOutput) P2SVpnServerConfigRadiusClientRootCertificates() P2SVpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) P2SVpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
		return v.P2SVpnServerConfigRadiusClientRootCertificates
	}).(P2SVpnServerConfigRadiusClientRootCertificateResponseArrayOutput)
}

func (o P2sVpnServerConfigurationOutput) P2SVpnServerConfigRadiusServerRootCertificates() P2SVpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) P2SVpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
		return v.P2SVpnServerConfigRadiusServerRootCertificates
	}).(P2SVpnServerConfigRadiusServerRootCertificateResponseArrayOutput)
}

func (o P2sVpnServerConfigurationOutput) P2SVpnServerConfigVpnClientRevokedCertificates() P2SVpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) P2SVpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
		return v.P2SVpnServerConfigVpnClientRevokedCertificates
	}).(P2SVpnServerConfigVpnClientRevokedCertificateResponseArrayOutput)
}

func (o P2sVpnServerConfigurationOutput) P2SVpnServerConfigVpnClientRootCertificates() P2SVpnServerConfigVpnClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) P2SVpnServerConfigVpnClientRootCertificateResponseArrayOutput {
		return v.P2SVpnServerConfigVpnClientRootCertificates
	}).(P2SVpnServerConfigVpnClientRootCertificateResponseArrayOutput)
}

func (o P2sVpnServerConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o P2sVpnServerConfigurationOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) pulumi.StringPtrOutput { return v.RadiusServerAddress }).(pulumi.StringPtrOutput)
}

func (o P2sVpnServerConfigurationOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) pulumi.StringPtrOutput { return v.RadiusServerSecret }).(pulumi.StringPtrOutput)
}

func (o P2sVpnServerConfigurationOutput) VpnClientIpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) IpsecPolicyResponseArrayOutput { return v.VpnClientIpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

func (o P2sVpnServerConfigurationOutput) VpnProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) pulumi.StringArrayOutput { return v.VpnProtocols }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(P2sVpnServerConfigurationOutput{})
}
