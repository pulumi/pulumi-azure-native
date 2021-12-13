


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GatewayCertificateAuthority struct {
	pulumi.CustomResourceState

	IsTrusted pulumi.BoolPtrOutput `pulumi:"isTrusted"`
	Name      pulumi.StringOutput  `pulumi:"name"`
	Type      pulumi.StringOutput  `pulumi:"type"`
}


func NewGatewayCertificateAuthority(ctx *pulumi.Context,
	name string, args *GatewayCertificateAuthorityArgs, opts ...pulumi.ResourceOption) (*GatewayCertificateAuthority, error) {
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
			Type: pulumi.String("azure-native:apimanagement:GatewayCertificateAuthority"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:GatewayCertificateAuthority"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:GatewayCertificateAuthority"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:GatewayCertificateAuthority"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:GatewayCertificateAuthority"),
		},
	})
	opts = append(opts, aliases)
	var resource GatewayCertificateAuthority
	err := ctx.RegisterResource("azure-native:apimanagement/v20210101preview:GatewayCertificateAuthority", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGatewayCertificateAuthority(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayCertificateAuthorityState, opts ...pulumi.ResourceOption) (*GatewayCertificateAuthority, error) {
	var resource GatewayCertificateAuthority
	err := ctx.ReadResource("azure-native:apimanagement/v20210101preview:GatewayCertificateAuthority", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type gatewayCertificateAuthorityState struct {
}

type GatewayCertificateAuthorityState struct {
}

func (GatewayCertificateAuthorityState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayCertificateAuthorityState)(nil)).Elem()
}

type gatewayCertificateAuthorityArgs struct {
	CertificateId     *string `pulumi:"certificateId"`
	GatewayId         string  `pulumi:"gatewayId"`
	IsTrusted         *bool   `pulumi:"isTrusted"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type GatewayCertificateAuthorityArgs struct {
	CertificateId     pulumi.StringPtrInput
	GatewayId         pulumi.StringInput
	IsTrusted         pulumi.BoolPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (GatewayCertificateAuthorityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayCertificateAuthorityArgs)(nil)).Elem()
}

type GatewayCertificateAuthorityInput interface {
	pulumi.Input

	ToGatewayCertificateAuthorityOutput() GatewayCertificateAuthorityOutput
	ToGatewayCertificateAuthorityOutputWithContext(ctx context.Context) GatewayCertificateAuthorityOutput
}

func (*GatewayCertificateAuthority) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCertificateAuthority)(nil)).Elem()
}

func (i *GatewayCertificateAuthority) ToGatewayCertificateAuthorityOutput() GatewayCertificateAuthorityOutput {
	return i.ToGatewayCertificateAuthorityOutputWithContext(context.Background())
}

func (i *GatewayCertificateAuthority) ToGatewayCertificateAuthorityOutputWithContext(ctx context.Context) GatewayCertificateAuthorityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCertificateAuthorityOutput)
}

type GatewayCertificateAuthorityOutput struct{ *pulumi.OutputState }

func (GatewayCertificateAuthorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCertificateAuthority)(nil)).Elem()
}

func (o GatewayCertificateAuthorityOutput) ToGatewayCertificateAuthorityOutput() GatewayCertificateAuthorityOutput {
	return o
}

func (o GatewayCertificateAuthorityOutput) ToGatewayCertificateAuthorityOutputWithContext(ctx context.Context) GatewayCertificateAuthorityOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GatewayCertificateAuthorityOutput{})
}
