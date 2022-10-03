


package v20160601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Certificate struct {
	pulumi.CustomResourceState

	ChangedTime       pulumi.StringOutput                   `pulumi:"changedTime"`
	CreatedTime       pulumi.StringOutput                   `pulumi:"createdTime"`
	Key               KeyVaultKeyReferenceResponsePtrOutput `pulumi:"key"`
	Location          pulumi.StringPtrOutput                `pulumi:"location"`
	Metadata          pulumi.AnyOutput                      `pulumi:"metadata"`
	Name              pulumi.StringOutput                   `pulumi:"name"`
	PublicCertificate pulumi.StringPtrOutput                `pulumi:"publicCertificate"`
	Tags              pulumi.StringMapOutput                `pulumi:"tags"`
	Type              pulumi.StringOutput                   `pulumi:"type"`
}


func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20150801preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:Certificate"),
		},
	})
	opts = append(opts, aliases)
	var resource Certificate
	err := ctx.RegisterResource("azure-native:logic/v20160601:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azure-native:logic/v20160601:Certificate", name, id, state, &resource, opts...)
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
	CertificateName        *string               `pulumi:"certificateName"`
	IntegrationAccountName string                `pulumi:"integrationAccountName"`
	Key                    *KeyVaultKeyReference `pulumi:"key"`
	Location               *string               `pulumi:"location"`
	Metadata               interface{}           `pulumi:"metadata"`
	PublicCertificate      *string               `pulumi:"publicCertificate"`
	ResourceGroupName      string                `pulumi:"resourceGroupName"`
	Tags                   map[string]string     `pulumi:"tags"`
}


type CertificateArgs struct {
	CertificateName        pulumi.StringPtrInput
	IntegrationAccountName pulumi.StringInput
	Key                    KeyVaultKeyReferencePtrInput
	Location               pulumi.StringPtrInput
	Metadata               pulumi.Input
	PublicCertificate      pulumi.StringPtrInput
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

func (o CertificateOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o CertificateOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o CertificateOutput) Key() KeyVaultKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v *Certificate) KeyVaultKeyReferenceResponsePtrOutput { return v.Key }).(KeyVaultKeyReferenceResponsePtrOutput)
}

func (o CertificateOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *Certificate) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

func (o CertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CertificateOutput) PublicCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.PublicCertificate }).(pulumi.StringPtrOutput)
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
