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
    /// A2A provider specific settings.
    /// </summary>
    [OutputType]
    public sealed class A2ACrossClusterMigrationReplicationDetailsResponse
    {
        /// <summary>
        /// The fabric specific object Id of the virtual machine.
        /// </summary>
        public readonly string? FabricObjectId;
        /// <summary>
        /// Gets the Instance type.
        /// Expected value is 'A2ACrossClusterMigration'.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// An id associated with the PE that survives actions like switch protection which change the backing PE/CPE objects internally.The lifecycle id gets carried forward to have a link/continuity in being able to have an Id that denotes the "same" protected item even though other internal Ids/ARM Id might be changing.
        /// </summary>
        public readonly string? LifecycleId;
        /// <summary>
        /// The type of operating system.
        /// </summary>
        public readonly string? OsType;
        /// <summary>
        /// Primary fabric location.
        /// </summary>
        public readonly string? PrimaryFabricLocation;
        /// <summary>
        /// The protection state for the vm.
        /// </summary>
        public readonly string? VmProtectionState;
        /// <summary>
        /// The protection state description for the vm.
        /// </summary>
        public readonly string? VmProtectionStateDescription;

        [OutputConstructor]
        private A2ACrossClusterMigrationReplicationDetailsResponse(
            string? fabricObjectId,

            string instanceType,

            string? lifecycleId,

            string? osType,

            string? primaryFabricLocation,

            string? vmProtectionState,

            string? vmProtectionStateDescription)
        {
            FabricObjectId = fabricObjectId;
            InstanceType = instanceType;
            LifecycleId = lifecycleId;
            OsType = osType;
            PrimaryFabricLocation = primaryFabricLocation;
            VmProtectionState = vmProtectionState;
            VmProtectionStateDescription = vmProtectionStateDescription;
        }
    }
}
