


package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DataSetMapping struct {
	pulumi.CustomResourceState

	Kind pulumi.StringOutput `pulumi:"kind"`
	Name pulumi.StringOutput `pulumi:"name"`
	Type pulumi.StringOutput `pulumi:"type"`
}


func NewDataSetMapping(ctx *pulumi.Context,
	name string, args *DataSetMappingArgs, opts ...pulumi.ResourceOption) (*DataSetMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'ShareSubscriptionName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datashare/v20181101preview:DataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:DataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare:DataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:DataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20191101:DataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:DataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20200901:DataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:DataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20201001preview:DataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:DataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20210801:DataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource DataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20181101preview:DataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSetMappingState, opts ...pulumi.ResourceOption) (*DataSetMapping, error) {
	var resource DataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20181101preview:DataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataSetMappingState struct {
}

type DataSetMappingState struct {
}

func (DataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSetMappingState)(nil)).Elem()
}

type dataSetMappingArgs struct {
	AccountName           string  `pulumi:"accountName"`
	DataSetMappingName    *string `pulumi:"dataSetMappingName"`
	Kind                  string  `pulumi:"kind"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
}


type DataSetMappingArgs struct {
	AccountName           pulumi.StringInput
	DataSetMappingName    pulumi.StringPtrInput
	Kind                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	ShareSubscriptionName pulumi.StringInput
}

func (DataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSetMappingArgs)(nil)).Elem()
}

type DataSetMappingInput interface {
	pulumi.Input

	ToDataSetMappingOutput() DataSetMappingOutput
	ToDataSetMappingOutputWithContext(ctx context.Context) DataSetMappingOutput
}

func (*DataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSetMapping)(nil))
}

func (i *DataSetMapping) ToDataSetMappingOutput() DataSetMappingOutput {
	return i.ToDataSetMappingOutputWithContext(context.Background())
}

func (i *DataSetMapping) ToDataSetMappingOutputWithContext(ctx context.Context) DataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSetMappingOutput)
}

type DataSetMappingOutput struct{ *pulumi.OutputState }

func (DataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSetMapping)(nil))
}

func (o DataSetMappingOutput) ToDataSetMappingOutput() DataSetMappingOutput {
	return o
}

func (o DataSetMappingOutput) ToDataSetMappingOutputWithContext(ctx context.Context) DataSetMappingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DataSetMappingOutput{})
}
