// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataReplication.Outputs
{

    /// <summary>
    /// VMware to AzStackHCI Protected item model custom properties.
    /// </summary>
    [OutputType]
    public sealed class VMwareToAzStackHCIProtectedItemModelCustomPropertiesResponse
    {
        /// <summary>
        /// Gets or sets the location of the protected item.
        /// </summary>
        public readonly string ActiveLocation;
        /// <summary>
        /// Gets or sets the location of Azure Arc HCI custom location resource.
        /// </summary>
        public readonly string CustomLocationRegion;
        /// <summary>
        /// Gets or sets the list of disks to replicate.
        /// </summary>
        public readonly ImmutableArray<Outputs.VMwareToAzStackHCIDiskInputResponse> DisksToInclude;
        /// <summary>
        /// Protected item dynamic memory config.
        /// </summary>
        public readonly Outputs.ProtectedItemDynamicMemoryConfigResponse? DynamicMemoryConfig;
        /// <summary>
        /// Gets or sets the ARM Id of the discovered machine.
        /// </summary>
        public readonly string FabricDiscoveryMachineId;
        /// <summary>
        /// Gets or sets the recovery point Id to which the VM was failed over.
        /// </summary>
        public readonly string FailoverRecoveryPointId;
        /// <summary>
        /// Gets or sets the firmware type.
        /// </summary>
        public readonly string FirmwareType;
        /// <summary>
        /// Gets or sets the hypervisor generation of the virtual machine possible values are 1,2.
        /// </summary>
        public readonly string HyperVGeneration;
        /// <summary>
        /// Gets or sets the initial replication progress percentage. This is calculated based on
        /// total bytes processed for all disks in the source VM.
        /// </summary>
        public readonly int InitialReplicationProgressPercentage;
        /// <summary>
        /// Gets or sets the instance type.
        /// Expected value is 'VMwareToAzStackHCI'.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// Gets or sets a value indicating whether memory is dynamical.
        /// </summary>
        public readonly bool? IsDynamicRam;
        /// <summary>
        /// Gets or sets the last recovery point Id.
        /// </summary>
        public readonly string LastRecoveryPointId;
        /// <summary>
        /// Gets or sets the last recovery point received time.
        /// </summary>
        public readonly string LastRecoveryPointReceived;
        /// <summary>
        /// Gets or sets the latest timestamp that replication status is updated.
        /// </summary>
        public readonly string LastReplicationUpdateTime;
        /// <summary>
        /// Gets or sets the migration progress percentage.
        /// </summary>
        public readonly int MigrationProgressPercentage;
        /// <summary>
        /// Gets or sets the list of VM NIC to replicate.
        /// </summary>
        public readonly ImmutableArray<Outputs.VMwareToAzStackHCINicInputResponse> NicsToInclude;
        /// <summary>
        /// Gets or sets the name of the OS.
        /// </summary>
        public readonly string OsName;
        /// <summary>
        /// Gets or sets the type of the OS.
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// Gets or sets a value indicating whether auto resync is to be done.
        /// </summary>
        public readonly bool? PerformAutoResync;
        /// <summary>
        /// Gets or sets the list of protected disks.
        /// </summary>
        public readonly ImmutableArray<Outputs.VMwareToAzStackHCIProtectedDiskPropertiesResponse> ProtectedDisks;
        /// <summary>
        /// Gets or sets the VM NIC details.
        /// </summary>
        public readonly ImmutableArray<Outputs.VMwareToAzStackHCIProtectedNicPropertiesResponse> ProtectedNics;
        /// <summary>
        /// Gets or sets the resume progress percentage.
        /// </summary>
        public readonly int ResumeProgressPercentage;
        /// <summary>
        /// Gets or sets the resume retry count.
        /// </summary>
        public readonly double ResumeRetryCount;
        /// <summary>
        /// Gets or sets the resync progress percentage. This is calculated based on total bytes
        /// processed for all disks in the source VM.
        /// </summary>
        public readonly int ResyncProgressPercentage;
        /// <summary>
        /// Gets or sets a value indicating whether resync is required.
        /// </summary>
        public readonly bool ResyncRequired;
        /// <summary>
        /// Gets or sets the resync retry count.
        /// </summary>
        public readonly double ResyncRetryCount;
        /// <summary>
        /// Gets or sets the resync state.
        /// </summary>
        public readonly string ResyncState;
        /// <summary>
        /// Gets or sets the run as account Id.
        /// </summary>
        public readonly string RunAsAccountId;
        /// <summary>
        /// Gets or sets the source appliance name.
        /// </summary>
        public readonly string SourceApplianceName;
        /// <summary>
        /// Gets or sets the source VM CPU cores.
        /// </summary>
        public readonly int SourceCpuCores;
        /// <summary>
        /// Gets or sets the source DRA name.
        /// </summary>
        public readonly string SourceDraName;
        /// <summary>
        /// Gets or sets the source VM ram memory size in megabytes.
        /// </summary>
        public readonly double SourceMemoryInMegaBytes;
        /// <summary>
        /// Gets or sets the source VM display name.
        /// </summary>
        public readonly string SourceVmName;
        /// <summary>
        /// Gets or sets the target storage container ARM Id.
        /// </summary>
        public readonly string StorageContainerId;
        /// <summary>
        /// Gets or sets the target appliance name.
        /// </summary>
        public readonly string TargetApplianceName;
        /// <summary>
        /// Gets or sets the Target Arc Cluster Custom Location ARM Id.
        /// </summary>
        public readonly string TargetArcClusterCustomLocationId;
        /// <summary>
        /// Gets or sets the Target AzStackHCI cluster name.
        /// </summary>
        public readonly string TargetAzStackHciClusterName;
        /// <summary>
        /// Gets or sets the target CPU cores.
        /// </summary>
        public readonly int? TargetCpuCores;
        /// <summary>
        /// Gets or sets the target DRA name.
        /// </summary>
        public readonly string TargetDraName;
        /// <summary>
        /// Gets or sets the Target HCI Cluster ARM Id.
        /// </summary>
        public readonly string TargetHciClusterId;
        /// <summary>
        /// Gets or sets the target location.
        /// </summary>
        public readonly string TargetLocation;
        /// <summary>
        /// Gets or sets the target memory in mega-bytes.
        /// </summary>
        public readonly int? TargetMemoryInMegaBytes;
        /// <summary>
        /// Gets or sets the target network Id within AzStackHCI Cluster.
        /// </summary>
        public readonly string? TargetNetworkId;
        /// <summary>
        /// Gets or sets the target resource group ARM Id.
        /// </summary>
        public readonly string TargetResourceGroupId;
        /// <summary>
        /// Gets or sets the BIOS Id of the target AzStackHCI VM.
        /// </summary>
        public readonly string TargetVmBiosId;
        /// <summary>
        /// Gets or sets the target VM display name.
        /// </summary>
        public readonly string? TargetVmName;
        /// <summary>
        /// Gets or sets the target test network Id within AzStackHCI Cluster.
        /// </summary>
        public readonly string? TestNetworkId;

        [OutputConstructor]
        private VMwareToAzStackHCIProtectedItemModelCustomPropertiesResponse(
            string activeLocation,

            string customLocationRegion,

            ImmutableArray<Outputs.VMwareToAzStackHCIDiskInputResponse> disksToInclude,

            Outputs.ProtectedItemDynamicMemoryConfigResponse? dynamicMemoryConfig,

            string fabricDiscoveryMachineId,

            string failoverRecoveryPointId,

            string firmwareType,

            string hyperVGeneration,

            int initialReplicationProgressPercentage,

            string instanceType,

            bool? isDynamicRam,

            string lastRecoveryPointId,

            string lastRecoveryPointReceived,

            string lastReplicationUpdateTime,

            int migrationProgressPercentage,

            ImmutableArray<Outputs.VMwareToAzStackHCINicInputResponse> nicsToInclude,

            string osName,

            string osType,

            bool? performAutoResync,

            ImmutableArray<Outputs.VMwareToAzStackHCIProtectedDiskPropertiesResponse> protectedDisks,

            ImmutableArray<Outputs.VMwareToAzStackHCIProtectedNicPropertiesResponse> protectedNics,

            int resumeProgressPercentage,

            double resumeRetryCount,

            int resyncProgressPercentage,

            bool resyncRequired,

            double resyncRetryCount,

            string resyncState,

            string runAsAccountId,

            string sourceApplianceName,

            int sourceCpuCores,

            string sourceDraName,

            double sourceMemoryInMegaBytes,

            string sourceVmName,

            string storageContainerId,

            string targetApplianceName,

            string targetArcClusterCustomLocationId,

            string targetAzStackHciClusterName,

            int? targetCpuCores,

            string targetDraName,

            string targetHciClusterId,

            string targetLocation,

            int? targetMemoryInMegaBytes,

            string? targetNetworkId,

            string targetResourceGroupId,

            string targetVmBiosId,

            string? targetVmName,

            string? testNetworkId)
        {
            ActiveLocation = activeLocation;
            CustomLocationRegion = customLocationRegion;
            DisksToInclude = disksToInclude;
            DynamicMemoryConfig = dynamicMemoryConfig;
            FabricDiscoveryMachineId = fabricDiscoveryMachineId;
            FailoverRecoveryPointId = failoverRecoveryPointId;
            FirmwareType = firmwareType;
            HyperVGeneration = hyperVGeneration;
            InitialReplicationProgressPercentage = initialReplicationProgressPercentage;
            InstanceType = instanceType;
            IsDynamicRam = isDynamicRam;
            LastRecoveryPointId = lastRecoveryPointId;
            LastRecoveryPointReceived = lastRecoveryPointReceived;
            LastReplicationUpdateTime = lastReplicationUpdateTime;
            MigrationProgressPercentage = migrationProgressPercentage;
            NicsToInclude = nicsToInclude;
            OsName = osName;
            OsType = osType;
            PerformAutoResync = performAutoResync;
            ProtectedDisks = protectedDisks;
            ProtectedNics = protectedNics;
            ResumeProgressPercentage = resumeProgressPercentage;
            ResumeRetryCount = resumeRetryCount;
            ResyncProgressPercentage = resyncProgressPercentage;
            ResyncRequired = resyncRequired;
            ResyncRetryCount = resyncRetryCount;
            ResyncState = resyncState;
            RunAsAccountId = runAsAccountId;
            SourceApplianceName = sourceApplianceName;
            SourceCpuCores = sourceCpuCores;
            SourceDraName = sourceDraName;
            SourceMemoryInMegaBytes = sourceMemoryInMegaBytes;
            SourceVmName = sourceVmName;
            StorageContainerId = storageContainerId;
            TargetApplianceName = targetApplianceName;
            TargetArcClusterCustomLocationId = targetArcClusterCustomLocationId;
            TargetAzStackHciClusterName = targetAzStackHciClusterName;
            TargetCpuCores = targetCpuCores;
            TargetDraName = targetDraName;
            TargetHciClusterId = targetHciClusterId;
            TargetLocation = targetLocation;
            TargetMemoryInMegaBytes = targetMemoryInMegaBytes;
            TargetNetworkId = targetNetworkId;
            TargetResourceGroupId = targetResourceGroupId;
            TargetVmBiosId = targetVmBiosId;
            TargetVmName = targetVmName;
            TestNetworkId = testNetworkId;
        }
    }
}
