// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Inputs
{

    /// <summary>
    /// AzureBackupServer (DPMVenus) workload-specific protection container.
    /// </summary>
    public sealed class AzureBackupServerContainerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of backup management for the container.
        /// </summary>
        [Input("backupManagementType")]
        public InputUnion<string, Pulumi.AzureNative.RecoveryServices.BackupManagementType>? BackupManagementType { get; set; }

        /// <summary>
        /// Specifies whether the container is re-registrable.
        /// </summary>
        [Input("canReRegister")]
        public Input<bool>? CanReRegister { get; set; }

        /// <summary>
        /// ID of container.
        /// </summary>
        [Input("containerId")]
        public Input<string>? ContainerId { get; set; }

        /// <summary>
        /// Type of the container. The value of this property for: 1. Compute Azure VM is Microsoft.Compute/virtualMachines 2.
        /// Classic Compute Azure VM is Microsoft.ClassicCompute/virtualMachines 3. Windows machines (like MAB, DPM etc) is
        /// Windows 4. Azure SQL instance is AzureSqlContainer. 5. Storage containers is StorageContainer. 6. Azure workload
        /// Backup is VMAppContainer
        /// Expected value is 'AzureBackupServerContainer'.
        /// </summary>
        [Input("containerType", required: true)]
        public Input<string> ContainerType { get; set; } = null!;

        /// <summary>
        /// Backup engine Agent version
        /// </summary>
        [Input("dpmAgentVersion")]
        public Input<string>? DpmAgentVersion { get; set; }

        [Input("dpmServers")]
        private InputList<string>? _dpmServers;

        /// <summary>
        /// List of BackupEngines protecting the container
        /// </summary>
        public InputList<string> DpmServers
        {
            get => _dpmServers ?? (_dpmServers = new InputList<string>());
            set => _dpmServers = value;
        }

        /// <summary>
        /// Extended Info of the container.
        /// </summary>
        [Input("extendedInfo")]
        public Input<Inputs.DPMContainerExtendedInfoArgs>? ExtendedInfo { get; set; }

        /// <summary>
        /// Friendly name of the container.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// Status of health of the container.
        /// </summary>
        [Input("healthStatus")]
        public Input<string>? HealthStatus { get; set; }

        /// <summary>
        /// Type of the protectable object associated with this container
        /// </summary>
        [Input("protectableObjectType")]
        public Input<string>? ProtectableObjectType { get; set; }

        /// <summary>
        /// Number of protected items in the BackupEngine
        /// </summary>
        [Input("protectedItemCount")]
        public Input<double>? ProtectedItemCount { get; set; }

        /// <summary>
        /// Protection status of the container.
        /// </summary>
        [Input("protectionStatus")]
        public Input<string>? ProtectionStatus { get; set; }

        /// <summary>
        /// Status of registration of the container with the Recovery Services Vault.
        /// </summary>
        [Input("registrationStatus")]
        public Input<string>? RegistrationStatus { get; set; }

        /// <summary>
        /// To check if upgrade available
        /// </summary>
        [Input("upgradeAvailable")]
        public Input<bool>? UpgradeAvailable { get; set; }

        public AzureBackupServerContainerArgs()
        {
        }
        public static new AzureBackupServerContainerArgs Empty => new AzureBackupServerContainerArgs();
    }
}
