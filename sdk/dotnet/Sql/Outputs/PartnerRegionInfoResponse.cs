// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql.Outputs
{

    /// <summary>
    /// Partner region information for the failover group.
    /// </summary>
    [OutputType]
    public sealed class PartnerRegionInfoResponse
    {
        /// <summary>
        /// Geo location of the partner managed instances.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Replication role of the partner managed instances.
        /// </summary>
        public readonly string ReplicationRole;

        [OutputConstructor]
        private PartnerRegionInfoResponse(
            string? location,

            string replicationRole)
        {
            Location = location;
            ReplicationRole = replicationRole;
        }
    }
}
