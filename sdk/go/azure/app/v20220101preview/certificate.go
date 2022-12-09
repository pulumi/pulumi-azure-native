


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Certificate struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                 `pulumi:"location"`
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties CertificateResponsePropertiesOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput            `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput              `pulumi:"tags"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedEnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedEnvironmentName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220301:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220601preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221001:Certificate"),
		},
	})
	opts = append(opts, aliases)
	var resource Certificate
	err := ctx.RegisterResource("azure-native:app/v20220101preview:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azure-native:app/v20220101preview:Certificate", name, id, state, &resource, opts...)
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
	Location               *string                `pulumi:"location"`
	ManagedEnvironmentName string                 `pulumi:"managedEnvironmentName"`
	Name                   *string                `pulumi:"name"`
	Properties             *CertificateProperties `pulumi:"properties"`
	ResourceGroupName      string                 `pulumi:"resourceGroupName"`
	Tags                   map[string]string      `pulumi:"tags"`
}


type CertificateArgs struct {
	Location               pulumi.StringPtrInput
	ManagedEnvironmentName pulumi.StringInput
	Name                   pulumi.StringPtrInput
	Properties             CertificatePropertiesPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
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

func (o CertificateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o CertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CertificateOutput) Properties() CertificateResponsePropertiesOutput {
	return o.ApplyT(func(v *Certificate) CertificateResponsePropertiesOutput { return v.Properties }).(CertificateResponsePropertiesOutput)
}

func (o CertificateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Certificate) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o CertificateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
}
