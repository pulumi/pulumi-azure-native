


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type CertificateCsr struct {
	pulumi.CustomResourceState

	CsrString          pulumi.StringPtrOutput `pulumi:"csrString"`
	DistinguishedName  pulumi.StringPtrOutput `pulumi:"distinguishedName"`
	HostingEnvironment pulumi.StringPtrOutput `pulumi:"hostingEnvironment"`
	Kind               pulumi.StringPtrOutput `pulumi:"kind"`
	Location           pulumi.StringOutput    `pulumi:"location"`
	Name               pulumi.StringPtrOutput `pulumi:"name"`
	Password           pulumi.StringPtrOutput `pulumi:"password"`
	PfxBlob            pulumi.StringPtrOutput `pulumi:"pfxBlob"`
	PublicKeyHash      pulumi.StringPtrOutput `pulumi:"publicKeyHash"`
	Tags               pulumi.StringMapOutput `pulumi:"tags"`
	Type               pulumi.StringPtrOutput `pulumi:"type"`
}


func NewCertificateCsr(ctx *pulumi.Context,
	name string, args *CertificateCsrArgs, opts ...pulumi.ResourceOption) (*CertificateCsr, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource CertificateCsr
	err := ctx.RegisterResource("azure-native:web/v20150801:CertificateCsr", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCertificateCsr(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateCsrState, opts ...pulumi.ResourceOption) (*CertificateCsr, error) {
	var resource CertificateCsr
	err := ctx.ReadResource("azure-native:web/v20150801:CertificateCsr", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type certificateCsrState struct {
}

type CertificateCsrState struct {
}

func (CertificateCsrState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateCsrState)(nil)).Elem()
}

type certificateCsrArgs struct {
	CsrString          *string           `pulumi:"csrString"`
	DistinguishedName  *string           `pulumi:"distinguishedName"`
	HostingEnvironment *string           `pulumi:"hostingEnvironment"`
	Id                 *string           `pulumi:"id"`
	Kind               *string           `pulumi:"kind"`
	Location           *string           `pulumi:"location"`
	Name               *string           `pulumi:"name"`
	Password           *string           `pulumi:"password"`
	PfxBlob            *string           `pulumi:"pfxBlob"`
	PublicKeyHash      *string           `pulumi:"publicKeyHash"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	Tags               map[string]string `pulumi:"tags"`
	Type               *string           `pulumi:"type"`
}


type CertificateCsrArgs struct {
	CsrString          pulumi.StringPtrInput
	DistinguishedName  pulumi.StringPtrInput
	HostingEnvironment pulumi.StringPtrInput
	Id                 pulumi.StringPtrInput
	Kind               pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	Password           pulumi.StringPtrInput
	PfxBlob            pulumi.StringPtrInput
	PublicKeyHash      pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	Tags               pulumi.StringMapInput
	Type               pulumi.StringPtrInput
}

func (CertificateCsrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateCsrArgs)(nil)).Elem()
}

type CertificateCsrInput interface {
	pulumi.Input

	ToCertificateCsrOutput() CertificateCsrOutput
	ToCertificateCsrOutputWithContext(ctx context.Context) CertificateCsrOutput
}

func (*CertificateCsr) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateCsr)(nil)).Elem()
}

func (i *CertificateCsr) ToCertificateCsrOutput() CertificateCsrOutput {
	return i.ToCertificateCsrOutputWithContext(context.Background())
}

func (i *CertificateCsr) ToCertificateCsrOutputWithContext(ctx context.Context) CertificateCsrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateCsrOutput)
}

type CertificateCsrOutput struct{ *pulumi.OutputState }

func (CertificateCsrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateCsr)(nil)).Elem()
}

func (o CertificateCsrOutput) ToCertificateCsrOutput() CertificateCsrOutput {
	return o
}

func (o CertificateCsrOutput) ToCertificateCsrOutputWithContext(ctx context.Context) CertificateCsrOutput {
	return o
}

func (o CertificateCsrOutput) CsrString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateCsr) pulumi.StringPtrOutput { return v.CsrString }).(pulumi.StringPtrOutput)
}

func (o CertificateCsrOutput) DistinguishedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateCsr) pulumi.StringPtrOutput { return v.DistinguishedName }).(pulumi.StringPtrOutput)
}

func (o CertificateCsrOutput) HostingEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateCsr) pulumi.StringPtrOutput { return v.HostingEnvironment }).(pulumi.StringPtrOutput)
}

func (o CertificateCsrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateCsr) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o CertificateCsrOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateCsr) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o CertificateCsrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateCsr) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CertificateCsrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateCsr) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o CertificateCsrOutput) PfxBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateCsr) pulumi.StringPtrOutput { return v.PfxBlob }).(pulumi.StringPtrOutput)
}

func (o CertificateCsrOutput) PublicKeyHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateCsr) pulumi.StringPtrOutput { return v.PublicKeyHash }).(pulumi.StringPtrOutput)
}

func (o CertificateCsrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CertificateCsr) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CertificateCsrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateCsr) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateCsrOutput{})
}
