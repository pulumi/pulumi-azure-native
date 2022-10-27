


package containerservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AgentPoolMode string

const (
	AgentPoolModeSystem = AgentPoolMode("System")
	AgentPoolModeUser   = AgentPoolMode("User")
)

type AgentPoolType string

const (
	AgentPoolTypeVirtualMachineScaleSets = AgentPoolType("VirtualMachineScaleSets")
	AgentPoolTypeAvailabilitySet         = AgentPoolType("AvailabilitySet")
)

type ConnectionStatus string

const (
	ConnectionStatusPending      = ConnectionStatus("Pending")
	ConnectionStatusApproved     = ConnectionStatus("Approved")
	ConnectionStatusRejected     = ConnectionStatus("Rejected")
	ConnectionStatusDisconnected = ConnectionStatus("Disconnected")
)

type Expander string

const (
	Expander_Least_Waste = Expander("least-waste")
	Expander_Most_Pods   = Expander("most-pods")
	ExpanderPriority     = Expander("priority")
	ExpanderRandom       = Expander("random")
)

type ExtendedLocationTypes string

const (
	ExtendedLocationTypesEdgeZone = ExtendedLocationTypes("EdgeZone")
)

type GPUInstanceProfile string

const (
	GPUInstanceProfileMIG1g = GPUInstanceProfile("MIG1g")
	GPUInstanceProfileMIG2g = GPUInstanceProfile("MIG2g")
	GPUInstanceProfileMIG3g = GPUInstanceProfile("MIG3g")
	GPUInstanceProfileMIG4g = GPUInstanceProfile("MIG4g")
	GPUInstanceProfileMIG7g = GPUInstanceProfile("MIG7g")
)

type KubeletDiskType string

const (
	KubeletDiskTypeOS        = KubeletDiskType("OS")
	KubeletDiskTypeTemporary = KubeletDiskType("Temporary")
)

type LicenseType string

const (
	LicenseTypeNone            = LicenseType("None")
	LicenseType_Windows_Server = LicenseType("Windows_Server")
)

type LoadBalancerSku string

const (
	LoadBalancerSkuStandard = LoadBalancerSku("standard")
	LoadBalancerSkuBasic    = LoadBalancerSku("basic")
)

type ManagedClusterSKUName string

const (
	ManagedClusterSKUNameBasic = ManagedClusterSKUName("Basic")
)

type ManagedClusterSKUTier string

const (
	ManagedClusterSKUTierPaid = ManagedClusterSKUTier("Paid")
	ManagedClusterSKUTierFree = ManagedClusterSKUTier("Free")
)

type NetworkMode string

const (
	NetworkModeTransparent = NetworkMode("transparent")
	NetworkModeBridge      = NetworkMode("bridge")
)

type NetworkPlugin string

const (
	NetworkPluginAzure   = NetworkPlugin("azure")
	NetworkPluginKubenet = NetworkPlugin("kubenet")
)

type NetworkPolicy string

const (
	NetworkPolicyCalico = NetworkPolicy("calico")
	NetworkPolicyAzure  = NetworkPolicy("azure")
)

type OSDiskType string

const (
	OSDiskTypeManaged   = OSDiskType("Managed")
	OSDiskTypeEphemeral = OSDiskType("Ephemeral")
)

type OSSKU string

const (
	OSSKUUbuntu     = OSSKU("Ubuntu")
	OSSKUCBLMariner = OSSKU("CBLMariner")
)

type OSType string

const (
	OSTypeLinux   = OSType("Linux")
	OSTypeWindows = OSType("Windows")
)

type OpenShiftAgentPoolProfileRole string

const (
	OpenShiftAgentPoolProfileRoleCompute = OpenShiftAgentPoolProfileRole("compute")
	OpenShiftAgentPoolProfileRoleInfra   = OpenShiftAgentPoolProfileRole("infra")
)

type OpenShiftContainerServiceVMSize string

const (
	OpenShiftContainerServiceVMSize_Standard_D2s_v3  = OpenShiftContainerServiceVMSize("Standard_D2s_v3")
	OpenShiftContainerServiceVMSize_Standard_D4s_v3  = OpenShiftContainerServiceVMSize("Standard_D4s_v3")
	OpenShiftContainerServiceVMSize_Standard_D8s_v3  = OpenShiftContainerServiceVMSize("Standard_D8s_v3")
	OpenShiftContainerServiceVMSize_Standard_D16s_v3 = OpenShiftContainerServiceVMSize("Standard_D16s_v3")
	OpenShiftContainerServiceVMSize_Standard_D32s_v3 = OpenShiftContainerServiceVMSize("Standard_D32s_v3")
	OpenShiftContainerServiceVMSize_Standard_D64s_v3 = OpenShiftContainerServiceVMSize("Standard_D64s_v3")
	OpenShiftContainerServiceVMSize_Standard_DS4_v2  = OpenShiftContainerServiceVMSize("Standard_DS4_v2")
	OpenShiftContainerServiceVMSize_Standard_DS5_v2  = OpenShiftContainerServiceVMSize("Standard_DS5_v2")
	OpenShiftContainerServiceVMSize_Standard_F8s_v2  = OpenShiftContainerServiceVMSize("Standard_F8s_v2")
	OpenShiftContainerServiceVMSize_Standard_F16s_v2 = OpenShiftContainerServiceVMSize("Standard_F16s_v2")
	OpenShiftContainerServiceVMSize_Standard_F32s_v2 = OpenShiftContainerServiceVMSize("Standard_F32s_v2")
	OpenShiftContainerServiceVMSize_Standard_F64s_v2 = OpenShiftContainerServiceVMSize("Standard_F64s_v2")
	OpenShiftContainerServiceVMSize_Standard_F72s_v2 = OpenShiftContainerServiceVMSize("Standard_F72s_v2")
	OpenShiftContainerServiceVMSize_Standard_F8s     = OpenShiftContainerServiceVMSize("Standard_F8s")
	OpenShiftContainerServiceVMSize_Standard_F16s    = OpenShiftContainerServiceVMSize("Standard_F16s")
	OpenShiftContainerServiceVMSize_Standard_E4s_v3  = OpenShiftContainerServiceVMSize("Standard_E4s_v3")
	OpenShiftContainerServiceVMSize_Standard_E8s_v3  = OpenShiftContainerServiceVMSize("Standard_E8s_v3")
	OpenShiftContainerServiceVMSize_Standard_E16s_v3 = OpenShiftContainerServiceVMSize("Standard_E16s_v3")
	OpenShiftContainerServiceVMSize_Standard_E20s_v3 = OpenShiftContainerServiceVMSize("Standard_E20s_v3")
	OpenShiftContainerServiceVMSize_Standard_E32s_v3 = OpenShiftContainerServiceVMSize("Standard_E32s_v3")
	OpenShiftContainerServiceVMSize_Standard_E64s_v3 = OpenShiftContainerServiceVMSize("Standard_E64s_v3")
	OpenShiftContainerServiceVMSize_Standard_GS2     = OpenShiftContainerServiceVMSize("Standard_GS2")
	OpenShiftContainerServiceVMSize_Standard_GS3     = OpenShiftContainerServiceVMSize("Standard_GS3")
	OpenShiftContainerServiceVMSize_Standard_GS4     = OpenShiftContainerServiceVMSize("Standard_GS4")
	OpenShiftContainerServiceVMSize_Standard_GS5     = OpenShiftContainerServiceVMSize("Standard_GS5")
	OpenShiftContainerServiceVMSize_Standard_DS12_v2 = OpenShiftContainerServiceVMSize("Standard_DS12_v2")
	OpenShiftContainerServiceVMSize_Standard_DS13_v2 = OpenShiftContainerServiceVMSize("Standard_DS13_v2")
	OpenShiftContainerServiceVMSize_Standard_DS14_v2 = OpenShiftContainerServiceVMSize("Standard_DS14_v2")
	OpenShiftContainerServiceVMSize_Standard_DS15_v2 = OpenShiftContainerServiceVMSize("Standard_DS15_v2")
	OpenShiftContainerServiceVMSize_Standard_L4s     = OpenShiftContainerServiceVMSize("Standard_L4s")
	OpenShiftContainerServiceVMSize_Standard_L8s     = OpenShiftContainerServiceVMSize("Standard_L8s")
	OpenShiftContainerServiceVMSize_Standard_L16s    = OpenShiftContainerServiceVMSize("Standard_L16s")
	OpenShiftContainerServiceVMSize_Standard_L32s    = OpenShiftContainerServiceVMSize("Standard_L32s")
)

type OutboundType string

const (
	OutboundTypeLoadBalancer       = OutboundType("loadBalancer")
	OutboundTypeUserDefinedRouting = OutboundType("userDefinedRouting")
)

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned   = ResourceIdentityType("UserAssigned")
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

type ScaleSetEvictionPolicy string

const (
	ScaleSetEvictionPolicyDelete     = ScaleSetEvictionPolicy("Delete")
	ScaleSetEvictionPolicyDeallocate = ScaleSetEvictionPolicy("Deallocate")
)

type ScaleSetPriority string

const (
	ScaleSetPrioritySpot    = ScaleSetPriority("Spot")
	ScaleSetPriorityRegular = ScaleSetPriority("Regular")
)

type SnapshotType string

const (
	// The snapshot is a snapshot of a node pool.
	SnapshotTypeNodePool = SnapshotType("NodePool")
)

type UpgradeChannel string

const (
	UpgradeChannelRapid       = UpgradeChannel("rapid")
	UpgradeChannelStable      = UpgradeChannel("stable")
	UpgradeChannelPatch       = UpgradeChannel("patch")
	UpgradeChannel_Node_Image = UpgradeChannel("node-image")
	UpgradeChannelNone        = UpgradeChannel("none")
)

type WeekDay string

const (
	WeekDaySunday    = WeekDay("Sunday")
	WeekDayMonday    = WeekDay("Monday")
	WeekDayTuesday   = WeekDay("Tuesday")
	WeekDayWednesday = WeekDay("Wednesday")
	WeekDayThursday  = WeekDay("Thursday")
	WeekDayFriday    = WeekDay("Friday")
	WeekDaySaturday  = WeekDay("Saturday")
)

func init() {
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
