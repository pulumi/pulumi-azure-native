


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
	return reflect.TypeOf((*DistributedAvailabilityGroup)(nil))
}

func (i *DistributedAvailabilityGroup) ToDistributedAvailabilityGroupOutput() DistributedAvailabilityGroupOutput {
	return i.ToDistributedAvailabilityGroupOutputWithContext(context.Background())
}

func (i *DistributedAvailabilityGroup) ToDistributedAvailabilityGroupOutputWithContext(ctx context.Context) DistributedAvailabilityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DistributedAvailabilityGroupOutput)
}

type DistributedAvailabilityGroupOutput struct{ *pulumi.OutputState }

func (DistributedAvailabilityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DistributedAvailabilityGroup)(nil))
}

func (o DistributedAvailabilityGroupOutput) ToDistributedAvailabilityGroupOutput() DistributedAvailabilityGroupOutput {
	return o
}

func (o DistributedAvailabilityGroupOutput) ToDistributedAvailabilityGroupOutputWithContext(ctx context.Context) DistributedAvailabilityGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DistributedAvailabilityGroupOutput{})
}
