


package eventhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SchemaRegistry struct {
	pulumi.CustomResourceState

	CreatedAtUtc        pulumi.StringOutput      `pulumi:"createdAtUtc"`
	ETag                pulumi.StringOutput      `pulumi:"eTag"`
	GroupProperties     pulumi.StringMapOutput   `pulumi:"groupProperties"`
	Location            pulumi.StringOutput      `pulumi:"location"`
	Name                pulumi.StringOutput      `pulumi:"name"`
	SchemaCompatibility pulumi.StringPtrOutput   `pulumi:"schemaCompatibility"`
	SchemaType          pulumi.StringPtrOutput   `pulumi:"schemaType"`
	SystemData          SystemDataResponseOutput `pulumi:"systemData"`
	Type                pulumi.StringOutput      `pulumi:"type"`
	UpdatedAtUtc        pulumi.StringOutput      `pulumi:"updatedAtUtc"`
}


func NewSchemaRegistry(ctx *pulumi.Context,
	name string, args *SchemaRegistryArgs, opts ...pulumi.ResourceOption) (*SchemaRegistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventhub/v20211101:SchemaRegistry"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20220101preview:SchemaRegistry"),
		},
	})
	opts = append(opts, aliases)
	var resource SchemaRegistry
	err := ctx.RegisterResource("azure-native:eventhub:SchemaRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSchemaRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaRegistryState, opts ...pulumi.ResourceOption) (*SchemaRegistry, error) {
	var resource SchemaRegistry
	err := ctx.ReadResource("azure-native:eventhub:SchemaRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type schemaRegistryState struct {
}

type SchemaRegistryState struct {
}

func (SchemaRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaRegistryState)(nil)).Elem()
}

type schemaRegistryArgs struct {
	GroupProperties     map[string]string `pulumi:"groupProperties"`
	NamespaceName       string            `pulumi:"namespaceName"`
	ResourceGroupName   string            `pulumi:"resourceGroupName"`
	SchemaCompatibility *string           `pulumi:"schemaCompatibility"`
	SchemaGroupName     *string           `pulumi:"schemaGroupName"`
	SchemaType          *string           `pulumi:"schemaType"`
}


type SchemaRegistryArgs struct {
	GroupProperties     pulumi.StringMapInput
	NamespaceName       pulumi.StringInput
	ResourceGroupName   pulumi.StringInput
	SchemaCompatibility pulumi.StringPtrInput
	SchemaGroupName     pulumi.StringPtrInput
	SchemaType          pulumi.StringPtrInput
}

func (SchemaRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaRegistryArgs)(nil)).Elem()
}

type SchemaRegistryInput interface {
	pulumi.Input

	ToSchemaRegistryOutput() SchemaRegistryOutput
	ToSchemaRegistryOutputWithContext(ctx context.Context) SchemaRegistryOutput
}

func (*SchemaRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaRegistry)(nil)).Elem()
}

func (i *SchemaRegistry) ToSchemaRegistryOutput() SchemaRegistryOutput {
	return i.ToSchemaRegistryOutputWithContext(context.Background())
}

func (i *SchemaRegistry) ToSchemaRegistryOutputWithContext(ctx context.Context) SchemaRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaRegistryOutput)
}

type SchemaRegistryOutput struct{ *pulumi.OutputState }

func (SchemaRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaRegistry)(nil)).Elem()
}

func (o SchemaRegistryOutput) ToSchemaRegistryOutput() SchemaRegistryOutput {
	return o
}

func (o SchemaRegistryOutput) ToSchemaRegistryOutputWithContext(ctx context.Context) SchemaRegistryOutput {
	return o
}

func (o SchemaRegistryOutput) CreatedAtUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringOutput { return v.CreatedAtUtc }).(pulumi.StringOutput)
}

func (o SchemaRegistryOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

func (o SchemaRegistryOutput) GroupProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringMapOutput { return v.GroupProperties }).(pulumi.StringMapOutput)
}

func (o SchemaRegistryOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SchemaRegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SchemaRegistryOutput) SchemaCompatibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringPtrOutput { return v.SchemaCompatibility }).(pulumi.StringPtrOutput)
}

func (o SchemaRegistryOutput) SchemaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringPtrOutput { return v.SchemaType }).(pulumi.StringPtrOutput)
}

func (o SchemaRegistryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SchemaRegistry) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SchemaRegistryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SchemaRegistryOutput) UpdatedAtUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringOutput { return v.UpdatedAtUtc }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SchemaRegistryOutput{})
}
