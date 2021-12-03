


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReplicationPolicy struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput         `pulumi:"location"`
	Name       pulumi.StringOutput            `pulumi:"name"`
	Properties PolicyPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput            `pulumi:"type"`
}


func NewReplicationPolicy(ctx *pulumi.Context,
	name string, args *ReplicationPolicyArgs, opts ...pulumi.ResourceOption) (*ReplicationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160810:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211101:ReplicationPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationPolicy
	err := ctx.RegisterResource("azure-native:recoveryservices/v20210601:ReplicationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReplicationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationPolicyState, opts ...pulumi.ResourceOption) (*ReplicationPolicy, error) {
	var resource ReplicationPolicy
	err := ctx.ReadResource("azure-native:recoveryservices/v20210601:ReplicationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type replicationPolicyState struct {
}

type ReplicationPolicyState struct {
}

func (ReplicationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationPolicyState)(nil)).Elem()
}

type replicationPolicyArgs struct {
	PolicyName        *string                      `pulumi:"policyName"`
	Properties        *CreatePolicyInputProperties `pulumi:"properties"`
	ResourceGroupName string                       `pulumi:"resourceGroupName"`
	ResourceName      string                       `pulumi:"resourceName"`
}


type ReplicationPolicyArgs struct {
	PolicyName        pulumi.StringPtrInput
	Properties        CreatePolicyInputPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
}

func (ReplicationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationPolicyArgs)(nil)).Elem()
}

type ReplicationPolicyInput interface {
	pulumi.Input

	ToReplicationPolicyOutput() ReplicationPolicyOutput
	ToReplicationPolicyOutputWithContext(ctx context.Context) ReplicationPolicyOutput
}

func (*ReplicationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationPolicy)(nil))
}

func (i *ReplicationPolicy) ToReplicationPolicyOutput() ReplicationPolicyOutput {
	return i.ToReplicationPolicyOutputWithContext(context.Background())
}

func (i *ReplicationPolicy) ToReplicationPolicyOutputWithContext(ctx context.Context) ReplicationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationPolicyOutput)
}

type ReplicationPolicyOutput struct{ *pulumi.OutputState }

func (ReplicationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationPolicy)(nil))
}

func (o ReplicationPolicyOutput) ToReplicationPolicyOutput() ReplicationPolicyOutput {
	return o
}

func (o ReplicationPolicyOutput) ToReplicationPolicyOutputWithContext(ctx context.Context) ReplicationPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReplicationPolicyOutput{})
}
