// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DocumentDB.V20240301Preview.Outputs
{

    /// <summary>
    /// The properties of a mongo cluster.
    /// </summary>
    [OutputType]
    public sealed class MongoClusterPropertiesResponse
    {
        /// <summary>
        /// The administrator's login for the mongo cluster.
        /// </summary>
        public readonly string? AdministratorLogin;
        /// <summary>
        /// The status of the mongo cluster.
        /// </summary>
        public readonly string ClusterStatus;
        /// <summary>
        /// The default mongo connection string for the cluster.
        /// </summary>
        public readonly string ConnectionString;
        /// <summary>
        /// Earliest restore timestamp in UTC ISO8601 format.
        /// </summary>
        public readonly string EarliestRestoreTime;
        /// <summary>
        /// The list of node group specs in the cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.NodeGroupSpecResponse> NodeGroupSpecs;
        /// <summary>
        /// List of private endpoint connections.
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionResponse> PrivateEndpointConnections;
        /// <summary>
        /// The provisioning state of the mongo cluster.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Whether or not public endpoint access is allowed for this mongo cluster.
        /// </summary>
        public readonly string? PublicNetworkAccess;
        /// <summary>
        /// The Mongo DB server version. Defaults to the latest available version if not specified.
        /// </summary>
        public readonly string? ServerVersion;

        [OutputConstructor]
        private MongoClusterPropertiesResponse(
            string? administratorLogin,

            string clusterStatus,

            string connectionString,

            string earliestRestoreTime,

            ImmutableArray<Outputs.NodeGroupSpecResponse> nodeGroupSpecs,

            ImmutableArray<Outputs.PrivateEndpointConnectionResponse> privateEndpointConnections,

            string provisioningState,

            string? publicNetworkAccess,

            string? serverVersion)
        {
            AdministratorLogin = administratorLogin;
            ClusterStatus = clusterStatus;
            ConnectionString = connectionString;
            EarliestRestoreTime = earliestRestoreTime;
            NodeGroupSpecs = nodeGroupSpecs;
            PrivateEndpointConnections = privateEndpointConnections;
            ProvisioningState = provisioningState;
            PublicNetworkAccess = publicNetworkAccess;
            ServerVersion = serverVersion;
        }
    }
}
