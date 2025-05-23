// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media.Inputs
{

    /// <summary>
    /// Select audio tracks from the input by specifying a track identifier.
    /// </summary>
    public sealed class SelectAudioTrackByIdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional designation for single channel audio tracks.  Can be used to combine the tracks into stereo or multi-channel audio tracks.
        /// </summary>
        [Input("channelMapping")]
        public InputUnion<string, Pulumi.AzureNative.Media.ChannelMapping>? ChannelMapping { get; set; }

        /// <summary>
        /// The discriminator for derived types.
        /// Expected value is '#Microsoft.Media.SelectAudioTrackById'.
        /// </summary>
        [Input("odataType", required: true)]
        public Input<string> OdataType { get; set; } = null!;

        /// <summary>
        /// Track identifier to select
        /// </summary>
        [Input("trackId", required: true)]
        public Input<double> TrackId { get; set; } = null!;

        public SelectAudioTrackByIdArgs()
        {
        }
        public static new SelectAudioTrackByIdArgs Empty => new SelectAudioTrackByIdArgs();
    }
}
