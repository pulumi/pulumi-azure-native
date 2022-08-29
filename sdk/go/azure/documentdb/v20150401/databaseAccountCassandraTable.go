


package v20150401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DatabaseAccountCassandraTable struct {
	pulumi.CustomResourceState

	DefaultTtl pulumi.IntPtrOutput              `pulumi:"defaultTtl"`
	Location   pulumi.StringPtrOutput           `pulumi:"location"`
	Name       pulumi.StringOutput              `pulumi:"name"`
	Schema     CassandraSchemaResponsePtrOutput `pulumi:"schema"`
	Tags       pulumi.StringMapOutput           `pulumi:"tags"`
	Type       pulumi.StringOutput              `pulumi:"type"`
}


func NewDatabaseAccountCassandraTable(ctx *pulumi.Context,
	name string, args *DatabaseAccountCassandraTableArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountCassandraTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.KeyspaceName == nil {
		return nil, errors.New("invalid value for required argument 'KeyspaceName'")
	}
	if args.Options == nil {
		return nil, errors.New("invalid value for required argument 'Options'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:DatabaseAccountCassandraTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:DatabaseAccountCassandraTable"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountCassandraTable
	err := ctx.RegisterResource("azure-native:documentdb/v20150401:DatabaseAccountCassandraTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseAccountCassandraTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountCassandraTableState, opts ...pulumi.ResourceOption) (*DatabaseAccountCassandraTable, error) {
	var resource DatabaseAccountCassandraTable
	err := ctx.ReadResource("azure-native:documentdb/v20150401:DatabaseAccountCassandraTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseAccountCassandraTableState struct {
}

type DatabaseAccountCassandraTableState struct {
}

func (DatabaseAccountCassandraTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountCassandraTableState)(nil)).Elem()
}

type databaseAccountCassandraTableArgs struct {
	AccountName       string                 `pulumi:"accountName"`
	KeyspaceName      string                 `pulumi:"keyspaceName"`
	Options           map[string]string      `pulumi:"options"`
	Resource          CassandraTableResource `pulumi:"resource"`
	ResourceGroupName string                 `pulumi:"resourceGroupName"`
	TableName         *string                `pulumi:"tableName"`
}


type DatabaseAccountCassandraTableArgs struct {
	AccountName       pulumi.StringInput
	KeyspaceName      pulumi.StringInput
	Options           pulumi.StringMapInput
	Resource          CassandraTableResourceInput
	ResourceGroupName pulumi.StringInput
	TableName         pulumi.StringPtrInput
}

func (DatabaseAccountCassandraTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountCassandraTableArgs)(nil)).Elem()
}

type DatabaseAccountCassandraTableInput interface {
	pulumi.Input

	ToDatabaseAccountCassandraTableOutput() DatabaseAccountCassandraTableOutput
	ToDatabaseAccountCassandraTableOutputWithContext(ctx context.Context) DatabaseAccountCassandraTableOutput
}

func (*DatabaseAccountCassandraTable) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountCassandraTable)(nil)).Elem()
}

func (i *DatabaseAccountCassandraTable) ToDatabaseAccountCassandraTableOutput() DatabaseAccountCassandraTableOutput {
	return i.ToDatabaseAccountCassandraTableOutputWithContext(context.Background())
}

func (i *DatabaseAccountCassandraTable) ToDatabaseAccountCassandraTableOutputWithContext(ctx context.Context) DatabaseAccountCassandraTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountCassandraTableOutput)
}

type DatabaseAccountCassandraTableOutput struct{ *pulumi.OutputState }

func (DatabaseAccountCassandraTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountCassandraTable)(nil)).Elem()
}

func (o DatabaseAccountCassandraTableOutput) ToDatabaseAccountCassandraTableOutput() DatabaseAccountCassandraTableOutput {
	return o
}

func (o DatabaseAccountCassandraTableOutput) ToDatabaseAccountCassandraTableOutputWithContext(ctx context.Context) DatabaseAccountCassandraTableOutput {
	return o
}

func (o DatabaseAccountCassandraTableOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraTable) pulumi.IntPtrOutput { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o DatabaseAccountCassandraTableOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraTable) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountCassandraTableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraTable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseAccountCassandraTableOutput) Schema() CassandraSchemaResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraTable) CassandraSchemaResponsePtrOutput { return v.Schema }).(CassandraSchemaResponsePtrOutput)
}

func (o DatabaseAccountCassandraTableOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraTable) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DatabaseAccountCassandraTableOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountCassandraTable) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountCassandraTableOutput{})
}
