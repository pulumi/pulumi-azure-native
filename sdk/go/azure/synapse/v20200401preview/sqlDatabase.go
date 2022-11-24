


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
	return reflect.TypeOf((**SqlDatabase)(nil)).Elem()
}

func (i *SqlDatabase) ToSqlDatabaseOutput() SqlDatabaseOutput {
	return i.ToSqlDatabaseOutputWithContext(context.Background())
}

func (i *SqlDatabase) ToSqlDatabaseOutputWithContext(ctx context.Context) SqlDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseOutput)
}

type SqlDatabaseOutput struct{ *pulumi.OutputState }

func (SqlDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDatabase)(nil)).Elem()
}

func (o SqlDatabaseOutput) ToSqlDatabaseOutput() SqlDatabaseOutput {
	return o
}

func (o SqlDatabaseOutput) ToSqlDatabaseOutputWithContext(ctx context.Context) SqlDatabaseOutput {
	return o
}

func (o SqlDatabaseOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabase) pulumi.StringPtrOutput { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseOutput) DataRetention() SqlDatabaseDataRetentionResponsePtrOutput {
	return o.ApplyT(func(v *SqlDatabase) SqlDatabaseDataRetentionResponsePtrOutput { return v.DataRetention }).(SqlDatabaseDataRetentionResponsePtrOutput)
}

func (o SqlDatabaseOutput) DatabaseGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDatabase) pulumi.StringOutput { return v.DatabaseGuid }).(pulumi.StringOutput)
}

func (o SqlDatabaseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDatabase) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SqlDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlDatabaseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDatabase) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SqlDatabaseOutput) StorageRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabase) pulumi.StringPtrOutput { return v.StorageRedundancy }).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlDatabase) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SqlDatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlDatabase) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlDatabaseOutput{})
}
