


package datamigration

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseMigrationsSqlDb struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                            `pulumi:"name"`
	Properties DatabaseMigrationPropertiesSqlDbResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                       `pulumi:"systemData"`
	Type       pulumi.StringOutput                            `pulumi:"type"`
}


func NewDatabaseMigrationsSqlDb(ctx *pulumi.Context,
	name string, args *DatabaseMigrationsSqlDbArgs, opts ...pulumi.ResourceOption) (*DatabaseMigrationsSqlDb, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SqlDbInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'SqlDbInstanceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datamigration/v20220330preview:DatabaseMigrationsSqlDb"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseMigrationsSqlDb
	err := ctx.RegisterResource("azure-native:datamigration:DatabaseMigrationsSqlDb", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseMigrationsSqlDb(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseMigrationsSqlDbState, opts ...pulumi.ResourceOption) (*DatabaseMigrationsSqlDb, error) {
	var resource DatabaseMigrationsSqlDb
	err := ctx.ReadResource("azure-native:datamigration:DatabaseMigrationsSqlDb", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseMigrationsSqlDbState struct {
}

type DatabaseMigrationsSqlDbState struct {
}

func (DatabaseMigrationsSqlDbState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseMigrationsSqlDbState)(nil)).Elem()
}

type databaseMigrationsSqlDbArgs struct {
	Properties        *DatabaseMigrationPropertiesSqlDb `pulumi:"properties"`
	ResourceGroupName string                            `pulumi:"resourceGroupName"`
	SqlDbInstanceName string                            `pulumi:"sqlDbInstanceName"`
	TargetDbName      *string                           `pulumi:"targetDbName"`
}


type DatabaseMigrationsSqlDbArgs struct {
	Properties        DatabaseMigrationPropertiesSqlDbPtrInput
	ResourceGroupName pulumi.StringInput
	SqlDbInstanceName pulumi.StringInput
	TargetDbName      pulumi.StringPtrInput
}

func (DatabaseMigrationsSqlDbArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseMigrationsSqlDbArgs)(nil)).Elem()
}

type DatabaseMigrationsSqlDbInput interface {
	pulumi.Input

	ToDatabaseMigrationsSqlDbOutput() DatabaseMigrationsSqlDbOutput
	ToDatabaseMigrationsSqlDbOutputWithContext(ctx context.Context) DatabaseMigrationsSqlDbOutput
}

func (*DatabaseMigrationsSqlDb) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseMigrationsSqlDb)(nil)).Elem()
}

func (i *DatabaseMigrationsSqlDb) ToDatabaseMigrationsSqlDbOutput() DatabaseMigrationsSqlDbOutput {
	return i.ToDatabaseMigrationsSqlDbOutputWithContext(context.Background())
}

func (i *DatabaseMigrationsSqlDb) ToDatabaseMigrationsSqlDbOutputWithContext(ctx context.Context) DatabaseMigrationsSqlDbOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMigrationsSqlDbOutput)
}

type DatabaseMigrationsSqlDbOutput struct{ *pulumi.OutputState }

func (DatabaseMigrationsSqlDbOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseMigrationsSqlDb)(nil)).Elem()
}

func (o DatabaseMigrationsSqlDbOutput) ToDatabaseMigrationsSqlDbOutput() DatabaseMigrationsSqlDbOutput {
	return o
}

func (o DatabaseMigrationsSqlDbOutput) ToDatabaseMigrationsSqlDbOutputWithContext(ctx context.Context) DatabaseMigrationsSqlDbOutput {
	return o
}

func (o DatabaseMigrationsSqlDbOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseMigrationsSqlDb) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseMigrationsSqlDbOutput) Properties() DatabaseMigrationPropertiesSqlDbResponseOutput {
	return o.ApplyT(func(v *DatabaseMigrationsSqlDb) DatabaseMigrationPropertiesSqlDbResponseOutput { return v.Properties }).(DatabaseMigrationPropertiesSqlDbResponseOutput)
}

func (o DatabaseMigrationsSqlDbOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DatabaseMigrationsSqlDb) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DatabaseMigrationsSqlDbOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseMigrationsSqlDb) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseMigrationsSqlDbOutput{})
}
