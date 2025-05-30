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
    /// A Sequence contains an ordered list of Clips where each clip is a JobInput.  The Sequence will be treated as a single input.
    /// </summary>
    [OutputType]
    public sealed class JobInputSequenceResponse
    {
        /// <summary>
        /// JobInputs that make up the timeline.
        /// </summary>
        public readonly ImmutableArray<Outputs.JobInputClipResponse> Inputs;
        /// <summary>
        /// The discriminator for derived types.
        /// Expected value is '#Microsoft.Media.JobInputSequence'.
        /// </summary>
        public readonly string OdataType;

        [OutputConstructor]
        private JobInputSequenceResponse(
            ImmutableArray<Outputs.JobInputClipResponse> inputs,

            string odataType)
        {
            Inputs = inputs;
            OdataType = odataType;
        }
    }
}
