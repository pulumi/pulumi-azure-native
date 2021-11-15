


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAgentPool(ctx *pulumi.Context, args *LookupAgentPoolArgs, opts ...pulumi.InvokeOption) (*LookupAgentPoolResult, error) {
	var rv LookupAgentPoolResult
	err := ctx.Invoke("azure-native:containerservice/v20201201:getAgentPool", args, &rv, opts...)
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
	EnableNodePublicIP        *bool                             `pulumi:"enableNodePublicIP"`
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
	NodeTaints                []string                          `pulumi:"nodeTaints"`
	OrchestratorVersion       *string                           `pulumi:"orchestratorVersion"`
	OsDiskSizeGB              *int                              `pulumi:"osDiskSizeGB"`
	OsDiskType                *string                           `pulumi:"osDiskType"`
	OsType                    *string                           `pulumi:"osType"`
	PodSubnetID               *string                           `pulumi:"podSubnetID"`
	PowerState                PowerStateResponse                `pulumi:"powerState"`
	ProvisioningState         string                            `pulumi:"provisioningState"`
	ProximityPlacementGroupID *string                           `pulumi:"proximityPlacementGroupID"`
	ScaleSetEvictionPolicy    *string                           `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority          *string                           `pulumi:"scaleSetPriority"`
	SpotMaxPrice              *float64                          `pulumi:"spotMaxPrice"`
	Tags                      map[string]string                 `pulumi:"tags"`
	Type                      string                            `pulumi:"type"`
	UpgradeSettings           *AgentPoolUpgradeSettingsResponse `pulumi:"upgradeSettings"`
	VmSize                    *string                           `pulumi:"vmSize"`
	VnetSubnetID              *string                           `pulumi:"vnetSubnetID"`
}
