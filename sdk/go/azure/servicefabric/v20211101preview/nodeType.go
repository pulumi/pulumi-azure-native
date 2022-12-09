


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NodeType struct {
	pulumi.CustomResourceState

	AdditionalDataDisks          VmssDataDiskResponseArrayOutput           `pulumi:"additionalDataDisks"`
	ApplicationPorts             EndpointRangeDescriptionResponsePtrOutput `pulumi:"applicationPorts"`
	Capacities                   pulumi.StringMapOutput                    `pulumi:"capacities"`
	DataDiskLetter               pulumi.StringPtrOutput                    `pulumi:"dataDiskLetter"`
	DataDiskSizeGB               pulumi.IntPtrOutput                       `pulumi:"dataDiskSizeGB"`
	DataDiskType                 pulumi.StringPtrOutput                    `pulumi:"dataDiskType"`
	EnableAcceleratedNetworking  pulumi.BoolPtrOutput                      `pulumi:"enableAcceleratedNetworking"`
	EnableEncryptionAtHost       pulumi.BoolPtrOutput                      `pulumi:"enableEncryptionAtHost"`
	EphemeralPorts               EndpointRangeDescriptionResponsePtrOutput `pulumi:"ephemeralPorts"`
	FrontendConfigurations       FrontendConfigurationResponseArrayOutput  `pulumi:"frontendConfigurations"`
	IsPrimary                    pulumi.BoolOutput                         `pulumi:"isPrimary"`
	IsStateless                  pulumi.BoolPtrOutput                      `pulumi:"isStateless"`
	MultiplePlacementGroups      pulumi.BoolPtrOutput                      `pulumi:"multiplePlacementGroups"`
	Name                         pulumi.StringOutput                       `pulumi:"name"`
	NetworkSecurityRules         NetworkSecurityRuleResponseArrayOutput    `pulumi:"networkSecurityRules"`
	PlacementProperties          pulumi.StringMapOutput                    `pulumi:"placementProperties"`
	ProvisioningState            pulumi.StringOutput                       `pulumi:"provisioningState"`
	Sku                          NodeTypeSkuResponsePtrOutput              `pulumi:"sku"`
	SystemData                   SystemDataResponseOutput                  `pulumi:"systemData"`
	Tags                         pulumi.StringMapOutput                    `pulumi:"tags"`
	Type                         pulumi.StringOutput                       `pulumi:"type"`
	UseDefaultPublicLoadBalancer pulumi.BoolPtrOutput                      `pulumi:"useDefaultPublicLoadBalancer"`
	UseTempDataDisk              pulumi.BoolPtrOutput                      `pulumi:"useTempDataDisk"`
	VmExtensions                 VMSSExtensionResponseArrayOutput          `pulumi:"vmExtensions"`
	VmImageOffer                 pulumi.StringPtrOutput                    `pulumi:"vmImageOffer"`
	VmImagePublisher             pulumi.StringPtrOutput                    `pulumi:"vmImagePublisher"`
	VmImageSku                   pulumi.StringPtrOutput                    `pulumi:"vmImageSku"`
	VmImageVersion               pulumi.StringPtrOutput                    `pulumi:"vmImageVersion"`
	VmInstanceCount              pulumi.IntOutput                          `pulumi:"vmInstanceCount"`
	VmManagedIdentity            VmManagedIdentityResponsePtrOutput        `pulumi:"vmManagedIdentity"`
	VmSecrets                    VaultSecretGroupResponseArrayOutput       `pulumi:"vmSecrets"`
	VmSize                       pulumi.StringPtrOutput                    `pulumi:"vmSize"`
}


func NewNodeType(ctx *pulumi.Context,
	name string, args *NodeTypeArgs, opts ...pulumi.ResourceOption) (*NodeType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
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
	if isZero(args.EnableEncryptionAtHost) {
		args.EnableEncryptionAtHost = pulumi.BoolPtr(false)
	}
	if isZero(args.IsStateless) {
		args.IsStateless = pulumi.BoolPtr(false)
	}
	if isZero(args.MultiplePlacementGroups) {
		args.MultiplePlacementGroups = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20200101preview:NodeType"),
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
			Type: pulumi.String("azure-native:servicefabric/v20220101:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220201preview:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220601preview:NodeType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220801preview:NodeType"),
		},
	})
	opts = append(opts, aliases)
	var resource NodeType
	err := ctx.RegisterResource("azure-native:servicefabric/v20211101preview:NodeType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNodeType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeTypeState, opts ...pulumi.ResourceOption) (*NodeType, error) {
	var resource NodeType
	err := ctx.ReadResource("azure-native:servicefabric/v20211101preview:NodeType", name, id, state, &resource, opts...)
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
	AdditionalDataDisks          []VmssDataDisk            `pulumi:"additionalDataDisks"`
	ApplicationPorts             *EndpointRangeDescription `pulumi:"applicationPorts"`
	Capacities                   map[string]string         `pulumi:"capacities"`
	ClusterName                  string                    `pulumi:"clusterName"`
	DataDiskLetter               *string                   `pulumi:"dataDiskLetter"`
	DataDiskSizeGB               *int                      `pulumi:"dataDiskSizeGB"`
	DataDiskType                 *string                   `pulumi:"dataDiskType"`
	EnableAcceleratedNetworking  *bool                     `pulumi:"enableAcceleratedNetworking"`
	EnableEncryptionAtHost       *bool                     `pulumi:"enableEncryptionAtHost"`
	EphemeralPorts               *EndpointRangeDescription `pulumi:"ephemeralPorts"`
	FrontendConfigurations       []FrontendConfiguration   `pulumi:"frontendConfigurations"`
	IsPrimary                    bool                      `pulumi:"isPrimary"`
	IsStateless                  *bool                     `pulumi:"isStateless"`
	MultiplePlacementGroups      *bool                     `pulumi:"multiplePlacementGroups"`
	NetworkSecurityRules         []NetworkSecurityRule     `pulumi:"networkSecurityRules"`
	NodeTypeName                 *string                   `pulumi:"nodeTypeName"`
	PlacementProperties          map[string]string         `pulumi:"placementProperties"`
	ResourceGroupName            string                    `pulumi:"resourceGroupName"`
	Sku                          *NodeTypeSku              `pulumi:"sku"`
	Tags                         map[string]string         `pulumi:"tags"`
	UseDefaultPublicLoadBalancer *bool                     `pulumi:"useDefaultPublicLoadBalancer"`
	UseTempDataDisk              *bool                     `pulumi:"useTempDataDisk"`
	VmExtensions                 []VMSSExtension           `pulumi:"vmExtensions"`
	VmImageOffer                 *string                   `pulumi:"vmImageOffer"`
	VmImagePublisher             *string                   `pulumi:"vmImagePublisher"`
	VmImageSku                   *string                   `pulumi:"vmImageSku"`
	VmImageVersion               *string                   `pulumi:"vmImageVersion"`
	VmInstanceCount              int                       `pulumi:"vmInstanceCount"`
	VmManagedIdentity            *VmManagedIdentity        `pulumi:"vmManagedIdentity"`
	VmSecrets                    []VaultSecretGroup        `pulumi:"vmSecrets"`
	VmSize                       *string                   `pulumi:"vmSize"`
}


type NodeTypeArgs struct {
	AdditionalDataDisks          VmssDataDiskArrayInput
	ApplicationPorts             EndpointRangeDescriptionPtrInput
	Capacities                   pulumi.StringMapInput
	ClusterName                  pulumi.StringInput
	DataDiskLetter               pulumi.StringPtrInput
	DataDiskSizeGB               pulumi.IntPtrInput
	DataDiskType                 pulumi.StringPtrInput
	EnableAcceleratedNetworking  pulumi.BoolPtrInput
	EnableEncryptionAtHost       pulumi.BoolPtrInput
	EphemeralPorts               EndpointRangeDescriptionPtrInput
	FrontendConfigurations       FrontendConfigurationArrayInput
	IsPrimary                    pulumi.BoolInput
	IsStateless                  pulumi.BoolPtrInput
	MultiplePlacementGroups      pulumi.BoolPtrInput
	NetworkSecurityRules         NetworkSecurityRuleArrayInput
	NodeTypeName                 pulumi.StringPtrInput
	PlacementProperties          pulumi.StringMapInput
	ResourceGroupName            pulumi.StringInput
	Sku                          NodeTypeSkuPtrInput
	Tags                         pulumi.StringMapInput
	UseDefaultPublicLoadBalancer pulumi.BoolPtrInput
	UseTempDataDisk              pulumi.BoolPtrInput
	VmExtensions                 VMSSExtensionArrayInput
	VmImageOffer                 pulumi.StringPtrInput
	VmImagePublisher             pulumi.StringPtrInput
	VmImageSku                   pulumi.StringPtrInput
	VmImageVersion               pulumi.StringPtrInput
	VmInstanceCount              pulumi.IntInput
	VmManagedIdentity            VmManagedIdentityPtrInput
	VmSecrets                    VaultSecretGroupArrayInput
	VmSize                       pulumi.StringPtrInput
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

func (o NodeTypeOutput) AdditionalDataDisks() VmssDataDiskResponseArrayOutput {
	return o.ApplyT(func(v *NodeType) VmssDataDiskResponseArrayOutput { return v.AdditionalDataDisks }).(VmssDataDiskResponseArrayOutput)
}

func (o NodeTypeOutput) ApplicationPorts() EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *NodeType) EndpointRangeDescriptionResponsePtrOutput { return v.ApplicationPorts }).(EndpointRangeDescriptionResponsePtrOutput)
}

func (o NodeTypeOutput) Capacities() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringMapOutput { return v.Capacities }).(pulumi.StringMapOutput)
}

func (o NodeTypeOutput) DataDiskLetter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringPtrOutput { return v.DataDiskLetter }).(pulumi.StringPtrOutput)
}

func (o NodeTypeOutput) DataDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeType) pulumi.IntPtrOutput { return v.DataDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o NodeTypeOutput) DataDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringPtrOutput { return v.DataDiskType }).(pulumi.StringPtrOutput)
}

func (o NodeTypeOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NodeType) pulumi.BoolPtrOutput { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o NodeTypeOutput) EnableEncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NodeType) pulumi.BoolPtrOutput { return v.EnableEncryptionAtHost }).(pulumi.BoolPtrOutput)
}

func (o NodeTypeOutput) EphemeralPorts() EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *NodeType) EndpointRangeDescriptionResponsePtrOutput { return v.EphemeralPorts }).(EndpointRangeDescriptionResponsePtrOutput)
}

func (o NodeTypeOutput) FrontendConfigurations() FrontendConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NodeType) FrontendConfigurationResponseArrayOutput { return v.FrontendConfigurations }).(FrontendConfigurationResponseArrayOutput)
}

func (o NodeTypeOutput) IsPrimary() pulumi.BoolOutput {
	return o.ApplyT(func(v *NodeType) pulumi.BoolOutput { return v.IsPrimary }).(pulumi.BoolOutput)
}

func (o NodeTypeOutput) IsStateless() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NodeType) pulumi.BoolPtrOutput { return v.IsStateless }).(pulumi.BoolPtrOutput)
}

func (o NodeTypeOutput) MultiplePlacementGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NodeType) pulumi.BoolPtrOutput { return v.MultiplePlacementGroups }).(pulumi.BoolPtrOutput)
}

func (o NodeTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NodeTypeOutput) NetworkSecurityRules() NetworkSecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v *NodeType) NetworkSecurityRuleResponseArrayOutput { return v.NetworkSecurityRules }).(NetworkSecurityRuleResponseArrayOutput)
}

func (o NodeTypeOutput) PlacementProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringMapOutput { return v.PlacementProperties }).(pulumi.StringMapOutput)
}

func (o NodeTypeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NodeTypeOutput) Sku() NodeTypeSkuResponsePtrOutput {
	return o.ApplyT(func(v *NodeType) NodeTypeSkuResponsePtrOutput { return v.Sku }).(NodeTypeSkuResponsePtrOutput)
}

func (o NodeTypeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NodeType) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o NodeTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NodeTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeType) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o NodeTypeOutput) UseDefaultPublicLoadBalancer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NodeType) pulumi.BoolPtrOutput { return v.UseDefaultPublicLoadBalancer }).(pulumi.BoolPtrOutput)
}

func (o NodeTypeOutput) UseTempDataDisk() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NodeType) pulumi.BoolPtrOutput { return v.UseTempDataDisk }).(pulumi.BoolPtrOutput)
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

func (o NodeTypeOutput) VmManagedIdentity() VmManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *NodeType) VmManagedIdentityResponsePtrOutput { return v.VmManagedIdentity }).(VmManagedIdentityResponsePtrOutput)
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
