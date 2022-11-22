


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServerTrustCertificate struct {
	pulumi.CustomResourceState

	CertificateName pulumi.StringOutput    `pulumi:"certificateName"`
	Name            pulumi.StringOutput    `pulumi:"name"`
	PublicBlob      pulumi.StringPtrOutput `pulumi:"publicBlob"`
	Thumbprint      pulumi.StringOutput    `pulumi:"thumbprint"`
	Type            pulumi.StringOutput    `pulumi:"type"`
}


func NewServerTrustCertificate(ctx *pulumi.Context,
	name string, args *ServerTrustCertificateArgs, opts ...pulumi.ResourceOption) (*ServerTrustCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ServerTrustCertificate"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerTrustCertificate"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ServerTrustCertificate"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ServerTrustCertificate"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ServerTrustCertificate"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ServerTrustCertificate"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerTrustCertificate
	err := ctx.RegisterResource("azure-native:sql/v20220501preview:ServerTrustCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServerTrustCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerTrustCertificateState, opts ...pulumi.ResourceOption) (*ServerTrustCertificate, error) {
	var resource ServerTrustCertificate
	err := ctx.ReadResource("azure-native:sql/v20220501preview:ServerTrustCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serverTrustCertificateState struct {
}

type ServerTrustCertificateState struct {
}

func (ServerTrustCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverTrustCertificateState)(nil)).Elem()
}

type serverTrustCertificateArgs struct {
	CertificateName     *string `pulumi:"certificateName"`
	ManagedInstanceName string  `pulumi:"managedInstanceName"`
	PublicBlob          *string `pulumi:"publicBlob"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
}


type ServerTrustCertificateArgs struct {
	CertificateName     pulumi.StringPtrInput
	ManagedInstanceName pulumi.StringInput
	PublicBlob          pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
}

func (ServerTrustCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverTrustCertificateArgs)(nil)).Elem()
}

type ServerTrustCertificateInput interface {
	pulumi.Input

	ToServerTrustCertificateOutput() ServerTrustCertificateOutput
	ToServerTrustCertificateOutputWithContext(ctx context.Context) ServerTrustCertificateOutput
}

func (*ServerTrustCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerTrustCertificate)(nil)).Elem()
}

func (i *ServerTrustCertificate) ToServerTrustCertificateOutput() ServerTrustCertificateOutput {
	return i.ToServerTrustCertificateOutputWithContext(context.Background())
}

func (i *ServerTrustCertificate) ToServerTrustCertificateOutputWithContext(ctx context.Context) ServerTrustCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerTrustCertificateOutput)
}

type ServerTrustCertificateOutput struct{ *pulumi.OutputState }

func (ServerTrustCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerTrustCertificate)(nil)).Elem()
}

func (o ServerTrustCertificateOutput) ToServerTrustCertificateOutput() ServerTrustCertificateOutput {
	return o
}

func (o ServerTrustCertificateOutput) ToServerTrustCertificateOutputWithContext(ctx context.Context) ServerTrustCertificateOutput {
	return o
}

func (o ServerTrustCertificateOutput) CertificateName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTrustCertificate) pulumi.StringOutput { return v.CertificateName }).(pulumi.StringOutput)
}

func (o ServerTrustCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTrustCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServerTrustCertificateOutput) PublicBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerTrustCertificate) pulumi.StringPtrOutput { return v.PublicBlob }).(pulumi.StringPtrOutput)
}

func (o ServerTrustCertificateOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTrustCertificate) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o ServerTrustCertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTrustCertificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerTrustCertificateOutput{})
}
