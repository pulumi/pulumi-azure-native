


package v20200101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DpsCertificate struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput                 `pulumi:"etag"`
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties CertificatePropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewDpsCertificate(ctx *pulumi.Context,
	name string, args *DpsCertificateArgs, opts ...pulumi.ResourceOption) (*DpsCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProvisioningServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ProvisioningServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devices/v20200101:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:devices:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20170821preview:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20170821preview:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20171115:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20171115:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180122:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20180122:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200301:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200301:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200901preview:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200901preview:DpsCertificate"),
		},
	})
	opts = append(opts, aliases)
	var resource DpsCertificate
	err := ctx.RegisterResource("azure-native:devices/v20200101:DpsCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDpsCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DpsCertificateState, opts ...pulumi.ResourceOption) (*DpsCertificate, error) {
	var resource DpsCertificate
	err := ctx.ReadResource("azure-native:devices/v20200101:DpsCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dpsCertificateState struct {
}

type DpsCertificateState struct {
}

func (DpsCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*dpsCertificateState)(nil)).Elem()
}

type dpsCertificateArgs struct {
	Certificate             *string `pulumi:"certificate"`
	CertificateName         *string `pulumi:"certificateName"`
	ProvisioningServiceName string  `pulumi:"provisioningServiceName"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
}


type DpsCertificateArgs struct {
	Certificate             pulumi.StringPtrInput
	CertificateName         pulumi.StringPtrInput
	ProvisioningServiceName pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
}

func (DpsCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dpsCertificateArgs)(nil)).Elem()
}

type DpsCertificateInput interface {
	pulumi.Input

	ToDpsCertificateOutput() DpsCertificateOutput
	ToDpsCertificateOutputWithContext(ctx context.Context) DpsCertificateOutput
}

func (*DpsCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((*DpsCertificate)(nil))
}

func (i *DpsCertificate) ToDpsCertificateOutput() DpsCertificateOutput {
	return i.ToDpsCertificateOutputWithContext(context.Background())
}

func (i *DpsCertificate) ToDpsCertificateOutputWithContext(ctx context.Context) DpsCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DpsCertificateOutput)
}

type DpsCertificateOutput struct{ *pulumi.OutputState }

func (DpsCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DpsCertificate)(nil))
}

func (o DpsCertificateOutput) ToDpsCertificateOutput() DpsCertificateOutput {
	return o
}

func (o DpsCertificateOutput) ToDpsCertificateOutputWithContext(ctx context.Context) DpsCertificateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DpsCertificateOutput{})
}
