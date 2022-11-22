


package v20170701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Certificate struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput                 `pulumi:"etag"`
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
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devices:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180122:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180401:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20181201preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190322:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190322preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190701preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20191104:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200301:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200401:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200615:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200710preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200801:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200831:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200831preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210201preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210303preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210331:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210701:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210701preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210702:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210702preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20220430preview:Certificate"),
		},
	})
	opts = append(opts, aliases)
	var resource Certificate
	err := ctx.RegisterResource("azure-native:devices/v20170701:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azure-native:devices/v20170701:Certificate", name, id, state, &resource, opts...)
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
	Certificate       *string `pulumi:"certificate"`
	CertificateName   *string `pulumi:"certificateName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceName      string  `pulumi:"resourceName"`
}


type CertificateArgs struct {
	Certificate       pulumi.StringPtrInput
	CertificateName   pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
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

func (o CertificateOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
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
