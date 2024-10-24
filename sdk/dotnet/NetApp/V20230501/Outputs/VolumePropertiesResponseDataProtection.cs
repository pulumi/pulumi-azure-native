// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetApp.V20230501.Outputs
{

    /// <summary>
    /// DataProtection type volumes include an object containing details of the replication
    /// </summary>
    [OutputType]
    public sealed class VolumePropertiesResponseDataProtection
    {
        /// <summary>
        /// Replication properties
        /// </summary>
        public readonly Outputs.ReplicationObjectResponse? Replication;
        /// <summary>
        /// Snapshot properties.
        /// </summary>
        public readonly Outputs.VolumeSnapshotPropertiesResponse? Snapshot;
        /// <summary>
        /// VolumeRelocation properties
        /// </summary>
        public readonly Outputs.VolumeRelocationPropertiesResponse? VolumeRelocation;

        [OutputConstructor]
        private VolumePropertiesResponseDataProtection(
            Outputs.ReplicationObjectResponse? replication,

            Outputs.VolumeSnapshotPropertiesResponse? snapshot,

            Outputs.VolumeRelocationPropertiesResponse? volumeRelocation)
        {
            Replication = replication;
            Snapshot = snapshot;
            VolumeRelocation = volumeRelocation;
        }
    }
}
