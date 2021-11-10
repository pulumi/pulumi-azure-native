


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KustoTableDataSet struct {
	pulumi.CustomResourceState

	DataSetId                   pulumi.StringOutput                       `pulumi:"dataSetId"`
	Kind                        pulumi.StringOutput                       `pulumi:"kind"`
	KustoDatabaseResourceId     pulumi.StringOutput                       `pulumi:"kustoDatabaseResourceId"`
	Location                    pulumi.StringOutput                       `pulumi:"location"`
	Name                        pulumi.StringOutput                       `pulumi:"name"`
	ProvisioningState           pulumi.StringOutput                       `pulumi:"provisioningState"`
	SystemData                  SystemDataResponseOutput                  `pulumi:"systemData"`
	TableLevelSharingProperties TableLevelSharingPropertiesResponseOutput `pulumi:"tableLevelSharingProperties"`
	Type                        pulumi.StringOutput                       `pulumi:"type"`
}


func NewKustoTableDataSet(ctx *pulumi.Context,
	name string, args *KustoTableDataSetArgs, opts ...pulumi.ResourceOption) (*KustoTableDataSet, error) {
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
	if args.TableLevelSharingProperties == nil {
		return nil, errors.New("invalid value for required argument 'TableLevelSharingProperties'")
	}
	args.Kind = pulumi.String("KustoTable")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:KustoTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:KustoTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:KustoTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:KustoTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:KustoTableDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoTableDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20210801:KustoTableDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKustoTableDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoTableDataSetState, opts ...pulumi.ResourceOption) (*KustoTableDataSet, error) {
	var resource KustoTableDataSet
	err := ctx.ReadResource("azure-native:datashare/v20210801:KustoTableDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type kustoTableDataSetState struct {
}

type KustoTableDataSetState struct {
}

func (KustoTableDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoTableDataSetState)(nil)).Elem()
}

type kustoTableDataSetArgs struct {
	AccountName                 string                      `pulumi:"accountName"`
	DataSetName                 *string                     `pulumi:"dataSetName"`
	Kind                        string                      `pulumi:"kind"`
	KustoDatabaseResourceId     string                      `pulumi:"kustoDatabaseResourceId"`
	ResourceGroupName           string                      `pulumi:"resourceGroupName"`
	ShareName                   string                      `pulumi:"shareName"`
	TableLevelSharingProperties TableLevelSharingProperties `pulumi:"tableLevelSharingProperties"`
}


type KustoTableDataSetArgs struct {
	AccountName                 pulumi.StringInput
	DataSetName                 pulumi.StringPtrInput
	Kind                        pulumi.StringInput
	KustoDatabaseResourceId     pulumi.StringInput
	ResourceGroupName           pulumi.StringInput
	ShareName                   pulumi.StringInput
	TableLevelSharingProperties TableLevelSharingPropertiesInput
}

func (KustoTableDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoTableDataSetArgs)(nil)).Elem()
}

type KustoTableDataSetInput interface {
	pulumi.Input

	ToKustoTableDataSetOutput() KustoTableDataSetOutput
	ToKustoTableDataSetOutputWithContext(ctx context.Context) KustoTableDataSetOutput
}

func (*KustoTableDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((*KustoTableDataSet)(nil))
}

func (i *KustoTableDataSet) ToKustoTableDataSetOutput() KustoTableDataSetOutput {
	return i.ToKustoTableDataSetOutputWithContext(context.Background())
}

func (i *KustoTableDataSet) ToKustoTableDataSetOutputWithContext(ctx context.Context) KustoTableDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoTableDataSetOutput)
}

type KustoTableDataSetOutput struct{ *pulumi.OutputState }

func (KustoTableDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KustoTableDataSet)(nil))
}

func (o KustoTableDataSetOutput) ToKustoTableDataSetOutput() KustoTableDataSetOutput {
	return o
}

func (o KustoTableDataSetOutput) ToKustoTableDataSetOutputWithContext(ctx context.Context) KustoTableDataSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KustoTableDataSetOutput{})
}
