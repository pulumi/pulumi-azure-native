// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20240301.Inputs
{

    /// <summary>
    /// Connectivity group item.
    /// </summary>
    public sealed class ConnectivityGroupItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Group connectivity type.
        /// </summary>
        [Input("groupConnectivity", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.V20240301.GroupConnectivity> GroupConnectivity { get; set; } = null!;

        /// <summary>
        /// Flag if global is supported.
        /// </summary>
        [Input("isGlobal")]
        public InputUnion<string, Pulumi.AzureNative.Network.V20240301.IsGlobal>? IsGlobal { get; set; }

        /// <summary>
        /// Network group Id.
        /// </summary>
        [Input("networkGroupId", required: true)]
        public Input<string> NetworkGroupId { get; set; } = null!;

        /// <summary>
        /// Flag if need to use hub gateway.
        /// </summary>
        [Input("useHubGateway")]
        public InputUnion<string, Pulumi.AzureNative.Network.V20240301.UseHubGateway>? UseHubGateway { get; set; }

        public ConnectivityGroupItemArgs()
        {
        }
        public static new ConnectivityGroupItemArgs Empty => new ConnectivityGroupItemArgs();
    }
}
