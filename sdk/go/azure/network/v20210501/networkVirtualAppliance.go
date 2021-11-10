


package v20210501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkVirtualAppliance struct {
	pulumi.CustomResourceState

	AddressPrefix               pulumi.StringOutput                              `pulumi:"addressPrefix"`
	BootStrapConfigurationBlobs pulumi.StringArrayOutput                         `pulumi:"bootStrapConfigurationBlobs"`
	CloudInitConfiguration      pulumi.StringPtrOutput                           `pulumi:"cloudInitConfiguration"`
	CloudInitConfigurationBlobs pulumi.StringArrayOutput                         `pulumi:"cloudInitConfigurationBlobs"`
	Etag                        pulumi.StringOutput                              `pulumi:"etag"`
	Identity                    ManagedServiceIdentityResponsePtrOutput          `pulumi:"identity"`
	InboundSecurityRules        SubResourceResponseArrayOutput                   `pulumi:"inboundSecurityRules"`
	Location                    pulumi.StringPtrOutput                           `pulumi:"location"`
	Name                        pulumi.StringOutput                              `pulumi:"name"`
	NvaSku                      VirtualApplianceSkuPropertiesResponsePtrOutput   `pulumi:"nvaSku"`
	ProvisioningState           pulumi.StringOutput                              `pulumi:"provisioningState"`
	SshPublicKey                pulumi.StringPtrOutput                           `pulumi:"sshPublicKey"`
	Tags                        pulumi.StringMapOutput                           `pulumi:"tags"`
	Type                        pulumi.StringOutput                              `pulumi:"type"`
	VirtualApplianceAsn         pulumi.Float64PtrOutput                          `pulumi:"virtualApplianceAsn"`
	VirtualApplianceNics        VirtualApplianceNicPropertiesResponseArrayOutput `pulumi:"virtualApplianceNics"`
	VirtualApplianceSites       SubResourceResponseArrayOutput                   `pulumi:"virtualApplianceSites"`
	VirtualHub                  SubResourceResponsePtrOutput                     `pulumi:"virtualHub"`
}


func NewNetworkVirtualAppliance(ctx *pulumi.Context,
	name string, args *NetworkVirtualApplianceArgs, opts ...pulumi.ResourceOption) (*NetworkVirtualAppliance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NetworkVirtualAppliance"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NetworkVirtualAppliance"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkVirtualAppliance
	err := ctx.RegisterResource("azure-native:network/v20210501:NetworkVirtualAppliance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkVirtualAppliance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkVirtualApplianceState, opts ...pulumi.ResourceOption) (*NetworkVirtualAppliance, error) {
	var resource NetworkVirtualAppliance
	err := ctx.ReadResource("azure-native:network/v20210501:NetworkVirtualAppliance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkVirtualApplianceState struct {
}

type NetworkVirtualApplianceState struct {
}

func (NetworkVirtualApplianceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVirtualApplianceState)(nil)).Elem()
}

type networkVirtualApplianceArgs struct {
	BootStrapConfigurationBlobs []string                       `pulumi:"bootStrapConfigurationBlobs"`
	CloudInitConfiguration      *string                        `pulumi:"cloudInitConfiguration"`
	CloudInitConfigurationBlobs []string                       `pulumi:"cloudInitConfigurationBlobs"`
	Id                          *string                        `pulumi:"id"`
	Identity                    *ManagedServiceIdentity        `pulumi:"identity"`
	Location                    *string                        `pulumi:"location"`
	NetworkVirtualApplianceName *string                        `pulumi:"networkVirtualApplianceName"`
	NvaSku                      *VirtualApplianceSkuProperties `pulumi:"nvaSku"`
	ResourceGroupName           string                         `pulumi:"resourceGroupName"`
	SshPublicKey                *string                        `pulumi:"sshPublicKey"`
	Tags                        map[string]string              `pulumi:"tags"`
	VirtualApplianceAsn         *float64                       `pulumi:"virtualApplianceAsn"`
	VirtualHub                  *SubResource                   `pulumi:"virtualHub"`
}


type NetworkVirtualApplianceArgs struct {
	BootStrapConfigurationBlobs pulumi.StringArrayInput
	CloudInitConfiguration      pulumi.StringPtrInput
	CloudInitConfigurationBlobs pulumi.StringArrayInput
	Id                          pulumi.StringPtrInput
	Identity                    ManagedServiceIdentityPtrInput
	Location                    pulumi.StringPtrInput
	NetworkVirtualApplianceName pulumi.StringPtrInput
	NvaSku                      VirtualApplianceSkuPropertiesPtrInput
	ResourceGroupName           pulumi.StringInput
	SshPublicKey                pulumi.StringPtrInput
	Tags                        pulumi.StringMapInput
	VirtualApplianceAsn         pulumi.Float64PtrInput
	VirtualHub                  SubResourcePtrInput
}

func (NetworkVirtualApplianceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVirtualApplianceArgs)(nil)).Elem()
}

type NetworkVirtualApplianceInput interface {
	pulumi.Input

	ToNetworkVirtualApplianceOutput() NetworkVirtualApplianceOutput
	ToNetworkVirtualApplianceOutputWithContext(ctx context.Context) NetworkVirtualApplianceOutput
}

func (*NetworkVirtualAppliance) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkVirtualAppliance)(nil))
}

func (i *NetworkVirtualAppliance) ToNetworkVirtualApplianceOutput() NetworkVirtualApplianceOutput {
	return i.ToNetworkVirtualApplianceOutputWithContext(context.Background())
}

func (i *NetworkVirtualAppliance) ToNetworkVirtualApplianceOutputWithContext(ctx context.Context) NetworkVirtualApplianceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkVirtualApplianceOutput)
}

type NetworkVirtualApplianceOutput struct{ *pulumi.OutputState }

func (NetworkVirtualApplianceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkVirtualAppliance)(nil))
}

func (o NetworkVirtualApplianceOutput) ToNetworkVirtualApplianceOutput() NetworkVirtualApplianceOutput {
	return o
}

func (o NetworkVirtualApplianceOutput) ToNetworkVirtualApplianceOutputWithContext(ctx context.Context) NetworkVirtualApplianceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkVirtualApplianceOutput{})
}
