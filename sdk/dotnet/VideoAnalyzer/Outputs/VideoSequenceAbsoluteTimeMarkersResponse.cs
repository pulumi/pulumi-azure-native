// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VideoAnalyzer.Outputs
{

    /// <summary>
    /// A sequence of absolute datetime ranges as a string. The datetime values should follow IS08601, and the sum of the ranges should add up to 24 hours or less. Currently, there can be only one range specified in the sequence.
    /// </summary>
    [OutputType]
    public sealed class VideoSequenceAbsoluteTimeMarkersResponse
    {
        /// <summary>
        /// The sequence of datetime ranges. Example: '[["2021-10-05T03:30:00Z", "2021-10-05T03:40:00Z"]]'.
        /// </summary>
        public readonly string Ranges;
        /// <summary>
        /// The discriminator for derived types.
        /// Expected value is '#Microsoft.VideoAnalyzer.VideoSequenceAbsoluteTimeMarkers'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private VideoSequenceAbsoluteTimeMarkersResponse(
            string ranges,

            string type)
        {
            Ranges = ranges;
            Type = type;
        }
    }
}
