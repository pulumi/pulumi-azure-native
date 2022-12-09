


package v20180801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ManagedCluster struct {
	pulumi.CustomResourceState

	AadProfile              ManagedClusterAADProfileResponsePtrOutput              `pulumi:"aadProfile"`
	AddonProfiles           ManagedClusterAddonProfileResponseMapOutput            `pulumi:"addonProfiles"`
	AgentPoolProfiles       ManagedClusterAgentPoolProfileResponseArrayOutput      `pulumi:"agentPoolProfiles"`
	DnsPrefix               pulumi.StringPtrOutput                                 `pulumi:"dnsPrefix"`
	EnableRBAC              pulumi.BoolPtrOutput                                   `pulumi:"enableRBAC"`
	Fqdn                    pulumi.StringOutput                                    `pulumi:"fqdn"`
	KubernetesVersion       pulumi.StringPtrOutput                                 `pulumi:"kubernetesVersion"`
	LinuxProfile            ContainerServiceLinuxProfileResponsePtrOutput          `pulumi:"linuxProfile"`
	Location                pulumi.StringOutput                                    `pulumi:"location"`
	Name                    pulumi.StringOutput                                    `pulumi:"name"`
	NetworkProfile          ContainerServiceNetworkProfileResponsePtrOutput        `pulumi:"networkProfile"`
	NodeResourceGroup       pulumi.StringOutput                                    `pulumi:"nodeResourceGroup"`
	ProvisioningState       pulumi.StringOutput                                    `pulumi:"provisioningState"`
	ServicePrincipalProfile ManagedClusterServicePrincipalProfileResponsePtrOutput `pulumi:"servicePrincipalProfile"`
	Tags                    pulumi.StringMapOutput                                 `pulumi:"tags"`
	Type                    pulumi.StringOutput                                    `pulumi:"type"`
}


func NewManagedCluster(ctx *pulumi.Context,
	name string, args *ManagedClusterArgs, opts ...pulumi.ResourceOption) (*ManagedCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.NetworkProfile != nil {
		args.NetworkProfile = args.NetworkProfile.ToContainerServiceNetworkProfilePtrOutput().ApplyT(func(v *ContainerServiceNetworkProfile) *ContainerServiceNetworkProfile { return v.Defaults() }).(ContainerServiceNetworkProfilePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20170831:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20180331:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190401:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190601:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190801:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191001:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200301:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200401:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200601:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200701:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200901:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210301:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210501:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210701:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210801:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210901:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211001:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211101preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220102preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220202preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220301:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220302preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220401:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220402preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220502preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220601:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220701:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220802preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220803preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220901:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220902preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221002preview:ManagedCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedCluster
	err := ctx.RegisterResource("azure-native:containerservice/v20180801preview:ManagedCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedClusterState, opts ...pulumi.ResourceOption) (*ManagedCluster, error) {
	var resource ManagedCluster
	err := ctx.ReadResource("azure-native:containerservice/v20180801preview:ManagedCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedClusterState struct {
}

type ManagedClusterState struct {
}

func (ManagedClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterState)(nil)).Elem()
}

type managedClusterArgs struct {
	AadProfile              *ManagedClusterAADProfile              `pulumi:"aadProfile"`
	AddonProfiles           map[string]ManagedClusterAddonProfile  `pulumi:"addonProfiles"`
	AgentPoolProfiles       []ManagedClusterAgentPoolProfile       `pulumi:"agentPoolProfiles"`
	DnsPrefix               *string                                `pulumi:"dnsPrefix"`
	EnableRBAC              *bool                                  `pulumi:"enableRBAC"`
	KubernetesVersion       *string                                `pulumi:"kubernetesVersion"`
	LinuxProfile            *ContainerServiceLinuxProfile          `pulumi:"linuxProfile"`
	Location                *string                                `pulumi:"location"`
	NetworkProfile          *ContainerServiceNetworkProfile        `pulumi:"networkProfile"`
	ResourceGroupName       string                                 `pulumi:"resourceGroupName"`
	ResourceName            *string                                `pulumi:"resourceName"`
	ServicePrincipalProfile *ManagedClusterServicePrincipalProfile `pulumi:"servicePrincipalProfile"`
	Tags                    map[string]string                      `pulumi:"tags"`
}


type ManagedClusterArgs struct {
	AadProfile              ManagedClusterAADProfilePtrInput
	AddonProfiles           ManagedClusterAddonProfileMapInput
	AgentPoolProfiles       ManagedClusterAgentPoolProfileArrayInput
	DnsPrefix               pulumi.StringPtrInput
	EnableRBAC              pulumi.BoolPtrInput
	KubernetesVersion       pulumi.StringPtrInput
	LinuxProfile            ContainerServiceLinuxProfilePtrInput
	Location                pulumi.StringPtrInput
	NetworkProfile          ContainerServiceNetworkProfilePtrInput
	ResourceGroupName       pulumi.StringInput
	ResourceName            pulumi.StringPtrInput
	ServicePrincipalProfile ManagedClusterServicePrincipalProfilePtrInput
	Tags                    pulumi.StringMapInput
}

func (ManagedClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterArgs)(nil)).Elem()
}

type ManagedClusterInput interface {
	pulumi.Input

	ToManagedClusterOutput() ManagedClusterOutput
	ToManagedClusterOutputWithContext(ctx context.Context) ManagedClusterOutput
}

func (*ManagedCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCluster)(nil)).Elem()
}

func (i *ManagedCluster) ToManagedClusterOutput() ManagedClusterOutput {
	return i.ToManagedClusterOutputWithContext(context.Background())
}

func (i *ManagedCluster) ToManagedClusterOutputWithContext(ctx context.Context) ManagedClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterOutput)
}

type ManagedClusterOutput struct{ *pulumi.OutputState }

func (ManagedClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCluster)(nil)).Elem()
}

func (o ManagedClusterOutput) ToManagedClusterOutput() ManagedClusterOutput {
	return o
}

func (o ManagedClusterOutput) ToManagedClusterOutputWithContext(ctx context.Context) ManagedClusterOutput {
	return o
}

func (o ManagedClusterOutput) AadProfile() ManagedClusterAADProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAADProfileResponsePtrOutput { return v.AadProfile }).(ManagedClusterAADProfileResponsePtrOutput)
}

func (o ManagedClusterOutput) AddonProfiles() ManagedClusterAddonProfileResponseMapOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAddonProfileResponseMapOutput { return v.AddonProfiles }).(ManagedClusterAddonProfileResponseMapOutput)
}

func (o ManagedClusterOutput) AgentPoolProfiles() ManagedClusterAgentPoolProfileResponseArrayOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterAgentPoolProfileResponseArrayOutput { return v.AgentPoolProfiles }).(ManagedClusterAgentPoolProfileResponseArrayOutput)
}

func (o ManagedClusterOutput) DnsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.DnsPrefix }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterOutput) EnableRBAC() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.BoolPtrOutput { return v.EnableRBAC }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) KubernetesVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringPtrOutput { return v.KubernetesVersion }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterOutput) LinuxProfile() ContainerServiceLinuxProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ContainerServiceLinuxProfileResponsePtrOutput { return v.LinuxProfile }).(ContainerServiceLinuxProfileResponsePtrOutput)
}

func (o ManagedClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) NetworkProfile() ContainerServiceNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ContainerServiceNetworkProfileResponsePtrOutput { return v.NetworkProfile }).(ContainerServiceNetworkProfileResponsePtrOutput)
}

func (o ManagedClusterOutput) NodeResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.NodeResourceGroup }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ManagedClusterOutput) ServicePrincipalProfile() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedCluster) ManagedClusterServicePrincipalProfileResponsePtrOutput {
		return v.ServicePrincipalProfile
	}).(ManagedClusterServicePrincipalProfileResponsePtrOutput)
}

func (o ManagedClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ManagedClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedClusterOutput{})
}
