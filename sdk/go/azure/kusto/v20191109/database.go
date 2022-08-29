


package v20191109

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Database struct {
	pulumi.CustomResourceState

	Kind     pulumi.StringOutput    `pulumi:"kind"`
	Location pulumi.StringPtrOutput `pulumi:"location"`
	Name     pulumi.StringOutput    `pulumi:"name"`
	Type     pulumi.StringOutput    `pulumi:"type"`
}


func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:Database"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20170907privatepreview:Database"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20180907preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190121:Database"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:Database"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:Database"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:Database"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:Database"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:Database"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:Database"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:Database"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:Database"),
		},
	})
	opts = append(opts, aliases)
	var resource Database
	err := ctx.RegisterResource("azure-native:kusto/v20191109:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("azure-native:kusto/v20191109:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseState struct {
}

type DatabaseState struct {
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	ClusterName       string  `pulumi:"clusterName"`
	DatabaseName      *string `pulumi:"databaseName"`
	Kind              string  `pulumi:"kind"`
	Location          *string `pulumi:"location"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type DatabaseArgs struct {
	ClusterName       pulumi.StringInput
	DatabaseName      pulumi.StringPtrInput
	Kind              pulumi.StringInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

func (o DatabaseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o DatabaseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseOutput{})
}
