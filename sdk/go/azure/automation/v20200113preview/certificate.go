


package v20200113preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Certificate struct {
	pulumi.CustomResourceState

	CreationTime     pulumi.StringOutput    `pulumi:"creationTime"`
	Description      pulumi.StringPtrOutput `pulumi:"description"`
	ExpiryTime       pulumi.StringOutput    `pulumi:"expiryTime"`
	IsExportable     pulumi.BoolOutput      `pulumi:"isExportable"`
	LastModifiedTime pulumi.StringOutput    `pulumi:"lastModifiedTime"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	Thumbprint       pulumi.StringOutput    `pulumi:"thumbprint"`
	Type             pulumi.StringOutput    `pulumi:"type"`
}


func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.Base64Value == nil {
		return nil, errors.New("invalid value for required argument 'Base64Value'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20220808:Certificate"),
		},
	})
	opts = append(opts, aliases)
	var resource Certificate
	err := ctx.RegisterResource("azure-native:automation/v20200113preview:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azure-native:automation/v20200113preview:Certificate", name, id, state, &resource, opts...)
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
	AutomationAccountName string  `pulumi:"automationAccountName"`
	Base64Value           string  `pulumi:"base64Value"`
	CertificateName       *string `pulumi:"certificateName"`
	Description           *string `pulumi:"description"`
	IsExportable          *bool   `pulumi:"isExportable"`
	Name                  string  `pulumi:"name"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	Thumbprint            *string `pulumi:"thumbprint"`
}


type CertificateArgs struct {
	AutomationAccountName pulumi.StringInput
	Base64Value           pulumi.StringInput
	CertificateName       pulumi.StringPtrInput
	Description           pulumi.StringPtrInput
	IsExportable          pulumi.BoolPtrInput
	Name                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	Thumbprint            pulumi.StringPtrInput
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

func (o CertificateOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o CertificateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CertificateOutput) ExpiryTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ExpiryTime }).(pulumi.StringOutput)
}

func (o CertificateOutput) IsExportable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Certificate) pulumi.BoolOutput { return v.IsExportable }).(pulumi.BoolOutput)
}

func (o CertificateOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o CertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CertificateOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o CertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
}
