


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KustoDatabaseDataSetMapping struct {
	pulumi.CustomResourceState

	DataSetId              pulumi.StringOutput      `pulumi:"dataSetId"`
	DataSetMappingStatus   pulumi.StringOutput      `pulumi:"dataSetMappingStatus"`
	Kind                   pulumi.StringOutput      `pulumi:"kind"`
	KustoClusterResourceId pulumi.StringOutput      `pulumi:"kustoClusterResourceId"`
	Location               pulumi.StringOutput      `pulumi:"location"`
	Name                   pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState      pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData             SystemDataResponseOutput `pulumi:"systemData"`
	Type                   pulumi.StringOutput      `pulumi:"type"`
}


func NewKustoDatabaseDataSetMapping(ctx *pulumi.Context,
	name string, args *KustoDatabaseDataSetMappingArgs, opts ...pulumi.ResourceOption) (*KustoDatabaseDataSetMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DataSetId == nil {
		return nil, errors.New("invalid value for required argument 'DataSetId'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.KustoClusterResourceId == nil {
		return nil, errors.New("invalid value for required argument 'KustoClusterResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'ShareSubscriptionName'")
	}
	args.Kind = pulumi.String("KustoDatabase")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datashare/v20200901:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20181101preview:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20191101:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20201001preview:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-nextgen:datashare/v20210801:KustoDatabaseDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoDatabaseDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20200901:KustoDatabaseDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKustoDatabaseDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoDatabaseDataSetMappingState, opts ...pulumi.ResourceOption) (*KustoDatabaseDataSetMapping, error) {
	var resource KustoDatabaseDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20200901:KustoDatabaseDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type kustoDatabaseDataSetMappingState struct {
}

type KustoDatabaseDataSetMappingState struct {
}

func (KustoDatabaseDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoDatabaseDataSetMappingState)(nil)).Elem()
}

type kustoDatabaseDataSetMappingArgs struct {
	AccountName            string  `pulumi:"accountName"`
	DataSetId              string  `pulumi:"dataSetId"`
	DataSetMappingName     *string `pulumi:"dataSetMappingName"`
	Kind                   string  `pulumi:"kind"`
	KustoClusterResourceId string  `pulumi:"kustoClusterResourceId"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName  string  `pulumi:"shareSubscriptionName"`
}


type KustoDatabaseDataSetMappingArgs struct {
	AccountName            pulumi.StringInput
	DataSetId              pulumi.StringInput
	DataSetMappingName     pulumi.StringPtrInput
	Kind                   pulumi.StringInput
	KustoClusterResourceId pulumi.StringInput
	ResourceGroupName      pulumi.StringInput
	ShareSubscriptionName  pulumi.StringInput
}

func (KustoDatabaseDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoDatabaseDataSetMappingArgs)(nil)).Elem()
}

type KustoDatabaseDataSetMappingInput interface {
	pulumi.Input

	ToKustoDatabaseDataSetMappingOutput() KustoDatabaseDataSetMappingOutput
	ToKustoDatabaseDataSetMappingOutputWithContext(ctx context.Context) KustoDatabaseDataSetMappingOutput
}

func (*KustoDatabaseDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((*KustoDatabaseDataSetMapping)(nil))
}

func (i *KustoDatabaseDataSetMapping) ToKustoDatabaseDataSetMappingOutput() KustoDatabaseDataSetMappingOutput {
	return i.ToKustoDatabaseDataSetMappingOutputWithContext(context.Background())
}

func (i *KustoDatabaseDataSetMapping) ToKustoDatabaseDataSetMappingOutputWithContext(ctx context.Context) KustoDatabaseDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoDatabaseDataSetMappingOutput)
}

type KustoDatabaseDataSetMappingOutput struct{ *pulumi.OutputState }

func (KustoDatabaseDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KustoDatabaseDataSetMapping)(nil))
}

func (o KustoDatabaseDataSetMappingOutput) ToKustoDatabaseDataSetMappingOutput() KustoDatabaseDataSetMappingOutput {
	return o
}

func (o KustoDatabaseDataSetMappingOutput) ToKustoDatabaseDataSetMappingOutputWithContext(ctx context.Context) KustoDatabaseDataSetMappingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KustoDatabaseDataSetMappingOutput{})
}
