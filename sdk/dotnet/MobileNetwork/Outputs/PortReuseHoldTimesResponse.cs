// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MobileNetwork.Outputs
{

    /// <summary>
    /// The minimum time (in seconds) that will pass before a port that was used by a closed pinhole can be recycled for use by another pinhole. All hold times must be minimum 1 second.
    /// </summary>
    [OutputType]
    public sealed class PortReuseHoldTimesResponse
    {
        /// <summary>
        /// Minimum time in seconds that will pass before a TCP port that was used by a closed pinhole can be reused. Default for TCP is 2 minutes.
        /// </summary>
        public readonly int? Tcp;
        /// <summary>
        /// Minimum time in seconds that will pass before a UDP port that was used by a closed pinhole can be reused. Default for UDP is 1 minute.
        /// </summary>
        public readonly int? Udp;

        [OutputConstructor]
        private PortReuseHoldTimesResponse(
            int? tcp,

            int? udp)
        {
            Tcp = tcp;
            Udp = udp;
        }
    }
}
