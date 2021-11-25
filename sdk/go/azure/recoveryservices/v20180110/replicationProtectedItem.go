


package v20180110

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReplicationProtectedItem struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                           `pulumi:"location"`
	Name       pulumi.StringOutput                              `pulumi:"name"`
	Properties ReplicationProtectedItemPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                              `pulumi:"type"`
}


func NewReplicationProtectedItem(ctx *pulumi.Context,
	name string, args *ReplicationProtectedItemArgs, opts ...pulumi.ResourceOption) (*ReplicationProtectedItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricName == nil {
		return nil, errors.New("invalid value for required argument 'FabricName'")
	}
	if args.ProtectionContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ProtectionContainerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160810:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationProtectedItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationProtectedItem"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationProtectedItem
	err := ctx.RegisterResource("azure-native:recoveryservices/v20180110:ReplicationProtectedItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReplicationProtectedItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationProtectedItemState, opts ...pulumi.ResourceOption) (*ReplicationProtectedItem, error) {
	var resource ReplicationProtectedItem
	err := ctx.ReadResource("azure-native:recoveryservices/v20180110:ReplicationProtectedItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type replicationProtectedItemState struct {
}

type ReplicationProtectedItemState struct {
}

func (ReplicationProtectedItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationProtectedItemState)(nil)).Elem()
}

type replicationProtectedItemArgs struct {
	FabricName                  string                           `pulumi:"fabricName"`
	Properties                  *EnableProtectionInputProperties `pulumi:"properties"`
	ProtectionContainerName     string                           `pulumi:"protectionContainerName"`
	ReplicatedProtectedItemName *string                          `pulumi:"replicatedProtectedItemName"`
	ResourceGroupName           string                           `pulumi:"resourceGroupName"`
	ResourceName                string                           `pulumi:"resourceName"`
}


type ReplicationProtectedItemArgs struct {
	FabricName                  pulumi.StringInput
	Properties                  EnableProtectionInputPropertiesPtrInput
	ProtectionContainerName     pulumi.StringInput
	ReplicatedProtectedItemName pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	ResourceName                pulumi.StringInput
}

func (ReplicationProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationProtectedItemArgs)(nil)).Elem()
}

type ReplicationProtectedItemInput interface {
	pulumi.Input

	ToReplicationProtectedItemOutput() ReplicationProtectedItemOutput
	ToReplicationProtectedItemOutputWithContext(ctx context.Context) ReplicationProtectedItemOutput
}

func (*ReplicationProtectedItem) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationProtectedItem)(nil))
}

func (i *ReplicationProtectedItem) ToReplicationProtectedItemOutput() ReplicationProtectedItemOutput {
	return i.ToReplicationProtectedItemOutputWithContext(context.Background())
}

func (i *ReplicationProtectedItem) ToReplicationProtectedItemOutputWithContext(ctx context.Context) ReplicationProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationProtectedItemOutput)
}

type ReplicationProtectedItemOutput struct{ *pulumi.OutputState }

func (ReplicationProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationProtectedItem)(nil))
}

func (o ReplicationProtectedItemOutput) ToReplicationProtectedItemOutput() ReplicationProtectedItemOutput {
	return o
}

func (o ReplicationProtectedItemOutput) ToReplicationProtectedItemOutputWithContext(ctx context.Context) ReplicationProtectedItemOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReplicationProtectedItemOutput{})
}
