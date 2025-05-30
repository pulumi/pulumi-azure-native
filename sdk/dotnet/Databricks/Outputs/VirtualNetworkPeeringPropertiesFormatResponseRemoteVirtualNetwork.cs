// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Databricks.Outputs
{

    /// <summary>
    ///  The remote virtual network should be in the same region. See here to learn more (https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/vnet-peering).
    /// </summary>
    [OutputType]
    public sealed class VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork
    {
        /// <summary>
        /// The Id of the remote virtual network.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork(string? id)
        {
            Id = id;
        }
    }
}
