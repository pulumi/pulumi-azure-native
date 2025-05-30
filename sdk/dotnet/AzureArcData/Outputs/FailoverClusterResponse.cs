// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData.Outputs
{

    /// <summary>
    /// Failover Cluster Instance properties.
    /// </summary>
    [OutputType]
    public sealed class FailoverClusterResponse
    {
        /// <summary>
        /// The host names which are part of the SQL FCI resource group.
        /// </summary>
        public readonly ImmutableArray<string> HostNames;
        /// <summary>
        /// The GUID of the SQL Server's underlying Failover Cluster.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The network name to connect to the SQL FCI.
        /// </summary>
        public readonly string NetworkName;
        /// <summary>
        /// The ARM IDs of the Arc SQL Server resources, belonging to the current server's Failover cluster.
        /// </summary>
        public readonly ImmutableArray<string> SqlInstanceIds;

        [OutputConstructor]
        private FailoverClusterResponse(
            ImmutableArray<string> hostNames,

            string id,

            string networkName,

            ImmutableArray<string> sqlInstanceIds)
        {
            HostNames = hostNames;
            Id = id;
            NetworkName = networkName;
            SqlInstanceIds = sqlInstanceIds;
        }
    }
}
