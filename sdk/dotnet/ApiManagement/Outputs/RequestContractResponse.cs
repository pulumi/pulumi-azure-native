// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.Outputs
{

    /// <summary>
    /// Operation request details.
    /// </summary>
    [OutputType]
    public sealed class RequestContractResponse
    {
        /// <summary>
        /// Operation request description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Collection of operation request headers.
        /// </summary>
        public readonly ImmutableArray<Outputs.ParameterContractResponse> Headers;
        /// <summary>
        /// Collection of operation request query parameters.
        /// </summary>
        public readonly ImmutableArray<Outputs.ParameterContractResponse> QueryParameters;
        /// <summary>
        /// Collection of operation request representations.
        /// </summary>
        public readonly ImmutableArray<Outputs.RepresentationContractResponse> Representations;

        [OutputConstructor]
        private RequestContractResponse(
            string? description,

            ImmutableArray<Outputs.ParameterContractResponse> headers,

            ImmutableArray<Outputs.ParameterContractResponse> queryParameters,

            ImmutableArray<Outputs.RepresentationContractResponse> representations)
        {
            Description = description;
            Headers = headers;
            QueryParameters = queryParameters;
            Representations = representations;
        }
    }
}
