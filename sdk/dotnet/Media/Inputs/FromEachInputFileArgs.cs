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
    /// An InputDefinition that looks at each input file provided to select tracks specified by the IncludedTracks property. Generally used with the AudioTrackByAttribute and VideoTrackByAttribute to select tracks from each file given.
    /// </summary>
    public sealed class FromEachInputFileArgs : global::Pulumi.ResourceArgs
    {
        [Input("includedTracks")]
        private InputList<object>? _includedTracks;

        /// <summary>
        /// The list of TrackDescriptors which define the metadata and selection of tracks in the input.
        /// </summary>
        public InputList<object> IncludedTracks
        {
            get => _includedTracks ?? (_includedTracks = new InputList<object>());
            set => _includedTracks = value;
        }

        /// <summary>
        /// The discriminator for derived types.
        /// Expected value is '#Microsoft.Media.FromEachInputFile'.
        /// </summary>
        [Input("odataType", required: true)]
        public Input<string> OdataType { get; set; } = null!;

        public FromEachInputFileArgs()
        {
        }
        public static new FromEachInputFileArgs Empty => new FromEachInputFileArgs();
    }
}
