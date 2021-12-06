


package recoveryservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReplicationRecoveryServicesProvider struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                           `pulumi:"location"`
	Name       pulumi.StringOutput                              `pulumi:"name"`
	Properties RecoveryServicesProviderPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                              `pulumi:"type"`
}


func NewReplicationRecoveryServicesProvider(ctx *pulumi.Context,
	name string, args *ReplicationRecoveryServicesProviderArgs, opts ...pulumi.ResourceOption) (*ReplicationRecoveryServicesProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricName == nil {
		return nil, errors.New("invalid value for required argument 'FabricName'")
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
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationRecoveryServicesProvider"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationRecoveryServicesProvider
	err := ctx.RegisterResource("azure-native:recoveryservices:ReplicationRecoveryServicesProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReplicationRecoveryServicesProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationRecoveryServicesProviderState, opts ...pulumi.ResourceOption) (*ReplicationRecoveryServicesProvider, error) {
	var resource ReplicationRecoveryServicesProvider
	err := ctx.ReadResource("azure-native:recoveryservices:ReplicationRecoveryServicesProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type replicationRecoveryServicesProviderState struct {
}

type ReplicationRecoveryServicesProviderState struct {
}

func (ReplicationRecoveryServicesProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationRecoveryServicesProviderState)(nil)).Elem()
}

type replicationRecoveryServicesProviderArgs struct {
	FabricName        string                                     `pulumi:"fabricName"`
	Properties        AddRecoveryServicesProviderInputProperties `pulumi:"properties"`
	ProviderName      *string                                    `pulumi:"providerName"`
	ResourceGroupName string                                     `pulumi:"resourceGroupName"`
	ResourceName      string                                     `pulumi:"resourceName"`
}


type ReplicationRecoveryServicesProviderArgs struct {
	FabricName        pulumi.StringInput
	Properties        AddRecoveryServicesProviderInputPropertiesInput
	ProviderName      pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
}

func (ReplicationRecoveryServicesProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationRecoveryServicesProviderArgs)(nil)).Elem()
}

type ReplicationRecoveryServicesProviderInput interface {
	pulumi.Input

	ToReplicationRecoveryServicesProviderOutput() ReplicationRecoveryServicesProviderOutput
	ToReplicationRecoveryServicesProviderOutputWithContext(ctx context.Context) ReplicationRecoveryServicesProviderOutput
}

func (*ReplicationRecoveryServicesProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationRecoveryServicesProvider)(nil))
}

func (i *ReplicationRecoveryServicesProvider) ToReplicationRecoveryServicesProviderOutput() ReplicationRecoveryServicesProviderOutput {
	return i.ToReplicationRecoveryServicesProviderOutputWithContext(context.Background())
}

func (i *ReplicationRecoveryServicesProvider) ToReplicationRecoveryServicesProviderOutputWithContext(ctx context.Context) ReplicationRecoveryServicesProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationRecoveryServicesProviderOutput)
}

type ReplicationRecoveryServicesProviderOutput struct{ *pulumi.OutputState }

func (ReplicationRecoveryServicesProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationRecoveryServicesProvider)(nil))
}

func (o ReplicationRecoveryServicesProviderOutput) ToReplicationRecoveryServicesProviderOutput() ReplicationRecoveryServicesProviderOutput {
	return o
}

func (o ReplicationRecoveryServicesProviderOutput) ToReplicationRecoveryServicesProviderOutputWithContext(ctx context.Context) ReplicationRecoveryServicesProviderOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReplicationRecoveryServicesProviderOutput{})
}
