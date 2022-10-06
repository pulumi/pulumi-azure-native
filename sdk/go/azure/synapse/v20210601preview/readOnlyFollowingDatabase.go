


package v20210601preview

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
	SystemData                        SystemDataResponseOutput         `pulumi:"systemData"`
	Type                              pulumi.StringOutput              `pulumi:"type"`
}


func NewReadOnlyFollowingDatabase(ctx *pulumi.Context,
	name string, args *ReadOnlyFollowingDatabaseArgs, opts ...pulumi.ResourceOption) (*ReadOnlyFollowingDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.KustoPoolName == nil {
		return nil, errors.New("invalid value for required argument 'KustoPoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("ReadOnlyFollowing")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:ReadOnlyFollowingDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource ReadOnlyFollowingDatabase
	err := ctx.RegisterResource("azure-native:synapse/v20210601preview:ReadOnlyFollowingDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReadOnlyFollowingDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReadOnlyFollowingDatabaseState, opts ...pulumi.ResourceOption) (*ReadOnlyFollowingDatabase, error) {
	var resource ReadOnlyFollowingDatabase
	err := ctx.ReadResource("azure-native:synapse/v20210601preview:ReadOnlyFollowingDatabase", name, id, state, &resource, opts...)
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
	DatabaseName      *string `pulumi:"databaseName"`
	HotCachePeriod    *string `pulumi:"hotCachePeriod"`
	Kind              string  `pulumi:"kind"`
	KustoPoolName     string  `pulumi:"kustoPoolName"`
	Location          *string `pulumi:"location"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type ReadOnlyFollowingDatabaseArgs struct {
	DatabaseName      pulumi.StringPtrInput
	HotCachePeriod    pulumi.StringPtrInput
	Kind              pulumi.StringInput
	KustoPoolName     pulumi.StringInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
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

func (o ReadOnlyFollowingDatabaseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ReadOnlyFollowingDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReadOnlyFollowingDatabaseOutput{})
}
