// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric.Inputs
{

    /// <summary>
    /// Neighbor Address properties.
    /// </summary>
    public sealed class NeighborAddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP Address.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        public NeighborAddressArgs()
        {
        }
        public static new NeighborAddressArgs Empty => new NeighborAddressArgs();
    }
}
