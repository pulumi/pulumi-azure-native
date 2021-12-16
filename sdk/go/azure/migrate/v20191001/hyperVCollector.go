


package v20191001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HyperVCollector struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput            `pulumi:"eTag"`
	Name       pulumi.StringOutput               `pulumi:"name"`
	Properties CollectorPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput               `pulumi:"type"`
}


func NewHyperVCollector(ctx *pulumi.Context,
	name string, args *HyperVCollectorArgs, opts ...pulumi.ResourceOption) (*HyperVCollector, error) {
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
			Type: pulumi.String("azure-native:migrate:HyperVCollector"),
		},
	})
	opts = append(opts, aliases)
	var resource HyperVCollector
	err := ctx.RegisterResource("azure-native:migrate/v20191001:HyperVCollector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHyperVCollector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HyperVCollectorState, opts ...pulumi.ResourceOption) (*HyperVCollector, error) {
	var resource HyperVCollector
	err := ctx.ReadResource("azure-native:migrate/v20191001:HyperVCollector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hyperVCollectorState struct {
}

type HyperVCollectorState struct {
}

func (HyperVCollectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*hyperVCollectorState)(nil)).Elem()
}

type hyperVCollectorArgs struct {
	ETag                *string              `pulumi:"eTag"`
	HyperVCollectorName *string              `pulumi:"hyperVCollectorName"`
	ProjectName         string               `pulumi:"projectName"`
	Properties          *CollectorProperties `pulumi:"properties"`
	ResourceGroupName   string               `pulumi:"resourceGroupName"`
}


type HyperVCollectorArgs struct {
	ETag                pulumi.StringPtrInput
	HyperVCollectorName pulumi.StringPtrInput
	ProjectName         pulumi.StringInput
	Properties          CollectorPropertiesPtrInput
	ResourceGroupName   pulumi.StringInput
}

func (HyperVCollectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hyperVCollectorArgs)(nil)).Elem()
}

type HyperVCollectorInput interface {
	pulumi.Input

	ToHyperVCollectorOutput() HyperVCollectorOutput
	ToHyperVCollectorOutputWithContext(ctx context.Context) HyperVCollectorOutput
}

func (*HyperVCollector) ElementType() reflect.Type {
	return reflect.TypeOf((**HyperVCollector)(nil)).Elem()
}

func (i *HyperVCollector) ToHyperVCollectorOutput() HyperVCollectorOutput {
	return i.ToHyperVCollectorOutputWithContext(context.Background())
}

func (i *HyperVCollector) ToHyperVCollectorOutputWithContext(ctx context.Context) HyperVCollectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVCollectorOutput)
}

type HyperVCollectorOutput struct{ *pulumi.OutputState }

func (HyperVCollectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HyperVCollector)(nil)).Elem()
}

func (o HyperVCollectorOutput) ToHyperVCollectorOutput() HyperVCollectorOutput {
	return o
}

func (o HyperVCollectorOutput) ToHyperVCollectorOutputWithContext(ctx context.Context) HyperVCollectorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HyperVCollectorOutput{})
}
