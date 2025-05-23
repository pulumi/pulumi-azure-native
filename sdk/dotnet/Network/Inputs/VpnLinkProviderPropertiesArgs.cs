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
    /// List of properties of a link provider.
    /// </summary>
    public sealed class VpnLinkProviderPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the link provider.
        /// </summary>
        [Input("linkProviderName")]
        public Input<string>? LinkProviderName { get; set; }

        /// <summary>
        /// Link speed.
        /// </summary>
        [Input("linkSpeedInMbps")]
        public Input<int>? LinkSpeedInMbps { get; set; }

        public VpnLinkProviderPropertiesArgs()
        {
        }
        public static new VpnLinkProviderPropertiesArgs Empty => new VpnLinkProviderPropertiesArgs();
    }
}
