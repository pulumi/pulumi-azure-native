// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HealthcareApis.Outputs
{

    /// <summary>
    /// The settings for the CORS configuration of the service instance.
    /// </summary>
    [OutputType]
    public sealed class ServiceCorsConfigurationInfoResponse
    {
        /// <summary>
        /// If credentials are allowed via CORS.
        /// </summary>
        public readonly bool? AllowCredentials;
        /// <summary>
        /// The headers to be allowed via CORS.
        /// </summary>
        public readonly ImmutableArray<string> Headers;
        /// <summary>
        /// The max age to be allowed via CORS.
        /// </summary>
        public readonly int? MaxAge;
        /// <summary>
        /// The methods to be allowed via CORS.
        /// </summary>
        public readonly ImmutableArray<string> Methods;
        /// <summary>
        /// The origins to be allowed via CORS.
        /// </summary>
        public readonly ImmutableArray<string> Origins;

        [OutputConstructor]
        private ServiceCorsConfigurationInfoResponse(
            bool? allowCredentials,

            ImmutableArray<string> headers,

            int? maxAge,

            ImmutableArray<string> methods,

            ImmutableArray<string> origins)
        {
            AllowCredentials = allowCredentials;
            Headers = headers;
            MaxAge = maxAge;
            Methods = methods;
            Origins = origins;
        }
    }
}
