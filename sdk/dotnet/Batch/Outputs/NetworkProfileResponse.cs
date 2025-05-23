// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.Outputs
{

    /// <summary>
    /// Network profile for Batch account, which contains network rule settings for each endpoint.
    /// </summary>
    [OutputType]
    public sealed class NetworkProfileResponse
    {
        /// <summary>
        /// Network access profile for batchAccount endpoint (Batch account data plane API).
        /// </summary>
        public readonly Outputs.EndpointAccessProfileResponse? AccountAccess;
        /// <summary>
        /// Network access profile for nodeManagement endpoint (Batch service managing compute nodes for Batch pools).
        /// </summary>
        public readonly Outputs.EndpointAccessProfileResponse? NodeManagementAccess;

        [OutputConstructor]
        private NetworkProfileResponse(
            Outputs.EndpointAccessProfileResponse? accountAccess,

            Outputs.EndpointAccessProfileResponse? nodeManagementAccess)
        {
            AccountAccess = accountAccess;
            NodeManagementAccess = nodeManagementAccess;
        }
    }
}
