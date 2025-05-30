// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Orbital.Outputs
{

    /// <summary>
    /// List of authorized spacecraft links per ground station and the expiration date of the authorization.
    /// </summary>
    [OutputType]
    public sealed class SpacecraftLinkResponse
    {
        /// <summary>
        /// Authorized Ground Stations
        /// </summary>
        public readonly ImmutableArray<Outputs.AuthorizedGroundstationResponse> Authorizations;
        /// <summary>
        /// Bandwidth in MHz.
        /// </summary>
        public readonly double BandwidthMHz;
        /// <summary>
        /// Center Frequency in MHz.
        /// </summary>
        public readonly double CenterFrequencyMHz;
        /// <summary>
        /// Direction (Uplink or Downlink).
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// Link name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Polarization. e.g. (RHCP, LHCP).
        /// </summary>
        public readonly string Polarization;

        [OutputConstructor]
        private SpacecraftLinkResponse(
            ImmutableArray<Outputs.AuthorizedGroundstationResponse> authorizations,

            double bandwidthMHz,

            double centerFrequencyMHz,

            string direction,

            string name,

            string polarization)
        {
            Authorizations = authorizations;
            BandwidthMHz = bandwidthMHz;
            CenterFrequencyMHz = centerFrequencyMHz;
            Direction = direction;
            Name = name;
            Polarization = polarization;
        }
    }
}
