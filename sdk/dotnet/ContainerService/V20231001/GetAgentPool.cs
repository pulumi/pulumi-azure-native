// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20231001
{
    public static class GetAgentPool
    {
        /// <summary>
        /// Agent Pool.
        /// </summary>
        public static Task<GetAgentPoolResult> InvokeAsync(GetAgentPoolArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAgentPoolResult>("azure-native:containerservice/v20231001:getAgentPool", args ?? new GetAgentPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Agent Pool.
        /// </summary>
        public static Output<GetAgentPoolResult> Invoke(GetAgentPoolInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAgentPoolResult>("azure-native:containerservice/v20231001:getAgentPool", args ?? new GetAgentPoolInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Agent Pool.
        /// </summary>
        public static Output<GetAgentPoolResult> Invoke(GetAgentPoolInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAgentPoolResult>("azure-native:containerservice/v20231001:getAgentPool", args ?? new GetAgentPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAgentPoolArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the agent pool.
        /// </summary>
        [Input("agentPoolName", required: true)]
        public string AgentPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the managed cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetAgentPoolArgs()
        {
        }
        public static new GetAgentPoolArgs Empty => new GetAgentPoolArgs();
    }

    public sealed class GetAgentPoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the agent pool.
        /// </summary>
        [Input("agentPoolName", required: true)]
        public Input<string> AgentPoolName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the managed cluster resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public GetAgentPoolInvokeArgs()
        {
        }
        public static new GetAgentPoolInvokeArgs Empty => new GetAgentPoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetAgentPoolResult
    {
        /// <summary>
        /// The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.
        /// </summary>
        public readonly ImmutableArray<string> AvailabilityZones;
        /// <summary>
        /// AKS will associate the specified agent pool with the Capacity Reservation Group.
        /// </summary>
        public readonly string? CapacityReservationGroupID;
        /// <summary>
        /// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using a snapshot.
        /// </summary>
        public readonly Outputs.CreationDataResponse? CreationData;
        /// <summary>
        /// If orchestratorVersion is a fully specified version &lt;major.minor.patch&gt;, this field will be exactly equal to it. If orchestratorVersion is &lt;major.minor&gt;, this field will contain the full &lt;major.minor.patch&gt; version being used.
        /// </summary>
        public readonly string CurrentOrchestratorVersion;
        /// <summary>
        /// Whether to enable auto-scaler
        /// </summary>
        public readonly bool? EnableAutoScaling;
        /// <summary>
        /// This is only supported on certain VM sizes and in certain Azure regions. For more information, see: https://docs.microsoft.com/azure/aks/enable-host-encryption
        /// </summary>
        public readonly bool? EnableEncryptionAtHost;
        /// <summary>
        /// See [Add a FIPS-enabled node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more details.
        /// </summary>
        public readonly bool? EnableFIPS;
        /// <summary>
        /// Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses. A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine to minimize hops. For more information see [assigning a public IP per node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The default is false.
        /// </summary>
        public readonly bool? EnableNodePublicIP;
        /// <summary>
        /// Whether to enable UltraSSD
        /// </summary>
        public readonly bool? EnableUltraSSD;
        /// <summary>
        /// GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.
        /// </summary>
        public readonly string? GpuInstanceProfile;
        /// <summary>
        /// This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}. For more information see [Azure dedicated hosts](https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts).
        /// </summary>
        public readonly string? HostGroupID;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Kubelet configuration on the agent pool nodes.
        /// </summary>
        public readonly Outputs.KubeletConfigResponse? KubeletConfig;
        /// <summary>
        /// Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.
        /// </summary>
        public readonly string? KubeletDiskType;
        /// <summary>
        /// The OS configuration of Linux agent nodes.
        /// </summary>
        public readonly Outputs.LinuxOSConfigResponse? LinuxOSConfig;
        /// <summary>
        /// The maximum number of nodes for auto-scaling
        /// </summary>
        public readonly int? MaxCount;
        /// <summary>
        /// The maximum number of pods that can run on a node.
        /// </summary>
        public readonly int? MaxPods;
        /// <summary>
        /// The minimum number of nodes for auto-scaling
        /// </summary>
        public readonly int? MinCount;
        /// <summary>
        /// A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool restrictions and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Network-related settings of an agent pool.
        /// </summary>
        public readonly Outputs.AgentPoolNetworkProfileResponse? NetworkProfile;
        /// <summary>
        /// The version of node image
        /// </summary>
        public readonly string NodeImageVersion;
        /// <summary>
        /// The node labels to be persisted across all nodes in agent pool.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? NodeLabels;
        /// <summary>
        /// This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}
        /// </summary>
        public readonly string? NodePublicIPPrefixID;
        /// <summary>
        /// The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
        /// </summary>
        public readonly ImmutableArray<string> NodeTaints;
        /// <summary>
        /// Both patch version &lt;major.minor.patch&gt; (e.g. 1.20.13) and &lt;major.minor&gt; (e.g. 1.20) are supported. When &lt;major.minor&gt; is specified, the latest supported GA patch version is chosen automatically. Updating the cluster with the same &lt;major.minor&gt; once it has been created (e.g. 1.14.x -&gt; 1.14) will not trigger an upgrade, even if a newer patch version is available. As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool).
        /// </summary>
        public readonly string? OrchestratorVersion;
        /// <summary>
        /// OS Disk Size in GB to be used to specify the disk size for every machine in the master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
        /// </summary>
        public readonly int? OsDiskSizeGB;
        /// <summary>
        /// The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise, defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).
        /// </summary>
        public readonly string? OsDiskType;
        /// <summary>
        /// Specifies the OS SKU used by the agent pool. The default is Ubuntu if OSType is Linux. The default is Windows2019 when Kubernetes &lt;= 1.24 or Windows2022 when Kubernetes &gt;= 1.25 if OSType is Windows.
        /// </summary>
        public readonly string? OsSKU;
        /// <summary>
        /// The operating system type. The default is Linux.
        /// </summary>
        public readonly string? OsType;
        /// <summary>
        /// If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
        /// </summary>
        public readonly string? PodSubnetID;
        /// <summary>
        /// When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only be stopped if it is Running and provisioning state is Succeeded
        /// </summary>
        public readonly Outputs.PowerStateResponse? PowerState;
        /// <summary>
        /// The current deployment or provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The ID for Proximity Placement Group.
        /// </summary>
        public readonly string? ProximityPlacementGroupID;
        /// <summary>
        /// This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.
        /// </summary>
        public readonly string? ScaleDownMode;
        /// <summary>
        /// This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is 'Delete'.
        /// </summary>
        public readonly string? ScaleSetEvictionPolicy;
        /// <summary>
        /// The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.
        /// </summary>
        public readonly string? ScaleSetPriority;
        /// <summary>
        /// Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any on-demand price. For more details on spot pricing, see [spot VMs pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)
        /// </summary>
        public readonly double? SpotMaxPrice;
        /// <summary>
        /// The tags to be persisted on the agent pool virtual machine scale set.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Settings for upgrading the agentpool
        /// </summary>
        public readonly Outputs.AgentPoolUpgradeSettingsResponse? UpgradeSettings;
        /// <summary>
        /// VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions
        /// </summary>
        public readonly string? VmSize;
        /// <summary>
        /// If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
        /// </summary>
        public readonly string? VnetSubnetID;
        /// <summary>
        /// Determines the type of workload a node can run.
        /// </summary>
        public readonly string? WorkloadRuntime;

        [OutputConstructor]
        private GetAgentPoolResult(
            ImmutableArray<string> availabilityZones,

            string? capacityReservationGroupID,

            int? count,

            Outputs.CreationDataResponse? creationData,

            string currentOrchestratorVersion,

            bool? enableAutoScaling,

            bool? enableEncryptionAtHost,

            bool? enableFIPS,

            bool? enableNodePublicIP,

            bool? enableUltraSSD,

            string? gpuInstanceProfile,

            string? hostGroupID,

            string id,

            Outputs.KubeletConfigResponse? kubeletConfig,

            string? kubeletDiskType,

            Outputs.LinuxOSConfigResponse? linuxOSConfig,

            int? maxCount,

            int? maxPods,

            int? minCount,

            string? mode,

            string name,

            Outputs.AgentPoolNetworkProfileResponse? networkProfile,

            string nodeImageVersion,

            ImmutableDictionary<string, string>? nodeLabels,

            string? nodePublicIPPrefixID,

            ImmutableArray<string> nodeTaints,

            string? orchestratorVersion,

            int? osDiskSizeGB,

            string? osDiskType,

            string? osSKU,

            string? osType,

            string? podSubnetID,

            Outputs.PowerStateResponse? powerState,

            string provisioningState,

            string? proximityPlacementGroupID,

            string? scaleDownMode,

            string? scaleSetEvictionPolicy,

            string? scaleSetPriority,

            double? spotMaxPrice,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.AgentPoolUpgradeSettingsResponse? upgradeSettings,

            string? vmSize,

            string? vnetSubnetID,

            string? workloadRuntime)
        {
            AvailabilityZones = availabilityZones;
            CapacityReservationGroupID = capacityReservationGroupID;
            Count = count;
            CreationData = creationData;
            CurrentOrchestratorVersion = currentOrchestratorVersion;
            EnableAutoScaling = enableAutoScaling;
            EnableEncryptionAtHost = enableEncryptionAtHost;
            EnableFIPS = enableFIPS;
            EnableNodePublicIP = enableNodePublicIP;
            EnableUltraSSD = enableUltraSSD;
            GpuInstanceProfile = gpuInstanceProfile;
            HostGroupID = hostGroupID;
            Id = id;
            KubeletConfig = kubeletConfig;
            KubeletDiskType = kubeletDiskType;
            LinuxOSConfig = linuxOSConfig;
            MaxCount = maxCount;
            MaxPods = maxPods;
            MinCount = minCount;
            Mode = mode;
            Name = name;
            NetworkProfile = networkProfile;
            NodeImageVersion = nodeImageVersion;
            NodeLabels = nodeLabels;
            NodePublicIPPrefixID = nodePublicIPPrefixID;
            NodeTaints = nodeTaints;
            OrchestratorVersion = orchestratorVersion;
            OsDiskSizeGB = osDiskSizeGB;
            OsDiskType = osDiskType;
            OsSKU = osSKU;
            OsType = osType;
            PodSubnetID = podSubnetID;
            PowerState = powerState;
            ProvisioningState = provisioningState;
            ProximityPlacementGroupID = proximityPlacementGroupID;
            ScaleDownMode = scaleDownMode;
            ScaleSetEvictionPolicy = scaleSetEvictionPolicy;
            ScaleSetPriority = scaleSetPriority;
            SpotMaxPrice = spotMaxPrice;
            Tags = tags;
            Type = type;
            UpgradeSettings = upgradeSettings;
            VmSize = vmSize;
            VnetSubnetID = vnetSubnetID;
            WorkloadRuntime = workloadRuntime;
        }
    }
}
