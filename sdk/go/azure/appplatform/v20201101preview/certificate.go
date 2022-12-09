


package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Certificate struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties CertificatePropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20200701:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210601preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:Certificate"),
		},
	})
	opts = append(opts, aliases)
	var resource Certificate
	err := ctx.RegisterResource("azure-native:appplatform/v20201101preview:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azure-native:appplatform/v20201101preview:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type certificateState struct {
}

type CertificateState struct {
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	CertificateName   *string                `pulumi:"certificateName"`
	Properties        *CertificateProperties `pulumi:"properties"`
	ResourceGroupName string                 `pulumi:"resourceGroupName"`
	ServiceName       string                 `pulumi:"serviceName"`
}


type CertificateArgs struct {
	CertificateName   pulumi.StringPtrInput
	Properties        CertificatePropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func (o CertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CertificateOutput) Properties() CertificatePropertiesResponseOutput {
	return o.ApplyT(func(v *Certificate) CertificatePropertiesResponseOutput { return v.Properties }).(CertificatePropertiesResponseOutput)
}

func (o CertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
}
