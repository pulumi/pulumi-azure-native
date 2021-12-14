


package synapse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlPoolSensitivityLabel struct {
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


func NewSqlPoolSensitivityLabel(ctx *pulumi.Context,
	name string, args *SqlPoolSensitivityLabelArgs, opts ...pulumi.ResourceOption) (*SqlPoolSensitivityLabel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ColumnName == nil {
		return nil, errors.New("invalid value for required argument 'ColumnName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SchemaName == nil {
		return nil, errors.New("invalid value for required argument 'SchemaName'")
	}
	if args.SqlPoolName == nil {
		return nil, errors.New("invalid value for required argument 'SqlPoolName'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:SqlPoolSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:SqlPoolSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:SqlPoolSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:SqlPoolSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:SqlPoolSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:SqlPoolSensitivityLabel"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:SqlPoolSensitivityLabel"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlPoolSensitivityLabel
	err := ctx.RegisterResource("azure-native:synapse:SqlPoolSensitivityLabel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlPoolSensitivityLabel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlPoolSensitivityLabelState, opts ...pulumi.ResourceOption) (*SqlPoolSensitivityLabel, error) {
	var resource SqlPoolSensitivityLabel
	err := ctx.ReadResource("azure-native:synapse:SqlPoolSensitivityLabel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlPoolSensitivityLabelState struct {
}

type SqlPoolSensitivityLabelState struct {
}

func (SqlPoolSensitivityLabelState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolSensitivityLabelState)(nil)).Elem()
}

type sqlPoolSensitivityLabelArgs struct {
	ColumnName             string                `pulumi:"columnName"`
	InformationType        *string               `pulumi:"informationType"`
	InformationTypeId      *string               `pulumi:"informationTypeId"`
	LabelId                *string               `pulumi:"labelId"`
	LabelName              *string               `pulumi:"labelName"`
	Rank                   *SensitivityLabelRank `pulumi:"rank"`
	ResourceGroupName      string                `pulumi:"resourceGroupName"`
	SchemaName             string                `pulumi:"schemaName"`
	SensitivityLabelSource *string               `pulumi:"sensitivityLabelSource"`
	SqlPoolName            string                `pulumi:"sqlPoolName"`
	TableName              string                `pulumi:"tableName"`
	WorkspaceName          string                `pulumi:"workspaceName"`
}


type SqlPoolSensitivityLabelArgs struct {
	ColumnName             pulumi.StringInput
	InformationType        pulumi.StringPtrInput
	InformationTypeId      pulumi.StringPtrInput
	LabelId                pulumi.StringPtrInput
	LabelName              pulumi.StringPtrInput
	Rank                   SensitivityLabelRankPtrInput
	ResourceGroupName      pulumi.StringInput
	SchemaName             pulumi.StringInput
	SensitivityLabelSource pulumi.StringPtrInput
	SqlPoolName            pulumi.StringInput
	TableName              pulumi.StringInput
	WorkspaceName          pulumi.StringInput
}

func (SqlPoolSensitivityLabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolSensitivityLabelArgs)(nil)).Elem()
}

type SqlPoolSensitivityLabelInput interface {
	pulumi.Input

	ToSqlPoolSensitivityLabelOutput() SqlPoolSensitivityLabelOutput
	ToSqlPoolSensitivityLabelOutputWithContext(ctx context.Context) SqlPoolSensitivityLabelOutput
}

func (*SqlPoolSensitivityLabel) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlPoolSensitivityLabel)(nil)).Elem()
}

func (i *SqlPoolSensitivityLabel) ToSqlPoolSensitivityLabelOutput() SqlPoolSensitivityLabelOutput {
	return i.ToSqlPoolSensitivityLabelOutputWithContext(context.Background())
}

func (i *SqlPoolSensitivityLabel) ToSqlPoolSensitivityLabelOutputWithContext(ctx context.Context) SqlPoolSensitivityLabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolSensitivityLabelOutput)
}

type SqlPoolSensitivityLabelOutput struct{ *pulumi.OutputState }

func (SqlPoolSensitivityLabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlPoolSensitivityLabel)(nil)).Elem()
}

func (o SqlPoolSensitivityLabelOutput) ToSqlPoolSensitivityLabelOutput() SqlPoolSensitivityLabelOutput {
	return o
}

func (o SqlPoolSensitivityLabelOutput) ToSqlPoolSensitivityLabelOutputWithContext(ctx context.Context) SqlPoolSensitivityLabelOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SqlPoolSensitivityLabelOutput{})
}
