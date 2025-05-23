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
    /// Quality of Service defines the traffic configuration between endpoints. Mandatory to have one marking.
    /// </summary>
    [OutputType]
    public sealed class QosDefinitionResponse
    {
        /// <summary>
        /// Destination IP ranges.
        /// </summary>
        public readonly ImmutableArray<Outputs.QosIpRangeResponse> DestinationIpRanges;
        /// <summary>
        /// Destination port ranges.
        /// </summary>
        public readonly ImmutableArray<Outputs.QosPortRangeResponse> DestinationPortRanges;
        /// <summary>
        /// List of markings to be used in the configuration.
        /// </summary>
        public readonly ImmutableArray<int> Markings;
        /// <summary>
        /// RNM supported protocol types.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// Source IP ranges.
        /// </summary>
        public readonly ImmutableArray<Outputs.QosIpRangeResponse> SourceIpRanges;
        /// <summary>
        /// Sources port ranges.
        /// </summary>
        public readonly ImmutableArray<Outputs.QosPortRangeResponse> SourcePortRanges;

        [OutputConstructor]
        private QosDefinitionResponse(
            ImmutableArray<Outputs.QosIpRangeResponse> destinationIpRanges,

            ImmutableArray<Outputs.QosPortRangeResponse> destinationPortRanges,

            ImmutableArray<int> markings,

            string? protocol,

            ImmutableArray<Outputs.QosIpRangeResponse> sourceIpRanges,

            ImmutableArray<Outputs.QosPortRangeResponse> sourcePortRanges)
        {
            DestinationIpRanges = destinationIpRanges;
            DestinationPortRanges = destinationPortRanges;
            Markings = markings;
            Protocol = protocol;
            SourceIpRanges = sourceIpRanges;
            SourcePortRanges = sourcePortRanges;
        }
    }
}
