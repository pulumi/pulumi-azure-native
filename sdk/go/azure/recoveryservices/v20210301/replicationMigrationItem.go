


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReplicationMigrationItem struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                `pulumi:"location"`
	Name       pulumi.StringOutput                   `pulumi:"name"`
	Properties MigrationItemPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                   `pulumi:"type"`
}


func NewReplicationMigrationItem(ctx *pulumi.Context,
	name string, args *ReplicationMigrationItemArgs, opts ...pulumi.ResourceOption) (*ReplicationMigrationItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricName == nil {
		return nil, errors.New("invalid value for required argument 'FabricName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
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
			Type: pulumi.String("azure-native:recoveryservices:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationMigrationItem"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationMigrationItem"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationMigrationItem
	err := ctx.RegisterResource("azure-native:recoveryservices/v20210301:ReplicationMigrationItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReplicationMigrationItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationMigrationItemState, opts ...pulumi.ResourceOption) (*ReplicationMigrationItem, error) {
	var resource ReplicationMigrationItem
	err := ctx.ReadResource("azure-native:recoveryservices/v20210301:ReplicationMigrationItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type replicationMigrationItemState struct {
}

type ReplicationMigrationItemState struct {
}

func (ReplicationMigrationItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationMigrationItemState)(nil)).Elem()
}

type replicationMigrationItemArgs struct {
	FabricName              string                         `pulumi:"fabricName"`
	MigrationItemName       *string                        `pulumi:"migrationItemName"`
	Properties              EnableMigrationInputProperties `pulumi:"properties"`
	ProtectionContainerName string                         `pulumi:"protectionContainerName"`
	ResourceGroupName       string                         `pulumi:"resourceGroupName"`
	ResourceName            string                         `pulumi:"resourceName"`
}


type ReplicationMigrationItemArgs struct {
	FabricName              pulumi.StringInput
	MigrationItemName       pulumi.StringPtrInput
	Properties              EnableMigrationInputPropertiesInput
	ProtectionContainerName pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
	ResourceName            pulumi.StringInput
}

func (ReplicationMigrationItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationMigrationItemArgs)(nil)).Elem()
}

type ReplicationMigrationItemInput interface {
	pulumi.Input

	ToReplicationMigrationItemOutput() ReplicationMigrationItemOutput
	ToReplicationMigrationItemOutputWithContext(ctx context.Context) ReplicationMigrationItemOutput
}

func (*ReplicationMigrationItem) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationMigrationItem)(nil)).Elem()
}

func (i *ReplicationMigrationItem) ToReplicationMigrationItemOutput() ReplicationMigrationItemOutput {
	return i.ToReplicationMigrationItemOutputWithContext(context.Background())
}

func (i *ReplicationMigrationItem) ToReplicationMigrationItemOutputWithContext(ctx context.Context) ReplicationMigrationItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationMigrationItemOutput)
}

type ReplicationMigrationItemOutput struct{ *pulumi.OutputState }

func (ReplicationMigrationItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationMigrationItem)(nil)).Elem()
}

func (o ReplicationMigrationItemOutput) ToReplicationMigrationItemOutput() ReplicationMigrationItemOutput {
	return o
}

func (o ReplicationMigrationItemOutput) ToReplicationMigrationItemOutputWithContext(ctx context.Context) ReplicationMigrationItemOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReplicationMigrationItemOutput{})
}
