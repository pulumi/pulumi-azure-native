


package kusto

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReadOnlyFollowingDatabase struct {
	pulumi.CustomResourceState

	AttachedDatabaseConfigurationName pulumi.StringOutput              `pulumi:"attachedDatabaseConfigurationName"`
	HotCachePeriod                    pulumi.StringPtrOutput           `pulumi:"hotCachePeriod"`
	Kind                              pulumi.StringOutput              `pulumi:"kind"`
	LeaderClusterResourceId           pulumi.StringOutput              `pulumi:"leaderClusterResourceId"`
	Location                          pulumi.StringPtrOutput           `pulumi:"location"`
	Name                              pulumi.StringOutput              `pulumi:"name"`
	PrincipalsModificationKind        pulumi.StringOutput              `pulumi:"principalsModificationKind"`
	ProvisioningState                 pulumi.StringOutput              `pulumi:"provisioningState"`
	SoftDeletePeriod                  pulumi.StringOutput              `pulumi:"softDeletePeriod"`
	Statistics                        DatabaseStatisticsResponseOutput `pulumi:"statistics"`
	Type                              pulumi.StringOutput              `pulumi:"type"`
}


func NewReadOnlyFollowingDatabase(ctx *pulumi.Context,
	name string, args *ReadOnlyFollowingDatabaseArgs, opts ...pulumi.ResourceOption) (*ReadOnlyFollowingDatabase, error) {
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
	args.Kind = pulumi.String("ReadOnlyFollowing")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto/v20170907privatepreview:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20180907preview:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190121:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220707:ReadOnlyFollowingDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource ReadOnlyFollowingDatabase
	err := ctx.RegisterResource("azure-native:kusto:ReadOnlyFollowingDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReadOnlyFollowingDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReadOnlyFollowingDatabaseState, opts ...pulumi.ResourceOption) (*ReadOnlyFollowingDatabase, error) {
	var resource ReadOnlyFollowingDatabase
	err := ctx.ReadResource("azure-native:kusto:ReadOnlyFollowingDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type readOnlyFollowingDatabaseState struct {
}

type ReadOnlyFollowingDatabaseState struct {
}

func (ReadOnlyFollowingDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*readOnlyFollowingDatabaseState)(nil)).Elem()
}

type readOnlyFollowingDatabaseArgs struct {
	ClusterName       string  `pulumi:"clusterName"`
	DatabaseName      *string `pulumi:"databaseName"`
	HotCachePeriod    *string `pulumi:"hotCachePeriod"`
	Kind              string  `pulumi:"kind"`
	Location          *string `pulumi:"location"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type ReadOnlyFollowingDatabaseArgs struct {
	ClusterName       pulumi.StringInput
	DatabaseName      pulumi.StringPtrInput
	HotCachePeriod    pulumi.StringPtrInput
	Kind              pulumi.StringInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (ReadOnlyFollowingDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*readOnlyFollowingDatabaseArgs)(nil)).Elem()
}

type ReadOnlyFollowingDatabaseInput interface {
	pulumi.Input

	ToReadOnlyFollowingDatabaseOutput() ReadOnlyFollowingDatabaseOutput
	ToReadOnlyFollowingDatabaseOutputWithContext(ctx context.Context) ReadOnlyFollowingDatabaseOutput
}

func (*ReadOnlyFollowingDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadOnlyFollowingDatabase)(nil)).Elem()
}

func (i *ReadOnlyFollowingDatabase) ToReadOnlyFollowingDatabaseOutput() ReadOnlyFollowingDatabaseOutput {
	return i.ToReadOnlyFollowingDatabaseOutputWithContext(context.Background())
}

func (i *ReadOnlyFollowingDatabase) ToReadOnlyFollowingDatabaseOutputWithContext(ctx context.Context) ReadOnlyFollowingDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadOnlyFollowingDatabaseOutput)
}

type ReadOnlyFollowingDatabaseOutput struct{ *pulumi.OutputState }

func (ReadOnlyFollowingDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadOnlyFollowingDatabase)(nil)).Elem()
}

func (o ReadOnlyFollowingDatabaseOutput) ToReadOnlyFollowingDatabaseOutput() ReadOnlyFollowingDatabaseOutput {
	return o
}

func (o ReadOnlyFollowingDatabaseOutput) ToReadOnlyFollowingDatabaseOutputWithContext(ctx context.Context) ReadOnlyFollowingDatabaseOutput {
	return o
}

func (o ReadOnlyFollowingDatabaseOutput) AttachedDatabaseConfigurationName() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.AttachedDatabaseConfigurationName }).(pulumi.StringOutput)
}

func (o ReadOnlyFollowingDatabaseOutput) HotCachePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringPtrOutput { return v.HotCachePeriod }).(pulumi.StringPtrOutput)
}

func (o ReadOnlyFollowingDatabaseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ReadOnlyFollowingDatabaseOutput) LeaderClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.LeaderClusterResourceId }).(pulumi.StringOutput)
}

func (o ReadOnlyFollowingDatabaseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ReadOnlyFollowingDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReadOnlyFollowingDatabaseOutput) PrincipalsModificationKind() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.PrincipalsModificationKind }).(pulumi.StringOutput)
}

func (o ReadOnlyFollowingDatabaseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ReadOnlyFollowingDatabaseOutput) SoftDeletePeriod() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.SoftDeletePeriod }).(pulumi.StringOutput)
}

func (o ReadOnlyFollowingDatabaseOutput) Statistics() DatabaseStatisticsResponseOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) DatabaseStatisticsResponseOutput { return v.Statistics }).(DatabaseStatisticsResponseOutput)
}

func (o ReadOnlyFollowingDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReadOnlyFollowingDatabaseOutput{})
}
