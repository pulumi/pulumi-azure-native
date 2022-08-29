


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Key struct {
	pulumi.CustomResourceState

	IsActiveCMK pulumi.BoolPtrOutput   `pulumi:"isActiveCMK"`
	KeyVaultUrl pulumi.StringPtrOutput `pulumi:"keyVaultUrl"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Type        pulumi.StringOutput    `pulumi:"type"`
}


func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:Key"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:Key"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:Key"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:Key"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:Key"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:Key"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:Key"),
		},
	})
	opts = append(opts, aliases)
	var resource Key
	err := ctx.RegisterResource("azure-native:synapse/v20210601:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("azure-native:synapse/v20210601:Key", name, id, state, &resource, opts...)
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
	IsActiveCMK       *bool   `pulumi:"isActiveCMK"`
	KeyName           *string `pulumi:"keyName"`
	KeyVaultUrl       *string `pulumi:"keyVaultUrl"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type KeyArgs struct {
	IsActiveCMK       pulumi.BoolPtrInput
	KeyName           pulumi.StringPtrInput
	KeyVaultUrl       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
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
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (i *Key) ToKeyOutput() KeyOutput {
	return i.ToKeyOutputWithContext(context.Background())
}

func (i *Key) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyOutput)
}

type KeyOutput struct{ *pulumi.OutputState }

func (KeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (o KeyOutput) ToKeyOutput() KeyOutput {
	return o
}

func (o KeyOutput) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return o
}

func (o KeyOutput) IsActiveCMK() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.BoolPtrOutput { return v.IsActiveCMK }).(pulumi.BoolPtrOutput)
}

func (o KeyOutput) KeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.KeyVaultUrl }).(pulumi.StringPtrOutput)
}

func (o KeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KeyOutput{})
}
