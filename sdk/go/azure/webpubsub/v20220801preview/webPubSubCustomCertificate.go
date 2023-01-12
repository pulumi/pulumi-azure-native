


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebPubSubCustomCertificate struct {
	pulumi.CustomResourceState

	KeyVaultBaseUri       pulumi.StringOutput      `pulumi:"keyVaultBaseUri"`
	KeyVaultSecretName    pulumi.StringOutput      `pulumi:"keyVaultSecretName"`
	KeyVaultSecretVersion pulumi.StringPtrOutput   `pulumi:"keyVaultSecretVersion"`
	Name                  pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData            SystemDataResponseOutput `pulumi:"systemData"`
	Type                  pulumi.StringOutput      `pulumi:"type"`
}


func NewWebPubSubCustomCertificate(ctx *pulumi.Context,
	name string, args *WebPubSubCustomCertificateArgs, opts ...pulumi.ResourceOption) (*WebPubSubCustomCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyVaultBaseUri == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultBaseUri'")
	}
	if args.KeyVaultSecretName == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultSecretName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	var resource WebPubSubCustomCertificate
	err := ctx.RegisterResource("azure-native:webpubsub/v20220801preview:WebPubSubCustomCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebPubSubCustomCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebPubSubCustomCertificateState, opts ...pulumi.ResourceOption) (*WebPubSubCustomCertificate, error) {
	var resource WebPubSubCustomCertificate
	err := ctx.ReadResource("azure-native:webpubsub/v20220801preview:WebPubSubCustomCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webPubSubCustomCertificateState struct {
}

type WebPubSubCustomCertificateState struct {
}

func (WebPubSubCustomCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubCustomCertificateState)(nil)).Elem()
}

type webPubSubCustomCertificateArgs struct {
	CertificateName       *string `pulumi:"certificateName"`
	KeyVaultBaseUri       string  `pulumi:"keyVaultBaseUri"`
	KeyVaultSecretName    string  `pulumi:"keyVaultSecretName"`
	KeyVaultSecretVersion *string `pulumi:"keyVaultSecretVersion"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ResourceName          string  `pulumi:"resourceName"`
}


type WebPubSubCustomCertificateArgs struct {
	CertificateName       pulumi.StringPtrInput
	KeyVaultBaseUri       pulumi.StringInput
	KeyVaultSecretName    pulumi.StringInput
	KeyVaultSecretVersion pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	ResourceName          pulumi.StringInput
}

func (WebPubSubCustomCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubCustomCertificateArgs)(nil)).Elem()
}

type WebPubSubCustomCertificateInput interface {
	pulumi.Input

	ToWebPubSubCustomCertificateOutput() WebPubSubCustomCertificateOutput
	ToWebPubSubCustomCertificateOutputWithContext(ctx context.Context) WebPubSubCustomCertificateOutput
}

func (*WebPubSubCustomCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubCustomCertificate)(nil)).Elem()
}

func (i *WebPubSubCustomCertificate) ToWebPubSubCustomCertificateOutput() WebPubSubCustomCertificateOutput {
	return i.ToWebPubSubCustomCertificateOutputWithContext(context.Background())
}

func (i *WebPubSubCustomCertificate) ToWebPubSubCustomCertificateOutputWithContext(ctx context.Context) WebPubSubCustomCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubCustomCertificateOutput)
}

type WebPubSubCustomCertificateOutput struct{ *pulumi.OutputState }

func (WebPubSubCustomCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubCustomCertificate)(nil)).Elem()
}

func (o WebPubSubCustomCertificateOutput) ToWebPubSubCustomCertificateOutput() WebPubSubCustomCertificateOutput {
	return o
}

func (o WebPubSubCustomCertificateOutput) ToWebPubSubCustomCertificateOutputWithContext(ctx context.Context) WebPubSubCustomCertificateOutput {
	return o
}

func (o WebPubSubCustomCertificateOutput) KeyVaultBaseUri() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomCertificate) pulumi.StringOutput { return v.KeyVaultBaseUri }).(pulumi.StringOutput)
}

func (o WebPubSubCustomCertificateOutput) KeyVaultSecretName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomCertificate) pulumi.StringOutput { return v.KeyVaultSecretName }).(pulumi.StringOutput)
}

func (o WebPubSubCustomCertificateOutput) KeyVaultSecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebPubSubCustomCertificate) pulumi.StringPtrOutput { return v.KeyVaultSecretVersion }).(pulumi.StringPtrOutput)
}

func (o WebPubSubCustomCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebPubSubCustomCertificateOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomCertificate) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WebPubSubCustomCertificateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebPubSubCustomCertificate) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WebPubSubCustomCertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomCertificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebPubSubCustomCertificateOutput{})
}
