// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Orbital.Inputs
{

    /// <summary>
    /// List of authorized spacecraft links per ground station and the expiration date of the authorization.
    /// </summary>
    public sealed class SpacecraftLinkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bandwidth in MHz.
        /// </summary>
        [Input("bandwidthMHz", required: true)]
        public Input<double> BandwidthMHz { get; set; } = null!;

        /// <summary>
        /// Center Frequency in MHz.
        /// </summary>
        [Input("centerFrequencyMHz", required: true)]
        public Input<double> CenterFrequencyMHz { get; set; } = null!;

        /// <summary>
        /// Direction (Uplink or Downlink).
        /// </summary>
        [Input("direction", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Orbital.Direction> Direction { get; set; } = null!;

        /// <summary>
        /// Link name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Polarization. e.g. (RHCP, LHCP).
        /// </summary>
        [Input("polarization", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Orbital.Polarization> Polarization { get; set; } = null!;

        public SpacecraftLinkArgs()
        {
        }
        public static new SpacecraftLinkArgs Empty => new SpacecraftLinkArgs();
    }
}
