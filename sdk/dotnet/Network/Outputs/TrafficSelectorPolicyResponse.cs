// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// An traffic selector policy for a virtual network gateway connection.
    /// </summary>
    [OutputType]
    public sealed class TrafficSelectorPolicyResponse
    {
        /// <summary>
        /// A collection of local address spaces in CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> LocalAddressRanges;
        /// <summary>
        /// A collection of remote address spaces in CIDR format.
        /// </summary>
        public readonly ImmutableArray<string> RemoteAddressRanges;

        [OutputConstructor]
        private TrafficSelectorPolicyResponse(
            ImmutableArray<string> localAddressRanges,

            ImmutableArray<string> remoteAddressRanges)
        {
            LocalAddressRanges = localAddressRanges;
            RemoteAddressRanges = remoteAddressRanges;
        }
    }
}
