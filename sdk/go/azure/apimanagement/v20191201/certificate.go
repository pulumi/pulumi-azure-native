


package v20191201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Certificate struct {
	pulumi.CustomResourceState

	ExpirationDate pulumi.StringOutput `pulumi:"expirationDate"`
	Name           pulumi.StringOutput `pulumi:"name"`
	Subject        pulumi.StringOutput `pulumi:"subject"`
	Thumbprint     pulumi.StringOutput `pulumi:"thumbprint"`
	Type           pulumi.StringOutput `pulumi:"type"`
}


func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Data == nil {
		return nil, errors.New("invalid value for required argument 'Data'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20160707:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20161010:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180601preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210801:Certificate"),
		},
	})
	opts = append(opts, aliases)
	var resource Certificate
	err := ctx.RegisterResource("azure-native:apimanagement/v20191201:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azure-native:apimanagement/v20191201:Certificate", name, id, state, &resource, opts...)
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
	CertificateId     *string `pulumi:"certificateId"`
	Data              string  `pulumi:"data"`
	Password          string  `pulumi:"password"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type CertificateArgs struct {
	CertificateId     pulumi.StringPtrInput
	Data              pulumi.StringInput
	Password          pulumi.StringInput
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
	return reflect.TypeOf((*Certificate)(nil))
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil))
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterOutputType(CertificateOutput{})
}
