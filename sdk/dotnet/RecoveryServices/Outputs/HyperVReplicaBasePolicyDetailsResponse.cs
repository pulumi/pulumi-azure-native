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
    /// Base class for HyperVReplica policy details.
    /// </summary>
    [OutputType]
    public sealed class HyperVReplicaBasePolicyDetailsResponse
    {
        /// <summary>
        /// A value indicating the authentication type.
        /// </summary>
        public readonly int? AllowedAuthenticationType;
        /// <summary>
        /// A value indicating the application consistent frequency.
        /// </summary>
        public readonly int? ApplicationConsistentSnapshotFrequencyInHours;
        /// <summary>
        /// A value indicating whether compression has to be enabled.
        /// </summary>
        public readonly string? Compression;
        /// <summary>
        /// A value indicating whether IR is online.
        /// </summary>
        public readonly string? InitialReplicationMethod;
        /// <summary>
        /// Gets the class type. Overridden in derived classes.
        /// Expected value is 'HyperVReplicaBasePolicyDetails'.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// A value indicating the offline IR export path.
        /// </summary>
        public readonly string? OfflineReplicationExportPath;
        /// <summary>
        /// A value indicating the offline IR import path.
        /// </summary>
        public readonly string? OfflineReplicationImportPath;
        /// <summary>
        /// A value indicating the online IR start time.
        /// </summary>
        public readonly string? OnlineReplicationStartTime;
        /// <summary>
        /// A value indicating the number of recovery points.
        /// </summary>
        public readonly int? RecoveryPoints;
        /// <summary>
        /// A value indicating whether the VM has to be auto deleted. Supported Values: String.Empty, None, OnRecoveryCloud.
        /// </summary>
        public readonly string? ReplicaDeletionOption;
        /// <summary>
        /// A value indicating the recovery HTTPS port.
        /// </summary>
        public readonly int? ReplicationPort;

        [OutputConstructor]
        private HyperVReplicaBasePolicyDetailsResponse(
            int? allowedAuthenticationType,

            int? applicationConsistentSnapshotFrequencyInHours,

            string? compression,

            string? initialReplicationMethod,

            string instanceType,

            string? offlineReplicationExportPath,

            string? offlineReplicationImportPath,

            string? onlineReplicationStartTime,

            int? recoveryPoints,

            string? replicaDeletionOption,

            int? replicationPort)
        {
            AllowedAuthenticationType = allowedAuthenticationType;
            ApplicationConsistentSnapshotFrequencyInHours = applicationConsistentSnapshotFrequencyInHours;
            Compression = compression;
            InitialReplicationMethod = initialReplicationMethod;
            InstanceType = instanceType;
            OfflineReplicationExportPath = offlineReplicationExportPath;
            OfflineReplicationImportPath = offlineReplicationImportPath;
            OnlineReplicationStartTime = onlineReplicationStartTime;
            RecoveryPoints = recoveryPoints;
            ReplicaDeletionOption = replicaDeletionOption;
            ReplicationPort = replicationPort;
        }
    }
}
