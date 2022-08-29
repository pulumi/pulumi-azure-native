


package v20210901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AgentPool struct {
	pulumi.CustomResourceState

	AvailabilityZones         pulumi.StringArrayOutput                  `pulumi:"availabilityZones"`
	Count                     pulumi.IntPtrOutput                       `pulumi:"count"`
	CreationData              CreationDataResponsePtrOutput             `pulumi:"creationData"`
	EnableAutoScaling         pulumi.BoolPtrOutput                      `pulumi:"enableAutoScaling"`
	EnableEncryptionAtHost    pulumi.BoolPtrOutput                      `pulumi:"enableEncryptionAtHost"`
	EnableFIPS                pulumi.BoolPtrOutput                      `pulumi:"enableFIPS"`
	EnableNodePublicIP        pulumi.BoolPtrOutput                      `pulumi:"enableNodePublicIP"`
	EnableUltraSSD            pulumi.BoolPtrOutput                      `pulumi:"enableUltraSSD"`
	GpuInstanceProfile        pulumi.StringPtrOutput                    `pulumi:"gpuInstanceProfile"`
	KubeletConfig             KubeletConfigResponsePtrOutput            `pulumi:"kubeletConfig"`
	KubeletDiskType           pulumi.StringPtrOutput                    `pulumi:"kubeletDiskType"`
	LinuxOSConfig             LinuxOSConfigResponsePtrOutput            `pulumi:"linuxOSConfig"`
	MaxCount                  pulumi.IntPtrOutput                       `pulumi:"maxCount"`
	MaxPods                   pulumi.IntPtrOutput                       `pulumi:"maxPods"`
	MinCount                  pulumi.IntPtrOutput                       `pulumi:"minCount"`
	Mode                      pulumi.StringPtrOutput                    `pulumi:"mode"`
	Name                      pulumi.StringOutput                       `pulumi:"name"`
	NodeImageVersion          pulumi.StringOutput                       `pulumi:"nodeImageVersion"`
	NodeLabels                pulumi.StringMapOutput                    `pulumi:"nodeLabels"`
	NodePublicIPPrefixID      pulumi.StringPtrOutput                    `pulumi:"nodePublicIPPrefixID"`
	NodeTaints                pulumi.StringArrayOutput                  `pulumi:"nodeTaints"`
	OrchestratorVersion       pulumi.StringPtrOutput                    `pulumi:"orchestratorVersion"`
	OsDiskSizeGB              pulumi.IntPtrOutput                       `pulumi:"osDiskSizeGB"`
	OsDiskType                pulumi.StringPtrOutput                    `pulumi:"osDiskType"`
	OsSKU                     pulumi.StringPtrOutput                    `pulumi:"osSKU"`
	OsType                    pulumi.StringPtrOutput                    `pulumi:"osType"`
	PodSubnetID               pulumi.StringPtrOutput                    `pulumi:"podSubnetID"`
	PowerState                PowerStateResponsePtrOutput               `pulumi:"powerState"`
	ProvisioningState         pulumi.StringOutput                       `pulumi:"provisioningState"`
	ProximityPlacementGroupID pulumi.StringPtrOutput                    `pulumi:"proximityPlacementGroupID"`
	ScaleDownMode             pulumi.StringPtrOutput                    `pulumi:"scaleDownMode"`
	ScaleSetEvictionPolicy    pulumi.StringPtrOutput                    `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority          pulumi.StringPtrOutput                    `pulumi:"scaleSetPriority"`
	SpotMaxPrice              pulumi.Float64PtrOutput                   `pulumi:"spotMaxPrice"`
	Tags                      pulumi.StringMapOutput                    `pulumi:"tags"`
	Type                      pulumi.StringOutput                       `pulumi:"type"`
	UpgradeSettings           AgentPoolUpgradeSettingsResponsePtrOutput `pulumi:"upgradeSettings"`
	VmSize                    pulumi.StringPtrOutput                    `pulumi:"vmSize"`
	VnetSubnetID              pulumi.StringPtrOutput                    `pulumi:"vnetSubnetID"`
	WorkloadRuntime           pulumi.StringPtrOutput                    `pulumi:"workloadRuntime"`
}


func NewAgentPool(ctx *pulumi.Context,
	name string, args *AgentPoolArgs, opts ...pulumi.ResourceOption) (*AgentPool, error) {
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
			Type: pulumi.String("azure-native:containerservice:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190801:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191001:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200901:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210501:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210801:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211001:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211101preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220102preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220202preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220302preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220402preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220502preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:AgentPool"),
		},
	})
	opts = append(opts, aliases)
	var resource AgentPool
	err := ctx.RegisterResource("azure-native:containerservice/v20210901:AgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentPoolState, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	var resource AgentPool
	err := ctx.ReadResource("azure-native:containerservice/v20210901:AgentPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type agentPoolState struct {
}

type AgentPoolState struct {
}

func (AgentPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentPoolState)(nil)).Elem()
}

type agentPoolArgs struct {
	AgentPoolName             *string                   `pulumi:"agentPoolName"`
	AvailabilityZones         []string                  `pulumi:"availabilityZones"`
	Count                     *int                      `pulumi:"count"`
	CreationData              *CreationData             `pulumi:"creationData"`
	EnableAutoScaling         *bool                     `pulumi:"enableAutoScaling"`
	EnableEncryptionAtHost    *bool                     `pulumi:"enableEncryptionAtHost"`
	EnableFIPS                *bool                     `pulumi:"enableFIPS"`
	EnableNodePublicIP        *bool                     `pulumi:"enableNodePublicIP"`
	EnableUltraSSD            *bool                     `pulumi:"enableUltraSSD"`
	GpuInstanceProfile        *string                   `pulumi:"gpuInstanceProfile"`
	KubeletConfig             *KubeletConfig            `pulumi:"kubeletConfig"`
	KubeletDiskType           *string                   `pulumi:"kubeletDiskType"`
	LinuxOSConfig             *LinuxOSConfig            `pulumi:"linuxOSConfig"`
	MaxCount                  *int                      `pulumi:"maxCount"`
	MaxPods                   *int                      `pulumi:"maxPods"`
	MinCount                  *int                      `pulumi:"minCount"`
	Mode                      *string                   `pulumi:"mode"`
	NodeLabels                map[string]string         `pulumi:"nodeLabels"`
	NodePublicIPPrefixID      *string                   `pulumi:"nodePublicIPPrefixID"`
	NodeTaints                []string                  `pulumi:"nodeTaints"`
	OrchestratorVersion       *string                   `pulumi:"orchestratorVersion"`
	OsDiskSizeGB              *int                      `pulumi:"osDiskSizeGB"`
	OsDiskType                *string                   `pulumi:"osDiskType"`
	OsSKU                     *string                   `pulumi:"osSKU"`
	OsType                    *string                   `pulumi:"osType"`
	PodSubnetID               *string                   `pulumi:"podSubnetID"`
	PowerState                *PowerState               `pulumi:"powerState"`
	ProximityPlacementGroupID *string                   `pulumi:"proximityPlacementGroupID"`
	ResourceGroupName         string                    `pulumi:"resourceGroupName"`
	ResourceName              string                    `pulumi:"resourceName"`
	ScaleDownMode             *string                   `pulumi:"scaleDownMode"`
	ScaleSetEvictionPolicy    *string                   `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority          *string                   `pulumi:"scaleSetPriority"`
	SpotMaxPrice              *float64                  `pulumi:"spotMaxPrice"`
	Tags                      map[string]string         `pulumi:"tags"`
	Type                      *string                   `pulumi:"type"`
	UpgradeSettings           *AgentPoolUpgradeSettings `pulumi:"upgradeSettings"`
	VmSize                    *string                   `pulumi:"vmSize"`
	VnetSubnetID              *string                   `pulumi:"vnetSubnetID"`
	WorkloadRuntime           *string                   `pulumi:"workloadRuntime"`
}


type AgentPoolArgs struct {
	AgentPoolName             pulumi.StringPtrInput
	AvailabilityZones         pulumi.StringArrayInput
	Count                     pulumi.IntPtrInput
	CreationData              CreationDataPtrInput
	EnableAutoScaling         pulumi.BoolPtrInput
	EnableEncryptionAtHost    pulumi.BoolPtrInput
	EnableFIPS                pulumi.BoolPtrInput
	EnableNodePublicIP        pulumi.BoolPtrInput
	EnableUltraSSD            pulumi.BoolPtrInput
	GpuInstanceProfile        pulumi.StringPtrInput
	KubeletConfig             KubeletConfigPtrInput
	KubeletDiskType           pulumi.StringPtrInput
	LinuxOSConfig             LinuxOSConfigPtrInput
	MaxCount                  pulumi.IntPtrInput
	MaxPods                   pulumi.IntPtrInput
	MinCount                  pulumi.IntPtrInput
	Mode                      pulumi.StringPtrInput
	NodeLabels                pulumi.StringMapInput
	NodePublicIPPrefixID      pulumi.StringPtrInput
	NodeTaints                pulumi.StringArrayInput
	OrchestratorVersion       pulumi.StringPtrInput
	OsDiskSizeGB              pulumi.IntPtrInput
	OsDiskType                pulumi.StringPtrInput
	OsSKU                     pulumi.StringPtrInput
	OsType                    pulumi.StringPtrInput
	PodSubnetID               pulumi.StringPtrInput
	PowerState                PowerStatePtrInput
	ProximityPlacementGroupID pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	ResourceName              pulumi.StringInput
	ScaleDownMode             pulumi.StringPtrInput
	ScaleSetEvictionPolicy    pulumi.StringPtrInput
	ScaleSetPriority          pulumi.StringPtrInput
	SpotMaxPrice              pulumi.Float64PtrInput
	Tags                      pulumi.StringMapInput
	Type                      pulumi.StringPtrInput
	UpgradeSettings           AgentPoolUpgradeSettingsPtrInput
	VmSize                    pulumi.StringPtrInput
	VnetSubnetID              pulumi.StringPtrInput
	WorkloadRuntime           pulumi.StringPtrInput
}

func (AgentPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentPoolArgs)(nil)).Elem()
}

type AgentPoolInput interface {
	pulumi.Input

	ToAgentPoolOutput() AgentPoolOutput
	ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput
}

func (*AgentPool) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPool)(nil)).Elem()
}

func (i *AgentPool) ToAgentPoolOutput() AgentPoolOutput {
	return i.ToAgentPoolOutputWithContext(context.Background())
}

func (i *AgentPool) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolOutput)
}

type AgentPoolOutput struct{ *pulumi.OutputState }

func (AgentPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPool)(nil)).Elem()
}

func (o AgentPoolOutput) ToAgentPoolOutput() AgentPoolOutput {
	return o
}

func (o AgentPoolOutput) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return o
}

func (o AgentPoolOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o AgentPoolOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.Count }).(pulumi.IntPtrOutput)
}

func (o AgentPoolOutput) CreationData() CreationDataResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) CreationDataResponsePtrOutput { return v.CreationData }).(CreationDataResponsePtrOutput)
}

func (o AgentPoolOutput) EnableAutoScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableAutoScaling }).(pulumi.BoolPtrOutput)
}

func (o AgentPoolOutput) EnableEncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableEncryptionAtHost }).(pulumi.BoolPtrOutput)
}

func (o AgentPoolOutput) EnableFIPS() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableFIPS }).(pulumi.BoolPtrOutput)
}

func (o AgentPoolOutput) EnableNodePublicIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableNodePublicIP }).(pulumi.BoolPtrOutput)
}

func (o AgentPoolOutput) EnableUltraSSD() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableUltraSSD }).(pulumi.BoolPtrOutput)
}

func (o AgentPoolOutput) GpuInstanceProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.GpuInstanceProfile }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) KubeletConfig() KubeletConfigResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) KubeletConfigResponsePtrOutput { return v.KubeletConfig }).(KubeletConfigResponsePtrOutput)
}

func (o AgentPoolOutput) KubeletDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.KubeletDiskType }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) LinuxOSConfig() LinuxOSConfigResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) LinuxOSConfigResponsePtrOutput { return v.LinuxOSConfig }).(LinuxOSConfigResponsePtrOutput)
}

func (o AgentPoolOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o AgentPoolOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o AgentPoolOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o AgentPoolOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AgentPoolOutput) NodeImageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.NodeImageVersion }).(pulumi.StringOutput)
}

func (o AgentPoolOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringMapOutput { return v.NodeLabels }).(pulumi.StringMapOutput)
}

func (o AgentPoolOutput) NodePublicIPPrefixID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.NodePublicIPPrefixID }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringArrayOutput { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

func (o AgentPoolOutput) OrchestratorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OrchestratorVersion }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o AgentPoolOutput) OsDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OsDiskType }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) OsSKU() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OsSKU }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) PodSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.PodSubnetID }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) PowerState() PowerStateResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) PowerStateResponsePtrOutput { return v.PowerState }).(PowerStateResponsePtrOutput)
}

func (o AgentPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AgentPoolOutput) ProximityPlacementGroupID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.ProximityPlacementGroupID }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) ScaleDownMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.ScaleDownMode }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) ScaleSetEvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.ScaleSetEvictionPolicy }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) ScaleSetPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.ScaleSetPriority }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) SpotMaxPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.Float64PtrOutput { return v.SpotMaxPrice }).(pulumi.Float64PtrOutput)
}

func (o AgentPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AgentPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AgentPoolOutput) UpgradeSettings() AgentPoolUpgradeSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) AgentPoolUpgradeSettingsResponsePtrOutput { return v.UpgradeSettings }).(AgentPoolUpgradeSettingsResponsePtrOutput)
}

func (o AgentPoolOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.VmSize }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) WorkloadRuntime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.WorkloadRuntime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentPoolOutput{})
}
