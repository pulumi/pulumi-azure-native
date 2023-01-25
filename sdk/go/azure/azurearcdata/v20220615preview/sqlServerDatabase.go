


package v20220615preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlServerDatabase struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                               `pulumi:"location"`
	Name       pulumi.StringOutput                               `pulumi:"name"`
	Properties SqlServerDatabaseResourcePropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                          `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                            `pulumi:"tags"`
	Type       pulumi.StringOutput                               `pulumi:"type"`
}


func NewSqlServerDatabase(ctx *pulumi.Context,
	name string, args *SqlServerDatabaseArgs, opts ...pulumi.ResourceOption) (*SqlServerDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SqlServerInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'SqlServerInstanceName'")
	}
	var resource SqlServerDatabase
	err := ctx.RegisterResource("azure-native:azurearcdata/v20220615preview:SqlServerDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlServerDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlServerDatabaseState, opts ...pulumi.ResourceOption) (*SqlServerDatabase, error) {
	var resource SqlServerDatabase
	err := ctx.ReadResource("azure-native:azurearcdata/v20220615preview:SqlServerDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlServerDatabaseState struct {
}

type SqlServerDatabaseState struct {
}

func (SqlServerDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerDatabaseState)(nil)).Elem()
}

type sqlServerDatabaseArgs struct {
	DatabaseName          *string                             `pulumi:"databaseName"`
	Location              *string                             `pulumi:"location"`
	Properties            SqlServerDatabaseResourceProperties `pulumi:"properties"`
	ResourceGroupName     string                              `pulumi:"resourceGroupName"`
	SqlServerInstanceName string                              `pulumi:"sqlServerInstanceName"`
	Tags                  map[string]string                   `pulumi:"tags"`
}


type SqlServerDatabaseArgs struct {
	DatabaseName          pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	Properties            SqlServerDatabaseResourcePropertiesInput
	ResourceGroupName     pulumi.StringInput
	SqlServerInstanceName pulumi.StringInput
	Tags                  pulumi.StringMapInput
}

func (SqlServerDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerDatabaseArgs)(nil)).Elem()
}

type SqlServerDatabaseInput interface {
	pulumi.Input

	ToSqlServerDatabaseOutput() SqlServerDatabaseOutput
	ToSqlServerDatabaseOutputWithContext(ctx context.Context) SqlServerDatabaseOutput
}

func (*SqlServerDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerDatabase)(nil)).Elem()
}

func (i *SqlServerDatabase) ToSqlServerDatabaseOutput() SqlServerDatabaseOutput {
	return i.ToSqlServerDatabaseOutputWithContext(context.Background())
}

func (i *SqlServerDatabase) ToSqlServerDatabaseOutputWithContext(ctx context.Context) SqlServerDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerDatabaseOutput)
}

type SqlServerDatabaseOutput struct{ *pulumi.OutputState }

func (SqlServerDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServerDatabase)(nil)).Elem()
}

func (o SqlServerDatabaseOutput) ToSqlServerDatabaseOutput() SqlServerDatabaseOutput {
	return o
}

func (o SqlServerDatabaseOutput) ToSqlServerDatabaseOutputWithContext(ctx context.Context) SqlServerDatabaseOutput {
	return o
}

func (o SqlServerDatabaseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServerDatabase) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SqlServerDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServerDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlServerDatabaseOutput) Properties() SqlServerDatabaseResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *SqlServerDatabase) SqlServerDatabaseResourcePropertiesResponseOutput { return v.Properties }).(SqlServerDatabaseResourcePropertiesResponseOutput)
}

func (o SqlServerDatabaseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlServerDatabase) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SqlServerDatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlServerDatabase) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlServerDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServerDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlServerDatabaseOutput{})
}
