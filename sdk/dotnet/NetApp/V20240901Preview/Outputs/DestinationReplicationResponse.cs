// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NetApp.V20240901Preview.Outputs
{

    /// <summary>
    /// Destination replication properties
    /// </summary>
    [OutputType]
    public sealed class DestinationReplicationResponse
    {
        /// <summary>
        /// The remote region for the destination volume.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// Indicates whether the replication is cross zone or cross region.
        /// </summary>
        public readonly string? ReplicationType;
        /// <summary>
        /// The resource ID of the remote volume
        /// </summary>
        public readonly string? ResourceId;
        /// <summary>
        /// The remote zone for the destination volume.
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private DestinationReplicationResponse(
            string? region,

            string? replicationType,

            string? resourceId,

            string? zone)
        {
            Region = region;
            ReplicationType = replicationType;
            ResourceId = resourceId;
            Zone = zone;
        }
    }
}
