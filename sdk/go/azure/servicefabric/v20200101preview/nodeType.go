


package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NodeType struct {
	pulumi.CustomResourceState

	ApplicationPorts    EndpointRangeDescriptionResponsePtrOutput `pulumi:"applicationPorts"`
	Capacities          pulumi.StringMapOutput                    `pulumi:"capacities"`
	DataDiskSizeGB      pulumi.IntOutput                          `pulumi:"dataDiskSizeGB"`
	EphemeralPorts      EndpointRangeDescriptionResponsePtrOutput `pulumi:"ephemeralPorts"`
	IsPrimary           pulumi.BoolOutput                         `pulumi:"isPrimary"`
	Name                pulumi.StringOutput                       `pulumi:"name"`
	PlacementProperties pulumi.StringMapOutput                    `pulumi:"placementProperties"`
	ProvisioningState   pulumi.StringOutput                       `pulumi:"provisioningState"`
	Tags                pulumi.StringMapOutput                    `pulumi:"tags"`
	Type                pulumi.StringOutput                       `pulumi:"type"`
	VmExtensions        VMSSExtensionResponseArrayOutput          `pulumi:"vmExtensions"`
	VmImageOffer        pulumi.StringPtrOutput                    `pulumi:"vmImageOffer"`
	VmImagePublisher    pulumi.StringPtrOutput                    `pulumi:"vmImagePublisher"`
	VmImageSku          pulumi.StringPtrOutput                    `pulumi:"vmImageSku"`
	VmImageVersion      pulumi.StringPtrOutput                    `pulumi:"vmImageVersion"`
	VmInstanceCount     pulumi.IntOutput                          `pulumi:"vmInstanceCount"`
	VmSecrets           VaultSecretGroupResponseArrayOutput       `pulumi:"vmSecrets"`
	VmSize              pulumi.StringPtrOutput                    `pulumi:"vmSize"`
}


func NewNodeType(ctx *pulumi.Context,
	name string, args *NodeTypeArgs, opts ...pulumi.ResourceOption) (*NodeType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.DataDiskSizeGB == nil {
		return nil, errors.New("invalid value for required argument 'DataDiskSizeGB'")
	}
	if args.IsPrimary == nil {
		return nil, errors.New("invalid value for required argument 'IsPrimary'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmInstanceCount == nil {
		return nil, errors.New("invalid value for required argument 'VmInstanceCount'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210101preview:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210501:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210701preview:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210901privatepreview:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20211101preview:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220101:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220201preview:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220601preview:NodeType"),
		},
	})
	opts = append(opts, aliases)
	var resource NodeType
	err := ctx.RegisterResource("azure-native:servicefabric/v20200101preview:NodeType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNodeType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeTypeState, opts ...pulumi.ResourceOption) (*NodeType, error) {
	var resource NodeType
	err := ctx.ReadResource("azure-native:servicefabric/v20200101preview:NodeType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type nodeTypeState struct {
}

type NodeTypeState struct {
}

func (NodeTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeTypeState)(nil)).Elem()
}

type nodeTypeArgs struct {
	ApplicationPorts    *EndpointRangeDescription `pulumi:"applicationPorts"`
	Capacities          map[string]string         `pulumi:"capacities"`
	ClusterName         string                    `pulumi:"clusterName"`
	DataDiskSizeGB      int                       `pulumi:"dataDiskSizeGB"`
	EphemeralPorts      *EndpointRangeDescription `pulumi:"ephemeralPorts"`
	IsPrimary           bool                      `pulumi:"isPrimary"`
	NodeTypeName        *string                   `pulumi:"nodeTypeName"`
	PlacementProperties map[string]string         `pulumi:"placementProperties"`
	ResourceGroupName   string                    `pulumi:"resourceGroupName"`
	Tags                map[string]string         `pulumi:"tags"`
	VmExtensions        []VMSSExtension           `pulumi:"vmExtensions"`
	VmImageOffer        *string                   `pulumi:"vmImageOffer"`
	VmImagePublisher    *string                   `pulumi:"vmImagePublisher"`
	VmImageSku          *string                   `pulumi:"vmImageSku"`
	VmImageVersion      *string                   `pulumi:"vmImageVersion"`
	VmInstanceCount     int                       `pulumi:"vmInstanceCount"`
	VmSecrets           []VaultSecretGroup        `pulumi:"vmSecrets"`
	VmSize              *string                   `pulumi:"vmSize"`
}


type NodeTypeArgs struct {
	ApplicationPorts    EndpointRangeDescriptionPtrInput
	Capacities          pulumi.StringMapInput
	ClusterName         pulumi.StringInput
	DataDiskSizeGB      pulumi.IntInput
	EphemeralPorts      EndpointRangeDescriptionPtrInput
	IsPrimary           pulumi.BoolInput
	NodeTypeName        pulumi.StringPtrInput
	PlacementProperties pulumi.StringMapInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
	VmExtensions        VMSSExtensionArrayInput
	VmImageOffer        pulumi.StringPtrInput
	VmImagePublisher    pulumi.StringPtrInput
	VmImageSku          pulumi.StringPtrInput
	VmImageVersion      pulumi.StringPtrInput
	VmInstanceCount     pulumi.IntInput
	VmSecrets           VaultSecretGroupArrayInput
	VmSize              pulumi.StringPtrInput
}

func (NodeTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeTypeArgs)(nil)).Elem()
}

type NodeTypeInput interface {
	pulumi.Input

	ToNodeTypeOutput() NodeTypeOutput
	ToNodeTypeOutputWithContext(ctx context.Context) NodeTypeOutput
}

func (*NodeType) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeType)(nil)).Elem()
}

func (i *NodeType) ToNodeTypeOutput() NodeTypeOutput {
	return i.ToNodeTypeOutputWithContext(context.Background())
}

func (i *NodeType) ToNodeTypeOutputWithContext(ctx context.Context) NodeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTypeOutput)
}

type NodeTypeOutput struct{ *pulumi.OutputState }

func (NodeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeType)(nil)).Elem()
}

func (o NodeTypeOutput) ToNodeTypeOutput() NodeTypeOutput {
	return o
}

func (o NodeTypeOutput) ToNodeTypeOutputWithContext(ctx context.Context) NodeTypeOutput {
	return o
}

func (o NodeTypeOutput) ApplicationPorts() EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *NodeType) EndpointRangeDescriptionResponsePtrOutput { return v.ApplicationPorts }).(EndpointRangeDescriptionResponsePtrOutput)
}

func (o NodeTypeOutput) Capacities() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringMapOutput { return v.Capacities }).(pulumi.StringMapOutput)
}

func (o NodeTypeOutput) DataDiskSizeGB() pulumi.IntOutput {
	return o.ApplyT(func(v *NodeType) pulumi.IntOutput { return v.DataDiskSizeGB }).(pulumi.IntOutput)
}

func (o NodeTypeOutput) EphemeralPorts() EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *NodeType) EndpointRangeDescriptionResponsePtrOutput { return v.EphemeralPorts }).(EndpointRangeDescriptionResponsePtrOutput)
}

func (o NodeTypeOutput) IsPrimary() pulumi.BoolOutput {
	return o.ApplyT(func(v *NodeType) pulumi.BoolOutput { return v.IsPrimary }).(pulumi.BoolOutput)
}

func (o NodeTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NodeTypeOutput) PlacementProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringMapOutput { return v.PlacementProperties }).(pulumi.StringMapOutput)
}

func (o NodeTypeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NodeTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NodeTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o NodeTypeOutput) VmExtensions() VMSSExtensionResponseArrayOutput {
	return o.ApplyT(func(v *NodeType) VMSSExtensionResponseArrayOutput { return v.VmExtensions }).(VMSSExtensionResponseArrayOutput)
}

func (o NodeTypeOutput) VmImageOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringPtrOutput { return v.VmImageOffer }).(pulumi.StringPtrOutput)
}

func (o NodeTypeOutput) VmImagePublisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringPtrOutput { return v.VmImagePublisher }).(pulumi.StringPtrOutput)
}

func (o NodeTypeOutput) VmImageSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringPtrOutput { return v.VmImageSku }).(pulumi.StringPtrOutput)
}

func (o NodeTypeOutput) VmImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringPtrOutput { return v.VmImageVersion }).(pulumi.StringPtrOutput)
}

func (o NodeTypeOutput) VmInstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v *NodeType) pulumi.IntOutput { return v.VmInstanceCount }).(pulumi.IntOutput)
}

func (o NodeTypeOutput) VmSecrets() VaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v *NodeType) VaultSecretGroupResponseArrayOutput { return v.VmSecrets }).(VaultSecretGroupResponseArrayOutput)
}

func (o NodeTypeOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringPtrOutput { return v.VmSize }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NodeTypeOutput{})
}
