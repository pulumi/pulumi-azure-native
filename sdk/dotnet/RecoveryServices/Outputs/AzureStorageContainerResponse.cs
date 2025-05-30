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
    /// Azure Storage Account workload-specific container.
    /// </summary>
    [OutputType]
    public sealed class AzureStorageContainerResponse
    {
        /// <summary>
        /// Whether storage account lock is to be acquired for this container or not.
        /// </summary>
        public readonly string? AcquireStorageAccountLock;
        /// <summary>
        /// Type of backup management for the container.
        /// </summary>
        public readonly string? BackupManagementType;
        /// <summary>
        /// Type of the container. The value of this property for: 1. Compute Azure VM is Microsoft.Compute/virtualMachines 2.
        /// Classic Compute Azure VM is Microsoft.ClassicCompute/virtualMachines 3. Windows machines (like MAB, DPM etc) is
        /// Windows 4. Azure SQL instance is AzureSqlContainer. 5. Storage containers is StorageContainer. 6. Azure workload
        /// Backup is VMAppContainer
        /// Expected value is 'StorageContainer'.
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
        /// Re-Do Operation
        /// </summary>
        public readonly string? OperationType;
        /// <summary>
        /// Type of the protectable object associated with this container
        /// </summary>
        public readonly string? ProtectableObjectType;
        /// <summary>
        /// Number of items backed up in this container.
        /// </summary>
        public readonly double? ProtectedItemCount;
        /// <summary>
        /// Status of registration of the container with the Recovery Services Vault.
        /// </summary>
        public readonly string? RegistrationStatus;
        /// <summary>
        /// Resource group name of Recovery Services Vault.
        /// </summary>
        public readonly string? ResourceGroup;
        /// <summary>
        /// Fully qualified ARM url.
        /// </summary>
        public readonly string? SourceResourceId;
        /// <summary>
        /// Storage account version.
        /// </summary>
        public readonly string? StorageAccountVersion;

        [OutputConstructor]
        private AzureStorageContainerResponse(
            string? acquireStorageAccountLock,

            string? backupManagementType,

            string containerType,

            string? friendlyName,

            string? healthStatus,

            string? operationType,

            string? protectableObjectType,

            double? protectedItemCount,

            string? registrationStatus,

            string? resourceGroup,

            string? sourceResourceId,

            string? storageAccountVersion)
        {
            AcquireStorageAccountLock = acquireStorageAccountLock;
            BackupManagementType = backupManagementType;
            ContainerType = containerType;
            FriendlyName = friendlyName;
            HealthStatus = healthStatus;
            OperationType = operationType;
            ProtectableObjectType = protectableObjectType;
            ProtectedItemCount = protectedItemCount;
            RegistrationStatus = registrationStatus;
            ResourceGroup = resourceGroup;
            SourceResourceId = sourceResourceId;
            StorageAccountVersion = storageAccountVersion;
        }
    }
}
