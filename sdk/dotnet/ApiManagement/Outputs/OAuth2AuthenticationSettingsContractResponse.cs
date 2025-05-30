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
    /// API OAuth2 Authentication settings details.
    /// </summary>
    [OutputType]
    public sealed class OAuth2AuthenticationSettingsContractResponse
    {
        /// <summary>
        /// OAuth authorization server identifier.
        /// </summary>
        public readonly string? AuthorizationServerId;
        /// <summary>
        /// operations scope.
        /// </summary>
        public readonly string? Scope;

        [OutputConstructor]
        private OAuth2AuthenticationSettingsContractResponse(
            string? authorizationServerId,

            string? scope)
        {
            AuthorizationServerId = authorizationServerId;
            Scope = scope;
        }
    }
}
