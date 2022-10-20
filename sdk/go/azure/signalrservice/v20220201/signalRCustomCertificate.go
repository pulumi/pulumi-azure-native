


package v20220201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SignalRCustomCertificate struct {
	pulumi.CustomResourceState

	KeyVaultBaseUri       pulumi.StringOutput      `pulumi:"keyVaultBaseUri"`
	KeyVaultSecretName    pulumi.StringOutput      `pulumi:"keyVaultSecretName"`
	KeyVaultSecretVersion pulumi.StringPtrOutput   `pulumi:"keyVaultSecretVersion"`
	Name                  pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData            SystemDataResponseOutput `pulumi:"systemData"`
	Type                  pulumi.StringOutput      `pulumi:"type"`
}


func NewSignalRCustomCertificate(ctx *pulumi.Context,
	name string, args *SignalRCustomCertificateArgs, opts ...pulumi.ResourceOption) (*SignalRCustomCertificate, error) {
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:signalrservice:SignalRCustomCertificate"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220801preview:SignalRCustomCertificate"),
		},
	})
	opts = append(opts, aliases)
	var resource SignalRCustomCertificate
	err := ctx.RegisterResource("azure-native:signalrservice/v20220201:SignalRCustomCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSignalRCustomCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignalRCustomCertificateState, opts ...pulumi.ResourceOption) (*SignalRCustomCertificate, error) {
	var resource SignalRCustomCertificate
	err := ctx.ReadResource("azure-native:signalrservice/v20220201:SignalRCustomCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type signalRCustomCertificateState struct {
}

type SignalRCustomCertificateState struct {
}

func (SignalRCustomCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRCustomCertificateState)(nil)).Elem()
}

type signalRCustomCertificateArgs struct {
	CertificateName       *string `pulumi:"certificateName"`
	KeyVaultBaseUri       string  `pulumi:"keyVaultBaseUri"`
	KeyVaultSecretName    string  `pulumi:"keyVaultSecretName"`
	KeyVaultSecretVersion *string `pulumi:"keyVaultSecretVersion"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ResourceName          string  `pulumi:"resourceName"`
}


type SignalRCustomCertificateArgs struct {
	CertificateName       pulumi.StringPtrInput
	KeyVaultBaseUri       pulumi.StringInput
	KeyVaultSecretName    pulumi.StringInput
	KeyVaultSecretVersion pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	ResourceName          pulumi.StringInput
}

func (SignalRCustomCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRCustomCertificateArgs)(nil)).Elem()
}

type SignalRCustomCertificateInput interface {
	pulumi.Input

	ToSignalRCustomCertificateOutput() SignalRCustomCertificateOutput
	ToSignalRCustomCertificateOutputWithContext(ctx context.Context) SignalRCustomCertificateOutput
}

func (*SignalRCustomCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCustomCertificate)(nil)).Elem()
}

func (i *SignalRCustomCertificate) ToSignalRCustomCertificateOutput() SignalRCustomCertificateOutput {
	return i.ToSignalRCustomCertificateOutputWithContext(context.Background())
}

func (i *SignalRCustomCertificate) ToSignalRCustomCertificateOutputWithContext(ctx context.Context) SignalRCustomCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCustomCertificateOutput)
}

type SignalRCustomCertificateOutput struct{ *pulumi.OutputState }

func (SignalRCustomCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCustomCertificate)(nil)).Elem()
}

func (o SignalRCustomCertificateOutput) ToSignalRCustomCertificateOutput() SignalRCustomCertificateOutput {
	return o
}

func (o SignalRCustomCertificateOutput) ToSignalRCustomCertificateOutputWithContext(ctx context.Context) SignalRCustomCertificateOutput {
	return o
}

func (o SignalRCustomCertificateOutput) KeyVaultBaseUri() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomCertificate) pulumi.StringOutput { return v.KeyVaultBaseUri }).(pulumi.StringOutput)
}

func (o SignalRCustomCertificateOutput) KeyVaultSecretName() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomCertificate) pulumi.StringOutput { return v.KeyVaultSecretName }).(pulumi.StringOutput)
}

func (o SignalRCustomCertificateOutput) KeyVaultSecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalRCustomCertificate) pulumi.StringPtrOutput { return v.KeyVaultSecretVersion }).(pulumi.StringPtrOutput)
}

func (o SignalRCustomCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SignalRCustomCertificateOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomCertificate) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SignalRCustomCertificateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SignalRCustomCertificate) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SignalRCustomCertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomCertificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SignalRCustomCertificateOutput{})
}
