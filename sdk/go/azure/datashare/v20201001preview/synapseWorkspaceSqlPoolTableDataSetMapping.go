


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SynapseWorkspaceSqlPoolTableDataSetMapping struct {
	pulumi.CustomResourceState

	DataSetId                              pulumi.StringOutput      `pulumi:"dataSetId"`
	DataSetMappingStatus                   pulumi.StringOutput      `pulumi:"dataSetMappingStatus"`
	Kind                                   pulumi.StringOutput      `pulumi:"kind"`
	Name                                   pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState                      pulumi.StringOutput      `pulumi:"provisioningState"`
	SynapseWorkspaceSqlPoolTableResourceId pulumi.StringOutput      `pulumi:"synapseWorkspaceSqlPoolTableResourceId"`
	SystemData                             SystemDataResponseOutput `pulumi:"systemData"`
	Type                                   pulumi.StringOutput      `pulumi:"type"`
}


func NewSynapseWorkspaceSqlPoolTableDataSetMapping(ctx *pulumi.Context,
	name string, args *SynapseWorkspaceSqlPoolTableDataSetMappingArgs, opts ...pulumi.ResourceOption) (*SynapseWorkspaceSqlPoolTableDataSetMapping, error) {
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
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'ShareSubscriptionName'")
	}
	if args.SynapseWorkspaceSqlPoolTableResourceId == nil {
		return nil, errors.New("invalid value for required argument 'SynapseWorkspaceSqlPoolTableResourceId'")
	}
	args.Kind = pulumi.String("SynapseWorkspaceSqlPoolTable")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:SynapseWorkspaceSqlPoolTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:SynapseWorkspaceSqlPoolTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:SynapseWorkspaceSqlPoolTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:SynapseWorkspaceSqlPoolTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:SynapseWorkspaceSqlPoolTableDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource SynapseWorkspaceSqlPoolTableDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:SynapseWorkspaceSqlPoolTableDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSynapseWorkspaceSqlPoolTableDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SynapseWorkspaceSqlPoolTableDataSetMappingState, opts ...pulumi.ResourceOption) (*SynapseWorkspaceSqlPoolTableDataSetMapping, error) {
	var resource SynapseWorkspaceSqlPoolTableDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:SynapseWorkspaceSqlPoolTableDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type synapseWorkspaceSqlPoolTableDataSetMappingState struct {
}

type SynapseWorkspaceSqlPoolTableDataSetMappingState struct {
}

func (SynapseWorkspaceSqlPoolTableDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*synapseWorkspaceSqlPoolTableDataSetMappingState)(nil)).Elem()
}

type synapseWorkspaceSqlPoolTableDataSetMappingArgs struct {
	AccountName                            string  `pulumi:"accountName"`
	DataSetId                              string  `pulumi:"dataSetId"`
	DataSetMappingName                     *string `pulumi:"dataSetMappingName"`
	Kind                                   string  `pulumi:"kind"`
	ResourceGroupName                      string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName                  string  `pulumi:"shareSubscriptionName"`
	SynapseWorkspaceSqlPoolTableResourceId string  `pulumi:"synapseWorkspaceSqlPoolTableResourceId"`
}


type SynapseWorkspaceSqlPoolTableDataSetMappingArgs struct {
	AccountName                            pulumi.StringInput
	DataSetId                              pulumi.StringInput
	DataSetMappingName                     pulumi.StringPtrInput
	Kind                                   pulumi.StringInput
	ResourceGroupName                      pulumi.StringInput
	ShareSubscriptionName                  pulumi.StringInput
	SynapseWorkspaceSqlPoolTableResourceId pulumi.StringInput
}

func (SynapseWorkspaceSqlPoolTableDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*synapseWorkspaceSqlPoolTableDataSetMappingArgs)(nil)).Elem()
}

type SynapseWorkspaceSqlPoolTableDataSetMappingInput interface {
	pulumi.Input

	ToSynapseWorkspaceSqlPoolTableDataSetMappingOutput() SynapseWorkspaceSqlPoolTableDataSetMappingOutput
	ToSynapseWorkspaceSqlPoolTableDataSetMappingOutputWithContext(ctx context.Context) SynapseWorkspaceSqlPoolTableDataSetMappingOutput
}

func (*SynapseWorkspaceSqlPoolTableDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**SynapseWorkspaceSqlPoolTableDataSetMapping)(nil)).Elem()
}

func (i *SynapseWorkspaceSqlPoolTableDataSetMapping) ToSynapseWorkspaceSqlPoolTableDataSetMappingOutput() SynapseWorkspaceSqlPoolTableDataSetMappingOutput {
	return i.ToSynapseWorkspaceSqlPoolTableDataSetMappingOutputWithContext(context.Background())
}

func (i *SynapseWorkspaceSqlPoolTableDataSetMapping) ToSynapseWorkspaceSqlPoolTableDataSetMappingOutputWithContext(ctx context.Context) SynapseWorkspaceSqlPoolTableDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynapseWorkspaceSqlPoolTableDataSetMappingOutput)
}

type SynapseWorkspaceSqlPoolTableDataSetMappingOutput struct{ *pulumi.OutputState }

func (SynapseWorkspaceSqlPoolTableDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SynapseWorkspaceSqlPoolTableDataSetMapping)(nil)).Elem()
}

func (o SynapseWorkspaceSqlPoolTableDataSetMappingOutput) ToSynapseWorkspaceSqlPoolTableDataSetMappingOutput() SynapseWorkspaceSqlPoolTableDataSetMappingOutput {
	return o
}

func (o SynapseWorkspaceSqlPoolTableDataSetMappingOutput) ToSynapseWorkspaceSqlPoolTableDataSetMappingOutputWithContext(ctx context.Context) SynapseWorkspaceSqlPoolTableDataSetMappingOutput {
	return o
}

func (o SynapseWorkspaceSqlPoolTableDataSetMappingOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *SynapseWorkspaceSqlPoolTableDataSetMapping) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

func (o SynapseWorkspaceSqlPoolTableDataSetMappingOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *SynapseWorkspaceSqlPoolTableDataSetMapping) pulumi.StringOutput { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o SynapseWorkspaceSqlPoolTableDataSetMappingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *SynapseWorkspaceSqlPoolTableDataSetMapping) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o SynapseWorkspaceSqlPoolTableDataSetMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SynapseWorkspaceSqlPoolTableDataSetMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SynapseWorkspaceSqlPoolTableDataSetMappingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SynapseWorkspaceSqlPoolTableDataSetMapping) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SynapseWorkspaceSqlPoolTableDataSetMappingOutput) SynapseWorkspaceSqlPoolTableResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SynapseWorkspaceSqlPoolTableDataSetMapping) pulumi.StringOutput {
		return v.SynapseWorkspaceSqlPoolTableResourceId
	}).(pulumi.StringOutput)
}

func (o SynapseWorkspaceSqlPoolTableDataSetMappingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SynapseWorkspaceSqlPoolTableDataSetMapping) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SynapseWorkspaceSqlPoolTableDataSetMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SynapseWorkspaceSqlPoolTableDataSetMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SynapseWorkspaceSqlPoolTableDataSetMappingOutput{})
}
