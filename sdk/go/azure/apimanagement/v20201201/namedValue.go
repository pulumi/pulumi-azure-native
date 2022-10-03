


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NamedValue struct {
	pulumi.CustomResourceState

	DisplayName pulumi.StringOutput                         `pulumi:"displayName"`
	KeyVault    KeyVaultContractPropertiesResponsePtrOutput `pulumi:"keyVault"`
	Name        pulumi.StringOutput                         `pulumi:"name"`
	Secret      pulumi.BoolPtrOutput                        `pulumi:"secret"`
	Tags        pulumi.StringArrayOutput                    `pulumi:"tags"`
	Type        pulumi.StringOutput                         `pulumi:"type"`
	Value       pulumi.StringPtrOutput                      `pulumi:"value"`
}


func NewNamedValue(ctx *pulumi.Context,
	name string, args *NamedValueArgs, opts ...pulumi.ResourceOption) (*NamedValue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:NamedValue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:NamedValue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:NamedValue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:NamedValue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:NamedValue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:NamedValue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:NamedValue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:NamedValue"),
		},
	})
	opts = append(opts, aliases)
	var resource NamedValue
	err := ctx.RegisterResource("azure-native:apimanagement/v20201201:NamedValue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNamedValue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamedValueState, opts ...pulumi.ResourceOption) (*NamedValue, error) {
	var resource NamedValue
	err := ctx.ReadResource("azure-native:apimanagement/v20201201:NamedValue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type namedValueState struct {
}

type NamedValueState struct {
}

func (NamedValueState) ElementType() reflect.Type {
	return reflect.TypeOf((*namedValueState)(nil)).Elem()
}

type namedValueArgs struct {
	DisplayName       string                            `pulumi:"displayName"`
	KeyVault          *KeyVaultContractCreateProperties `pulumi:"keyVault"`
	NamedValueId      *string                           `pulumi:"namedValueId"`
	ResourceGroupName string                            `pulumi:"resourceGroupName"`
	Secret            *bool                             `pulumi:"secret"`
	ServiceName       string                            `pulumi:"serviceName"`
	Tags              []string                          `pulumi:"tags"`
	Value             *string                           `pulumi:"value"`
}


type NamedValueArgs struct {
	DisplayName       pulumi.StringInput
	KeyVault          KeyVaultContractCreatePropertiesPtrInput
	NamedValueId      pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Secret            pulumi.BoolPtrInput
	ServiceName       pulumi.StringInput
	Tags              pulumi.StringArrayInput
	Value             pulumi.StringPtrInput
}

func (NamedValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namedValueArgs)(nil)).Elem()
}

type NamedValueInput interface {
	pulumi.Input

	ToNamedValueOutput() NamedValueOutput
	ToNamedValueOutputWithContext(ctx context.Context) NamedValueOutput
}

func (*NamedValue) ElementType() reflect.Type {
	return reflect.TypeOf((**NamedValue)(nil)).Elem()
}

func (i *NamedValue) ToNamedValueOutput() NamedValueOutput {
	return i.ToNamedValueOutputWithContext(context.Background())
}

func (i *NamedValue) ToNamedValueOutputWithContext(ctx context.Context) NamedValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedValueOutput)
}

type NamedValueOutput struct{ *pulumi.OutputState }

func (NamedValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamedValue)(nil)).Elem()
}

func (o NamedValueOutput) ToNamedValueOutput() NamedValueOutput {
	return o
}

func (o NamedValueOutput) ToNamedValueOutputWithContext(ctx context.Context) NamedValueOutput {
	return o
}

func (o NamedValueOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *NamedValue) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o NamedValueOutput) KeyVault() KeyVaultContractPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *NamedValue) KeyVaultContractPropertiesResponsePtrOutput { return v.KeyVault }).(KeyVaultContractPropertiesResponsePtrOutput)
}

func (o NamedValueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NamedValue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NamedValueOutput) Secret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NamedValue) pulumi.BoolPtrOutput { return v.Secret }).(pulumi.BoolPtrOutput)
}

func (o NamedValueOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NamedValue) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o NamedValueOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NamedValue) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o NamedValueOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamedValue) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NamedValueOutput{})
}
