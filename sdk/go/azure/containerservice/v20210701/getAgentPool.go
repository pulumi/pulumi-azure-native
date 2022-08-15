


package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAgentPool(ctx *pulumi.Context, args *LookupAgentPoolArgs, opts ...pulumi.InvokeOption) (*LookupAgentPoolResult, error) {
	var rv LookupAgentPoolResult
	err := ctx.Invoke("azure-native:containerservice/v20210701:getAgentPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAgentPoolArgs struct {
	AgentPoolName     string `pulumi:"agentPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupAgentPoolResult struct {
	AvailabilityZones         []string                          `pulumi:"availabilityZones"`
	Count                     *int                              `pulumi:"count"`
	EnableAutoScaling         *bool                             `pulumi:"enableAutoScaling"`
	EnableEncryptionAtHost    *bool                             `pulumi:"enableEncryptionAtHost"`
	EnableFIPS                *bool                             `pulumi:"enableFIPS"`
	EnableNodePublicIP        *bool                             `pulumi:"enableNodePublicIP"`
	EnableUltraSSD            *bool                             `pulumi:"enableUltraSSD"`
	GpuInstanceProfile        *string                           `pulumi:"gpuInstanceProfile"`
	Id                        string                            `pulumi:"id"`
	KubeletConfig             *KubeletConfigResponse            `pulumi:"kubeletConfig"`
	KubeletDiskType           *string                           `pulumi:"kubeletDiskType"`
	LinuxOSConfig             *LinuxOSConfigResponse            `pulumi:"linuxOSConfig"`
	MaxCount                  *int                              `pulumi:"maxCount"`
	MaxPods                   *int                              `pulumi:"maxPods"`
	MinCount                  *int                              `pulumi:"minCount"`
	Mode                      *string                           `pulumi:"mode"`
	Name                      string                            `pulumi:"name"`
	NodeImageVersion          string                            `pulumi:"nodeImageVersion"`
	NodeLabels                map[string]string                 `pulumi:"nodeLabels"`
	NodePublicIPPrefixID      *string                           `pulumi:"nodePublicIPPrefixID"`
	NodeTaints                []string                          `pulumi:"nodeTaints"`
	OrchestratorVersion       *string                           `pulumi:"orchestratorVersion"`
	OsDiskSizeGB              *int                              `pulumi:"osDiskSizeGB"`
	OsDiskType                *string                           `pulumi:"osDiskType"`
	OsSKU                     *string                           `pulumi:"osSKU"`
	OsType                    *string                           `pulumi:"osType"`
	PodSubnetID               *string                           `pulumi:"podSubnetID"`
	PowerState                PowerStateResponse                `pulumi:"powerState"`
	ProvisioningState         string                            `pulumi:"provisioningState"`
	ProximityPlacementGroupID *string                           `pulumi:"proximityPlacementGroupID"`
	ScaleDownMode             *string                           `pulumi:"scaleDownMode"`
	ScaleSetEvictionPolicy    *string                           `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority          *string                           `pulumi:"scaleSetPriority"`
	SpotMaxPrice              *float64                          `pulumi:"spotMaxPrice"`
	Tags                      map[string]string                 `pulumi:"tags"`
	Type                      string                            `pulumi:"type"`
	UpgradeSettings           *AgentPoolUpgradeSettingsResponse `pulumi:"upgradeSettings"`
	VmSize                    *string                           `pulumi:"vmSize"`
	VnetSubnetID              *string                           `pulumi:"vnetSubnetID"`
}

func LookupAgentPoolOutput(ctx *pulumi.Context, args LookupAgentPoolOutputArgs, opts ...pulumi.InvokeOption) LookupAgentPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAgentPoolResult, error) {
			args := v.(LookupAgentPoolArgs)
			r, err := LookupAgentPool(ctx, &args, opts...)
			var s LookupAgentPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAgentPoolResultOutput)
}

type LookupAgentPoolOutputArgs struct {
	AgentPoolName     pulumi.StringInput `pulumi:"agentPoolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupAgentPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentPoolArgs)(nil)).Elem()
}


type LookupAgentPoolResultOutput struct{ *pulumi.OutputState }

func (LookupAgentPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentPoolResult)(nil)).Elem()
}

func (o LookupAgentPoolResultOutput) ToLookupAgentPoolResultOutput() LookupAgentPoolResultOutput {
	return o
}

func (o LookupAgentPoolResultOutput) ToLookupAgentPoolResultOutputWithContext(ctx context.Context) LookupAgentPoolResultOutput {
	return o
}

func (o LookupAgentPoolResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o LookupAgentPoolResultOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o LookupAgentPoolResultOutput) EnableAutoScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *bool { return v.EnableAutoScaling }).(pulumi.BoolPtrOutput)
}

func (o LookupAgentPoolResultOutput) EnableEncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *bool { return v.EnableEncryptionAtHost }).(pulumi.BoolPtrOutput)
}

func (o LookupAgentPoolResultOutput) EnableFIPS() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *bool { return v.EnableFIPS }).(pulumi.BoolPtrOutput)
}

func (o LookupAgentPoolResultOutput) EnableNodePublicIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *bool { return v.EnableNodePublicIP }).(pulumi.BoolPtrOutput)
}

func (o LookupAgentPoolResultOutput) EnableUltraSSD() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *bool { return v.EnableUltraSSD }).(pulumi.BoolPtrOutput)
}

func (o LookupAgentPoolResultOutput) GpuInstanceProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.GpuInstanceProfile }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAgentPoolResultOutput) KubeletConfig() KubeletConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *KubeletConfigResponse { return v.KubeletConfig }).(KubeletConfigResponsePtrOutput)
}

func (o LookupAgentPoolResultOutput) KubeletDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.KubeletDiskType }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) LinuxOSConfig() LinuxOSConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *LinuxOSConfigResponse { return v.LinuxOSConfig }).(LinuxOSConfigResponsePtrOutput)
}

func (o LookupAgentPoolResultOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o LookupAgentPoolResultOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o LookupAgentPoolResultOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o LookupAgentPoolResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAgentPoolResultOutput) NodeImageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.NodeImageVersion }).(pulumi.StringOutput)
}

func (o LookupAgentPoolResultOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) map[string]string { return v.NodeLabels }).(pulumi.StringMapOutput)
}

func (o LookupAgentPoolResultOutput) NodePublicIPPrefixID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.NodePublicIPPrefixID }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) []string { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

func (o LookupAgentPoolResultOutput) OrchestratorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.OrchestratorVersion }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o LookupAgentPoolResultOutput) OsDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.OsDiskType }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) OsSKU() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.OsSKU }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) PodSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.PodSubnetID }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) PowerState() PowerStateResponseOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) PowerStateResponse { return v.PowerState }).(PowerStateResponseOutput)
}

func (o LookupAgentPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAgentPoolResultOutput) ProximityPlacementGroupID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.ProximityPlacementGroupID }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) ScaleDownMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.ScaleDownMode }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) ScaleSetEvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.ScaleSetEvictionPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) ScaleSetPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.ScaleSetPriority }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) SpotMaxPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *float64 { return v.SpotMaxPrice }).(pulumi.Float64PtrOutput)
}

func (o LookupAgentPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAgentPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAgentPoolResultOutput) UpgradeSettings() AgentPoolUpgradeSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *AgentPoolUpgradeSettingsResponse { return v.UpgradeSettings }).(AgentPoolUpgradeSettingsResponsePtrOutput)
}

func (o LookupAgentPoolResultOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAgentPoolResultOutput{})
}
