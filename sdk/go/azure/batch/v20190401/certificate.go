


package v20190401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Certificate struct {
	pulumi.CustomResourceState

	DeleteCertificateError                  DeleteCertificateErrorResponseOutput `pulumi:"deleteCertificateError"`
	Etag                                    pulumi.StringOutput                  `pulumi:"etag"`
	Format                                  pulumi.StringPtrOutput               `pulumi:"format"`
	Name                                    pulumi.StringOutput                  `pulumi:"name"`
	PreviousProvisioningState               pulumi.StringOutput                  `pulumi:"previousProvisioningState"`
	PreviousProvisioningStateTransitionTime pulumi.StringOutput                  `pulumi:"previousProvisioningStateTransitionTime"`
	ProvisioningState                       pulumi.StringOutput                  `pulumi:"provisioningState"`
	ProvisioningStateTransitionTime         pulumi.StringOutput                  `pulumi:"provisioningStateTransitionTime"`
	PublicData                              pulumi.StringOutput                  `pulumi:"publicData"`
	Thumbprint                              pulumi.StringPtrOutput               `pulumi:"thumbprint"`
	ThumbprintAlgorithm                     pulumi.StringPtrOutput               `pulumi:"thumbprintAlgorithm"`
	Type                                    pulumi.StringOutput                  `pulumi:"type"`
}


func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Data == nil {
		return nil, errors.New("invalid value for required argument 'Data'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:batch:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20170901:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20181201:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190801:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200301:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200501:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200901:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210101:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210601:Certificate"),
		},
	})
	opts = append(opts, aliases)
	var resource Certificate
	err := ctx.RegisterResource("azure-native:batch/v20190401:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azure-native:batch/v20190401:Certificate", name, id, state, &resource, opts...)
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
	AccountName         string             `pulumi:"accountName"`
	CertificateName     *string            `pulumi:"certificateName"`
	Data                string             `pulumi:"data"`
	Format              *CertificateFormat `pulumi:"format"`
	Password            *string            `pulumi:"password"`
	ResourceGroupName   string             `pulumi:"resourceGroupName"`
	Thumbprint          *string            `pulumi:"thumbprint"`
	ThumbprintAlgorithm *string            `pulumi:"thumbprintAlgorithm"`
}


type CertificateArgs struct {
	AccountName         pulumi.StringInput
	CertificateName     pulumi.StringPtrInput
	Data                pulumi.StringInput
	Format              CertificateFormatPtrInput
	Password            pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Thumbprint          pulumi.StringPtrInput
	ThumbprintAlgorithm pulumi.StringPtrInput
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

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
}
