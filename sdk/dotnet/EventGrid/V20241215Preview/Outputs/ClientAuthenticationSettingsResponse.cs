// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.V20241215Preview.Outputs
{

    /// <summary>
    /// Client authentication settings for namespace resource.
    /// </summary>
    [OutputType]
    public sealed class ClientAuthenticationSettingsResponse
    {
        /// <summary>
        /// Alternative authentication name sources related to client authentication settings for namespace resource.
        /// </summary>
        public readonly ImmutableArray<string> AlternativeAuthenticationNameSources;
        /// <summary>
        /// Custom JWT authentication settings for namespace resource.
        /// </summary>
        public readonly Outputs.CustomJwtAuthenticationSettingsResponse? CustomJwtAuthentication;

        [OutputConstructor]
        private ClientAuthenticationSettingsResponse(
            ImmutableArray<string> alternativeAuthenticationNameSources,

            Outputs.CustomJwtAuthenticationSettingsResponse? customJwtAuthentication)
        {
            AlternativeAuthenticationNameSources = alternativeAuthenticationNameSources;
            CustomJwtAuthentication = customJwtAuthentication;
        }
    }
}
