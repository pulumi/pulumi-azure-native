


package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationAccountCertificate struct {
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


func NewIntegrationAccountCertificate(ctx *pulumi.Context,
	name string, args *IntegrationAccountCertificateArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountCertificate, error) {
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
			Type: pulumi.String("azure-nextgen:logic/v20180701preview:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-native:logic:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20150801preview:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20150801preview:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20160601:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20190501:IntegrationAccountCertificate"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountCertificate
	err := ctx.RegisterResource("azure-native:logic/v20180701preview:IntegrationAccountCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationAccountCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountCertificateState, opts ...pulumi.ResourceOption) (*IntegrationAccountCertificate, error) {
	var resource IntegrationAccountCertificate
	err := ctx.ReadResource("azure-native:logic/v20180701preview:IntegrationAccountCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type integrationAccountCertificateState struct {
}

type IntegrationAccountCertificateState struct {
}

func (IntegrationAccountCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountCertificateState)(nil)).Elem()
}

type integrationAccountCertificateArgs struct {
	CertificateName        *string               `pulumi:"certificateName"`
	IntegrationAccountName string                `pulumi:"integrationAccountName"`
	Key                    *KeyVaultKeyReference `pulumi:"key"`
	Location               *string               `pulumi:"location"`
	Metadata               interface{}           `pulumi:"metadata"`
	PublicCertificate      *string               `pulumi:"publicCertificate"`
	ResourceGroupName      string                `pulumi:"resourceGroupName"`
	Tags                   map[string]string     `pulumi:"tags"`
}


type IntegrationAccountCertificateArgs struct {
	CertificateName        pulumi.StringPtrInput
	IntegrationAccountName pulumi.StringInput
	Key                    KeyVaultKeyReferencePtrInput
	Location               pulumi.StringPtrInput
	Metadata               pulumi.Input
	PublicCertificate      pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (IntegrationAccountCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountCertificateArgs)(nil)).Elem()
}

type IntegrationAccountCertificateInput interface {
	pulumi.Input

	ToIntegrationAccountCertificateOutput() IntegrationAccountCertificateOutput
	ToIntegrationAccountCertificateOutputWithContext(ctx context.Context) IntegrationAccountCertificateOutput
}

func (*IntegrationAccountCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountCertificate)(nil))
}

func (i *IntegrationAccountCertificate) ToIntegrationAccountCertificateOutput() IntegrationAccountCertificateOutput {
	return i.ToIntegrationAccountCertificateOutputWithContext(context.Background())
}

func (i *IntegrationAccountCertificate) ToIntegrationAccountCertificateOutputWithContext(ctx context.Context) IntegrationAccountCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountCertificateOutput)
}

type IntegrationAccountCertificateOutput struct{ *pulumi.OutputState }

func (IntegrationAccountCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountCertificate)(nil))
}

func (o IntegrationAccountCertificateOutput) ToIntegrationAccountCertificateOutput() IntegrationAccountCertificateOutput {
	return o
}

func (o IntegrationAccountCertificateOutput) ToIntegrationAccountCertificateOutputWithContext(ctx context.Context) IntegrationAccountCertificateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountCertificateOutput{})
}
