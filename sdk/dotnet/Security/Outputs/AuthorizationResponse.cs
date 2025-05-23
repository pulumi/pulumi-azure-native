// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Outputs
{

    /// <summary>
    /// Authorization payload.
    /// </summary>
    [OutputType]
    public sealed class AuthorizationResponse
    {
        /// <summary>
        /// Gets or sets one-time OAuth code to exchange for refresh and access tokens.
        /// 
        /// Only used during PUT/PATCH operations. The secret is cleared during GET.
        /// </summary>
        public readonly string? Code;

        [OutputConstructor]
        private AuthorizationResponse(string? code)
        {
            Code = code;
        }
    }
}
