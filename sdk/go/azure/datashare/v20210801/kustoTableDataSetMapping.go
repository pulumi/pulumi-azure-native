


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KustoTableDataSetMapping struct {
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


func NewKustoTableDataSetMapping(ctx *pulumi.Context,
	name string, args *KustoTableDataSetMappingArgs, opts ...pulumi.ResourceOption) (*KustoTableDataSetMapping, error) {
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
	args.Kind = pulumi.String("KustoTable")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:KustoTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:KustoTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:KustoTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:KustoTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:KustoTableDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoTableDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20210801:KustoTableDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKustoTableDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoTableDataSetMappingState, opts ...pulumi.ResourceOption) (*KustoTableDataSetMapping, error) {
	var resource KustoTableDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20210801:KustoTableDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type kustoTableDataSetMappingState struct {
}

type KustoTableDataSetMappingState struct {
}

func (KustoTableDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoTableDataSetMappingState)(nil)).Elem()
}

type kustoTableDataSetMappingArgs struct {
	AccountName            string  `pulumi:"accountName"`
	DataSetId              string  `pulumi:"dataSetId"`
	DataSetMappingName     *string `pulumi:"dataSetMappingName"`
	Kind                   string  `pulumi:"kind"`
	KustoClusterResourceId string  `pulumi:"kustoClusterResourceId"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName  string  `pulumi:"shareSubscriptionName"`
}


type KustoTableDataSetMappingArgs struct {
	AccountName            pulumi.StringInput
	DataSetId              pulumi.StringInput
	DataSetMappingName     pulumi.StringPtrInput
	Kind                   pulumi.StringInput
	KustoClusterResourceId pulumi.StringInput
	ResourceGroupName      pulumi.StringInput
	ShareSubscriptionName  pulumi.StringInput
}

func (KustoTableDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoTableDataSetMappingArgs)(nil)).Elem()
}

type KustoTableDataSetMappingInput interface {
	pulumi.Input

	ToKustoTableDataSetMappingOutput() KustoTableDataSetMappingOutput
	ToKustoTableDataSetMappingOutputWithContext(ctx context.Context) KustoTableDataSetMappingOutput
}

func (*KustoTableDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoTableDataSetMapping)(nil)).Elem()
}

func (i *KustoTableDataSetMapping) ToKustoTableDataSetMappingOutput() KustoTableDataSetMappingOutput {
	return i.ToKustoTableDataSetMappingOutputWithContext(context.Background())
}

func (i *KustoTableDataSetMapping) ToKustoTableDataSetMappingOutputWithContext(ctx context.Context) KustoTableDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoTableDataSetMappingOutput)
}

type KustoTableDataSetMappingOutput struct{ *pulumi.OutputState }

func (KustoTableDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoTableDataSetMapping)(nil)).Elem()
}

func (o KustoTableDataSetMappingOutput) ToKustoTableDataSetMappingOutput() KustoTableDataSetMappingOutput {
	return o
}

func (o KustoTableDataSetMappingOutput) ToKustoTableDataSetMappingOutputWithContext(ctx context.Context) KustoTableDataSetMappingOutput {
	return o
}

func (o KustoTableDataSetMappingOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoTableDataSetMapping) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

func (o KustoTableDataSetMappingOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoTableDataSetMapping) pulumi.StringOutput { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o KustoTableDataSetMappingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoTableDataSetMapping) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o KustoTableDataSetMappingOutput) KustoClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoTableDataSetMapping) pulumi.StringOutput { return v.KustoClusterResourceId }).(pulumi.StringOutput)
}

func (o KustoTableDataSetMappingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoTableDataSetMapping) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o KustoTableDataSetMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoTableDataSetMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KustoTableDataSetMappingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoTableDataSetMapping) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o KustoTableDataSetMappingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KustoTableDataSetMapping) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o KustoTableDataSetMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoTableDataSetMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KustoTableDataSetMappingOutput{})
}
