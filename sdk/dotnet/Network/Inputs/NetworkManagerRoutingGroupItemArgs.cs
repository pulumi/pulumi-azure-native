// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Inputs
{

    /// <summary>
    /// Network manager routing group item.
    /// </summary>
    public sealed class NetworkManagerRoutingGroupItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Network manager group Id.
        /// </summary>
        [Input("networkGroupId", required: true)]
        public Input<string> NetworkGroupId { get; set; } = null!;

        public NetworkManagerRoutingGroupItemArgs()
        {
        }
        public static new NetworkManagerRoutingGroupItemArgs Empty => new NetworkManagerRoutingGroupItemArgs();
    }
}
