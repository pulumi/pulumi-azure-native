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
    /// Class to select a track
    /// </summary>
    [OutputType]
    public sealed class TrackSelectionResponse
    {
        /// <summary>
        /// TrackSelections is a track property condition list which can specify track(s)
        /// </summary>
        public readonly ImmutableArray<Outputs.TrackPropertyConditionResponse> TrackSelections;

        [OutputConstructor]
        private TrackSelectionResponse(ImmutableArray<Outputs.TrackPropertyConditionResponse> trackSelections)
        {
            TrackSelections = trackSelections;
        }
    }
}
