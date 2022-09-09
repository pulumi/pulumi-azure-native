


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Schema struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput `pulumi:"description"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	SchemaType  pulumi.StringOutput    `pulumi:"schemaType"`
	Type        pulumi.StringOutput    `pulumi:"type"`
	Value       pulumi.StringPtrOutput `pulumi:"value"`
}


func NewSchema(ctx *pulumi.Context,
	name string, args *SchemaArgs, opts ...pulumi.ResourceOption) (*Schema, error) {
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
			Type: pulumi.String("azure-native:apimanagement:Schema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Schema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Schema"),
		},
	})
	opts = append(opts, aliases)
	var resource Schema
	err := ctx.RegisterResource("azure-native:apimanagement/v20210401preview:Schema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaState, opts ...pulumi.ResourceOption) (*Schema, error) {
	var resource Schema
	err := ctx.ReadResource("azure-native:apimanagement/v20210401preview:Schema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type schemaState struct {
}

type SchemaState struct {
}

func (SchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaState)(nil)).Elem()
}

type schemaArgs struct {
	Description       *string `pulumi:"description"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SchemaId          *string `pulumi:"schemaId"`
	SchemaType        string  `pulumi:"schemaType"`
	ServiceName       string  `pulumi:"serviceName"`
	Value             *string `pulumi:"value"`
}


type SchemaArgs struct {
	Description       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	SchemaId          pulumi.StringPtrInput
	SchemaType        pulumi.StringInput
	ServiceName       pulumi.StringInput
	Value             pulumi.StringPtrInput
}

func (SchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaArgs)(nil)).Elem()
}

type SchemaInput interface {
	pulumi.Input

	ToSchemaOutput() SchemaOutput
	ToSchemaOutputWithContext(ctx context.Context) SchemaOutput
}

func (*Schema) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (i *Schema) ToSchemaOutput() SchemaOutput {
	return i.ToSchemaOutputWithContext(context.Background())
}

func (i *Schema) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaOutput)
}

type SchemaOutput struct{ *pulumi.OutputState }

func (SchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (o SchemaOutput) ToSchemaOutput() SchemaOutput {
	return o
}

func (o SchemaOutput) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return o
}

func (o SchemaOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SchemaOutput) SchemaType() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.SchemaType }).(pulumi.StringOutput)
}

func (o SchemaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SchemaOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SchemaOutput{})
}
