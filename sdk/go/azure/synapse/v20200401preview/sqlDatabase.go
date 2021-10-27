


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlDatabase struct {
	pulumi.CustomResourceState

	Collation         pulumi.StringPtrOutput                    `pulumi:"collation"`
	DataRetention     SqlDatabaseDataRetentionResponsePtrOutput `pulumi:"dataRetention"`
	DatabaseGuid      pulumi.StringOutput                       `pulumi:"databaseGuid"`
	Location          pulumi.StringOutput                       `pulumi:"location"`
	Name              pulumi.StringOutput                       `pulumi:"name"`
	Status            pulumi.StringOutput                       `pulumi:"status"`
	StorageRedundancy pulumi.StringPtrOutput                    `pulumi:"storageRedundancy"`
	SystemData        SystemDataResponseOutput                  `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                    `pulumi:"tags"`
	Type              pulumi.StringOutput                       `pulumi:"type"`
}


func NewSqlDatabase(ctx *pulumi.Context,
	name string, args *SqlDatabaseArgs, opts ...pulumi.ResourceOption) (*SqlDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:synapse/v20200401preview:SqlDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlDatabase
	err := ctx.RegisterResource("azure-native:synapse/v20200401preview:SqlDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlDatabaseState, opts ...pulumi.ResourceOption) (*SqlDatabase, error) {
	var resource SqlDatabase
	err := ctx.ReadResource("azure-native:synapse/v20200401preview:SqlDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlDatabaseState struct {
}

type SqlDatabaseState struct {
}

func (SqlDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDatabaseState)(nil)).Elem()
}

type sqlDatabaseArgs struct {
	Collation         *string                   `pulumi:"collation"`
	DataRetention     *SqlDatabaseDataRetention `pulumi:"dataRetention"`
	Location          *string                   `pulumi:"location"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	SqlDatabaseName   *string                   `pulumi:"sqlDatabaseName"`
	StorageRedundancy *string                   `pulumi:"storageRedundancy"`
	Tags              map[string]string         `pulumi:"tags"`
	WorkspaceName     string                    `pulumi:"workspaceName"`
}


type SqlDatabaseArgs struct {
	Collation         pulumi.StringPtrInput
	DataRetention     SqlDatabaseDataRetentionPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	SqlDatabaseName   pulumi.StringPtrInput
	StorageRedundancy pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	WorkspaceName     pulumi.StringInput
}

func (SqlDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDatabaseArgs)(nil)).Elem()
}

type SqlDatabaseInput interface {
	pulumi.Input

	ToSqlDatabaseOutput() SqlDatabaseOutput
	ToSqlDatabaseOutputWithContext(ctx context.Context) SqlDatabaseOutput
}

func (*SqlDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabase)(nil))
}

func (i *SqlDatabase) ToSqlDatabaseOutput() SqlDatabaseOutput {
	return i.ToSqlDatabaseOutputWithContext(context.Background())
}

func (i *SqlDatabase) ToSqlDatabaseOutputWithContext(ctx context.Context) SqlDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseOutput)
}

type SqlDatabaseOutput struct{ *pulumi.OutputState }

func (SqlDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabase)(nil))
}

func (o SqlDatabaseOutput) ToSqlDatabaseOutput() SqlDatabaseOutput {
	return o
}

func (o SqlDatabaseOutput) ToSqlDatabaseOutputWithContext(ctx context.Context) SqlDatabaseOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SqlDatabaseInput)(nil)).Elem(), &SqlDatabase{})
	pulumi.RegisterOutputType(SqlDatabaseOutput{})
}
