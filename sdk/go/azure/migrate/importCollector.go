


package migrate

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ImportCollector struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput                  `pulumi:"eTag"`
	Name       pulumi.StringOutput                     `pulumi:"name"`
	Properties ImportCollectorPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewImportCollector(ctx *pulumi.Context,
	name string, args *ImportCollectorArgs, opts ...pulumi.ResourceOption) (*ImportCollector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20191001:ImportCollector"),
		},
	})
	opts = append(opts, aliases)
	var resource ImportCollector
	err := ctx.RegisterResource("azure-native:migrate:ImportCollector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetImportCollector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImportCollectorState, opts ...pulumi.ResourceOption) (*ImportCollector, error) {
	var resource ImportCollector
	err := ctx.ReadResource("azure-native:migrate:ImportCollector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type importCollectorState struct {
}

type ImportCollectorState struct {
}

func (ImportCollectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*importCollectorState)(nil)).Elem()
}

type importCollectorArgs struct {
	ETag                *string                    `pulumi:"eTag"`
	ImportCollectorName *string                    `pulumi:"importCollectorName"`
	ProjectName         string                     `pulumi:"projectName"`
	Properties          *ImportCollectorProperties `pulumi:"properties"`
	ResourceGroupName   string                     `pulumi:"resourceGroupName"`
}


type ImportCollectorArgs struct {
	ETag                pulumi.StringPtrInput
	ImportCollectorName pulumi.StringPtrInput
	ProjectName         pulumi.StringInput
	Properties          ImportCollectorPropertiesPtrInput
	ResourceGroupName   pulumi.StringInput
}

func (ImportCollectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*importCollectorArgs)(nil)).Elem()
}

type ImportCollectorInput interface {
	pulumi.Input

	ToImportCollectorOutput() ImportCollectorOutput
	ToImportCollectorOutputWithContext(ctx context.Context) ImportCollectorOutput
}

func (*ImportCollector) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportCollector)(nil)).Elem()
}

func (i *ImportCollector) ToImportCollectorOutput() ImportCollectorOutput {
	return i.ToImportCollectorOutputWithContext(context.Background())
}

func (i *ImportCollector) ToImportCollectorOutputWithContext(ctx context.Context) ImportCollectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportCollectorOutput)
}

type ImportCollectorOutput struct{ *pulumi.OutputState }

func (ImportCollectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportCollector)(nil)).Elem()
}

func (o ImportCollectorOutput) ToImportCollectorOutput() ImportCollectorOutput {
	return o
}

func (o ImportCollectorOutput) ToImportCollectorOutputWithContext(ctx context.Context) ImportCollectorOutput {
	return o
}

func (o ImportCollectorOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportCollector) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o ImportCollectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ImportCollector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ImportCollectorOutput) Properties() ImportCollectorPropertiesResponseOutput {
	return o.ApplyT(func(v *ImportCollector) ImportCollectorPropertiesResponseOutput { return v.Properties }).(ImportCollectorPropertiesResponseOutput)
}

func (o ImportCollectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ImportCollector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ImportCollectorOutput{})
}
