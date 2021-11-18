


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppPublicCertificate struct {
	pulumi.CustomResourceState

	Blob                      pulumi.StringPtrOutput `pulumi:"blob"`
	Kind                      pulumi.StringPtrOutput `pulumi:"kind"`
	Name                      pulumi.StringOutput    `pulumi:"name"`
	PublicCertificateLocation pulumi.StringPtrOutput `pulumi:"publicCertificateLocation"`
	Thumbprint                pulumi.StringOutput    `pulumi:"thumbprint"`
	Type                      pulumi.StringOutput    `pulumi:"type"`
}


func NewWebAppPublicCertificate(ctx *pulumi.Context,
	name string, args *WebAppPublicCertificateArgs, opts ...pulumi.ResourceOption) (*WebAppPublicCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppPublicCertificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppPublicCertificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppPublicCertificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppPublicCertificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppPublicCertificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppPublicCertificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppPublicCertificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppPublicCertificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppPublicCertificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppPublicCertificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppPublicCertificate"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppPublicCertificate
	err := ctx.RegisterResource("azure-native:web/v20210201:WebAppPublicCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppPublicCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppPublicCertificateState, opts ...pulumi.ResourceOption) (*WebAppPublicCertificate, error) {
	var resource WebAppPublicCertificate
	err := ctx.ReadResource("azure-native:web/v20210201:WebAppPublicCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppPublicCertificateState struct {
}

type WebAppPublicCertificateState struct {
}

func (WebAppPublicCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPublicCertificateState)(nil)).Elem()
}

type webAppPublicCertificateArgs struct {
	Blob                      *string                    `pulumi:"blob"`
	Kind                      *string                    `pulumi:"kind"`
	Name                      string                     `pulumi:"name"`
	PublicCertificateLocation *PublicCertificateLocation `pulumi:"publicCertificateLocation"`
	PublicCertificateName     *string                    `pulumi:"publicCertificateName"`
	ResourceGroupName         string                     `pulumi:"resourceGroupName"`
}


type WebAppPublicCertificateArgs struct {
	Blob                      pulumi.StringPtrInput
	Kind                      pulumi.StringPtrInput
	Name                      pulumi.StringInput
	PublicCertificateLocation PublicCertificateLocationPtrInput
	PublicCertificateName     pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
}

func (WebAppPublicCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPublicCertificateArgs)(nil)).Elem()
}

type WebAppPublicCertificateInput interface {
	pulumi.Input

	ToWebAppPublicCertificateOutput() WebAppPublicCertificateOutput
	ToWebAppPublicCertificateOutputWithContext(ctx context.Context) WebAppPublicCertificateOutput
}

func (*WebAppPublicCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppPublicCertificate)(nil))
}

func (i *WebAppPublicCertificate) ToWebAppPublicCertificateOutput() WebAppPublicCertificateOutput {
	return i.ToWebAppPublicCertificateOutputWithContext(context.Background())
}

func (i *WebAppPublicCertificate) ToWebAppPublicCertificateOutputWithContext(ctx context.Context) WebAppPublicCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppPublicCertificateOutput)
}

type WebAppPublicCertificateOutput struct{ *pulumi.OutputState }

func (WebAppPublicCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppPublicCertificate)(nil))
}

func (o WebAppPublicCertificateOutput) ToWebAppPublicCertificateOutput() WebAppPublicCertificateOutput {
	return o
}

func (o WebAppPublicCertificateOutput) ToWebAppPublicCertificateOutputWithContext(ctx context.Context) WebAppPublicCertificateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppPublicCertificateOutput{})
}
