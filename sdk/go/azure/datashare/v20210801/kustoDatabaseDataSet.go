


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KustoDatabaseDataSet struct {
	pulumi.CustomResourceState

	DataSetId               pulumi.StringOutput      `pulumi:"dataSetId"`
	Kind                    pulumi.StringOutput      `pulumi:"kind"`
	KustoDatabaseResourceId pulumi.StringOutput      `pulumi:"kustoDatabaseResourceId"`
	Location                pulumi.StringOutput      `pulumi:"location"`
	Name                    pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState       pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData              SystemDataResponseOutput `pulumi:"systemData"`
	Type                    pulumi.StringOutput      `pulumi:"type"`
}


func NewKustoDatabaseDataSet(ctx *pulumi.Context,
	name string, args *KustoDatabaseDataSetArgs, opts ...pulumi.ResourceOption) (*KustoDatabaseDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.KustoDatabaseResourceId == nil {
		return nil, errors.New("invalid value for required argument 'KustoDatabaseResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	args.Kind = pulumi.String("KustoDatabase")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:KustoDatabaseDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:KustoDatabaseDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:KustoDatabaseDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:KustoDatabaseDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:KustoDatabaseDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoDatabaseDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20210801:KustoDatabaseDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKustoDatabaseDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoDatabaseDataSetState, opts ...pulumi.ResourceOption) (*KustoDatabaseDataSet, error) {
	var resource KustoDatabaseDataSet
	err := ctx.ReadResource("azure-native:datashare/v20210801:KustoDatabaseDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type kustoDatabaseDataSetState struct {
}

type KustoDatabaseDataSetState struct {
}

func (KustoDatabaseDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoDatabaseDataSetState)(nil)).Elem()
}

type kustoDatabaseDataSetArgs struct {
	AccountName             string  `pulumi:"accountName"`
	DataSetName             *string `pulumi:"dataSetName"`
	Kind                    string  `pulumi:"kind"`
	KustoDatabaseResourceId string  `pulumi:"kustoDatabaseResourceId"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
	ShareName               string  `pulumi:"shareName"`
}


type KustoDatabaseDataSetArgs struct {
	AccountName             pulumi.StringInput
	DataSetName             pulumi.StringPtrInput
	Kind                    pulumi.StringInput
	KustoDatabaseResourceId pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
	ShareName               pulumi.StringInput
}

func (KustoDatabaseDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoDatabaseDataSetArgs)(nil)).Elem()
}

type KustoDatabaseDataSetInput interface {
	pulumi.Input

	ToKustoDatabaseDataSetOutput() KustoDatabaseDataSetOutput
	ToKustoDatabaseDataSetOutputWithContext(ctx context.Context) KustoDatabaseDataSetOutput
}

func (*KustoDatabaseDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((*KustoDatabaseDataSet)(nil))
}

func (i *KustoDatabaseDataSet) ToKustoDatabaseDataSetOutput() KustoDatabaseDataSetOutput {
	return i.ToKustoDatabaseDataSetOutputWithContext(context.Background())
}

func (i *KustoDatabaseDataSet) ToKustoDatabaseDataSetOutputWithContext(ctx context.Context) KustoDatabaseDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoDatabaseDataSetOutput)
}

type KustoDatabaseDataSetOutput struct{ *pulumi.OutputState }

func (KustoDatabaseDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KustoDatabaseDataSet)(nil))
}

func (o KustoDatabaseDataSetOutput) ToKustoDatabaseDataSetOutput() KustoDatabaseDataSetOutput {
	return o
}

func (o KustoDatabaseDataSetOutput) ToKustoDatabaseDataSetOutputWithContext(ctx context.Context) KustoDatabaseDataSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KustoDatabaseDataSetOutput{})
}
