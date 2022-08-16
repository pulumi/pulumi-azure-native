


package v20160331

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DatabaseAccountGremlinDatabase struct {
	pulumi.CustomResourceState

	Etag     pulumi.StringPtrOutput `pulumi:"etag"`
	Location pulumi.StringPtrOutput `pulumi:"location"`
	Name     pulumi.StringOutput    `pulumi:"name"`
	Rid      pulumi.StringPtrOutput `pulumi:"rid"`
	Tags     pulumi.StringMapOutput `pulumi:"tags"`
	Ts       pulumi.AnyOutput       `pulumi:"ts"`
	Type     pulumi.StringOutput    `pulumi:"type"`
}


func NewDatabaseAccountGremlinDatabase(ctx *pulumi.Context,
	name string, args *DatabaseAccountGremlinDatabaseArgs, opts ...pulumi.ResourceOption) (*DatabaseAccountGremlinDatabase, error) {
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
			Type: pulumi.String("azure-native:documentdb:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:DatabaseAccountGremlinDatabase"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:DatabaseAccountGremlinDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccountGremlinDatabase
	err := ctx.RegisterResource("azure-native:documentdb/v20160331:DatabaseAccountGremlinDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseAccountGremlinDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountGremlinDatabaseState, opts ...pulumi.ResourceOption) (*DatabaseAccountGremlinDatabase, error) {
	var resource DatabaseAccountGremlinDatabase
	err := ctx.ReadResource("azure-native:documentdb/v20160331:DatabaseAccountGremlinDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseAccountGremlinDatabaseState struct {
}

type DatabaseAccountGremlinDatabaseState struct {
}

func (DatabaseAccountGremlinDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountGremlinDatabaseState)(nil)).Elem()
}

type databaseAccountGremlinDatabaseArgs struct {
	AccountName       string                  `pulumi:"accountName"`
	DatabaseName      *string                 `pulumi:"databaseName"`
	Options           map[string]string       `pulumi:"options"`
	Resource          GremlinDatabaseResource `pulumi:"resource"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
}


type DatabaseAccountGremlinDatabaseArgs struct {
	AccountName       pulumi.StringInput
	DatabaseName      pulumi.StringPtrInput
	Options           pulumi.StringMapInput
	Resource          GremlinDatabaseResourceInput
	ResourceGroupName pulumi.StringInput
}

func (DatabaseAccountGremlinDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountGremlinDatabaseArgs)(nil)).Elem()
}

type DatabaseAccountGremlinDatabaseInput interface {
	pulumi.Input

	ToDatabaseAccountGremlinDatabaseOutput() DatabaseAccountGremlinDatabaseOutput
	ToDatabaseAccountGremlinDatabaseOutputWithContext(ctx context.Context) DatabaseAccountGremlinDatabaseOutput
}

func (*DatabaseAccountGremlinDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountGremlinDatabase)(nil)).Elem()
}

func (i *DatabaseAccountGremlinDatabase) ToDatabaseAccountGremlinDatabaseOutput() DatabaseAccountGremlinDatabaseOutput {
	return i.ToDatabaseAccountGremlinDatabaseOutputWithContext(context.Background())
}

func (i *DatabaseAccountGremlinDatabase) ToDatabaseAccountGremlinDatabaseOutputWithContext(ctx context.Context) DatabaseAccountGremlinDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountGremlinDatabaseOutput)
}

type DatabaseAccountGremlinDatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseAccountGremlinDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountGremlinDatabase)(nil)).Elem()
}

func (o DatabaseAccountGremlinDatabaseOutput) ToDatabaseAccountGremlinDatabaseOutput() DatabaseAccountGremlinDatabaseOutput {
	return o
}

func (o DatabaseAccountGremlinDatabaseOutput) ToDatabaseAccountGremlinDatabaseOutputWithContext(ctx context.Context) DatabaseAccountGremlinDatabaseOutput {
	return o
}

func (o DatabaseAccountGremlinDatabaseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountGremlinDatabase) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountGremlinDatabaseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountGremlinDatabase) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountGremlinDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountGremlinDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseAccountGremlinDatabaseOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccountGremlinDatabase) pulumi.StringPtrOutput { return v.Rid }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountGremlinDatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseAccountGremlinDatabase) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DatabaseAccountGremlinDatabaseOutput) Ts() pulumi.AnyOutput {
	return o.ApplyT(func(v *DatabaseAccountGremlinDatabase) pulumi.AnyOutput { return v.Ts }).(pulumi.AnyOutput)
}

func (o DatabaseAccountGremlinDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccountGremlinDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountGremlinDatabaseOutput{})
}
