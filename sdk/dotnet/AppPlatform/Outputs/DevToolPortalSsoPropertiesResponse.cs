// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform.Outputs
{

    /// <summary>
    /// Single sign-on related configuration
    /// </summary>
    [OutputType]
    public sealed class DevToolPortalSsoPropertiesResponse
    {
        /// <summary>
        /// The public identifier for the application
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// The secret known only to the application and the authorization server
        /// </summary>
        public readonly string? ClientSecret;
        /// <summary>
        /// The URI of a JSON file with generic OIDC provider configuration.
        /// </summary>
        public readonly string? MetadataUrl;
        /// <summary>
        /// It defines the specific actions applications can be allowed to do on a user's behalf
        /// </summary>
        public readonly ImmutableArray<string> Scopes;

        [OutputConstructor]
        private DevToolPortalSsoPropertiesResponse(
            string? clientId,

            string? clientSecret,

            string? metadataUrl,

            ImmutableArray<string> scopes)
        {
            ClientId = clientId;
            ClientSecret = clientSecret;
            MetadataUrl = metadataUrl;
            Scopes = scopes;
        }
    }
}
