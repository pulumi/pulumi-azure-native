// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AVS.Outputs
{

    /// <summary>
    /// The properties of a management cluster
    /// </summary>
    [OutputType]
    public sealed class ManagementClusterResponse
    {
        /// <summary>
        /// The identity
        /// </summary>
        public readonly int ClusterId;
        /// <summary>
        /// The cluster size
        /// </summary>
        public readonly int? ClusterSize;
        /// <summary>
        /// The hosts
        /// </summary>
        public readonly ImmutableArray<string> Hosts;
        /// <summary>
        /// The state of the cluster provisioning
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Name of the vsan datastore associated with the cluster
        /// </summary>
        public readonly string? VsanDatastoreName;

        [OutputConstructor]
        private ManagementClusterResponse(
            int clusterId,

            int? clusterSize,

            ImmutableArray<string> hosts,

            string provisioningState,

            string? vsanDatastoreName)
        {
            ClusterId = clusterId;
            ClusterSize = clusterSize;
            Hosts = hosts;
            ProvisioningState = provisioningState;
            VsanDatastoreName = vsanDatastoreName;
        }
    }
}
