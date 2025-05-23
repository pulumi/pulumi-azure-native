// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media.Outputs
{

    /// <summary>
    /// An InputDefinition that looks across all of the files provided to select tracks specified by the IncludedTracks property. Generally used with the AudioTrackByAttribute and VideoTrackByAttribute to allow selection of a single track across a set of input files.
    /// </summary>
    [OutputType]
    public sealed class FromAllInputFileResponse
    {
        /// <summary>
        /// The list of TrackDescriptors which define the metadata and selection of tracks in the input.
        /// </summary>
        public readonly ImmutableArray<object> IncludedTracks;
        /// <summary>
        /// The discriminator for derived types.
        /// Expected value is '#Microsoft.Media.FromAllInputFile'.
        /// </summary>
        public readonly string OdataType;

        [OutputConstructor]
        private FromAllInputFileResponse(
            ImmutableArray<object> includedTracks,

            string odataType)
        {
            IncludedTracks = includedTracks;
            OdataType = odataType;
        }
    }
}
