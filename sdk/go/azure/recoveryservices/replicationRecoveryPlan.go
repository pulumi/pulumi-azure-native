


package recoveryservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReplicationRecoveryPlan struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput               `pulumi:"location"`
	Name       pulumi.StringOutput                  `pulumi:"name"`
	Properties RecoveryPlanPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                  `pulumi:"type"`
}


func NewReplicationRecoveryPlan(ctx *pulumi.Context,
	name string, args *ReplicationRecoveryPlanArgs, opts ...pulumi.ResourceOption) (*ReplicationRecoveryPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160810:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationRecoveryPlan"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationRecoveryPlan"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationRecoveryPlan
	err := ctx.RegisterResource("azure-native:recoveryservices:ReplicationRecoveryPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReplicationRecoveryPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationRecoveryPlanState, opts ...pulumi.ResourceOption) (*ReplicationRecoveryPlan, error) {
	var resource ReplicationRecoveryPlan
	err := ctx.ReadResource("azure-native:recoveryservices:ReplicationRecoveryPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type replicationRecoveryPlanState struct {
}

type ReplicationRecoveryPlanState struct {
}

func (ReplicationRecoveryPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationRecoveryPlanState)(nil)).Elem()
}

type replicationRecoveryPlanArgs struct {
	Properties        CreateRecoveryPlanInputProperties `pulumi:"properties"`
	RecoveryPlanName  *string                           `pulumi:"recoveryPlanName"`
	ResourceGroupName string                            `pulumi:"resourceGroupName"`
	ResourceName      string                            `pulumi:"resourceName"`
}


type ReplicationRecoveryPlanArgs struct {
	Properties        CreateRecoveryPlanInputPropertiesInput
	RecoveryPlanName  pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
}

func (ReplicationRecoveryPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationRecoveryPlanArgs)(nil)).Elem()
}

type ReplicationRecoveryPlanInput interface {
	pulumi.Input

	ToReplicationRecoveryPlanOutput() ReplicationRecoveryPlanOutput
	ToReplicationRecoveryPlanOutputWithContext(ctx context.Context) ReplicationRecoveryPlanOutput
}

func (*ReplicationRecoveryPlan) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationRecoveryPlan)(nil))
}

func (i *ReplicationRecoveryPlan) ToReplicationRecoveryPlanOutput() ReplicationRecoveryPlanOutput {
	return i.ToReplicationRecoveryPlanOutputWithContext(context.Background())
}

func (i *ReplicationRecoveryPlan) ToReplicationRecoveryPlanOutputWithContext(ctx context.Context) ReplicationRecoveryPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationRecoveryPlanOutput)
}

type ReplicationRecoveryPlanOutput struct{ *pulumi.OutputState }

func (ReplicationRecoveryPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationRecoveryPlan)(nil))
}

func (o ReplicationRecoveryPlanOutput) ToReplicationRecoveryPlanOutput() ReplicationRecoveryPlanOutput {
	return o
}

func (o ReplicationRecoveryPlanOutput) ToReplicationRecoveryPlanOutputWithContext(ctx context.Context) ReplicationRecoveryPlanOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReplicationRecoveryPlanOutput{})
}
