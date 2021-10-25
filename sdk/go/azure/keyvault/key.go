


package keyvault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Key struct {
	pulumi.CustomResourceState

	Attributes        KeyAttributesResponsePtrOutput `pulumi:"attributes"`
	CurveName         pulumi.StringPtrOutput         `pulumi:"curveName"`
	KeyOps            pulumi.StringArrayOutput       `pulumi:"keyOps"`
	KeySize           pulumi.IntPtrOutput            `pulumi:"keySize"`
	KeyUri            pulumi.StringOutput            `pulumi:"keyUri"`
	KeyUriWithVersion pulumi.StringOutput            `pulumi:"keyUriWithVersion"`
	Kty               pulumi.StringPtrOutput         `pulumi:"kty"`
	Location          pulumi.StringOutput            `pulumi:"location"`
	Name              pulumi.StringOutput            `pulumi:"name"`
	Tags              pulumi.StringMapOutput         `pulumi:"tags"`
	Type              pulumi.StringOutput            `pulumi:"type"`
}


func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:keyvault:Key"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20190901:Key"),
		},
		{
			Type: pulumi.String("azure-nextgen:keyvault/v20190901:Key"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20200401preview:Key"),
		},
		{
			Type: pulumi.String("azure-nextgen:keyvault/v20200401preview:Key"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20210401preview:Key"),
		},
		{
			Type: pulumi.String("azure-nextgen:keyvault/v20210401preview:Key"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20210601preview:Key"),
		},
		{
			Type: pulumi.String("azure-nextgen:keyvault/v20210601preview:Key"),
		},
	})
	opts = append(opts, aliases)
	var resource Key
	err := ctx.RegisterResource("azure-native:keyvault:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("azure-native:keyvault:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type keyState struct {
}

type KeyState struct {
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	KeyName           *string           `pulumi:"keyName"`
	Properties        KeyProperties     `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VaultName         string            `pulumi:"vaultName"`
}


type KeyArgs struct {
	KeyName           pulumi.StringPtrInput
	Properties        KeyPropertiesInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VaultName         pulumi.StringInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}

type KeyInput interface {
	pulumi.Input

	ToKeyOutput() KeyOutput
	ToKeyOutputWithContext(ctx context.Context) KeyOutput
}

func (*Key) ElementType() reflect.Type {
	return reflect.TypeOf((*Key)(nil))
}

func (i *Key) ToKeyOutput() KeyOutput {
	return i.ToKeyOutputWithContext(context.Background())
}

func (i *Key) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyOutput)
}

type KeyOutput struct{ *pulumi.OutputState }

func (KeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Key)(nil))
}

func (o KeyOutput) ToKeyOutput() KeyOutput {
	return o
}

func (o KeyOutput) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KeyOutput{})
}
