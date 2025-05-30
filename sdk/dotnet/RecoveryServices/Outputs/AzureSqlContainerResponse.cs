// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.Outputs
{

    /// <summary>
    /// Azure Sql workload-specific container.
    /// </summary>
    [OutputType]
    public sealed class AzureSqlContainerResponse
    {
        /// <summary>
        /// Type of backup management for the container.
        /// </summary>
        public readonly string? BackupManagementType;
        /// <summary>
        /// Type of the container. The value of this property for: 1. Compute Azure VM is Microsoft.Compute/virtualMachines 2.
        /// Classic Compute Azure VM is Microsoft.ClassicCompute/virtualMachines 3. Windows machines (like MAB, DPM etc) is
        /// Windows 4. Azure SQL instance is AzureSqlContainer. 5. Storage containers is StorageContainer. 6. Azure workload
        /// Backup is VMAppContainer
        /// Expected value is 'AzureSqlContainer'.
        /// </summary>
        public readonly string ContainerType;
        /// <summary>
        /// Friendly name of the container.
        /// </summary>
        public readonly string? FriendlyName;
        /// <summary>
        /// Status of health of the container.
        /// </summary>
        public readonly string? HealthStatus;
        /// <summary>
        /// Type of the protectable object associated with this container
        /// </summary>
        public readonly string? ProtectableObjectType;
        /// <summary>
        /// Status of registration of the container with the Recovery Services Vault.
        /// </summary>
        public readonly string? RegistrationStatus;

        [OutputConstructor]
        private AzureSqlContainerResponse(
            string? backupManagementType,

            string containerType,

            string? friendlyName,

            string? healthStatus,

            string? protectableObjectType,

            string? registrationStatus)
        {
            BackupManagementType = backupManagementType;
            ContainerType = containerType;
            FriendlyName = friendlyName;
            HealthStatus = healthStatus;
            ProtectableObjectType = protectableObjectType;
            RegistrationStatus = registrationStatus;
        }
    }
}
