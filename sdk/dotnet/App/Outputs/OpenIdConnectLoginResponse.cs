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
    /// The configuration settings of the login flow of the custom Open ID Connect provider.
    /// </summary>
    [OutputType]
    public sealed class OpenIdConnectLoginResponse
    {
        /// <summary>
        /// The name of the claim that contains the users name.
        /// </summary>
        public readonly string? NameClaimType;
        /// <summary>
        /// A list of the scopes that should be requested while authenticating.
        /// </summary>
        public readonly ImmutableArray<string> Scopes;

        [OutputConstructor]
        private OpenIdConnectLoginResponse(
            string? nameClaimType,

            ImmutableArray<string> scopes)
        {
            NameClaimType = nameClaimType;
            Scopes = scopes;
        }
    }
}
