


package v20170907privatepreview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Database struct {
	pulumi.CustomResourceState

	Etag                   pulumi.StringOutput              `pulumi:"etag"`
	HotCachePeriodInDays   pulumi.IntPtrOutput              `pulumi:"hotCachePeriodInDays"`
	Location               pulumi.StringOutput              `pulumi:"location"`
	Name                   pulumi.StringOutput              `pulumi:"name"`
	ProvisioningState      pulumi.StringOutput              `pulumi:"provisioningState"`
	SoftDeletePeriodInDays pulumi.IntOutput                 `pulumi:"softDeletePeriodInDays"`
	Statistics             DatabaseStatisticsResponseOutput `pulumi:"statistics"`
	Tags                   pulumi.StringMapOutput           `pulumi:"tags"`
	Type                   pulumi.StringOutput              `pulumi:"type"`
}


func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SoftDeletePeriodInDays == nil {
		return nil, errors.New("invalid value for required argument 'SoftDeletePeriodInDays'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:Database"),
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
			Type: pulumi.String("azure-native:kusto/v20191109:Database"),
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
		{
			Type: pulumi.String("azure-native:kusto/v20220707:Database"),
		},
	})
	opts = append(opts, aliases)
	var resource Database
	err := ctx.RegisterResource("azure-native:kusto/v20170907privatepreview:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("azure-native:kusto/v20170907privatepreview:Database", name, id, state, &resource, opts...)
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
	ClusterName            string            `pulumi:"clusterName"`
	DatabaseName           *string           `pulumi:"databaseName"`
	HotCachePeriodInDays   *int              `pulumi:"hotCachePeriodInDays"`
	Location               *string           `pulumi:"location"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	SoftDeletePeriodInDays int               `pulumi:"softDeletePeriodInDays"`
	Tags                   map[string]string `pulumi:"tags"`
}


type DatabaseArgs struct {
	ClusterName            pulumi.StringInput
	DatabaseName           pulumi.StringPtrInput
	HotCachePeriodInDays   pulumi.IntPtrInput
	Location               pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	SoftDeletePeriodInDays pulumi.IntInput
	Tags                   pulumi.StringMapInput
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

func (o DatabaseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DatabaseOutput) HotCachePeriodInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.IntPtrOutput { return v.HotCachePeriodInDays }).(pulumi.IntPtrOutput)
}

func (o DatabaseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DatabaseOutput) SoftDeletePeriodInDays() pulumi.IntOutput {
	return o.ApplyT(func(v *Database) pulumi.IntOutput { return v.SoftDeletePeriodInDays }).(pulumi.IntOutput)
}

func (o DatabaseOutput) Statistics() DatabaseStatisticsResponseOutput {
	return o.ApplyT(func(v *Database) DatabaseStatisticsResponseOutput { return v.Statistics }).(DatabaseStatisticsResponseOutput)
}

func (o DatabaseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Database) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseOutput{})
}
