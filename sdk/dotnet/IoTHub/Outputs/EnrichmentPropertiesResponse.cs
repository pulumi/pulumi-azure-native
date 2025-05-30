// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTHub.Outputs
{

    /// <summary>
    /// The properties of an enrichment that your IoT hub applies to messages delivered to endpoints.
    /// </summary>
    [OutputType]
    public sealed class EnrichmentPropertiesResponse
    {
        /// <summary>
        /// The list of endpoints for which the enrichment is applied to the message.
        /// </summary>
        public readonly ImmutableArray<string> EndpointNames;
        /// <summary>
        /// The key or name for the enrichment property.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The value for the enrichment property.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private EnrichmentPropertiesResponse(
            ImmutableArray<string> endpointNames,

            string key,

            string value)
        {
            EndpointNames = endpointNames;
            Key = key;
            Value = value;
        }
    }
}
