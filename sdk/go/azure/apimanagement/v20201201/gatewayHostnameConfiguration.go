


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GatewayHostnameConfiguration struct {
	pulumi.CustomResourceState

	CertificateId              pulumi.StringPtrOutput `pulumi:"certificateId"`
	Hostname                   pulumi.StringPtrOutput `pulumi:"hostname"`
	Http2Enabled               pulumi.BoolPtrOutput   `pulumi:"http2Enabled"`
	Name                       pulumi.StringOutput    `pulumi:"name"`
	NegotiateClientCertificate pulumi.BoolPtrOutput   `pulumi:"negotiateClientCertificate"`
	Tls10Enabled               pulumi.BoolPtrOutput   `pulumi:"tls10Enabled"`
	Tls11Enabled               pulumi.BoolPtrOutput   `pulumi:"tls11Enabled"`
	Type                       pulumi.StringOutput    `pulumi:"type"`
}


func NewGatewayHostnameConfiguration(ctx *pulumi.Context,
	name string, args *GatewayHostnameConfigurationArgs, opts ...pulumi.ResourceOption) (*GatewayHostnameConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:GatewayHostnameConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:GatewayHostnameConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:GatewayHostnameConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:GatewayHostnameConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:GatewayHostnameConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:GatewayHostnameConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:GatewayHostnameConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource GatewayHostnameConfiguration
	err := ctx.RegisterResource("azure-native:apimanagement/v20201201:GatewayHostnameConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGatewayHostnameConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayHostnameConfigurationState, opts ...pulumi.ResourceOption) (*GatewayHostnameConfiguration, error) {
	var resource GatewayHostnameConfiguration
	err := ctx.ReadResource("azure-native:apimanagement/v20201201:GatewayHostnameConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type gatewayHostnameConfigurationState struct {
}

type GatewayHostnameConfigurationState struct {
}

func (GatewayHostnameConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayHostnameConfigurationState)(nil)).Elem()
}

type gatewayHostnameConfigurationArgs struct {
	CertificateId              *string `pulumi:"certificateId"`
	GatewayId                  string  `pulumi:"gatewayId"`
	HcId                       *string `pulumi:"hcId"`
	Hostname                   *string `pulumi:"hostname"`
	Http2Enabled               *bool   `pulumi:"http2Enabled"`
	NegotiateClientCertificate *bool   `pulumi:"negotiateClientCertificate"`
	ResourceGroupName          string  `pulumi:"resourceGroupName"`
	ServiceName                string  `pulumi:"serviceName"`
	Tls10Enabled               *bool   `pulumi:"tls10Enabled"`
	Tls11Enabled               *bool   `pulumi:"tls11Enabled"`
}


type GatewayHostnameConfigurationArgs struct {
	CertificateId              pulumi.StringPtrInput
	GatewayId                  pulumi.StringInput
	HcId                       pulumi.StringPtrInput
	Hostname                   pulumi.StringPtrInput
	Http2Enabled               pulumi.BoolPtrInput
	NegotiateClientCertificate pulumi.BoolPtrInput
	ResourceGroupName          pulumi.StringInput
	ServiceName                pulumi.StringInput
	Tls10Enabled               pulumi.BoolPtrInput
	Tls11Enabled               pulumi.BoolPtrInput
}

func (GatewayHostnameConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayHostnameConfigurationArgs)(nil)).Elem()
}

type GatewayHostnameConfigurationInput interface {
	pulumi.Input

	ToGatewayHostnameConfigurationOutput() GatewayHostnameConfigurationOutput
	ToGatewayHostnameConfigurationOutputWithContext(ctx context.Context) GatewayHostnameConfigurationOutput
}

func (*GatewayHostnameConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayHostnameConfiguration)(nil))
}

func (i *GatewayHostnameConfiguration) ToGatewayHostnameConfigurationOutput() GatewayHostnameConfigurationOutput {
	return i.ToGatewayHostnameConfigurationOutputWithContext(context.Background())
}

func (i *GatewayHostnameConfiguration) ToGatewayHostnameConfigurationOutputWithContext(ctx context.Context) GatewayHostnameConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayHostnameConfigurationOutput)
}

type GatewayHostnameConfigurationOutput struct{ *pulumi.OutputState }

func (GatewayHostnameConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayHostnameConfiguration)(nil))
}

func (o GatewayHostnameConfigurationOutput) ToGatewayHostnameConfigurationOutput() GatewayHostnameConfigurationOutput {
	return o
}

func (o GatewayHostnameConfigurationOutput) ToGatewayHostnameConfigurationOutputWithContext(ctx context.Context) GatewayHostnameConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GatewayHostnameConfigurationOutput{})
}
