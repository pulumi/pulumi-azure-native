


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DistributedAvailabilityGroup struct {
	pulumi.CustomResourceState

	DistributedAvailabilityGroupId pulumi.StringOutput    `pulumi:"distributedAvailabilityGroupId"`
	LastHardenedLsn                pulumi.StringOutput    `pulumi:"lastHardenedLsn"`
	LinkState                      pulumi.StringOutput    `pulumi:"linkState"`
	Name                           pulumi.StringOutput    `pulumi:"name"`
	PrimaryAvailabilityGroupName   pulumi.StringPtrOutput `pulumi:"primaryAvailabilityGroupName"`
	ReplicationMode                pulumi.StringPtrOutput `pulumi:"replicationMode"`
	SecondaryAvailabilityGroupName pulumi.StringPtrOutput `pulumi:"secondaryAvailabilityGroupName"`
	SourceEndpoint                 pulumi.StringPtrOutput `pulumi:"sourceEndpoint"`
	SourceReplicaId                pulumi.StringOutput    `pulumi:"sourceReplicaId"`
	TargetDatabase                 pulumi.StringPtrOutput `pulumi:"targetDatabase"`
	TargetReplicaId                pulumi.StringOutput    `pulumi:"targetReplicaId"`
	Type                           pulumi.StringOutput    `pulumi:"type"`
}


func NewDistributedAvailabilityGroup(ctx *pulumi.Context,
	name string, args *DistributedAvailabilityGroupArgs, opts ...pulumi.ResourceOption) (*DistributedAvailabilityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:DistributedAvailabilityGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:DistributedAvailabilityGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:DistributedAvailabilityGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:DistributedAvailabilityGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:DistributedAvailabilityGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:DistributedAvailabilityGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource DistributedAvailabilityGroup
	err := ctx.RegisterResource("azure-native:sql/v20210501preview:DistributedAvailabilityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDistributedAvailabilityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DistributedAvailabilityGroupState, opts ...pulumi.ResourceOption) (*DistributedAvailabilityGroup, error) {
	var resource DistributedAvailabilityGroup
	err := ctx.ReadResource("azure-native:sql/v20210501preview:DistributedAvailabilityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type distributedAvailabilityGroupState struct {
}

type DistributedAvailabilityGroupState struct {
}

func (DistributedAvailabilityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*distributedAvailabilityGroupState)(nil)).Elem()
}

type distributedAvailabilityGroupArgs struct {
	DistributedAvailabilityGroupName *string `pulumi:"distributedAvailabilityGroupName"`
	ManagedInstanceName              string  `pulumi:"managedInstanceName"`
	PrimaryAvailabilityGroupName     *string `pulumi:"primaryAvailabilityGroupName"`
	ReplicationMode                  *string `pulumi:"replicationMode"`
	ResourceGroupName                string  `pulumi:"resourceGroupName"`
	SecondaryAvailabilityGroupName   *string `pulumi:"secondaryAvailabilityGroupName"`
	SourceEndpoint                   *string `pulumi:"sourceEndpoint"`
	TargetDatabase                   *string `pulumi:"targetDatabase"`
}


type DistributedAvailabilityGroupArgs struct {
	DistributedAvailabilityGroupName pulumi.StringPtrInput
	ManagedInstanceName              pulumi.StringInput
	PrimaryAvailabilityGroupName     pulumi.StringPtrInput
	ReplicationMode                  pulumi.StringPtrInput
	ResourceGroupName                pulumi.StringInput
	SecondaryAvailabilityGroupName   pulumi.StringPtrInput
	SourceEndpoint                   pulumi.StringPtrInput
	TargetDatabase                   pulumi.StringPtrInput
}

func (DistributedAvailabilityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*distributedAvailabilityGroupArgs)(nil)).Elem()
}

type DistributedAvailabilityGroupInput interface {
	pulumi.Input

	ToDistributedAvailabilityGroupOutput() DistributedAvailabilityGroupOutput
	ToDistributedAvailabilityGroupOutputWithContext(ctx context.Context) DistributedAvailabilityGroupOutput
}

func (*DistributedAvailabilityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributedAvailabilityGroup)(nil)).Elem()
}

func (i *DistributedAvailabilityGroup) ToDistributedAvailabilityGroupOutput() DistributedAvailabilityGroupOutput {
	return i.ToDistributedAvailabilityGroupOutputWithContext(context.Background())
}

func (i *DistributedAvailabilityGroup) ToDistributedAvailabilityGroupOutputWithContext(ctx context.Context) DistributedAvailabilityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributedAvailabilityGroupOutput)
}

type DistributedAvailabilityGroupOutput struct{ *pulumi.OutputState }

func (DistributedAvailabilityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributedAvailabilityGroup)(nil)).Elem()
}

func (o DistributedAvailabilityGroupOutput) ToDistributedAvailabilityGroupOutput() DistributedAvailabilityGroupOutput {
	return o
}

func (o DistributedAvailabilityGroupOutput) ToDistributedAvailabilityGroupOutputWithContext(ctx context.Context) DistributedAvailabilityGroupOutput {
	return o
}

func (o DistributedAvailabilityGroupOutput) DistributedAvailabilityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.DistributedAvailabilityGroupId }).(pulumi.StringOutput)
}

func (o DistributedAvailabilityGroupOutput) LastHardenedLsn() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.LastHardenedLsn }).(pulumi.StringOutput)
}

func (o DistributedAvailabilityGroupOutput) LinkState() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.LinkState }).(pulumi.StringOutput)
}

func (o DistributedAvailabilityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DistributedAvailabilityGroupOutput) PrimaryAvailabilityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.PrimaryAvailabilityGroupName }).(pulumi.StringPtrOutput)
}

func (o DistributedAvailabilityGroupOutput) ReplicationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.ReplicationMode }).(pulumi.StringPtrOutput)
}

func (o DistributedAvailabilityGroupOutput) SecondaryAvailabilityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.SecondaryAvailabilityGroupName }).(pulumi.StringPtrOutput)
}

func (o DistributedAvailabilityGroupOutput) SourceEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.SourceEndpoint }).(pulumi.StringPtrOutput)
}

func (o DistributedAvailabilityGroupOutput) SourceReplicaId() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.SourceReplicaId }).(pulumi.StringOutput)
}

func (o DistributedAvailabilityGroupOutput) TargetDatabase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringPtrOutput { return v.TargetDatabase }).(pulumi.StringPtrOutput)
}

func (o DistributedAvailabilityGroupOutput) TargetReplicaId() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.TargetReplicaId }).(pulumi.StringOutput)
}

func (o DistributedAvailabilityGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DistributedAvailabilityGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DistributedAvailabilityGroupOutput{})
}
