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
    /// The private IP addresses/IP ranges to which traffic will not be SNAT.
    /// </summary>
    public sealed class FirewallPolicySNATArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The operation mode for automatically learning private ranges to not be SNAT
        /// </summary>
        [Input("autoLearnPrivateRanges")]
        public InputUnion<string, Pulumi.AzureNative.Network.AutoLearnPrivateRangesMode>? AutoLearnPrivateRanges { get; set; }

        [Input("privateRanges")]
        private InputList<string>? _privateRanges;

        /// <summary>
        /// List of private IP addresses/IP address ranges to not be SNAT.
        /// </summary>
        public InputList<string> PrivateRanges
        {
            get => _privateRanges ?? (_privateRanges = new InputList<string>());
            set => _privateRanges = value;
        }

        public FirewallPolicySNATArgs()
        {
        }
        public static new FirewallPolicySNATArgs Empty => new FirewallPolicySNATArgs();
    }
}
