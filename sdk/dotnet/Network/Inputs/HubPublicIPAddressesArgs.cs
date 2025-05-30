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
    /// Public IP addresses associated with azure firewall.
    /// </summary>
    public sealed class HubPublicIPAddressesArgs : global::Pulumi.ResourceArgs
    {
        [Input("addresses")]
        private InputList<Inputs.AzureFirewallPublicIPAddressArgs>? _addresses;

        /// <summary>
        /// The list of Public IP addresses associated with azure firewall or IP addresses to be retained.
        /// </summary>
        public InputList<Inputs.AzureFirewallPublicIPAddressArgs> Addresses
        {
            get => _addresses ?? (_addresses = new InputList<Inputs.AzureFirewallPublicIPAddressArgs>());
            set => _addresses = value;
        }

        /// <summary>
        /// The number of Public IP addresses associated with azure firewall.
        /// </summary>
        [Input("count")]
        public Input<int>? Count { get; set; }

        public HubPublicIPAddressesArgs()
        {
        }
        public static new HubPublicIPAddressesArgs Empty => new HubPublicIPAddressesArgs();
    }
}
