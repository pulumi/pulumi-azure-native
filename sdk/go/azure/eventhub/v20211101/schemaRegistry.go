


package v20211101

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
			Type: pulumi.String("azure-nextgen:eventhub/v20211101:SchemaRegistry"),
		},
		{
			Type: pulumi.String("azure-native:eventhub:SchemaRegistry"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventhub:SchemaRegistry"),
		},
	})
	opts = append(opts, aliases)
	var resource SchemaRegistry
	err := ctx.RegisterResource("azure-native:eventhub/v20211101:SchemaRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSchemaRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaRegistryState, opts ...pulumi.ResourceOption) (*SchemaRegistry, error) {
	var resource SchemaRegistry
	err := ctx.ReadResource("azure-native:eventhub/v20211101:SchemaRegistry", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*SchemaRegistry)(nil))
}

func (i *SchemaRegistry) ToSchemaRegistryOutput() SchemaRegistryOutput {
	return i.ToSchemaRegistryOutputWithContext(context.Background())
}

func (i *SchemaRegistry) ToSchemaRegistryOutputWithContext(ctx context.Context) SchemaRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaRegistryOutput)
}

type SchemaRegistryOutput struct{ *pulumi.OutputState }

func (SchemaRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaRegistry)(nil))
}

func (o SchemaRegistryOutput) ToSchemaRegistryOutput() SchemaRegistryOutput {
	return o
}

func (o SchemaRegistryOutput) ToSchemaRegistryOutputWithContext(ctx context.Context) SchemaRegistryOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaRegistryInput)(nil)).Elem(), &SchemaRegistry{})
	pulumi.RegisterOutputType(SchemaRegistryOutput{})
}
