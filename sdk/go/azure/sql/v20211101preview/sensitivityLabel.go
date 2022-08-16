


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SensitivityLabel struct {
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


func NewSensitivityLabel(ctx *pulumi.Context,
	name string, args *SensitivityLabelArgs, opts ...pulumi.ResourceOption) (*SensitivityLabel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ColumnName == nil {
		return nil, errors.New("invalid value for required argument 'ColumnName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SchemaName == nil {
		return nil, errors.New("invalid value for required argument 'SchemaName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:SensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:SensitivityLabel"),
		},
	})
	opts = append(opts, aliases)
	var resource SensitivityLabel
	err := ctx.RegisterResource("azure-native:sql/v20211101preview:SensitivityLabel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSensitivityLabel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SensitivityLabelState, opts ...pulumi.ResourceOption) (*SensitivityLabel, error) {
	var resource SensitivityLabel
	err := ctx.ReadResource("azure-native:sql/v20211101preview:SensitivityLabel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sensitivityLabelState struct {
}

type SensitivityLabelState struct {
}

func (SensitivityLabelState) ElementType() reflect.Type {
	return reflect.TypeOf((*sensitivityLabelState)(nil)).Elem()
}

type sensitivityLabelArgs struct {
	ColumnName             string                `pulumi:"columnName"`
	DatabaseName           string                `pulumi:"databaseName"`
	InformationType        *string               `pulumi:"informationType"`
	InformationTypeId      *string               `pulumi:"informationTypeId"`
	LabelId                *string               `pulumi:"labelId"`
	LabelName              *string               `pulumi:"labelName"`
	Rank                   *SensitivityLabelRank `pulumi:"rank"`
	ResourceGroupName      string                `pulumi:"resourceGroupName"`
	SchemaName             string                `pulumi:"schemaName"`
	SensitivityLabelSource *string               `pulumi:"sensitivityLabelSource"`
	ServerName             string                `pulumi:"serverName"`
	TableName              string                `pulumi:"tableName"`
}


type SensitivityLabelArgs struct {
	ColumnName             pulumi.StringInput
	DatabaseName           pulumi.StringInput
	InformationType        pulumi.StringPtrInput
	InformationTypeId      pulumi.StringPtrInput
	LabelId                pulumi.StringPtrInput
	LabelName              pulumi.StringPtrInput
	Rank                   SensitivityLabelRankPtrInput
	ResourceGroupName      pulumi.StringInput
	SchemaName             pulumi.StringInput
	SensitivityLabelSource pulumi.StringPtrInput
	ServerName             pulumi.StringInput
	TableName              pulumi.StringInput
}

func (SensitivityLabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sensitivityLabelArgs)(nil)).Elem()
}

type SensitivityLabelInput interface {
	pulumi.Input

	ToSensitivityLabelOutput() SensitivityLabelOutput
	ToSensitivityLabelOutputWithContext(ctx context.Context) SensitivityLabelOutput
}

func (*SensitivityLabel) ElementType() reflect.Type {
	return reflect.TypeOf((**SensitivityLabel)(nil)).Elem()
}

func (i *SensitivityLabel) ToSensitivityLabelOutput() SensitivityLabelOutput {
	return i.ToSensitivityLabelOutputWithContext(context.Background())
}

func (i *SensitivityLabel) ToSensitivityLabelOutputWithContext(ctx context.Context) SensitivityLabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SensitivityLabelOutput)
}

type SensitivityLabelOutput struct{ *pulumi.OutputState }

func (SensitivityLabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SensitivityLabel)(nil)).Elem()
}

func (o SensitivityLabelOutput) ToSensitivityLabelOutput() SensitivityLabelOutput {
	return o
}

func (o SensitivityLabelOutput) ToSensitivityLabelOutputWithContext(ctx context.Context) SensitivityLabelOutput {
	return o
}

func (o SensitivityLabelOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringOutput { return v.ColumnName }).(pulumi.StringOutput)
}

func (o SensitivityLabelOutput) InformationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringPtrOutput { return v.InformationType }).(pulumi.StringPtrOutput)
}

func (o SensitivityLabelOutput) InformationTypeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringPtrOutput { return v.InformationTypeId }).(pulumi.StringPtrOutput)
}

func (o SensitivityLabelOutput) IsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.BoolOutput { return v.IsDisabled }).(pulumi.BoolOutput)
}

func (o SensitivityLabelOutput) LabelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringPtrOutput { return v.LabelId }).(pulumi.StringPtrOutput)
}

func (o SensitivityLabelOutput) LabelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringPtrOutput { return v.LabelName }).(pulumi.StringPtrOutput)
}

func (o SensitivityLabelOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

func (o SensitivityLabelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SensitivityLabelOutput) Rank() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringPtrOutput { return v.Rank }).(pulumi.StringPtrOutput)
}

func (o SensitivityLabelOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringOutput { return v.SchemaName }).(pulumi.StringOutput)
}

func (o SensitivityLabelOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

func (o SensitivityLabelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SensitivityLabel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SensitivityLabelOutput{})
}
