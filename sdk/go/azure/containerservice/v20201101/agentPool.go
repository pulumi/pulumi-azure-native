


package v20201101

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
	EnableAutoScaling         pulumi.BoolPtrOutput                      `pulumi:"enableAutoScaling"`
	EnableNodePublicIP        pulumi.BoolPtrOutput                      `pulumi:"enableNodePublicIP"`
	KubeletConfig             KubeletConfigResponsePtrOutput            `pulumi:"kubeletConfig"`
	LinuxOSConfig             LinuxOSConfigResponsePtrOutput            `pulumi:"linuxOSConfig"`
	MaxCount                  pulumi.IntPtrOutput                       `pulumi:"maxCount"`
	MaxPods                   pulumi.IntPtrOutput                       `pulumi:"maxPods"`
	MinCount                  pulumi.IntPtrOutput                       `pulumi:"minCount"`
	Mode                      pulumi.StringPtrOutput                    `pulumi:"mode"`
	Name                      pulumi.StringOutput                       `pulumi:"name"`
	NodeImageVersion          pulumi.StringOutput                       `pulumi:"nodeImageVersion"`
	NodeLabels                pulumi.StringMapOutput                    `pulumi:"nodeLabels"`
	NodeTaints                pulumi.StringArrayOutput                  `pulumi:"nodeTaints"`
	OrchestratorVersion       pulumi.StringPtrOutput                    `pulumi:"orchestratorVersion"`
	OsDiskSizeGB              pulumi.IntPtrOutput                       `pulumi:"osDiskSizeGB"`
	OsDiskType                pulumi.StringPtrOutput                    `pulumi:"osDiskType"`
	OsType                    pulumi.StringPtrOutput                    `pulumi:"osType"`
	PodSubnetID               pulumi.StringPtrOutput                    `pulumi:"podSubnetID"`
	PowerState                PowerStateResponseOutput                  `pulumi:"powerState"`
	ProvisioningState         pulumi.StringOutput                       `pulumi:"provisioningState"`
	ProximityPlacementGroupID pulumi.StringPtrOutput                    `pulumi:"proximityPlacementGroupID"`
	ScaleSetEvictionPolicy    pulumi.StringPtrOutput                    `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority          pulumi.StringPtrOutput                    `pulumi:"scaleSetPriority"`
	SpotMaxPrice              pulumi.Float64PtrOutput                   `pulumi:"spotMaxPrice"`
	Tags                      pulumi.StringMapOutput                    `pulumi:"tags"`
	Type                      pulumi.StringOutput                       `pulumi:"type"`
	UpgradeSettings           AgentPoolUpgradeSettingsResponsePtrOutput `pulumi:"upgradeSettings"`
	VmSize                    pulumi.StringPtrOutput                    `pulumi:"vmSize"`
	VnetSubnetID              pulumi.StringPtrOutput                    `pulumi:"vnetSubnetID"`
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
			Type: pulumi.String("azure-nextgen:containerservice/v20201101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20190201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20190401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20190601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190801:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20190801:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191001:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20191001:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20191101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200901:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200901:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20201201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210501:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210501:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210801:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210801:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210901:AgentPool"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210901:AgentPool"),
		},
	})
	opts = append(opts, aliases)
	var resource AgentPool
	err := ctx.RegisterResource("azure-native:containerservice/v20201101:AgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentPoolState, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	var resource AgentPool
	err := ctx.ReadResource("azure-native:containerservice/v20201101:AgentPool", name, id, state, &resource, opts...)
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
	EnableAutoScaling         *bool                     `pulumi:"enableAutoScaling"`
	EnableNodePublicIP        *bool                     `pulumi:"enableNodePublicIP"`
	KubeletConfig             *KubeletConfig            `pulumi:"kubeletConfig"`
	LinuxOSConfig             *LinuxOSConfig            `pulumi:"linuxOSConfig"`
	MaxCount                  *int                      `pulumi:"maxCount"`
	MaxPods                   *int                      `pulumi:"maxPods"`
	MinCount                  *int                      `pulumi:"minCount"`
	Mode                      *string                   `pulumi:"mode"`
	NodeLabels                map[string]string         `pulumi:"nodeLabels"`
	NodeTaints                []string                  `pulumi:"nodeTaints"`
	OrchestratorVersion       *string                   `pulumi:"orchestratorVersion"`
	OsDiskSizeGB              *int                      `pulumi:"osDiskSizeGB"`
	OsDiskType                *string                   `pulumi:"osDiskType"`
	OsType                    *string                   `pulumi:"osType"`
	PodSubnetID               *string                   `pulumi:"podSubnetID"`
	ProximityPlacementGroupID *string                   `pulumi:"proximityPlacementGroupID"`
	ResourceGroupName         string                    `pulumi:"resourceGroupName"`
	ResourceName              string                    `pulumi:"resourceName"`
	ScaleSetEvictionPolicy    *string                   `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority          *string                   `pulumi:"scaleSetPriority"`
	SpotMaxPrice              *float64                  `pulumi:"spotMaxPrice"`
	Tags                      map[string]string         `pulumi:"tags"`
	Type                      *string                   `pulumi:"type"`
	UpgradeSettings           *AgentPoolUpgradeSettings `pulumi:"upgradeSettings"`
	VmSize                    *string                   `pulumi:"vmSize"`
	VnetSubnetID              *string                   `pulumi:"vnetSubnetID"`
}


type AgentPoolArgs struct {
	AgentPoolName             pulumi.StringPtrInput
	AvailabilityZones         pulumi.StringArrayInput
	Count                     pulumi.IntPtrInput
	EnableAutoScaling         pulumi.BoolPtrInput
	EnableNodePublicIP        pulumi.BoolPtrInput
	KubeletConfig             KubeletConfigPtrInput
	LinuxOSConfig             LinuxOSConfigPtrInput
	MaxCount                  pulumi.IntPtrInput
	MaxPods                   pulumi.IntPtrInput
	MinCount                  pulumi.IntPtrInput
	Mode                      pulumi.StringPtrInput
	NodeLabels                pulumi.StringMapInput
	NodeTaints                pulumi.StringArrayInput
	OrchestratorVersion       pulumi.StringPtrInput
	OsDiskSizeGB              pulumi.IntPtrInput
	OsDiskType                pulumi.StringPtrInput
	OsType                    pulumi.StringPtrInput
	PodSubnetID               pulumi.StringPtrInput
	ProximityPlacementGroupID pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	ResourceName              pulumi.StringInput
	ScaleSetEvictionPolicy    pulumi.StringPtrInput
	ScaleSetPriority          pulumi.StringPtrInput
	SpotMaxPrice              pulumi.Float64PtrInput
	Tags                      pulumi.StringMapInput
	Type                      pulumi.StringPtrInput
	UpgradeSettings           AgentPoolUpgradeSettingsPtrInput
	VmSize                    pulumi.StringPtrInput
	VnetSubnetID              pulumi.StringPtrInput
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
	return reflect.TypeOf((*AgentPool)(nil))
}

func (i *AgentPool) ToAgentPoolOutput() AgentPoolOutput {
	return i.ToAgentPoolOutputWithContext(context.Background())
}

func (i *AgentPool) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolOutput)
}

type AgentPoolOutput struct{ *pulumi.OutputState }

func (AgentPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPool)(nil))
}

func (o AgentPoolOutput) ToAgentPoolOutput() AgentPoolOutput {
	return o
}

func (o AgentPoolOutput) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AgentPoolOutput{})
}
