


package v20150408

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseAccountTable struct {
	pulumi.CustomResourceState

	Location pulumi.StringPtrOutput `pulumi:"location"`
	Name     pulumi.StringOutput    `pulumi:"name"`
	Tags     pulumi.StringMapOutput `pulumi:"tags"`
	Type     pulumi.StringOutput    `pulumi:"type"`
}


func NewDatabaseAccountTable(ctx *pulumi.Context,
	name string, args *DatabaseAccountTableArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
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
			Type: pulumi.String("azure-native:documentdb:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:DatabaseAccountTable"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:DatabaseAccountTable"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountTable
	err := ctx.RegisterResource("azure-native:documentdb/v20150408:DatabaseAccountTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseAccountTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountTableState, opts ...pulumi.ResourceOption) (*DatabaseAccountTable, error) {
	var resource DatabaseAccountTable
	err := ctx.ReadResource("azure-native:documentdb/v20150408:DatabaseAccountTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseAccountTableState struct {
}

type DatabaseAccountTableState struct {
}

func (DatabaseAccountTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountTableState)(nil)).Elem()
}

type databaseAccountTableArgs struct {
	AccountName       string            `pulumi:"accountName"`
	Options           map[string]string `pulumi:"options"`
	Resource          TableResource     `pulumi:"resource"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	TableName         *string           `pulumi:"tableName"`
}


type DatabaseAccountTableArgs struct {
	AccountName       pulumi.StringInput
	Options           pulumi.StringMapInput
	Resource          TableResourceInput
	ResourceGroupName pulumi.StringInput
	TableName         pulumi.StringPtrInput
}

func (DatabaseAccountTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountTableArgs)(nil)).Elem()
}

type DatabaseAccountTableInput interface {
	pulumi.Input

	ToDatabaseAccountTableOutput() DatabaseAccountTableOutput
	ToDatabaseAccountTableOutputWithContext(ctx context.Context) DatabaseAccountTableOutput
}

func (*DatabaseAccountTable) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountTable)(nil)).Elem()
}

func (i *DatabaseAccountTable) ToDatabaseAccountTableOutput() DatabaseAccountTableOutput {
	return i.ToDatabaseAccountTableOutputWithContext(context.Background())
}

func (i *DatabaseAccountTable) ToDatabaseAccountTableOutputWithContext(ctx context.Context) DatabaseAccountTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountTableOutput)
}

type DatabaseAccountTableOutput struct{ *pulumi.OutputState }

func (DatabaseAccountTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountTable)(nil)).Elem()
}

func (o DatabaseAccountTableOutput) ToDatabaseAccountTableOutput() DatabaseAccountTableOutput {
	return o
}

func (o DatabaseAccountTableOutput) ToDatabaseAccountTableOutputWithContext(ctx context.Context) DatabaseAccountTableOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountTableOutput{})
}
