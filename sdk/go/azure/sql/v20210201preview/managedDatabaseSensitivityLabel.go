


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedDatabaseSensitivityLabel struct {
	pulumi.CustomResourceState

	ColumnName        pulumi.StringOutput    `pulumi:"columnName"`
	InformationType   pulumi.StringPtrOutput `pulumi:"informationType"`
	InformationTypeId pulumi.StringPtrOutput `pulumi:"informationTypeId"`
	IsDisabled        pulumi.BoolOutput      `pulumi:"isDisabled"`
	LabelId           pulumi.StringPtrOutput `pulumi:"labelId"`
	LabelName         pulumi.StringPtrOutput `pulumi:"labelName"`
	ManagedBy         pulumi.StringOutput    `pulumi:"managedBy"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	Rank              pulumi.StringPtrOutput `pulumi:"rank"`
	SchemaName        pulumi.StringOutput    `pulumi:"schemaName"`
	TableName         pulumi.StringOutput    `pulumi:"tableName"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewManagedDatabaseSensitivityLabel(ctx *pulumi.Context,
	name string, args *ManagedDatabaseSensitivityLabelArgs, opts ...pulumi.ResourceOption) (*ManagedDatabaseSensitivityLabel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ColumnName == nil {
		return nil, errors.New("invalid value for required argument 'ColumnName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SchemaName == nil {
		return nil, errors.New("invalid value for required argument 'SchemaName'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20180601preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ManagedDatabaseSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ManagedDatabaseSensitivityLabel"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedDatabaseSensitivityLabel
	err := ctx.RegisterResource("azure-native:sql/v20210201preview:ManagedDatabaseSensitivityLabel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedDatabaseSensitivityLabel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedDatabaseSensitivityLabelState, opts ...pulumi.ResourceOption) (*ManagedDatabaseSensitivityLabel, error) {
	var resource ManagedDatabaseSensitivityLabel
	err := ctx.ReadResource("azure-native:sql/v20210201preview:ManagedDatabaseSensitivityLabel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedDatabaseSensitivityLabelState struct {
}

type ManagedDatabaseSensitivityLabelState struct {
}

func (ManagedDatabaseSensitivityLabelState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedDatabaseSensitivityLabelState)(nil)).Elem()
}

type managedDatabaseSensitivityLabelArgs struct {
	ColumnName             string                `pulumi:"columnName"`
	DatabaseName           string                `pulumi:"databaseName"`
	InformationType        *string               `pulumi:"informationType"`
	InformationTypeId      *string               `pulumi:"informationTypeId"`
	LabelId                *string               `pulumi:"labelId"`
	LabelName              *string               `pulumi:"labelName"`
	ManagedInstanceName    string                `pulumi:"managedInstanceName"`
	Rank                   *SensitivityLabelRank `pulumi:"rank"`
	ResourceGroupName      string                `pulumi:"resourceGroupName"`
	SchemaName             string                `pulumi:"schemaName"`
	SensitivityLabelSource *string               `pulumi:"sensitivityLabelSource"`
	TableName              string                `pulumi:"tableName"`
}


type ManagedDatabaseSensitivityLabelArgs struct {
	ColumnName             pulumi.StringInput
	DatabaseName           pulumi.StringInput
	InformationType        pulumi.StringPtrInput
	InformationTypeId      pulumi.StringPtrInput
	LabelId                pulumi.StringPtrInput
	LabelName              pulumi.StringPtrInput
	ManagedInstanceName    pulumi.StringInput
	Rank                   SensitivityLabelRankPtrInput
	ResourceGroupName      pulumi.StringInput
	SchemaName             pulumi.StringInput
	SensitivityLabelSource pulumi.StringPtrInput
	TableName              pulumi.StringInput
}

func (ManagedDatabaseSensitivityLabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedDatabaseSensitivityLabelArgs)(nil)).Elem()
}

type ManagedDatabaseSensitivityLabelInput interface {
	pulumi.Input

	ToManagedDatabaseSensitivityLabelOutput() ManagedDatabaseSensitivityLabelOutput
	ToManagedDatabaseSensitivityLabelOutputWithContext(ctx context.Context) ManagedDatabaseSensitivityLabelOutput
}

func (*ManagedDatabaseSensitivityLabel) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDatabaseSensitivityLabel)(nil)).Elem()
}

func (i *ManagedDatabaseSensitivityLabel) ToManagedDatabaseSensitivityLabelOutput() ManagedDatabaseSensitivityLabelOutput {
	return i.ToManagedDatabaseSensitivityLabelOutputWithContext(context.Background())
}

func (i *ManagedDatabaseSensitivityLabel) ToManagedDatabaseSensitivityLabelOutputWithContext(ctx context.Context) ManagedDatabaseSensitivityLabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDatabaseSensitivityLabelOutput)
}

type ManagedDatabaseSensitivityLabelOutput struct{ *pulumi.OutputState }

func (ManagedDatabaseSensitivityLabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDatabaseSensitivityLabel)(nil)).Elem()
}

func (o ManagedDatabaseSensitivityLabelOutput) ToManagedDatabaseSensitivityLabelOutput() ManagedDatabaseSensitivityLabelOutput {
	return o
}

func (o ManagedDatabaseSensitivityLabelOutput) ToManagedDatabaseSensitivityLabelOutputWithContext(ctx context.Context) ManagedDatabaseSensitivityLabelOutput {
	return o
}

func (o ManagedDatabaseSensitivityLabelOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringOutput { return v.ColumnName }).(pulumi.StringOutput)
}

func (o ManagedDatabaseSensitivityLabelOutput) InformationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringPtrOutput { return v.InformationType }).(pulumi.StringPtrOutput)
}

func (o ManagedDatabaseSensitivityLabelOutput) InformationTypeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringPtrOutput { return v.InformationTypeId }).(pulumi.StringPtrOutput)
}

func (o ManagedDatabaseSensitivityLabelOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.BoolOutput { return v.IsDisabled }).(pulumi.BoolOutput)
}

func (o ManagedDatabaseSensitivityLabelOutput) LabelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringPtrOutput { return v.LabelId }).(pulumi.StringPtrOutput)
}

func (o ManagedDatabaseSensitivityLabelOutput) LabelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringPtrOutput { return v.LabelName }).(pulumi.StringPtrOutput)
}

func (o ManagedDatabaseSensitivityLabelOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

func (o ManagedDatabaseSensitivityLabelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedDatabaseSensitivityLabelOutput) Rank() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringPtrOutput { return v.Rank }).(pulumi.StringPtrOutput)
}

func (o ManagedDatabaseSensitivityLabelOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringOutput { return v.SchemaName }).(pulumi.StringOutput)
}

func (o ManagedDatabaseSensitivityLabelOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

func (o ManagedDatabaseSensitivityLabelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedDatabaseSensitivityLabel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedDatabaseSensitivityLabelOutput{})
}
