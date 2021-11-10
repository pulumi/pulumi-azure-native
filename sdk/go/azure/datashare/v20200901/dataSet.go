


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DataSet struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewDataSet(ctx *pulumi.Context,
	name string, args *DataSetArgs, opts ...pulumi.ResourceOption) (*DataSet, error) {
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
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:DataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:DataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:DataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:DataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:DataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource DataSet
	err := ctx.RegisterResource("azure-native:datashare/v20200901:DataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSetState, opts ...pulumi.ResourceOption) (*DataSet, error) {
	var resource DataSet
	err := ctx.ReadResource("azure-native:datashare/v20200901:DataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataSetState struct {
}

type DataSetState struct {
}

func (DataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSetState)(nil)).Elem()
}

type dataSetArgs struct {
	AccountName       string  `pulumi:"accountName"`
	DataSetName       *string `pulumi:"dataSetName"`
	Kind              string  `pulumi:"kind"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ShareName         string  `pulumi:"shareName"`
}


type DataSetArgs struct {
	AccountName       pulumi.StringInput
	DataSetName       pulumi.StringPtrInput
	Kind              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ShareName         pulumi.StringInput
}

func (DataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSetArgs)(nil)).Elem()
}

type DataSetInput interface {
	pulumi.Input

	ToDataSetOutput() DataSetOutput
	ToDataSetOutputWithContext(ctx context.Context) DataSetOutput
}

func (*DataSet) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSet)(nil))
}

func (i *DataSet) ToDataSetOutput() DataSetOutput {
	return i.ToDataSetOutputWithContext(context.Background())
}

func (i *DataSet) ToDataSetOutputWithContext(ctx context.Context) DataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSetOutput)
}

type DataSetOutput struct{ *pulumi.OutputState }

func (DataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSet)(nil))
}

func (o DataSetOutput) ToDataSetOutput() DataSetOutput {
	return o
}

func (o DataSetOutput) ToDataSetOutputWithContext(ctx context.Context) DataSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DataSetOutput{})
}
