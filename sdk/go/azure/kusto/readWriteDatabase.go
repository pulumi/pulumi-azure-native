


package kusto

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReadWriteDatabase struct {
	pulumi.CustomResourceState

	HotCachePeriod    pulumi.StringPtrOutput           `pulumi:"hotCachePeriod"`
	IsFollowed        pulumi.BoolOutput                `pulumi:"isFollowed"`
	Kind              pulumi.StringOutput              `pulumi:"kind"`
	Location          pulumi.StringPtrOutput           `pulumi:"location"`
	Name              pulumi.StringOutput              `pulumi:"name"`
	ProvisioningState pulumi.StringOutput              `pulumi:"provisioningState"`
	SoftDeletePeriod  pulumi.StringPtrOutput           `pulumi:"softDeletePeriod"`
	Statistics        DatabaseStatisticsResponseOutput `pulumi:"statistics"`
	Type              pulumi.StringOutput              `pulumi:"type"`
}


func NewReadWriteDatabase(ctx *pulumi.Context,
	name string, args *ReadWriteDatabaseArgs, opts ...pulumi.ResourceOption) (*ReadWriteDatabase, error) {
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
	args.Kind = pulumi.String("ReadWrite")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto/v20170907privatepreview:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20180907preview:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190121:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:ReadWriteDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource ReadWriteDatabase
	err := ctx.RegisterResource("azure-native:kusto:ReadWriteDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReadWriteDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReadWriteDatabaseState, opts ...pulumi.ResourceOption) (*ReadWriteDatabase, error) {
	var resource ReadWriteDatabase
	err := ctx.ReadResource("azure-native:kusto:ReadWriteDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type readWriteDatabaseState struct {
}

type ReadWriteDatabaseState struct {
}

func (ReadWriteDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*readWriteDatabaseState)(nil)).Elem()
}

type readWriteDatabaseArgs struct {
	ClusterName       string  `pulumi:"clusterName"`
	DatabaseName      *string `pulumi:"databaseName"`
	HotCachePeriod    *string `pulumi:"hotCachePeriod"`
	Kind              string  `pulumi:"kind"`
	Location          *string `pulumi:"location"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SoftDeletePeriod  *string `pulumi:"softDeletePeriod"`
}


type ReadWriteDatabaseArgs struct {
	ClusterName       pulumi.StringInput
	DatabaseName      pulumi.StringPtrInput
	HotCachePeriod    pulumi.StringPtrInput
	Kind              pulumi.StringInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	SoftDeletePeriod  pulumi.StringPtrInput
}

func (ReadWriteDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*readWriteDatabaseArgs)(nil)).Elem()
}

type ReadWriteDatabaseInput interface {
	pulumi.Input

	ToReadWriteDatabaseOutput() ReadWriteDatabaseOutput
	ToReadWriteDatabaseOutputWithContext(ctx context.Context) ReadWriteDatabaseOutput
}

func (*ReadWriteDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadWriteDatabase)(nil)).Elem()
}

func (i *ReadWriteDatabase) ToReadWriteDatabaseOutput() ReadWriteDatabaseOutput {
	return i.ToReadWriteDatabaseOutputWithContext(context.Background())
}

func (i *ReadWriteDatabase) ToReadWriteDatabaseOutputWithContext(ctx context.Context) ReadWriteDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadWriteDatabaseOutput)
}

type ReadWriteDatabaseOutput struct{ *pulumi.OutputState }

func (ReadWriteDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadWriteDatabase)(nil)).Elem()
}

func (o ReadWriteDatabaseOutput) ToReadWriteDatabaseOutput() ReadWriteDatabaseOutput {
	return o
}

func (o ReadWriteDatabaseOutput) ToReadWriteDatabaseOutputWithContext(ctx context.Context) ReadWriteDatabaseOutput {
	return o
}

func (o ReadWriteDatabaseOutput) HotCachePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.StringPtrOutput { return v.HotCachePeriod }).(pulumi.StringPtrOutput)
}

func (o ReadWriteDatabaseOutput) IsFollowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.BoolOutput { return v.IsFollowed }).(pulumi.BoolOutput)
}

func (o ReadWriteDatabaseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ReadWriteDatabaseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ReadWriteDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReadWriteDatabaseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ReadWriteDatabaseOutput) SoftDeletePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.StringPtrOutput { return v.SoftDeletePeriod }).(pulumi.StringPtrOutput)
}

func (o ReadWriteDatabaseOutput) Statistics() DatabaseStatisticsResponseOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) DatabaseStatisticsResponseOutput { return v.Statistics }).(DatabaseStatisticsResponseOutput)
}

func (o ReadWriteDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadWriteDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReadWriteDatabaseOutput{})
}
