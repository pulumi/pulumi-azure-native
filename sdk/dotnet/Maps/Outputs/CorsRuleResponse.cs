// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Maps.Outputs
{

    /// <summary>
    /// Specifies a CORS rule for the Map Account.
    /// </summary>
    [OutputType]
    public sealed class CorsRuleResponse
    {
        /// <summary>
        /// Required if CorsRule element is present. A list of origin domains that will be allowed via CORS, or "*" to allow all domains
        /// </summary>
        public readonly ImmutableArray<string> AllowedOrigins;

        [OutputConstructor]
        private CorsRuleResponse(ImmutableArray<string> allowedOrigins)
        {
            AllowedOrigins = allowedOrigins;
        }
    }
}
