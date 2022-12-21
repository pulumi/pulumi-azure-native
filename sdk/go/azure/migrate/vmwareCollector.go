


package migrate

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VMwareCollector struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput            `pulumi:"eTag"`
	Name       pulumi.StringOutput               `pulumi:"name"`
	Properties CollectorPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput               `pulumi:"type"`
}


func NewVMwareCollector(ctx *pulumi.Context,
	name string, args *VMwareCollectorArgs, opts ...pulumi.ResourceOption) (*VMwareCollector, error) {
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
			Type: pulumi.String("azure-native:migrate/v20191001:VMwareCollector"),
		},
	})
	opts = append(opts, aliases)
	var resource VMwareCollector
	err := ctx.RegisterResource("azure-native:migrate:VMwareCollector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVMwareCollector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VMwareCollectorState, opts ...pulumi.ResourceOption) (*VMwareCollector, error) {
	var resource VMwareCollector
	err := ctx.ReadResource("azure-native:migrate:VMwareCollector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type vmwareCollectorState struct {
}

type VMwareCollectorState struct {
}

func (VMwareCollectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmwareCollectorState)(nil)).Elem()
}

type vmwareCollectorArgs struct {
	ETag                *string              `pulumi:"eTag"`
	ProjectName         string               `pulumi:"projectName"`
	Properties          *CollectorProperties `pulumi:"properties"`
	ResourceGroupName   string               `pulumi:"resourceGroupName"`
	VmWareCollectorName *string              `pulumi:"vmWareCollectorName"`
}


type VMwareCollectorArgs struct {
	ETag                pulumi.StringPtrInput
	ProjectName         pulumi.StringInput
	Properties          CollectorPropertiesPtrInput
	ResourceGroupName   pulumi.StringInput
	VmWareCollectorName pulumi.StringPtrInput
}

func (VMwareCollectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmwareCollectorArgs)(nil)).Elem()
}

type VMwareCollectorInput interface {
	pulumi.Input

	ToVMwareCollectorOutput() VMwareCollectorOutput
	ToVMwareCollectorOutputWithContext(ctx context.Context) VMwareCollectorOutput
}

func (*VMwareCollector) ElementType() reflect.Type {
	return reflect.TypeOf((**VMwareCollector)(nil)).Elem()
}

func (i *VMwareCollector) ToVMwareCollectorOutput() VMwareCollectorOutput {
	return i.ToVMwareCollectorOutputWithContext(context.Background())
}

func (i *VMwareCollector) ToVMwareCollectorOutputWithContext(ctx context.Context) VMwareCollectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCollectorOutput)
}

type VMwareCollectorOutput struct{ *pulumi.OutputState }

func (VMwareCollectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMwareCollector)(nil)).Elem()
}

func (o VMwareCollectorOutput) ToVMwareCollectorOutput() VMwareCollectorOutput {
	return o
}

func (o VMwareCollectorOutput) ToVMwareCollectorOutputWithContext(ctx context.Context) VMwareCollectorOutput {
	return o
}

func (o VMwareCollectorOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCollector) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o VMwareCollectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VMwareCollector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VMwareCollectorOutput) Properties() CollectorPropertiesResponseOutput {
	return o.ApplyT(func(v *VMwareCollector) CollectorPropertiesResponseOutput { return v.Properties }).(CollectorPropertiesResponseOutput)
}

func (o VMwareCollectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VMwareCollector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VMwareCollectorOutput{})
}
