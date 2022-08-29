


package datashare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlDBTableDataSet struct {
	pulumi.CustomResourceState

	DataSetId           pulumi.StringOutput      `pulumi:"dataSetId"`
	DatabaseName        pulumi.StringOutput      `pulumi:"databaseName"`
	Kind                pulumi.StringOutput      `pulumi:"kind"`
	Name                pulumi.StringOutput      `pulumi:"name"`
	SchemaName          pulumi.StringOutput      `pulumi:"schemaName"`
	SqlServerResourceId pulumi.StringOutput      `pulumi:"sqlServerResourceId"`
	SystemData          SystemDataResponseOutput `pulumi:"systemData"`
	TableName           pulumi.StringOutput      `pulumi:"tableName"`
	Type                pulumi.StringOutput      `pulumi:"type"`
}


func NewSqlDBTableDataSet(ctx *pulumi.Context,
	name string, args *SqlDBTableDataSetArgs, opts ...pulumi.ResourceOption) (*SqlDBTableDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SchemaName == nil {
		return nil, errors.New("invalid value for required argument 'SchemaName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	if args.SqlServerResourceId == nil {
		return nil, errors.New("invalid value for required argument 'SqlServerResourceId'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	args.Kind = pulumi.String("SqlDBTable")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:SqlDBTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:SqlDBTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:SqlDBTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:SqlDBTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:SqlDBTableDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlDBTableDataSet
	err := ctx.RegisterResource("azure-native:datashare:SqlDBTableDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlDBTableDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlDBTableDataSetState, opts ...pulumi.ResourceOption) (*SqlDBTableDataSet, error) {
	var resource SqlDBTableDataSet
	err := ctx.ReadResource("azure-native:datashare:SqlDBTableDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlDBTableDataSetState struct {
}

type SqlDBTableDataSetState struct {
}

func (SqlDBTableDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDBTableDataSetState)(nil)).Elem()
}

type sqlDBTableDataSetArgs struct {
	AccountName         string  `pulumi:"accountName"`
	DataSetName         *string `pulumi:"dataSetName"`
	DatabaseName        string  `pulumi:"databaseName"`
	Kind                string  `pulumi:"kind"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
	SchemaName          string  `pulumi:"schemaName"`
	ShareName           string  `pulumi:"shareName"`
	SqlServerResourceId string  `pulumi:"sqlServerResourceId"`
	TableName           string  `pulumi:"tableName"`
}


type SqlDBTableDataSetArgs struct {
	AccountName         pulumi.StringInput
	DataSetName         pulumi.StringPtrInput
	DatabaseName        pulumi.StringInput
	Kind                pulumi.StringInput
	ResourceGroupName   pulumi.StringInput
	SchemaName          pulumi.StringInput
	ShareName           pulumi.StringInput
	SqlServerResourceId pulumi.StringInput
	TableName           pulumi.StringInput
}

func (SqlDBTableDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDBTableDataSetArgs)(nil)).Elem()
}

type SqlDBTableDataSetInput interface {
	pulumi.Input

	ToSqlDBTableDataSetOutput() SqlDBTableDataSetOutput
	ToSqlDBTableDataSetOutputWithContext(ctx context.Context) SqlDBTableDataSetOutput
}

func (*SqlDBTableDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDBTableDataSet)(nil)).Elem()
}

func (i *SqlDBTableDataSet) ToSqlDBTableDataSetOutput() SqlDBTableDataSetOutput {
	return i.ToSqlDBTableDataSetOutputWithContext(context.Background())
}

func (i *SqlDBTableDataSet) ToSqlDBTableDataSetOutputWithContext(ctx context.Context) SqlDBTableDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDBTableDataSetOutput)
}

type SqlDBTableDataSetOutput struct{ *pulumi.OutputState }

func (SqlDBTableDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDBTableDataSet)(nil)).Elem()
}

func (o SqlDBTableDataSetOutput) ToSqlDBTableDataSetOutput() SqlDBTableDataSetOutput {
	return o
}

func (o SqlDBTableDataSetOutput) ToSqlDBTableDataSetOutputWithContext(ctx context.Context) SqlDBTableDataSetOutput {
	return o
}

func (o SqlDBTableDataSetOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSet) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSet) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSet) pulumi.StringOutput { return v.SchemaName }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetOutput) SqlServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSet) pulumi.StringOutput { return v.SqlServerResourceId }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlDBTableDataSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SqlDBTableDataSetOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSet) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlDBTableDataSetOutput{})
}
