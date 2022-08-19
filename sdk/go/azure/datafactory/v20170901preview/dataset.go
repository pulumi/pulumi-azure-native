


package v20170901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Dataset struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput `pulumi:"etag"`
	Name       pulumi.StringOutput `pulumi:"name"`
	Properties pulumi.AnyOutput    `pulumi:"properties"`
	Type       pulumi.StringOutput `pulumi:"type"`
}


func NewDataset(ctx *pulumi.Context,
	name string, args *DatasetArgs, opts ...pulumi.ResourceOption) (*Dataset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FactoryName == nil {
		return nil, errors.New("invalid value for required argument 'FactoryName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datafactory:Dataset"),
		},
		{
			Type: pulumi.String("azure-native:datafactory/v20180601:Dataset"),
		},
	})
	opts = append(opts, aliases)
	var resource Dataset
	err := ctx.RegisterResource("azure-native:datafactory/v20170901preview:Dataset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetState, opts ...pulumi.ResourceOption) (*Dataset, error) {
	var resource Dataset
	err := ctx.ReadResource("azure-native:datafactory/v20170901preview:Dataset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type datasetState struct {
}

type DatasetState struct {
}

func (DatasetState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetState)(nil)).Elem()
}

type datasetArgs struct {
	DatasetName       *string     `pulumi:"datasetName"`
	FactoryName       string      `pulumi:"factoryName"`
	Properties        interface{} `pulumi:"properties"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
}


type DatasetArgs struct {
	DatasetName       pulumi.StringPtrInput
	FactoryName       pulumi.StringInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
}

func (DatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetArgs)(nil)).Elem()
}

type DatasetInput interface {
	pulumi.Input

	ToDatasetOutput() DatasetOutput
	ToDatasetOutputWithContext(ctx context.Context) DatasetOutput
}

func (*Dataset) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (i *Dataset) ToDatasetOutput() DatasetOutput {
	return i.ToDatasetOutputWithContext(context.Background())
}

func (i *Dataset) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetOutput)
}

type DatasetOutput struct{ *pulumi.OutputState }

func (DatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (o DatasetOutput) ToDatasetOutput() DatasetOutput {
	return o
}

func (o DatasetOutput) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return o
}

func (o DatasetOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DatasetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatasetOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Dataset) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o DatasetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatasetOutput{})
}
