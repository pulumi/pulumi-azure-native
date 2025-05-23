// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VoiceServices.Outputs
{

    /// <summary>
    /// Properties of Custom SIP Headers.
    /// </summary>
    [OutputType]
    public sealed class CustomSipHeadersPropertiesResponse
    {
        /// <summary>
        /// The Custom SIP Headers to apply to the calls which traverse the Communications Gateway
        /// </summary>
        public readonly ImmutableArray<Outputs.CustomSipHeaderResponse> Headers;

        [OutputConstructor]
        private CustomSipHeadersPropertiesResponse(ImmutableArray<Outputs.CustomSipHeaderResponse> headers)
        {
            Headers = headers;
        }
    }
}
