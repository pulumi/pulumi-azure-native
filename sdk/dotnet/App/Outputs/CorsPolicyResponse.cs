// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Outputs
{

    /// <summary>
    /// Cross-Origin-Resource-Sharing policy
    /// </summary>
    [OutputType]
    public sealed class CorsPolicyResponse
    {
        /// <summary>
        /// Specifies whether the resource allows credentials
        /// </summary>
        public readonly bool? AllowCredentials;
        /// <summary>
        /// Specifies the content for the access-control-allow-headers header
        /// </summary>
        public readonly ImmutableArray<string> AllowedHeaders;
        /// <summary>
        /// Specifies the content for the access-control-allow-methods header
        /// </summary>
        public readonly ImmutableArray<string> AllowedMethods;
        /// <summary>
        /// Specifies the content for the access-control-allow-origins header
        /// </summary>
        public readonly ImmutableArray<string> AllowedOrigins;
        /// <summary>
        /// Specifies the content for the access-control-expose-headers header 
        /// </summary>
        public readonly ImmutableArray<string> ExposeHeaders;
        /// <summary>
        /// Specifies the content for the access-control-max-age header
        /// </summary>
        public readonly int? MaxAge;

        [OutputConstructor]
        private CorsPolicyResponse(
            bool? allowCredentials,

            ImmutableArray<string> allowedHeaders,

            ImmutableArray<string> allowedMethods,

            ImmutableArray<string> allowedOrigins,

            ImmutableArray<string> exposeHeaders,

            int? maxAge)
        {
            AllowCredentials = allowCredentials;
            AllowedHeaders = allowedHeaders;
            AllowedMethods = allowedMethods;
            AllowedOrigins = allowedOrigins;
            ExposeHeaders = exposeHeaders;
            MaxAge = maxAge;
        }
    }
}
