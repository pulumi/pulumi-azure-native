


package v20180110

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ReplicationNetworkMapping struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                 `pulumi:"location"`
	Name       pulumi.StringOutput                    `pulumi:"name"`
	Properties NetworkMappingPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                    `pulumi:"type"`
}


func NewReplicationNetworkMapping(ctx *pulumi.Context,
	name string, args *ReplicationNetworkMappingArgs, opts ...pulumi.ResourceOption) (*ReplicationNetworkMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricName == nil {
		return nil, errors.New("invalid value for required argument 'FabricName'")
	}
	if args.NetworkName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160810:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211101:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220501:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220801:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220910:ReplicationNetworkMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ReplicationNetworkMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationNetworkMapping
	err := ctx.RegisterResource("azure-native:recoveryservices/v20180110:ReplicationNetworkMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReplicationNetworkMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationNetworkMappingState, opts ...pulumi.ResourceOption) (*ReplicationNetworkMapping, error) {
	var resource ReplicationNetworkMapping
	err := ctx.ReadResource("azure-native:recoveryservices/v20180110:ReplicationNetworkMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type replicationNetworkMappingState struct {
}

type ReplicationNetworkMappingState struct {
}

func (ReplicationNetworkMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationNetworkMappingState)(nil)).Elem()
}

type replicationNetworkMappingArgs struct {
	FabricName         string                               `pulumi:"fabricName"`
	NetworkMappingName *string                              `pulumi:"networkMappingName"`
	NetworkName        string                               `pulumi:"networkName"`
	Properties         *CreateNetworkMappingInputProperties `pulumi:"properties"`
	ResourceGroupName  string                               `pulumi:"resourceGroupName"`
	ResourceName       string                               `pulumi:"resourceName"`
}


type ReplicationNetworkMappingArgs struct {
	FabricName         pulumi.StringInput
	NetworkMappingName pulumi.StringPtrInput
	NetworkName        pulumi.StringInput
	Properties         CreateNetworkMappingInputPropertiesPtrInput
	ResourceGroupName  pulumi.StringInput
	ResourceName       pulumi.StringInput
}

func (ReplicationNetworkMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationNetworkMappingArgs)(nil)).Elem()
}

type ReplicationNetworkMappingInput interface {
	pulumi.Input

	ToReplicationNetworkMappingOutput() ReplicationNetworkMappingOutput
	ToReplicationNetworkMappingOutputWithContext(ctx context.Context) ReplicationNetworkMappingOutput
}

func (*ReplicationNetworkMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationNetworkMapping)(nil)).Elem()
}

func (i *ReplicationNetworkMapping) ToReplicationNetworkMappingOutput() ReplicationNetworkMappingOutput {
	return i.ToReplicationNetworkMappingOutputWithContext(context.Background())
}

func (i *ReplicationNetworkMapping) ToReplicationNetworkMappingOutputWithContext(ctx context.Context) ReplicationNetworkMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationNetworkMappingOutput)
}

type ReplicationNetworkMappingOutput struct{ *pulumi.OutputState }

func (ReplicationNetworkMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationNetworkMapping)(nil)).Elem()
}

func (o ReplicationNetworkMappingOutput) ToReplicationNetworkMappingOutput() ReplicationNetworkMappingOutput {
	return o
}

func (o ReplicationNetworkMappingOutput) ToReplicationNetworkMappingOutputWithContext(ctx context.Context) ReplicationNetworkMappingOutput {
	return o
}

func (o ReplicationNetworkMappingOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationNetworkMapping) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ReplicationNetworkMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationNetworkMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReplicationNetworkMappingOutput) Properties() NetworkMappingPropertiesResponseOutput {
	return o.ApplyT(func(v *ReplicationNetworkMapping) NetworkMappingPropertiesResponseOutput { return v.Properties }).(NetworkMappingPropertiesResponseOutput)
}

func (o ReplicationNetworkMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationNetworkMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationNetworkMappingOutput{})
}
