


package v20211201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReplicationProtectionContainerMapping struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                             `pulumi:"location"`
	Name       pulumi.StringOutput                                `pulumi:"name"`
	Properties ProtectionContainerMappingPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                                `pulumi:"type"`
}


func NewReplicationProtectionContainerMapping(ctx *pulumi.Context,
	name string, args *ReplicationProtectionContainerMappingArgs, opts ...pulumi.ResourceOption) (*ReplicationProtectionContainerMapping, error) {
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
			Type: pulumi.String("azure-native:recoveryservices:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160810:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211101:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220501:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220801:ReplicationProtectionContainerMapping"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220910:ReplicationProtectionContainerMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationProtectionContainerMapping
	err := ctx.RegisterResource("azure-native:recoveryservices/v20211201:ReplicationProtectionContainerMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReplicationProtectionContainerMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationProtectionContainerMappingState, opts ...pulumi.ResourceOption) (*ReplicationProtectionContainerMapping, error) {
	var resource ReplicationProtectionContainerMapping
	err := ctx.ReadResource("azure-native:recoveryservices/v20211201:ReplicationProtectionContainerMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type replicationProtectionContainerMappingState struct {
}

type ReplicationProtectionContainerMappingState struct {
}

func (ReplicationProtectionContainerMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationProtectionContainerMappingState)(nil)).Elem()
}

type replicationProtectionContainerMappingArgs struct {
	FabricName              string                                           `pulumi:"fabricName"`
	MappingName             *string                                          `pulumi:"mappingName"`
	Properties              *CreateProtectionContainerMappingInputProperties `pulumi:"properties"`
	ProtectionContainerName string                                           `pulumi:"protectionContainerName"`
	ResourceGroupName       string                                           `pulumi:"resourceGroupName"`
	ResourceName            string                                           `pulumi:"resourceName"`
}


type ReplicationProtectionContainerMappingArgs struct {
	FabricName              pulumi.StringInput
	MappingName             pulumi.StringPtrInput
	Properties              CreateProtectionContainerMappingInputPropertiesPtrInput
	ProtectionContainerName pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
	ResourceName            pulumi.StringInput
}

func (ReplicationProtectionContainerMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationProtectionContainerMappingArgs)(nil)).Elem()
}

type ReplicationProtectionContainerMappingInput interface {
	pulumi.Input

	ToReplicationProtectionContainerMappingOutput() ReplicationProtectionContainerMappingOutput
	ToReplicationProtectionContainerMappingOutputWithContext(ctx context.Context) ReplicationProtectionContainerMappingOutput
}

func (*ReplicationProtectionContainerMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationProtectionContainerMapping)(nil)).Elem()
}

func (i *ReplicationProtectionContainerMapping) ToReplicationProtectionContainerMappingOutput() ReplicationProtectionContainerMappingOutput {
	return i.ToReplicationProtectionContainerMappingOutputWithContext(context.Background())
}

func (i *ReplicationProtectionContainerMapping) ToReplicationProtectionContainerMappingOutputWithContext(ctx context.Context) ReplicationProtectionContainerMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationProtectionContainerMappingOutput)
}

type ReplicationProtectionContainerMappingOutput struct{ *pulumi.OutputState }

func (ReplicationProtectionContainerMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationProtectionContainerMapping)(nil)).Elem()
}

func (o ReplicationProtectionContainerMappingOutput) ToReplicationProtectionContainerMappingOutput() ReplicationProtectionContainerMappingOutput {
	return o
}

func (o ReplicationProtectionContainerMappingOutput) ToReplicationProtectionContainerMappingOutputWithContext(ctx context.Context) ReplicationProtectionContainerMappingOutput {
	return o
}

func (o ReplicationProtectionContainerMappingOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectionContainerMapping) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectionContainerMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationProtectionContainerMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReplicationProtectionContainerMappingOutput) Properties() ProtectionContainerMappingPropertiesResponseOutput {
	return o.ApplyT(func(v *ReplicationProtectionContainerMapping) ProtectionContainerMappingPropertiesResponseOutput {
		return v.Properties
	}).(ProtectionContainerMappingPropertiesResponseOutput)
}

func (o ReplicationProtectionContainerMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationProtectionContainerMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationProtectionContainerMappingOutput{})
}
