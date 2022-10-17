


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GlobalSchema struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput `pulumi:"description"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	SchemaType  pulumi.StringOutput    `pulumi:"schemaType"`
	Type        pulumi.StringOutput    `pulumi:"type"`
	Value       pulumi.AnyOutput       `pulumi:"value"`
}


func NewGlobalSchema(ctx *pulumi.Context,
	name string, args *GlobalSchemaArgs, opts ...pulumi.ResourceOption) (*GlobalSchema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SchemaType == nil {
		return nil, errors.New("invalid value for required argument 'SchemaType'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:GlobalSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:GlobalSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:GlobalSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:GlobalSchema"),
		},
	})
	opts = append(opts, aliases)
	var resource GlobalSchema
	err := ctx.RegisterResource("azure-native:apimanagement/v20211201preview:GlobalSchema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGlobalSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalSchemaState, opts ...pulumi.ResourceOption) (*GlobalSchema, error) {
	var resource GlobalSchema
	err := ctx.ReadResource("azure-native:apimanagement/v20211201preview:GlobalSchema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type globalSchemaState struct {
}

type GlobalSchemaState struct {
}

func (GlobalSchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalSchemaState)(nil)).Elem()
}

type globalSchemaArgs struct {
	Description       *string     `pulumi:"description"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
	SchemaId          *string     `pulumi:"schemaId"`
	SchemaType        string      `pulumi:"schemaType"`
	ServiceName       string      `pulumi:"serviceName"`
	Value             interface{} `pulumi:"value"`
}


type GlobalSchemaArgs struct {
	Description       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	SchemaId          pulumi.StringPtrInput
	SchemaType        pulumi.StringInput
	ServiceName       pulumi.StringInput
	Value             pulumi.Input
}

func (GlobalSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalSchemaArgs)(nil)).Elem()
}

type GlobalSchemaInput interface {
	pulumi.Input

	ToGlobalSchemaOutput() GlobalSchemaOutput
	ToGlobalSchemaOutputWithContext(ctx context.Context) GlobalSchemaOutput
}

func (*GlobalSchema) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalSchema)(nil)).Elem()
}

func (i *GlobalSchema) ToGlobalSchemaOutput() GlobalSchemaOutput {
	return i.ToGlobalSchemaOutputWithContext(context.Background())
}

func (i *GlobalSchema) ToGlobalSchemaOutputWithContext(ctx context.Context) GlobalSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalSchemaOutput)
}

type GlobalSchemaOutput struct{ *pulumi.OutputState }

func (GlobalSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalSchema)(nil)).Elem()
}

func (o GlobalSchemaOutput) ToGlobalSchemaOutput() GlobalSchemaOutput {
	return o
}

func (o GlobalSchemaOutput) ToGlobalSchemaOutputWithContext(ctx context.Context) GlobalSchemaOutput {
	return o
}

func (o GlobalSchemaOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlobalSchema) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GlobalSchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalSchema) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GlobalSchemaOutput) SchemaType() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalSchema) pulumi.StringOutput { return v.SchemaType }).(pulumi.StringOutput)
}

func (o GlobalSchemaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalSchema) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o GlobalSchemaOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v *GlobalSchema) pulumi.AnyOutput { return v.Value }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(GlobalSchemaOutput{})
}
