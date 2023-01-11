


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Certificate struct {
	pulumi.CustomResourceState

	CerBlob                   pulumi.StringPtrOutput                     `pulumi:"cerBlob"`
	ExpirationDate            pulumi.StringPtrOutput                     `pulumi:"expirationDate"`
	FriendlyName              pulumi.StringPtrOutput                     `pulumi:"friendlyName"`
	HostNames                 pulumi.StringArrayOutput                   `pulumi:"hostNames"`
	HostingEnvironmentProfile HostingEnvironmentProfileResponsePtrOutput `pulumi:"hostingEnvironmentProfile"`
	IssueDate                 pulumi.StringPtrOutput                     `pulumi:"issueDate"`
	Issuer                    pulumi.StringPtrOutput                     `pulumi:"issuer"`
	Kind                      pulumi.StringPtrOutput                     `pulumi:"kind"`
	Location                  pulumi.StringOutput                        `pulumi:"location"`
	Name                      pulumi.StringPtrOutput                     `pulumi:"name"`
	Password                  pulumi.StringPtrOutput                     `pulumi:"password"`
	PfxBlob                   pulumi.StringPtrOutput                     `pulumi:"pfxBlob"`
	PublicKeyHash             pulumi.StringPtrOutput                     `pulumi:"publicKeyHash"`
	SelfLink                  pulumi.StringPtrOutput                     `pulumi:"selfLink"`
	SiteName                  pulumi.StringPtrOutput                     `pulumi:"siteName"`
	SubjectName               pulumi.StringPtrOutput                     `pulumi:"subjectName"`
	Tags                      pulumi.StringMapOutput                     `pulumi:"tags"`
	Thumbprint                pulumi.StringPtrOutput                     `pulumi:"thumbprint"`
	Type                      pulumi.StringPtrOutput                     `pulumi:"type"`
	Valid                     pulumi.BoolPtrOutput                       `pulumi:"valid"`
}


func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160301:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:Certificate"),
		},
	})
	opts = append(opts, aliases)
	var resource Certificate
	err := ctx.RegisterResource("azure-native:web/v20150801:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azure-native:web/v20150801:Certificate", name, id, state, &resource, opts...)
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
	CerBlob                   *string                    `pulumi:"cerBlob"`
	ExpirationDate            *string                    `pulumi:"expirationDate"`
	FriendlyName              *string                    `pulumi:"friendlyName"`
	HostNames                 []string                   `pulumi:"hostNames"`
	HostingEnvironmentProfile *HostingEnvironmentProfile `pulumi:"hostingEnvironmentProfile"`
	Id                        *string                    `pulumi:"id"`
	IssueDate                 *string                    `pulumi:"issueDate"`
	Issuer                    *string                    `pulumi:"issuer"`
	Kind                      *string                    `pulumi:"kind"`
	Location                  *string                    `pulumi:"location"`
	Name                      *string                    `pulumi:"name"`
	Password                  *string                    `pulumi:"password"`
	PfxBlob                   *string                    `pulumi:"pfxBlob"`
	PublicKeyHash             *string                    `pulumi:"publicKeyHash"`
	ResourceGroupName         string                     `pulumi:"resourceGroupName"`
	SelfLink                  *string                    `pulumi:"selfLink"`
	SiteName                  *string                    `pulumi:"siteName"`
	SubjectName               *string                    `pulumi:"subjectName"`
	Tags                      map[string]string          `pulumi:"tags"`
	Thumbprint                *string                    `pulumi:"thumbprint"`
	Type                      *string                    `pulumi:"type"`
	Valid                     *bool                      `pulumi:"valid"`
}


type CertificateArgs struct {
	CerBlob                   pulumi.StringPtrInput
	ExpirationDate            pulumi.StringPtrInput
	FriendlyName              pulumi.StringPtrInput
	HostNames                 pulumi.StringArrayInput
	HostingEnvironmentProfile HostingEnvironmentProfilePtrInput
	Id                        pulumi.StringPtrInput
	IssueDate                 pulumi.StringPtrInput
	Issuer                    pulumi.StringPtrInput
	Kind                      pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	Password                  pulumi.StringPtrInput
	PfxBlob                   pulumi.StringPtrInput
	PublicKeyHash             pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	SelfLink                  pulumi.StringPtrInput
	SiteName                  pulumi.StringPtrInput
	SubjectName               pulumi.StringPtrInput
	Tags                      pulumi.StringMapInput
	Thumbprint                pulumi.StringPtrInput
	Type                      pulumi.StringPtrInput
	Valid                     pulumi.BoolPtrInput
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

func (o CertificateOutput) CerBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.CerBlob }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) HostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringArrayOutput { return v.HostNames }).(pulumi.StringArrayOutput)
}

func (o CertificateOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyT(func(v *Certificate) HostingEnvironmentProfileResponsePtrOutput { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponsePtrOutput)
}

func (o CertificateOutput) IssueDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.IssueDate }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Issuer }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o CertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) PfxBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.PfxBlob }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) PublicKeyHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.PublicKeyHash }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) SelfLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.SelfLink }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) SiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.SiteName }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) SubjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.SubjectName }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) Valid() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.BoolPtrOutput { return v.Valid }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
}
