


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KustoClusterDataSetMapping struct {
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


func NewKustoClusterDataSetMapping(ctx *pulumi.Context,
	name string, args *KustoClusterDataSetMappingArgs, opts ...pulumi.ResourceOption) (*KustoClusterDataSetMapping, error) {
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
	args.Kind = pulumi.String("KustoCluster")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:KustoClusterDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:KustoClusterDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:KustoClusterDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:KustoClusterDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:KustoClusterDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoClusterDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20200901:KustoClusterDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKustoClusterDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoClusterDataSetMappingState, opts ...pulumi.ResourceOption) (*KustoClusterDataSetMapping, error) {
	var resource KustoClusterDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20200901:KustoClusterDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type kustoClusterDataSetMappingState struct {
}

type KustoClusterDataSetMappingState struct {
}

func (KustoClusterDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoClusterDataSetMappingState)(nil)).Elem()
}

type kustoClusterDataSetMappingArgs struct {
	AccountName            string  `pulumi:"accountName"`
	DataSetId              string  `pulumi:"dataSetId"`
	DataSetMappingName     *string `pulumi:"dataSetMappingName"`
	Kind                   string  `pulumi:"kind"`
	KustoClusterResourceId string  `pulumi:"kustoClusterResourceId"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName  string  `pulumi:"shareSubscriptionName"`
}


type KustoClusterDataSetMappingArgs struct {
	AccountName            pulumi.StringInput
	DataSetId              pulumi.StringInput
	DataSetMappingName     pulumi.StringPtrInput
	Kind                   pulumi.StringInput
	KustoClusterResourceId pulumi.StringInput
	ResourceGroupName      pulumi.StringInput
	ShareSubscriptionName  pulumi.StringInput
}

func (KustoClusterDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoClusterDataSetMappingArgs)(nil)).Elem()
}

type KustoClusterDataSetMappingInput interface {
	pulumi.Input

	ToKustoClusterDataSetMappingOutput() KustoClusterDataSetMappingOutput
	ToKustoClusterDataSetMappingOutputWithContext(ctx context.Context) KustoClusterDataSetMappingOutput
}

func (*KustoClusterDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoClusterDataSetMapping)(nil)).Elem()
}

func (i *KustoClusterDataSetMapping) ToKustoClusterDataSetMappingOutput() KustoClusterDataSetMappingOutput {
	return i.ToKustoClusterDataSetMappingOutputWithContext(context.Background())
}

func (i *KustoClusterDataSetMapping) ToKustoClusterDataSetMappingOutputWithContext(ctx context.Context) KustoClusterDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoClusterDataSetMappingOutput)
}

type KustoClusterDataSetMappingOutput struct{ *pulumi.OutputState }

func (KustoClusterDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoClusterDataSetMapping)(nil)).Elem()
}

func (o KustoClusterDataSetMappingOutput) ToKustoClusterDataSetMappingOutput() KustoClusterDataSetMappingOutput {
	return o
}

func (o KustoClusterDataSetMappingOutput) ToKustoClusterDataSetMappingOutputWithContext(ctx context.Context) KustoClusterDataSetMappingOutput {
	return o
}

func (o KustoClusterDataSetMappingOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

func (o KustoClusterDataSetMappingOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o KustoClusterDataSetMappingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o KustoClusterDataSetMappingOutput) KustoClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.KustoClusterResourceId }).(pulumi.StringOutput)
}

func (o KustoClusterDataSetMappingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o KustoClusterDataSetMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KustoClusterDataSetMappingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o KustoClusterDataSetMappingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o KustoClusterDataSetMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSetMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KustoClusterDataSetMappingOutput{})
}
