


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectedEnvironmentsCertificate struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                 `pulumi:"location"`
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties CertificateResponsePropertiesOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput            `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput              `pulumi:"tags"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewConnectedEnvironmentsCertificate(ctx *pulumi.Context,
	name string, args *ConnectedEnvironmentsCertificateArgs, opts ...pulumi.ResourceOption) (*ConnectedEnvironmentsCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectedEnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectedEnvironmentName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app/v20221001:ConnectedEnvironmentsCertificate"),
		},
	})
	opts = append(opts, aliases)
	var resource ConnectedEnvironmentsCertificate
	err := ctx.RegisterResource("azure-native:app/v20220601preview:ConnectedEnvironmentsCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConnectedEnvironmentsCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectedEnvironmentsCertificateState, opts ...pulumi.ResourceOption) (*ConnectedEnvironmentsCertificate, error) {
	var resource ConnectedEnvironmentsCertificate
	err := ctx.ReadResource("azure-native:app/v20220601preview:ConnectedEnvironmentsCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type connectedEnvironmentsCertificateState struct {
}

type ConnectedEnvironmentsCertificateState struct {
}

func (ConnectedEnvironmentsCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentsCertificateState)(nil)).Elem()
}

type connectedEnvironmentsCertificateArgs struct {
	CertificateName          *string                `pulumi:"certificateName"`
	ConnectedEnvironmentName string                 `pulumi:"connectedEnvironmentName"`
	Location                 *string                `pulumi:"location"`
	Properties               *CertificateProperties `pulumi:"properties"`
	ResourceGroupName        string                 `pulumi:"resourceGroupName"`
	Tags                     map[string]string      `pulumi:"tags"`
}


type ConnectedEnvironmentsCertificateArgs struct {
	CertificateName          pulumi.StringPtrInput
	ConnectedEnvironmentName pulumi.StringInput
	Location                 pulumi.StringPtrInput
	Properties               CertificatePropertiesPtrInput
	ResourceGroupName        pulumi.StringInput
	Tags                     pulumi.StringMapInput
}

func (ConnectedEnvironmentsCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentsCertificateArgs)(nil)).Elem()
}

type ConnectedEnvironmentsCertificateInput interface {
	pulumi.Input

	ToConnectedEnvironmentsCertificateOutput() ConnectedEnvironmentsCertificateOutput
	ToConnectedEnvironmentsCertificateOutputWithContext(ctx context.Context) ConnectedEnvironmentsCertificateOutput
}

func (*ConnectedEnvironmentsCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironmentsCertificate)(nil)).Elem()
}

func (i *ConnectedEnvironmentsCertificate) ToConnectedEnvironmentsCertificateOutput() ConnectedEnvironmentsCertificateOutput {
	return i.ToConnectedEnvironmentsCertificateOutputWithContext(context.Background())
}

func (i *ConnectedEnvironmentsCertificate) ToConnectedEnvironmentsCertificateOutputWithContext(ctx context.Context) ConnectedEnvironmentsCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedEnvironmentsCertificateOutput)
}

type ConnectedEnvironmentsCertificateOutput struct{ *pulumi.OutputState }

func (ConnectedEnvironmentsCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironmentsCertificate)(nil)).Elem()
}

func (o ConnectedEnvironmentsCertificateOutput) ToConnectedEnvironmentsCertificateOutput() ConnectedEnvironmentsCertificateOutput {
	return o
}

func (o ConnectedEnvironmentsCertificateOutput) ToConnectedEnvironmentsCertificateOutputWithContext(ctx context.Context) ConnectedEnvironmentsCertificateOutput {
	return o
}

func (o ConnectedEnvironmentsCertificateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsCertificate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ConnectedEnvironmentsCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectedEnvironmentsCertificateOutput) Properties() CertificateResponsePropertiesOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsCertificate) CertificateResponsePropertiesOutput { return v.Properties }).(CertificateResponsePropertiesOutput)
}

func (o ConnectedEnvironmentsCertificateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsCertificate) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConnectedEnvironmentsCertificateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsCertificate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ConnectedEnvironmentsCertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsCertificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectedEnvironmentsCertificateOutput{})
}
