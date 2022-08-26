


package v20180710

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReplicationFabric struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput         `pulumi:"location"`
	Name       pulumi.StringOutput            `pulumi:"name"`
	Properties FabricPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput            `pulumi:"type"`
}


func NewReplicationFabric(ctx *pulumi.Context,
	name string, args *ReplicationFabricArgs, opts ...pulumi.ResourceOption) (*ReplicationFabric, error) {
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
			Type: pulumi.String("azure-native:recoveryservices:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160810:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211101:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220501:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220801:ReplicationFabric"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationFabric
	err := ctx.RegisterResource("azure-native:recoveryservices/v20180710:ReplicationFabric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReplicationFabric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationFabricState, opts ...pulumi.ResourceOption) (*ReplicationFabric, error) {
	var resource ReplicationFabric
	err := ctx.ReadResource("azure-native:recoveryservices/v20180710:ReplicationFabric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type replicationFabricState struct {
}

type ReplicationFabricState struct {
}

func (ReplicationFabricState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationFabricState)(nil)).Elem()
}

type replicationFabricArgs struct {
	FabricName        *string                        `pulumi:"fabricName"`
	Properties        *FabricCreationInputProperties `pulumi:"properties"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	ResourceName      string                         `pulumi:"resourceName"`
}


type ReplicationFabricArgs struct {
	FabricName        pulumi.StringPtrInput
	Properties        FabricCreationInputPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
}

func (ReplicationFabricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationFabricArgs)(nil)).Elem()
}

type ReplicationFabricInput interface {
	pulumi.Input

	ToReplicationFabricOutput() ReplicationFabricOutput
	ToReplicationFabricOutputWithContext(ctx context.Context) ReplicationFabricOutput
}

func (*ReplicationFabric) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationFabric)(nil)).Elem()
}

func (i *ReplicationFabric) ToReplicationFabricOutput() ReplicationFabricOutput {
	return i.ToReplicationFabricOutputWithContext(context.Background())
}

func (i *ReplicationFabric) ToReplicationFabricOutputWithContext(ctx context.Context) ReplicationFabricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationFabricOutput)
}

type ReplicationFabricOutput struct{ *pulumi.OutputState }

func (ReplicationFabricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationFabric)(nil)).Elem()
}

func (o ReplicationFabricOutput) ToReplicationFabricOutput() ReplicationFabricOutput {
	return o
}

func (o ReplicationFabricOutput) ToReplicationFabricOutputWithContext(ctx context.Context) ReplicationFabricOutput {
	return o
}

func (o ReplicationFabricOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationFabric) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ReplicationFabricOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationFabric) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReplicationFabricOutput) Properties() FabricPropertiesResponseOutput {
	return o.ApplyT(func(v *ReplicationFabric) FabricPropertiesResponseOutput { return v.Properties }).(FabricPropertiesResponseOutput)
}

func (o ReplicationFabricOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationFabric) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationFabricOutput{})
}
