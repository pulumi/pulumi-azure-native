


package v20191212

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlResourceSqlDatabase struct {
	pulumi.CustomResourceState

	Location pulumi.StringPtrOutput                            `pulumi:"location"`
	Name     pulumi.StringOutput                               `pulumi:"name"`
	Resource SqlDatabaseGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	Tags     pulumi.StringMapOutput                            `pulumi:"tags"`
	Type     pulumi.StringOutput                               `pulumi:"type"`
}


func NewSqlResourceSqlDatabase(ctx *pulumi.Context,
	name string, args *SqlResourceSqlDatabaseArgs, opts ...pulumi.ResourceOption) (*SqlResourceSqlDatabase, error) {
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
			Type: pulumi.String("azure-nextgen:documentdb/v20191212:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150401:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150408:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20151106:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160319:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160331:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20190801:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200301:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200401:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200601preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200901:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210115:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210301preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210315:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210401preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210415:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210515:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210615:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:SqlResourceSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20210701preview:SqlResourceSqlDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlResourceSqlDatabase
	err := ctx.RegisterResource("azure-native:documentdb/v20191212:SqlResourceSqlDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlResourceSqlDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlResourceSqlDatabaseState, opts ...pulumi.ResourceOption) (*SqlResourceSqlDatabase, error) {
	var resource SqlResourceSqlDatabase
	err := ctx.ReadResource("azure-native:documentdb/v20191212:SqlResourceSqlDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlResourceSqlDatabaseState struct {
}

type SqlResourceSqlDatabaseState struct {
}

func (SqlResourceSqlDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlDatabaseState)(nil)).Elem()
}

type sqlResourceSqlDatabaseArgs struct {
	AccountName       string              `pulumi:"accountName"`
	DatabaseName      *string             `pulumi:"databaseName"`
	Location          *string             `pulumi:"location"`
	Options           CreateUpdateOptions `pulumi:"options"`
	Resource          SqlDatabaseResource `pulumi:"resource"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	Tags              map[string]string   `pulumi:"tags"`
}


type SqlResourceSqlDatabaseArgs struct {
	AccountName       pulumi.StringInput
	DatabaseName      pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Options           CreateUpdateOptionsInput
	Resource          SqlDatabaseResourceInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (SqlResourceSqlDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlDatabaseArgs)(nil)).Elem()
}

type SqlResourceSqlDatabaseInput interface {
	pulumi.Input

	ToSqlResourceSqlDatabaseOutput() SqlResourceSqlDatabaseOutput
	ToSqlResourceSqlDatabaseOutputWithContext(ctx context.Context) SqlResourceSqlDatabaseOutput
}

func (*SqlResourceSqlDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlResourceSqlDatabase)(nil))
}

func (i *SqlResourceSqlDatabase) ToSqlResourceSqlDatabaseOutput() SqlResourceSqlDatabaseOutput {
	return i.ToSqlResourceSqlDatabaseOutputWithContext(context.Background())
}

func (i *SqlResourceSqlDatabase) ToSqlResourceSqlDatabaseOutputWithContext(ctx context.Context) SqlResourceSqlDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlResourceSqlDatabaseOutput)
}

type SqlResourceSqlDatabaseOutput struct{ *pulumi.OutputState }

func (SqlResourceSqlDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlResourceSqlDatabase)(nil))
}

func (o SqlResourceSqlDatabaseOutput) ToSqlResourceSqlDatabaseOutput() SqlResourceSqlDatabaseOutput {
	return o
}

func (o SqlResourceSqlDatabaseOutput) ToSqlResourceSqlDatabaseOutputWithContext(ctx context.Context) SqlResourceSqlDatabaseOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SqlResourceSqlDatabaseOutput{})
}
