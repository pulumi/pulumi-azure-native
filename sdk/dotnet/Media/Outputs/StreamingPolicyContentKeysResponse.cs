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
    /// Class to specify properties of all content keys in Streaming Policy
    /// </summary>
    [OutputType]
    public sealed class StreamingPolicyContentKeysResponse
    {
        /// <summary>
        /// Default content key for an encryption scheme
        /// </summary>
        public readonly Outputs.DefaultKeyResponse? DefaultKey;
        /// <summary>
        /// Representing tracks needs separate content key
        /// </summary>
        public readonly ImmutableArray<Outputs.StreamingPolicyContentKeyResponse> KeyToTrackMappings;

        [OutputConstructor]
        private StreamingPolicyContentKeysResponse(
            Outputs.DefaultKeyResponse? defaultKey,

            ImmutableArray<Outputs.StreamingPolicyContentKeyResponse> keyToTrackMappings)
        {
            DefaultKey = defaultKey;
            KeyToTrackMappings = keyToTrackMappings;
        }
    }
}
