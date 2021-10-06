


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NodeType struct {
	pulumi.CustomResourceState

	ApplicationPorts        EndpointRangeDescriptionResponsePtrOutput `pulumi:"applicationPorts"`
	Capacities              pulumi.StringMapOutput                    `pulumi:"capacities"`
	DataDiskSizeGB          pulumi.IntOutput                          `pulumi:"dataDiskSizeGB"`
	DataDiskType            pulumi.StringPtrOutput                    `pulumi:"dataDiskType"`
	EphemeralPorts          EndpointRangeDescriptionResponsePtrOutput `pulumi:"ephemeralPorts"`
	FrontendConfigurations  FrontendConfigurationResponseArrayOutput  `pulumi:"frontendConfigurations"`
	IsPrimary               pulumi.BoolOutput                         `pulumi:"isPrimary"`
	IsStateless             pulumi.BoolPtrOutput                      `pulumi:"isStateless"`
	MultiplePlacementGroups pulumi.BoolPtrOutput                      `pulumi:"multiplePlacementGroups"`
	Name                    pulumi.StringOutput                       `pulumi:"name"`
	NetworkSecurityRules    NetworkSecurityRuleResponseArrayOutput    `pulumi:"networkSecurityRules"`
	PlacementProperties     pulumi.StringMapOutput                    `pulumi:"placementProperties"`
	ProvisioningState       pulumi.StringOutput                       `pulumi:"provisioningState"`
	Sku                     NodeTypeSkuResponsePtrOutput              `pulumi:"sku"`
	SystemData              SystemDataResponseOutput                  `pulumi:"systemData"`
	Tags                    pulumi.StringMapOutput                    `pulumi:"tags"`
	Type                    pulumi.StringOutput                       `pulumi:"type"`
	VmExtensions            VMSSExtensionResponseArrayOutput          `pulumi:"vmExtensions"`
	VmImageOffer            pulumi.StringPtrOutput                    `pulumi:"vmImageOffer"`
	VmImagePublisher        pulumi.StringPtrOutput                    `pulumi:"vmImagePublisher"`
	VmImageSku              pulumi.StringPtrOutput                    `pulumi:"vmImageSku"`
	VmImageVersion          pulumi.StringPtrOutput                    `pulumi:"vmImageVersion"`
	VmInstanceCount         pulumi.IntOutput                          `pulumi:"vmInstanceCount"`
	VmManagedIdentity       VmManagedIdentityResponsePtrOutput        `pulumi:"vmManagedIdentity"`
	VmSecrets               VaultSecretGroupResponseArrayOutput       `pulumi:"vmSecrets"`
	VmSize                  pulumi.StringPtrOutput                    `pulumi:"vmSize"`
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
	if args.IsStateless == nil {
		args.IsStateless = pulumi.BoolPtr(false)
	}
	if args.MultiplePlacementGroups == nil {
		args.MultiplePlacementGroups = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:servicefabric/v20210701preview:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric:NodeType"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicefabric:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20200101preview:NodeType"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicefabric/v20200101preview:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210101preview:NodeType"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicefabric/v20210101preview:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210501:NodeType"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicefabric/v20210501:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210901privatepreview:NodeType"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicefabric/v20210901privatepreview:NodeType"),
		},
	})
	opts = append(opts, aliases)
	var resource NodeType
	err := ctx.RegisterResource("azure-native:servicefabric/v20210701preview:NodeType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNodeType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeTypeState, opts ...pulumi.ResourceOption) (*NodeType, error) {
	var resource NodeType
	err := ctx.ReadResource("azure-native:servicefabric/v20210701preview:NodeType", name, id, state, &resource, opts...)
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
	ApplicationPorts        *EndpointRangeDescription `pulumi:"applicationPorts"`
	Capacities              map[string]string         `pulumi:"capacities"`
	ClusterName             string                    `pulumi:"clusterName"`
	DataDiskSizeGB          int                       `pulumi:"dataDiskSizeGB"`
	DataDiskType            *string                   `pulumi:"dataDiskType"`
	EphemeralPorts          *EndpointRangeDescription `pulumi:"ephemeralPorts"`
	FrontendConfigurations  []FrontendConfiguration   `pulumi:"frontendConfigurations"`
	IsPrimary               bool                      `pulumi:"isPrimary"`
	IsStateless             *bool                     `pulumi:"isStateless"`
	MultiplePlacementGroups *bool                     `pulumi:"multiplePlacementGroups"`
	NetworkSecurityRules    []NetworkSecurityRule     `pulumi:"networkSecurityRules"`
	NodeTypeName            *string                   `pulumi:"nodeTypeName"`
	PlacementProperties     map[string]string         `pulumi:"placementProperties"`
	ResourceGroupName       string                    `pulumi:"resourceGroupName"`
	Sku                     *NodeTypeSku              `pulumi:"sku"`
	Tags                    map[string]string         `pulumi:"tags"`
	VmExtensions            []VMSSExtension           `pulumi:"vmExtensions"`
	VmImageOffer            *string                   `pulumi:"vmImageOffer"`
	VmImagePublisher        *string                   `pulumi:"vmImagePublisher"`
	VmImageSku              *string                   `pulumi:"vmImageSku"`
	VmImageVersion          *string                   `pulumi:"vmImageVersion"`
	VmInstanceCount         int                       `pulumi:"vmInstanceCount"`
	VmManagedIdentity       *VmManagedIdentity        `pulumi:"vmManagedIdentity"`
	VmSecrets               []VaultSecretGroup        `pulumi:"vmSecrets"`
	VmSize                  *string                   `pulumi:"vmSize"`
}


type NodeTypeArgs struct {
	ApplicationPorts        EndpointRangeDescriptionPtrInput
	Capacities              pulumi.StringMapInput
	ClusterName             pulumi.StringInput
	DataDiskSizeGB          pulumi.IntInput
	DataDiskType            pulumi.StringPtrInput
	EphemeralPorts          EndpointRangeDescriptionPtrInput
	FrontendConfigurations  FrontendConfigurationArrayInput
	IsPrimary               pulumi.BoolInput
	IsStateless             pulumi.BoolPtrInput
	MultiplePlacementGroups pulumi.BoolPtrInput
	NetworkSecurityRules    NetworkSecurityRuleArrayInput
	NodeTypeName            pulumi.StringPtrInput
	PlacementProperties     pulumi.StringMapInput
	ResourceGroupName       pulumi.StringInput
	Sku                     NodeTypeSkuPtrInput
	Tags                    pulumi.StringMapInput
	VmExtensions            VMSSExtensionArrayInput
	VmImageOffer            pulumi.StringPtrInput
	VmImagePublisher        pulumi.StringPtrInput
	VmImageSku              pulumi.StringPtrInput
	VmImageVersion          pulumi.StringPtrInput
	VmInstanceCount         pulumi.IntInput
	VmManagedIdentity       VmManagedIdentityPtrInput
	VmSecrets               VaultSecretGroupArrayInput
	VmSize                  pulumi.StringPtrInput
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
	return reflect.TypeOf((*NodeType)(nil))
}

func (i *NodeType) ToNodeTypeOutput() NodeTypeOutput {
	return i.ToNodeTypeOutputWithContext(context.Background())
}

func (i *NodeType) ToNodeTypeOutputWithContext(ctx context.Context) NodeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTypeOutput)
}

type NodeTypeOutput struct{ *pulumi.OutputState }

func (NodeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeType)(nil))
}

func (o NodeTypeOutput) ToNodeTypeOutput() NodeTypeOutput {
	return o
}

func (o NodeTypeOutput) ToNodeTypeOutputWithContext(ctx context.Context) NodeTypeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NodeTypeOutput{})
}
