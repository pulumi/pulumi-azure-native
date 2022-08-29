


package v20150408

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DatabaseAccountSqlDatabase struct {
	pulumi.CustomResourceState

	Colls    pulumi.StringPtrOutput `pulumi:"colls"`
	Etag     pulumi.StringPtrOutput `pulumi:"etag"`
	Location pulumi.StringPtrOutput `pulumi:"location"`
	Name     pulumi.StringOutput    `pulumi:"name"`
	Rid      pulumi.StringPtrOutput `pulumi:"rid"`
	Tags     pulumi.StringMapOutput `pulumi:"tags"`
	Ts       pulumi.AnyOutput       `pulumi:"ts"`
	Type     pulumi.StringOutput    `pulumi:"type"`
	Users    pulumi.StringPtrOutput `pulumi:"users"`
}


func NewDatabaseAccountSqlDatabase(ctx *pulumi.Context,
	name string, args *DatabaseAccountSqlDatabaseArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountSqlDatabase, error) {
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
			Type: pulumi.String("azure-native:documentdb:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:DatabaseAccountSqlDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:DatabaseAccountSqlDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountSqlDatabase
	err := ctx.RegisterResource("azure-native:documentdb/v20150408:DatabaseAccountSqlDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseAccountSqlDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountSqlDatabaseState, opts ...pulumi.ResourceOption) (*DatabaseAccountSqlDatabase, error) {
	var resource DatabaseAccountSqlDatabase
	err := ctx.ReadResource("azure-native:documentdb/v20150408:DatabaseAccountSqlDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseAccountSqlDatabaseState struct {
}

type DatabaseAccountSqlDatabaseState struct {
}

func (DatabaseAccountSqlDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountSqlDatabaseState)(nil)).Elem()
}

type databaseAccountSqlDatabaseArgs struct {
	AccountName       string              `pulumi:"accountName"`
	DatabaseName      *string             `pulumi:"databaseName"`
	Options           map[string]string   `pulumi:"options"`
	Resource          SqlDatabaseResource `pulumi:"resource"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
}


type DatabaseAccountSqlDatabaseArgs struct {
	AccountName       pulumi.StringInput
	DatabaseName      pulumi.StringPtrInput
	Options           pulumi.StringMapInput
	Resource          SqlDatabaseResourceInput
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountSqlDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountSqlDatabaseArgs)(nil)).Elem()
}

type DatabaseAccountSqlDatabaseInput interface {
	pulumi.Input

	ToDatabaseAccountSqlDatabaseOutput() DatabaseAccountSqlDatabaseOutput
	ToDatabaseAccountSqlDatabaseOutputWithContext(ctx context.Context) DatabaseAccountSqlDatabaseOutput
}

func (*DatabaseAccountSqlDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountSqlDatabase)(nil)).Elem()
}

func (i *DatabaseAccountSqlDatabase) ToDatabaseAccountSqlDatabaseOutput() DatabaseAccountSqlDatabaseOutput {
	return i.ToDatabaseAccountSqlDatabaseOutputWithContext(context.Background())
}

func (i *DatabaseAccountSqlDatabase) ToDatabaseAccountSqlDatabaseOutputWithContext(ctx context.Context) DatabaseAccountSqlDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountSqlDatabaseOutput)
}

type DatabaseAccountSqlDatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseAccountSqlDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountSqlDatabase)(nil)).Elem()
}

func (o DatabaseAccountSqlDatabaseOutput) ToDatabaseAccountSqlDatabaseOutput() DatabaseAccountSqlDatabaseOutput {
	return o
}

func (o DatabaseAccountSqlDatabaseOutput) ToDatabaseAccountSqlDatabaseOutputWithContext(ctx context.Context) DatabaseAccountSqlDatabaseOutput {
	return o
}

func (o DatabaseAccountSqlDatabaseOutput) Colls() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlDatabase) pulumi.StringPtrOutput { return v.Colls }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountSqlDatabaseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlDatabase) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountSqlDatabaseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlDatabase) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountSqlDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseAccountSqlDatabaseOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlDatabase) pulumi.StringPtrOutput { return v.Rid }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountSqlDatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlDatabase) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DatabaseAccountSqlDatabaseOutput) Ts() pulumi.AnyOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlDatabase) pulumi.AnyOutput { return v.Ts }).(pulumi.AnyOutput)
}

func (o DatabaseAccountSqlDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DatabaseAccountSqlDatabaseOutput) Users() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountSqlDatabase) pulumi.StringPtrOutput { return v.Users }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountSqlDatabaseOutput{})
}
