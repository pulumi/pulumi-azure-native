


package v20220301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SshPublicKey struct {
	pulumi.CustomResourceState

	Location  pulumi.StringOutput    `pulumi:"location"`
	Name      pulumi.StringOutput    `pulumi:"name"`
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
	Tags      pulumi.StringMapOutput `pulumi:"tags"`
	Type      pulumi.StringOutput    `pulumi:"type"`
}


func NewSshPublicKey(ctx *pulumi.Context,
	name string, args *SshPublicKeyArgs, opts ...pulumi.ResourceOption) (*SshPublicKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:SshPublicKey"),
		},
	})
	opts = append(opts, aliases)
	var resource SshPublicKey
	err := ctx.RegisterResource("azure-native:compute/v20220301:SshPublicKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSshPublicKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SshPublicKeyState, opts ...pulumi.ResourceOption) (*SshPublicKey, error) {
	var resource SshPublicKey
	err := ctx.ReadResource("azure-native:compute/v20220301:SshPublicKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sshPublicKeyState struct {
}

type SshPublicKeyState struct {
}

func (SshPublicKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sshPublicKeyState)(nil)).Elem()
}

type sshPublicKeyArgs struct {
	Location          *string           `pulumi:"location"`
	PublicKey         *string           `pulumi:"publicKey"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	SshPublicKeyName  *string           `pulumi:"sshPublicKeyName"`
	Tags              map[string]string `pulumi:"tags"`
}


type SshPublicKeyArgs struct {
	Location          pulumi.StringPtrInput
	PublicKey         pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	SshPublicKeyName  pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (SshPublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sshPublicKeyArgs)(nil)).Elem()
}

type SshPublicKeyInput interface {
	pulumi.Input

	ToSshPublicKeyOutput() SshPublicKeyOutput
	ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput
}

func (*SshPublicKey) ElementType() reflect.Type {
	return reflect.TypeOf((**SshPublicKey)(nil)).Elem()
}

func (i *SshPublicKey) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return i.ToSshPublicKeyOutputWithContext(context.Background())
}

func (i *SshPublicKey) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyOutput)
}

type SshPublicKeyOutput struct{ *pulumi.OutputState }

func (SshPublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshPublicKey)(nil)).Elem()
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return o
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return o
}

func (o SshPublicKeyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SshPublicKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SshPublicKeyOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

func (o SshPublicKeyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SshPublicKeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SshPublicKeyOutput{})
}
